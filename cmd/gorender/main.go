package main

import (
	"github.com/mirzakhany/pkg/signals"
	"github.com/sirupsen/logrus"
	"github.com/mirzakhany/gorender/internal/app/gorender"
)

const appName = "gorender"

func main() {
	defer gorender.InitGoRender(appName)()
	sig := signals.WaitExitSignal()
	logrus.Infof("%s received, exiting...", sig.String())
}
