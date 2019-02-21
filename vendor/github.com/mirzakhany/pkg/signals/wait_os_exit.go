package signals

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitExitSignal get os signals
func WaitExitSignal() os.Signal {
	quit := make(chan os.Signal, 6)
	signal.Notify(quit, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	return <-quit
}
