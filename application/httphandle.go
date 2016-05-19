package application

import (
	log "github.com/cihub/seelog"
	"net/http"
	"github.com/songshenyi/go-media-server/server"
)

func AddHandle(httpServer *server.HttpServer){
	httpServer.HandleMap["/live/"] = LiveHandler
}

func LiveHandler(w http.ResponseWriter, r *http.Request){
	log.Debug(r.Method)
	buf := make([]byte, 10240)
	for{
		len, err := r.Body.Read(buf)
		if err !=nil{
			log.Debug(len)
			log.Error(err)
			break;
		}
		log.Debugf("%d, %d",len, buf[0])
	}
}
