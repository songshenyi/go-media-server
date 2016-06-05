package application

import (
	"net/http"
	"github.com/songshenyi/go-media-server/server"
	"github.com/songshenyi/go-media-server/logger"
)

func AddHandle(httpServer *server.HttpServer){
	httpServer.HandleMap["/live/{name}"] = LiveHandler
}

func LiveHandler(w http.ResponseWriter, r *http.Request){
	logger.Debug(r.Method)
	//var buf1 bytes.Buffer

	buf := make([]byte, 10240)
	for{
		len, err := r.Body.Read(buf)
		if err !=nil{
			logger.Debug(len)
			logger.Error(err)
			break;
		}

	//	log.Debugf("%d, %d",len, buf[0])
	}
}
