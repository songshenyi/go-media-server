package agent

import (
	"github.com/songshenyi/go-media-server/core"
	"net/http"
	"github.com/songshenyi/go-media-server/avformat"
	"github.com/songshenyi/go-media-server/logger"
	"sync"
)



type HttpFlvPublishAgent struct{
	ctx core.Context
	req *http.Request

	dest Agent

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

func ( v*HttpFlvPublishAgent)RegisterSource(source Agent) (err error){
	return
}

func (v *HttpFlvPublishAgent)UnRegisterSource(source Agent) (err error){
	return
}

func (v*HttpFlvPublishAgent)GetSource() (source Agent){
	return
}

func (v*HttpFlvPublishAgent)RegisterDest(dest Agent) (err error){
	v.dest = dest
	return
}

func (v*HttpFlvPublishAgent)UnRegisterDest(dest Agent) (err error){
	return
}

type HttpFlvPlayAgent struct{
	ctx core.Context
	req *http.Request
	writer http.ResponseWriter
	source Agent


	writeBuf chan *avformat.FlvMessage
	writeLock sync.Mutex
}

func NewFLVPlayAgent(ctx core.Context, r *http.Request, w http.ResponseWriter) (*HttpFlvPlayAgent){
	return &HttpFlvPlayAgent{
		ctx: ctx,
		req: r,
		writer: w,
		writeBuf: make(chan *avformat.FlvMessage, 10),
	}
}

func (v *HttpFlvPlayAgent)Open() (err error){
	logger.Info("Open http flv play agent")
	return
}

func (v *HttpFlvPlayAgent)Close() (err error){
	return
}

func (v *HttpFlvPlayAgent)Pump() (err error){
	for {
		select {
		case m := <-v.writeBuf:
			if (m.Header != nil) {
				data, _ := m.Header.ToData()
				logger.Info("write flv header")
				v.writer.Write(data)
			}
			if (m.Tag != nil) {
				logger.Trace("write flv tag", m.Tag.TimeStamp)
				if tagHeader, err := m.Tag.TagHeaderBytes(); err !=nil{
					logger.Warn("get tag header failed", err)
					return err
				}else{
					v.writer.Write(tagHeader)
				}

				v.writer.Write(m.Tag.Payload)

				if preTagSize, err := m.Tag.PreTagSizeBytes(); err != nil{
					logger.Warn("get tag pretagsize failed", err)
					return err
				}else{
					v.writer.Write(preTagSize)
				}
			}
		}
	}
	return
}

func (v *HttpFlvPlayAgent)Write(m *avformat.FlvMessage) (err error){
	select {
	case v.writeBuf <- m:
	default:
	}
	return
}

func (v *HttpFlvPlayAgent)RegisterSource(source Agent) (err error){
	v.source = source
	return source.RegisterDest(v)
}

func (v *HttpFlvPlayAgent)UnRegisterSource(source Agent) (err error){
	return
}

func (v *HttpFlvPlayAgent)GetSource() (source Agent){
	return
}

func (v *HttpFlvPlayAgent)RegisterDest(dest Agent) (err error){
	return
}

func (v *HttpFlvPlayAgent)UnRegisterDest(dest Agent) (err error){
	return
}

type FLVIngestAgent struct {

}

type FLVForwardAgent struct {

}


