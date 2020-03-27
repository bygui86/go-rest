package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-rest/blogpost"
	"github.com/bygui86/go-rest/logging"
)

func main() {
	logging.Log.Info("Start go-rest")

	blogPostServer := startBlogPostServer()

	logging.Log.Info("go-rest up&running")

	startSysCallChannel()

	shutdownAndWait(blogPostServer, 3)
}

func startBlogPostServer() *blogpost.Server {
	logging.Log.Debug("Start BlogPost server")

	server := blogpost.NewRestServer()
	logging.Log.Debug("BlogPost server successfully created")

	server.Start()
	logging.Log.Debug("BlogPost server successfully started")

	return server
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(restServer *blogpost.Server, timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)
	restServer.Shutdown(timeout)
	time.Sleep(time.Duration(timeout+1) * time.Second)
}
