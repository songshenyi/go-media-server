package application

import (
	"net/http"
	"github.com/songshenyi/go-media-server/server"
	"github.com/songshenyi/go-media-server/logger"
	"github.com/songshenyi/go-media-server/avformat"
	"github.com/songshenyi/go-media-server/agent"
	"github.com/songshenyi/go-media-server/core"
)

func AddHandle(httpServer *server.HttpServer){
	httpServer.HandleMap["/live/{name}"] = LiveHandler
}

func LiveHandler(w http.ResponseWriter, r *http.Request){
	logger.Debug(r.Method)
	ctx := core.NewContext()
	//var buf1 bytes.Buffer

	//buf := make([]byte, 10240)

	if(r.Method == "PUT" || r.Method == "POST"){
		if _, err := agent.Manager.NewHttpFlvPublishAgent(ctx, r); err != nil{
			logger.Warn("create HttpFlvPublishAgent failed", err)
			return
		}
	}else if r.Method == "GET"{

	}

	header, err := avformat.ReadFlvHeader(r.Body)
	if  err != nil{
		logger.Warn(err)
	}
	logger.Info(header)

	for{
		tag, err := avformat.ReadFlvTag(r.Body)
		if  err != nil{
			logger.Warn(err)
			break;
		}
		logger.Info(tag.TagType, tag.TimeStamp, tag.DataSize)

		//len, err := r.Body.Read(buf)
		//if err !=nil{
		//	logger.Debug(len)
		//	logger.Error(err)
		//	break;
		//}

	//	log.Debugf("%d, %d",len, buf[0])
	}
}
