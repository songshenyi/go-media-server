package application

import (
	"net/http"
	"github.com/songshenyi/go-media-server/server"
	"github.com/songshenyi/go-media-server/logger"
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
	var httpAgent agent.Agent
	var err error
	if(r.Method == "PUT" || r.Method == "POST"){
		httpAgent, err = agent.Manager.NewHttpFlvPublishAgent(ctx, r, w)
		if err != nil{
			logger.Warn("create HttpFlvPublishAgent failed", err)
			return
		}
	}else if r.Method == "GET" {
		httpAgent, err = agent.Manager.NewHttpFlvPlayAgent(ctx, r, w)
		if err != nil {
			logger.Warn("create HttpFlvPlayAgent failed", err)
			return
		}
	}

	err = httpAgent.Pump()
	if err != nil{
		logger.Info("agent exit", err)
		return
	}

	return
}
