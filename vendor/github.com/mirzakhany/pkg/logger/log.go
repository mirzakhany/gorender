package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	reset   = string([]byte{27, 91, 48, 109})

	// LogAccess is log server request log
	LogAccess *logrus.Logger
	// LogError is log server error log
	LogError *logrus.Logger

	// LogToSentry flag to send errors to sentry or not
	LogToSentry bool

	isTerm bool

	settings LogSettings
)

// LogReq is http request log
type LogReq struct {
	URI         string `json:"uri"`
	Method      string `json:"method"`
	IP          string `json:"ip"`
	ContentType string `json:"content_type"`
	Agent       string `json:"agent"`
}

type LogSettings struct {
	// LogFormat output log format
	LogFormat string
	// AccessLevel level of access log
	AccessLevel string
	// AccessLog output of access log
	AccessLog string
	// ErrorLevel level of error log
	ErrorLevel string
	// ErrorLog output of error log
	ErrorLog string
	// SentryDNS url of sentry
	SentryDNS string
}

func init() {
	isTerm = isatty.IsTerminal(os.Stdout.Fd())
}

// InitLog use for initial log module
func InitLog(logSettings LogSettings) error {
	var err error
	settings = logSettings
	if settings.SentryDNS != "" {
		err := raven.SetDSN(settings.SentryDNS)
		if err != nil {
			return err
		}
		LogToSentry = true
	}

	// init logger
	LogAccess = logrus.New()
	LogError = logrus.New()

	LogAccess.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
		FullTimestamp:   true,
	}

	LogError.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
		FullTimestamp:   true,
	}

	// set logger
	if err = SetLogLevel(LogAccess, logSettings.AccessLevel); err != nil {
		return errors.New("Set access log level error: " + err.Error())
	}

	if err = SetLogLevel(LogError, logSettings.ErrorLevel); err != nil {
		return errors.New("Set error log level error: " + err.Error())
	}

	if err = SetLogOut(LogAccess, logSettings.AccessLog); err != nil {
		return errors.New("Set access log path error: " + err.Error())
	}

	if err = SetLogOut(LogError, logSettings.ErrorLog); err != nil {
		return errors.New("Set error log path error: " + err.Error())
	}

	return nil
}

// SetLogOut provide log stdout and stderr output
func SetLogOut(log *logrus.Logger, outString string) error {
	switch outString {
	case "stdout":
		log.Out = os.Stdout
	case "stderr":
		log.Out = os.Stderr
	default:
		l := &lumberjack.Logger{
			Filename:   outString,
			MaxSize:    50,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   true,
		}
		log.Out = l
	}

	return nil
}

// SetLogLevel is define log level what you want
// log level: panic, fatal, error, warn, info and debug
func SetLogLevel(log *logrus.Logger, levelString string) error {
	level, err := logrus.ParseLevel(levelString)

	if err != nil {
		return err
	}

	log.Level = level

	return nil
}

// LogRequest record http request
func LogRequest(uri string, method string, ip string, contentType string, agent string) {
	var output string
	log := &LogReq{
		URI:         uri,
		Method:      method,
		IP:          ip,
		ContentType: contentType,
		Agent:       agent,
	}

	if settings.LogFormat == "json" {
		logJSON, _ := json.Marshal(log)

		output = string(logJSON)
	} else {
		var headerColor, resetColor string

		if isTerm {
			headerColor = magenta
			resetColor = reset
		}

		// format is string
		output = fmt.Sprintf("|%s header %s| %s %s %s %s %s",
			headerColor, resetColor,
			log.Method,
			log.URI,
			log.IP,
			log.ContentType,
			log.Agent,
		)
	}

	LogAccess.Info(output)
}

// LogMiddleware provide gin router handler.
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		LogRequest(c.Request.URL.Path, c.Request.Method, c.ClientIP(), c.ContentType(), c.GetHeader("User-Agent"))
		c.Next()
	}
}
