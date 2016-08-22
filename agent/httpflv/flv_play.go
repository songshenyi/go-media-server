package httpflv

import (
	"github.com/songshenyi/go-media-server/core"
	"net/http"
	"github.com/songshenyi/go-media-server/avformat"
	"github.com/songshenyi/go-media-server/logger"
	"github.com/songshenyi/go-media-server/agent"
	"sync"
)

type HttpFlvPlayAgent struct{
	ctx core.Context
	req *http.Request
	writer http.ResponseWriter
	source agent.Agent


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

func (v *HttpFlvPlayAgent)RegisterSource(source agent.Agent) (err error){
	v.source = source
	return source.RegisterDest(v)
}

func (v *HttpFlvPlayAgent)UnRegisterSource(source agent.Agent) (err error){
	return
}

func (v *HttpFlvPlayAgent)GetSource() (source agent.Agent){
	return
}

func (v *HttpFlvPlayAgent)RegisterDest(dest agent.Agent) (err error){
	return
}

func (v *HttpFlvPlayAgent)UnRegisterDest(dest agent.Agent) (err error){
	return
}

type FLVIngestAgent struct {

}

type FLVForwardAgent struct {

}


