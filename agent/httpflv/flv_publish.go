package httpflv

import (
	"github.com/songshenyi/go-media-server/core"
	"net/http"
	"github.com/songshenyi/go-media-server/agent"
	"github.com/songshenyi/go-media-server/avformat"
	"github.com/songshenyi/go-media-server/logger"
)

type HttpFlvPublishAgent struct{
	ctx core.Context
	req *http.Request

	dest agent.Agent

	header *avformat.FlvMessage
}

func NewFLVPublishAgent(ctx core.Context, r *http.Request, w http.ResponseWriter) (*HttpFlvPublishAgent){
	return &HttpFlvPublishAgent{
		ctx: ctx,
		req: r,
	}
}

func (v *HttpFlvPublishAgent)Open() (err error){
	logger.Info("Open http flv publish agent")
	return
}

func (v*HttpFlvPublishAgent)Close() (err error){
	logger.Info("Close http flv publish agent")
	return
}

func (v *HttpFlvPublishAgent)Pump() (err error){
	logger.Info("Pump http flv publish agent")
	header, err := avformat.ReadFlvHeader(v.req.Body)
	if  err != nil{
		logger.Warn(err)
		return err
	}

	v.header, err = header.ToMessage()
	v.dest.Write(v.header)

	for {
		tag, err := avformat.ReadFlvTag(v.req.Body)
		if  err != nil{
			logger.Warn(err)
			return err
		}
		logger.Trace(tag.TagType, tag.TimeStamp, tag.DataSize)
		msg, err := tag.ToMessage()
		if (err != nil){
			return err
		}

		v.dest.Write(msg)
	}
	return
}

func (v *HttpFlvPublishAgent)Write(m *avformat.FlvMessage) (err error){
	return
}

func ( v*HttpFlvPublishAgent)RegisterSource(source agent.Agent) (err error){
	return
}

func (v *HttpFlvPublishAgent)UnRegisterSource(source agent.Agent) (err error){
	return
}

func (v*HttpFlvPublishAgent)GetSource() (source agent.Agent){
	return
}

func (v*HttpFlvPublishAgent)RegisterDest(dest agent.Agent) (err error){
	v.dest = dest
	return
}

func (v*HttpFlvPublishAgent)UnRegisterDest(dest agent.Agent) (err error){
	return
}
