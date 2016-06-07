package agent

import (
	"github.com/songshenyi/go-media-server/core"
	"net/http"
	"github.com/songshenyi/go-media-server/avformat"
)

type HttpFlvPublishAgent struct{
	ctx core.Context
	req *http.Request

	header *avformat.FlvHeader
}

func NewFLVPublishAgent(ctx core.Context, r *http.Request) (*HttpFlvPublishAgent){
	return &HttpFlvPublishAgent{
		ctx: ctx,
		req: r,
	}
}

func (v*HttpFlvPublishAgent)Open() (err error){

	return
}

func (v*HttpFlvPublishAgent)Close() (err error){
	return
}

func (v*HttpFlvPublishAgent)Pump() (err error){
	return
}

func (v*HttpFlvPublishAgent)Write(m Message) (err error){
	return
}

func (v*HttpFlvPublishAgent)RegisterSource(source Agent) (err error){
	return
}

func (v*HttpFlvPublishAgent)UnRegisterSource(source Agent) (err error){
	return
}

func (v*HttpFlvPublishAgent)GetSource() (source Agent){
	return
}

func (v*HttpFlvPublishAgent)RegisterDest(dest Agent) (err error){
	return
}

func (v*HttpFlvPublishAgent)UnRegisterDest(dest Agent) (err error){
	return
}

type FLVPlayAgent struct{

}

type FLVIngestAgent struct {

}

type FLVForwardAgent struct {

}


