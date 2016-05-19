package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/songshenyi/go-media-server/logger"
	"time"
	"github.com/cihub/seelog"
)

type HttpServer struct{
	Port uint16
	HandleMap map[string]func(http.ResponseWriter,*http.Request)
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
		ReadTimeout: 6 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			seelog.Error(err)
		}
	}()
}