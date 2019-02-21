package gorender

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mirzakhany/pkg/logger"
	"github.com/mirzakhany/pkg/version"
	"github.com/mirzakhany/gorender/internal/global"
	"github.com/mirzakhany/gorender/internal/pkg/config"
	"github.com/mirzakhany/gorender/internal/pkg/framework/router"
	"golang.org/x/sync/errgroup"
)

// version params
var (
	LongHash    string
	ShortHash   string
	CommitDate  string
	CommitCount string
	BuildDate   string
	VersionStr  string
)

// InitGoRender init goRender
func InitGoRender(appName string) func() {
	_, cnl := context.WithCancel(context.Background())
	err := initModules(appName)
	if err != nil {
		log.Fatalf("init modules failed with error: %v", err)
	}
	return func() {
		cnl()
		<-time.After(1 * time.Second)
	}
}

func printAppLogo(versionCode string) {
	var appLogo = fmt.Sprintf(`
  ▄████  ▒█████   ██▀███  ▓█████  ███▄    █ ▓█████▄ ▓█████  ██▀███  
 ██▒ ▀█▒▒██▒  ██▒▓██ ▒ ██▒▓█   ▀  ██ ▀█   █ ▒██▀ ██▌▓█   ▀ ▓██ ▒ ██▒
▒██░▄▄▄░▒██░  ██▒▓██ ░▄█ ▒▒███   ▓██  ▀█ ██▒░██   █▌▒███   ▓██ ░▄█ ▒
░▓█  ██▓▒██   ██░▒██▀▀█▄  ▒▓█  ▄ ▓██▒  ▐▌██▒░▓█▄   ▌▒▓█  ▄ ▒██▀▀█▄  
░▒▓███▀▒░ ████▓▒░░██▓ ▒██▒░▒████▒▒██░   ▓██░░▒████▓ ░▒████▒░██▓ ▒██▒
 ░▒   ▒ ░ ▒░▒░▒░ ░ ▒▓ ░▒▓░░░ ▒░ ░░ ▒░   ▒ ▒  ▒▒▓  ▒ ░░ ▒░ ░░ ▒▓ ░▒▓░
  ░   ░   ░ ▒ ▒░   ░▒ ░ ▒░ ░ ░  ░░ ░░   ░ ▒░ ░ ▒  ▒  ░ ░  ░  ░▒ ░ ▒░
░ ░   ░ ░ ░ ░ ▒    ░░   ░    ░      ░   ░ ░  ░ ░  ░    ░     ░░   ░ 
      ░     ░ ░     ░        ░  ░         ░    ░       ░  ░   ░     
                                             ░
  GoRender %s`, versionCode)
	fmt.Println(appLogo)
}

func initModules(appName string) error {

	versionSettings := &version.Settings{
		CopyRightYear: "2019",
		LongHash:      LongHash,
		ShortHash:     ShortHash,
		CommitDate:    CommitDate,
		CommitCount:   CommitCount,
		BuildDate:     BuildDate,
		Version:       VersionStr,
		ServiceName:   appName,
		CompanyName:   "mirzakhany",
	}

	version.SetupVersion(versionSettings)
	printAppLogo(version.GetVersion())
	version.PrintServiceVersion()

	var err error

	global.AppConf, err = config.LoadConf(appName, "")
	if err != nil {
		log.Fatalf("Load json config file error: '%v'", err)
		return err
	}

	logSetting := logger.LogSettings{
		LogFormat:   global.AppConf.Log.Format,
		AccessLevel: global.AppConf.Log.AccessLevel,
		AccessLog:   global.AppConf.Log.AccessLog,
		ErrorLevel:  global.AppConf.Log.ErrorLevel,
		ErrorLog:    global.AppConf.Log.ErrorLog,
		SentryDNS:   global.AppConf.Log.SentryDSN,
	}

	if err = logger.InitLog(logSetting); err != nil {
		log.Fatalf("Can't load log module, error: %v", err)
		return err
	}

	var g errgroup.Group

	g.Go(func() error {
		return router.InitRouter()
	})

	fmt.Println(fmt.Sprintf("[BOOT]Server Start On %s:%d", global.AppConf.Core.Address, global.AppConf.Core.Port))
	if err = g.Wait(); err != nil {
		logger.Fatal(err)
	}

	return err
}
