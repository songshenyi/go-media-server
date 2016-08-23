package httpflv

import (
	"github.com/songshenyi/go-media-server/core"
	"github.com/songshenyi/go-media-server/avformat"
	"github.com/songshenyi/go-media-server/logger"
	"github.com/songshenyi/go-media-server/agent"
	"net/http"
	"net"
	"time"
	"fmt"
)
type HttpFlvIngestAgent struct{
	ctx        core.Context
	stream_url string
	client     *http.Client

	dest       agent.Agent

	header     *avformat.FlvMessage
}

func NewFLVIngestAgent(ctx core.Context, url string)(*HttpFlvIngestAgent){
	return &HttpFlvIngestAgent{
		ctx: ctx,
		stream_url: url,
	}
}


func (v *HttpFlvIngestAgent)Open() (err error){
	logger.Info("Open http flv ingest agent")
	v.client = &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string)(net.Conn, error){
				conn, err := net.DialTimeout(network, addr, time.Second * 10)
				if err != nil{
					return nil, err
				}
				return conn, nil
			},
			DisableKeepAlives: true,
		},
	}
	return
}

func (v*HttpFlvIngestAgent)Close() (err error){
	logger.Info("Close http flv ingest agent")
	return
}

func (v *HttpFlvIngestAgent)Pump() (err error){
	logger.Info("Pump http flv ingest agent")
	req, _ := http.NewRequest(http.MethodGet, v.stream_url, nil)

	go func(){
		timer := time.NewTicker(time.Second * 10)
		for{
			select{
			case <-timer.C:
				//v.client.Transport.(*http.Transport).CancelRequest(req)
			}
		}
	}()

	resp, err := v.client.Do(req)
	if err!= nil {
		logger.Warn(err)
		return
	}

	header, err := avformat.ReadFlvHeader(resp.Body)
	if  err != nil{
		logger.Warn(err)
		return err
	}

	v.header, err = header.ToMessage()
	v.dest.Write(v.header)

	for {
		tag, err := avformat.ReadFlvTag(resp.Body)
		if  err != nil{
			logger.Warn(err)
			return err
		}
		logger.Trace(tag.TagType, tag.TimeStamp, tag.DataSize)
		msg, err := tag.ToMessage()

		if tag.TagType == 9 {
			frameType := avformat.RtmpAVCFrame((tag.Payload[0] >> 4) & 0x0f)
			if frameType == avformat.RtmpKeyFrame {
				fmt.Println(time.Now().UnixNano()/ 1000000, frameType, tag.TimeStamp, tag.DataSize)
			}
		}

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
