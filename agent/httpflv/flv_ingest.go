package httpflv

import (
	"github.com/songshenyi/go-media-server/core"
	"github.com/songshenyi/go-media-server/avformat"
	"github.com/songshenyi/go-media-server/logger"
	"github.com/songshenyi/go-media-server/agent"
	"sync"
	"net/http"
)
type HttpFlvIngestAgent struct{
	ctx core.Context
	url string
	client *http.Client
	req *http.Request

	dest agent.Agent

	writeBuf chan *avformat.FlvMessage
	writeLock sync.Mutex
}

func NewFLVIngestAgent(ctx core.Context, url string)(*HttpFlvIngestAgent){
	return nil
}


func (v *HttpFlvIngestAgent)Open() (err error){
	logger.Info("Open http flv ingest agent")
	return
}

func (v*HttpFlvIngestAgent)Close() (err error){
	logger.Info("Close http flv ingest agent")
	return
}

func (v *HttpFlvIngestAgent)Pump() (err error){
	logger.Info("Pump http flv ingest agent")
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

func (v *HttpFlvIngestAgent)Write(m *avformat.FlvMessage) (err error){
	return
}

func ( v*HttpFlvIngestAgent)RegisterSource(source agent.Agent) (err error){
	return
}

func (v *HttpFlvIngestAgent)UnRegisterSource(source agent.Agent) (err error){
	return
}

func (v*HttpFlvIngestAgent)GetSource() (source agent.Agent){
	return
}

func (v*HttpFlvIngestAgent)RegisterDest(dest agent.Agent) (err error){
	v.dest = dest
	return
}

func (v*HttpFlvIngestAgent)UnRegisterDest(dest agent.Agent) (err error){
	return
}
