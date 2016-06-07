package agent

import (
	"github.com/songshenyi/go-media-server/core"
	"github.com/songshenyi/go-media-server/avformat"
)

type CopyAgent struct{
	ctx core.Context
	source Agent
	dest []Agent

	header *avformat.FlvHeader
	metadata *avformat.FlvMessage
	videoSequenceHeader *avformat.FlvMessage
	audioSequenceHeader *avformat.FlvMessage
}

func NewCopyAgent(ctx core.Context) Agent{
	return &CopyAgent{
		ctx: ctx,
		dest: make([]Agent, 0),
	}
}

func (v* CopyAgent)Open() (err error){
	return
}

func (v* CopyAgent)Close() (err error){
	return
}

func (v* CopyAgent)Pump() (err error){
	return
}

func (v* CopyAgent)Write(m Message) (err error){
	return
}

func (v* CopyAgent)RegisterSource(source Agent) (err error){
	v.source = source
	return source.RegisterDest(v)
}

func (v* CopyAgent)UnRegisterSource(source Agent) (err error){
	return
}

func (v* CopyAgent)GetSource() (source Agent){
	return v.source
}

func (v* CopyAgent)RegisterDest(dest Agent) (err error){
	return
}

func (v* CopyAgent)UnRegisterDest(dest Agent) (err error){
	return
}