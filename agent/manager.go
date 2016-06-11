package agent

import (
	"github.com/songshenyi/go-media-server/core"
	"sync"
	"net/http"
	"github.com/songshenyi/go-media-server/logger"
)

type FlvAgentManager struct {
	ctx core.Context
	sources map[string]Agent
	lock sync.Mutex
}

var Manager *FlvAgentManager

func NewManager(ctx core.Context) *FlvAgentManager {
	return &FlvAgentManager{
		ctx: ctx,
		sources: make(map[string]Agent),
	}
}

func (v *FlvAgentManager)Close(){

}

func (v *FlvAgentManager) NewHttpFlvPublishAgent(ctx core.Context, r *http.Request, w http.ResponseWriter)(publishAgent Agent,  err error){
	v.lock.Lock()
	defer v.lock.Unlock()

	var copyAgent Agent
	if copyAgent, err = v.getCopyAgent(ctx, r.RequestURI); err != nil{
		logger.Warnf("getCopyAgent failed %s", err.Error())
		return publishAgent, err
	}

	if copyAgent.GetSource() != nil{
		err = PublishConfilictError
		logger.Warnf("conflict %s", err)
		return publishAgent, err
	}

	publishAgent = NewFLVPublishAgent(ctx, r, w)

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

func (v *FlvAgentManager) NewHttpFlvPlayAgent(ctx core.Context, r *http.Request, w http.ResponseWriter)(playAgent Agent, err error){
	v.lock.Lock()
	defer v.lock.Unlock()

	var copyAgent Agent
	if copyAgent, err = v.getCopyAgent(ctx, r.RequestURI); err != nil{
		logger.Warnf("getCopyAgent %s", err.Error())
		return playAgent, err
	}

	playAgent = NewFLVPlayAgent(ctx, r, w)
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

func (v *FlvAgentManager) getCopyAgent(ctx core.Context, uri string) (copyAgent Agent, err error){
	var ok bool
	if copyAgent, ok = v.sources[uri]; !ok {
		copyAgent = NewCopyAgent(ctx)
		v.sources[uri] = copyAgent
	}
	return
}