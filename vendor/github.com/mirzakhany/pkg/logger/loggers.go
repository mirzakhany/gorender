package logger

import (
	"fmt"

	"github.com/getsentry/raven-go"
	"runtime/debug"
)

func addStack(msg string) string {
	stack := debug.Stack()
	return msg + "\n" + string(stack)
}

func Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if LogToSentry {
		raven.CaptureMessage(addStack(msg), nil)
	}
	LogAccess.Debug(msg)
}

func Infof(format string, args ...interface{}) {
	LogAccess.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	LogAccess.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if LogToSentry {
		raven.CaptureMessage(addStack(msg), nil)
	}
	LogError.Warn(msg)
}

func Warningf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if LogToSentry {
		raven.CaptureMessage(addStack(msg), nil)
	}
	LogError.Warning(msg)
}

func Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if LogToSentry {
		raven.CaptureMessage(addStack(msg), nil)
	}
	LogError.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if LogToSentry {
		raven.CaptureMessage(addStack(msg), nil)
	}
	LogError.Fatal(msg)
}

func Panicf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if LogToSentry {
		raven.CaptureMessage(addStack(msg), nil)
	}
	LogError.Panic(msg)
}

func Debug(args ...interface{}) {
	LogAccess.Debug(args)
}

func Info(args ...interface{}) {
	LogAccess.Info(args)
}

func Print(args ...interface{}) {
	LogAccess.Print(args)
}

func Warn(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogError.Warn(args)
}

func Warning(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogError.Warning(args)
}

func Error(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogError.Error(args)
}

func Fatal(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogError.Fatal(args)
}

func Panic(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogError.Panic(args)
}

func Debugln(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogAccess.Debugln(args)
}

func Infoln(args ...interface{}) {
	LogAccess.Infoln(args...)
}

func Println(args ...interface{}) {
	LogAccess.Println(args...)
}

func Warnln(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogError.Warnln(args)
}

func Warningln(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogError.Warningln(args)
}

func Errorln(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogError.Errorln(args)
}

func Fatalln(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogError.Fatalln(args)
}

func Panicln(args ...interface{}) {
	if LogToSentry {
		raven.CaptureMessage(addStack(fmt.Sprint(args...)), nil)
	}
	LogError.Panicln(args)
}
