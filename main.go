package main

import (
	log "github.com/cihub/seelog"
	"github.com/songshenyi/go-media-server/server"
	"os"
	"os/signal"
	"syscall"
	"runtime"
	"github.com/songshenyi/go-media-server/logger"
	"github.com/songshenyi/go-media-server/application"
)

func signalHandle(){
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan)
	for{
		s:= <- signalChan
		log.Infof("recv signal %d", s)
		switch s{
		case syscall.SIGTERM:
			fallthrough
		case syscall.SIGQUIT:
			buf :=make([]byte, 1<<20)
			runtime.Stack(buf, true)
			log.Infof("killed by signal %d", s)
			log.Infof("goroutine stack \n%s", buf)
			return
		}
	}
}

func main(){
	log.Info("Server Start")
	logger.InitAccessLog("config/access.xml")
	httpServer := server.NewHttpServer(8888)
	application.AddHandle(httpServer)
	httpServer.Start()
	signalHandle()
}
