package application

import (
	"net/http"
	"github.com/songshenyi/go-media-server/server"
	"github.com/songshenyi/go-media-server/logger"
	"github.com/songshenyi/go-media-server/agent"
	agent_manager "github.com/songshenyi/go-media-server/agent/manager"
	"github.com/songshenyi/go-media-server/core"
//	"io/ioutil"
)

func AddHandle(httpServer *server.HttpServer){
	httpServer.HandleMap["/live/{name}"] = LiveHandler
	httpServer.HandleMap["/live/songshenyi/debug"] = DebugHandler
}

func DebugHandler(writer http.ResponseWriter, reader *http.Request){
	//defer reader.Body.Close()
	//body, err := ioutil.ReadAll(reader.Body)
	//if err != nil {
	//	logger.Error(err.Error())
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	ctx := core.NewContext()
	httpAgent, err := agent_manager.Manager.NewHttpFlvIngestAgent(ctx,
		"")
	if err != nil{
		logger.Warn(err)
		return
	}
	err = httpAgent.Pump()
	if err != nil{
		logger.Info("agent exit", err)
		return
	}

	return
}

func LiveHandler(w http.ResponseWriter, r *http.Request){
	logger.Debug(r.Method)
	ctx := core.NewContext()
	//var buf1 bytes.Buffer

	//buf := make([]byte, 10240)
	var httpAgent agent.Agent
	var err error
	if(r.Method == "PUT" || r.Method == "POST"){
		httpAgent, err = agent_manager.Manager.NewHttpFlvPublishAgent(ctx, r, w)
		if err != nil{
			logger.Warn("create HttpFlvPublishAgent failed", err)
			return
		}
	}else if r.Method == "GET" {
		httpAgent, err = agent_manager.Manager.NewHttpFlvPlayAgent(ctx, r, w)
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
