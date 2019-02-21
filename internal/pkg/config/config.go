package config

import (
	"strings"

	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	configName = "config"
)

var defaultConf = []byte(`
{
  "debug": false,
  "address": "0.0.0.0",
  "port": 8040,
  "log": {
    "format": "string",
    "access_log": "stdout",
    "access_level": "debug",
    "error_log": "stderr",
    "error_level": "error",
    "sentry_dsn": ""
  }
}
`)

// ConfJSON is config struct
type ConfJSON struct {
	Core SectionCore `json:"core"`
	Log  SectionLog  `json:"log"`
}

// SectionCore is sub section of config
type SectionCore struct {
	Debug   bool   `json:"debug"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// SectionLog is sub section of config.
type SectionLog struct {
	Format      string `json:"format"`
	AccessLog   string `json:"access_log"`
	AccessLevel string `json:"access_level"`
	ErrorLog    string `json:"error_log"`
	ErrorLevel  string `json:"error_level"`
	SentryDSN   string `json:"sentry_dsn"`
}

// KV key value type
type KV map[string]string

// LoadConf load the config settings
func LoadConf(prefix string, configPath string) (ConfJSON, error) {

	var confJSON ConfJSON
	lowerPrefix := strings.ToLower(prefix)
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	viper.SetEnvPrefix(lowerPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if configPath != "" {
		content, err := ioutil.ReadFile(configPath)
		if err != nil {
			return confJSON, err
		}
		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return confJSON, err
		}
	} else {

		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/" + lowerPrefix + "/")
		viper.AddConfigPath("$HOME/." + lowerPrefix)
		viper.SetConfigName(configName)

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())

			viper.WatchConfig()
			viper.OnConfigChange(func(e fsnotify.Event) {
				fmt.Println("Config file changed:", e.Name)
			})

		} else {
			fmt.Println("load default config ...")
			// load default config
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return confJSON, err
			}
		}
	}
	confJSON.Core.Debug = viper.GetBool("debug")
	confJSON.Core.Address = viper.GetString("address")
	confJSON.Core.Port = viper.GetInt("port")
	confJSON.Log.Format = viper.GetString("log.format")
	confJSON.Log.AccessLog = viper.GetString("log.access_log")
	confJSON.Log.AccessLevel = viper.GetString("log.access_level")
	confJSON.Log.ErrorLog = viper.GetString("log.error_log")
	confJSON.Log.ErrorLevel = viper.GetString("log.error_level")
	confJSON.Log.SentryDSN = viper.GetString("log.sentry_dsn")

	return confJSON, nil
}
