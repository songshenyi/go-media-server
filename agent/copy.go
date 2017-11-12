package agent

import (
	"github.com/songshenyi/go-media-server/core"
	"github.com/songshenyi/go-media-server/avformat"
	"github.com/songshenyi/go-media-server/logger"
)

type CopyAgent struct{
	ctx core.Context
	source Agent
	dest []Agent

	header *avformat.FlvMessage
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

func (v* CopyAgent)Write(m *avformat.FlvMessage) (err error){

	if m.Header != nil{
		v.header = m.Copy()
		logger.Debug("write flv header")
	}else if m.MetaData{
		v.metadata= m.Copy()
		logger.Debug("write flv metadata")
	}else if m.AudioSequenceHeader{
		v.audioSequenceHeader = m.Copy()
		logger.Debug("write flv aac0")
	} else if m.VideoSequenceHeader{
		v.videoSequenceHeader = m.Copy()
		logger.Debug("write flv avc0")
	}

	for _, d :=range v.dest{
		d.Write(m.Copy())
	}

	// for single core, manually sched to send more.
	//if Workers == 1 {
	//	runtime.Gosched()
	//}

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
	logger.Debug("register dest")
	v.dest = append(v.dest, dest)

	if v.header != nil{
		if err = dest.Write(v.header.Copy()); err != nil{
			logger.Warn(err); return
		}
		logger.Debug("copy flv header")
	}
	if v.metadata != nil{
		if err = dest.Write(v.metadata.Copy()); err != nil{
			logger.Warn(err); return
		}
		logger.Debug("copy flv metadata tag")
	}
	if v.audioSequenceHeader != nil{
		if err = dest.Write(v.audioSequenceHeader.Copy()); err != nil{
			logger.Warn(err); return
		}
		logger.Debug("copy flv aac0")
	}
	if v.videoSequenceHeader != nil{
		if err = dest.Write(v.videoSequenceHeader.Copy()); err !=nil{
			logger.Warn(err); return
		}
		logger.Debug("copy flv avc0")
	}

	return
}

func (v* CopyAgent)UnRegisterDest(dest Agent) (err error){
	return
}