package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-rest/logging"
	"github.com/bygui86/go-rest/rest"
)

func main() {
	logging.Log.Info("Start go-rest")

	restServer := startRestServer()

	logging.Log.Info("go-rest up&running")

	startSysCallChannel()

	shutdownAndWait(restServer, 3)
}

func startRestServer() *rest.Server {
	logging.Log.Debug("Start REST server")

	server := rest.NewRestServer()
	logging.Log.Debug("REST server successfully created")

	server.Start()
	logging.Log.Debug("REST server successfully started")

	return server
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(restServer *rest.Server, timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)
	restServer.Shutdown(timeout)
	time.Sleep(time.Duration(timeout+1) * time.Second)
}
