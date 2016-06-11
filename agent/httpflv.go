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

func NewFLVPublishAgent(ctx core.Context, r *http.Request) (*HttpFlvPublishAgent){
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
	return
}

func (v *HttpFlvPublishAgent)Pump() (err error){
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
	req *http.Request
	writer http.ResponseWriter
	source Agent

	writeBuf []*avformat.FlvMessage
	writeLock sync.Mutex
}

func (v *HttpFlvPlayAgent)Open() (err error){
	logger.Info("Open http flv play agent")
	return
}

func (v *HttpFlvPlayAgent)Close() (err error){
	return
}

func (v *HttpFlvPlayAgent)Pump() (err error){
	v.writeLock.Lock()
	writeBuf := v.writeBuf[:]
	v.writeBuf = v.writeBuf[0:0]
	v.writeLock.Unlock()

	for _, m := range writeBuf{
		if(m.Header != nil){
			data, _ := m.Header.ToData()
			v.writer.Write(data)
		}
		if (m.Tag != nil) {
			v.writer.Write(m.Tag.Payload)
		}
	}
	return
}

func (v *HttpFlvPlayAgent)Write(m *avformat.FlvMessage) (err error){
	return
}

func (v *HttpFlvPlayAgent)RegisterSource(source Agent) (err error){
	return
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


