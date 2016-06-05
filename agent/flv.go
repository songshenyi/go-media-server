package agent

import (
	"github.com/songshenyi/go-media-server/core"
	"net/http"
	"github.com/songshenyi/go-media-server/avformat"
)

type FLVPublishAgent struct{
	ctx core.Context
	req *http.Request

	header *avformat.FLVHeader
}

func NewFLVPublishAgent(ctx core.Context, r *http.Request) (*FLVPublishAgent){
	return &FLVPublishAgent{
		ctx: ctx,
		req: r,
	}
}

func (v* FLVPublishAgent)Open() (err error){

	return
}

func (v* FLVPublishAgent)Close() (err error){
	return
}

func (v* FLVPublishAgent)Pump() (err error){
	return
}

func (v* FLVPublishAgent)Write(m Message) (err error){
	return
}

func (v* FLVPublishAgent)RegisterSource(source Agent) (err error){
	return
}

func (v* FLVPublishAgent)UnRegisterSource(source Agent) (err error){
	return
}

func (v* FLVPublishAgent)GetSource() (source Agent){
	return
}

func (v* FLVPublishAgent)RegisterDest(dest Agent) (err error){
	return
}

func (v* FLVPublishAgent)UnRegisterDest(dest Agent) (err error){
	return
}

type FLVPlayAgent struct{

}

type FLVIngestAgent struct {

}

type FLVForwardAgent struct {

}


