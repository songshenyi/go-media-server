package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/songshenyi/go-media-server/logger"
	"time"
	log "github.com/cihub/seelog"
	"net"
)

type HttpServer struct{
	Port uint16
	HandleMap map[string]func(http.ResponseWriter,*http.Request)
}

func NewHttpServer(port uint16)(*HttpServer){
	return &HttpServer{
		Port: port,
		HandleMap: make( map[string]func(http.ResponseWriter,*http.Request)),
	}
}

func ConnState(c net.Conn, cs http.ConnState){
	idleTimeout:=time.Second * 10
	//activeTimeout := time.Second * 0
	log.Tracef("%s, %s", c.RemoteAddr(), cs)
	switch cs {
	case http.StateIdle, http.StateNew:
		c.SetReadDeadline(time.Now().Add(idleTimeout) )
	case http.StateActive:
		c.SetReadDeadline(time.Time{})
	}
}

func (s *HttpServer)Start(){
	muxHandler := mux.NewRouter()
	for path, f :=range s.HandleMap{
		muxHandler.HandleFunc(path, f)
	}

	addr := fmt.Sprintf(":%d", s.Port)
	server := http.Server{
		Addr:	addr,
		Handler: logger.LoggingHandler(&logger.Access{}, muxHandler),
		ReadTimeout: 0,
		ConnState: ConnState,
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Error(err)
		}
	}()
}