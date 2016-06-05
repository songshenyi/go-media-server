package agent

import (
	"github.com/songshenyi/go-media-server/core"
	"github.com/songshenyi/go-media-server/avformat"
)

type CopyAgent struct{
	ctx core.Context
	source Agent
	dest []Agent

	metadata *avformat.FLVMessage
	videoSequenceHeader *avformat.FLVMessage
	audioSequenceHeader *avformat.FLVMessage
}

func NewCopyAgent(ctx core.Context) Agent{
	return &CopyAgent{
		ctx: ctx,
		dest: make([]Agent, 0),
	}
}

