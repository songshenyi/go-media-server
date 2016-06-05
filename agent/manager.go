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

func (v *FlvAgentManager) NewHttpFlvPublishAgent(ctx core.Context, r *http.Request)(publishAgent Agent,  err error){
	var copyAgent Agent
	if copyAgent, err = v.getCopyAgent(ctx, r.RequestURI()); err != nil{
		logger.Warnf()
	}
}

func (v *FlvAgentManager) getCopyAgent(ctx core.Context, uri string) (copyAgent Agent, err error){
	if copyAgent, err = v.sources[uri]; err != nil{
		copyAgent = NewCopyAgent(ctx)
		v.sources[uri] = copyAgent
	}
	return
}