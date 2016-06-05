package agent

import (
	"github.com/songshenyi/go-media-server/core"
	"sync"
	"net/http"
)

type FlvAgentManager struct {
	ctx core.Context
	source map[string]Agent
	lock sync.Mutex
}

var Manager *FlvAgentManager

func NewManager(ctx core.Context) *FlvAgentManager {
	return &FlvAgentManager{
		ctx: ctx,
		source: make(map[string]Agent),
	}
}

func (v *FlvAgentManager)Close(){

}

func (v *FlvAgentManager) NewHttpFlvPublishAgent(ctx core.Context, r *http.Request)(newAgent Agent,  err error){

}