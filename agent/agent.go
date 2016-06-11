package agent

import (
	"fmt"
	"github.com/songshenyi/go-media-server/avformat"
)

// the muxer of message type.
type MessageMuxer uint8

const (
	MuxerRtmp MessageMuxer = iota
	MuxerFlv

)

// the common structure for RTMP/FLV/HLS/MP4 or any
// message, it can be media message or control message.
// the message flow from agent to another agent.
type Message interface {
	fmt.Stringer

	// the muxer of message.
	Muxer() MessageMuxer
}



type Agent interface {

	Open() (err error)
	Close() (err error)

	// do agent jobs, to pump messages
	// from source to dest.
	Pump() (err error)
	// write to dest, from source.
	Write(m *avformat.FlvMessage) (err error)

	RegisterSource(source Agent) (err error)

	UnRegisterSource(source Agent) (err error)

	GetSource() (source Agent)

	RegisterDest(dest Agent) (err error)

	UnRegisterDest(dest Agent) (err error)
}
