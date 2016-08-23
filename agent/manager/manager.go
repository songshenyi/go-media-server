package manager

import (
	"github.com/songshenyi/go-media-server/core"
	"sync"
	"net/http"
	"github.com/songshenyi/go-media-server/logger"
	"github.com/songshenyi/go-media-server/agent/httpflv"
	"github.com/songshenyi/go-media-server/agent"
)

type FlvAgentManager struct {
	ctx core.Context
	sources map[string]agent.Agent
	lock sync.Mutex
}

var Manager *FlvAgentManager

func NewManager(ctx core.Context) *FlvAgentManager {
	return &FlvAgentManager{
		ctx: ctx,
		sources: make(map[string]agent.Agent),
	}
}

func (v *FlvAgentManager)Close(){

}

func (v *FlvAgentManager) NewHttpFlvIngestAgent(ctx core.Context, url string)(ingestAgent agent.Agent, err error){
	v.lock.Lock()
	defer v.lock.Unlock()

	var copyAgent agent.Agent
	if copyAgent, err = v.getCopyAgent(ctx, url); err != nil{
		logger.Warnf("getCopyAgent failed %s", err.Error())
		return ingestAgent, err
	}

	if copyAgent.GetSource() != nil{
		err = agent.PublishConfilictError
		logger.Warnf("conflict %s", err)
		return ingestAgent, err
	}

	ingestAgent = httpflv.NewFLVIngestAgent(ctx, url)
	if err = ingestAgent.Open();  err != nil{
		logger.Warn("open agent failed", err)
		return ingestAgent, err
	}

	if err = copyAgent.RegisterSource(ingestAgent); err != nil{
		logger.Warn("RegisterSource failed", err)
		return ingestAgent, err
	}

	return ingestAgent, err
}

func (v *FlvAgentManager) NewHttpFlvPublishAgent(ctx core.Context, r *http.Request, w http.ResponseWriter)(publishAgent agent.Agent,  err error){
	v.lock.Lock()
	defer v.lock.Unlock()

	var copyAgent agent.Agent
	if copyAgent, err = v.getCopyAgent(ctx, r.RequestURI); err != nil{
		logger.Warnf("getCopyAgent failed %s", err.Error())
		return publishAgent, err
	}

	if copyAgent.GetSource() != nil{
		err = agent.PublishConfilictError
		logger.Warnf("conflict %s", err)
		return publishAgent, err
	}

	publishAgent = httpflv.NewFLVPublishAgent(ctx, r, w)
	if err = publishAgent.Open();  err != nil{
		logger.Warn("open agent failed", err)
		return publishAgent, err
	}

	if err = copyAgent.RegisterSource(publishAgent); err != nil{
		logger.Warn("RegisterSource failed", err)
		return publishAgent, err
	}

	return publishAgent, err
}

func (v *FlvAgentManager) NewHttpFlvPlayAgent(ctx core.Context, r *http.Request, w http.ResponseWriter)(playAgent agent.Agent, err error){
	v.lock.Lock()
	defer v.lock.Unlock()

	var copyAgent agent.Agent
	if copyAgent, err = v.getCopyAgent(ctx, r.RequestURI); err != nil{
		logger.Warnf("getCopyAgent %s", err.Error())
		return playAgent, err
	}

	playAgent = httpflv.NewFLVPlayAgent(ctx, r, w)
	if err = playAgent.Open(); err != nil{
		logger.Warn("open play agent failed", err)
		return playAgent, err
	}

	if err = playAgent.RegisterSource(copyAgent); err != nil{
		logger.Warn("register source failed", err)
		return playAgent, err
	}
	return playAgent, err
}

func (v *FlvAgentManager) getCopyAgent(ctx core.Context, uri string) (copyAgent agent.Agent, err error){
	var ok bool
	if copyAgent, ok = v.sources[uri]; !ok {
		copyAgent = agent.NewCopyAgent(ctx)
		v.sources[uri] = copyAgent
	}
	return
}