package agent

import (
	"github.com/songshenyi/go-media-server/core"
	"net/http"
)

type FLVPublishAgent struct{
	ctx core.Context
	req *http.Request
}

func NewFLVPublishAgent(ctx core.Context, r *http.Request) (*FLVPublishAgent){
	return &FLVPublishAgent{
		ctx: ctx,
		req: r,
	}
}

type FLVPlayAgent struct{

}

type FLVIngestAgent struct {

}

type FLVForwardAgent struct {

}


