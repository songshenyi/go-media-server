package agent

import (
	"io"
	"fmt"
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

// the opener to open the resource.
type Opener interface {
	// open the resource.
	Open() error
}

// the open and closer for resource manage.
type OpenCloser interface {
	Opener
	io.Closer
}

// the agent contains a source
// which ingest message from upstream sink
// write message to channel
// finally delivery to downstream sink.
//
// the arch for agent is:
//      +-----upstream----+           +---downstream----+
//    --+-source => sink--+--(tie->)--+-source => sink--+--
//      +-----------------+           +-----------------+
//
// @remark all method is sync, user should never assume it's async.
type Agent interface {
	// an agent is a resource manager.
	OpenCloser

	// do agent jobs, to pump messages
	// from source to sink.
	Pump() (err error)
	// write to source, from upstream sink.
	Write(m Message) (err error)

	// source tie to the upstream sink.
	Tie(sink Agent) (err error)
	// destroy the link between source and upstream sink.
	UnTie(sink Agent) (err error)
	// get the tied upstream sink of source.
	TiedSink() (sink Agent)

	// sink flow to the downstream source.
	// @remark internal api, sink.Flow(source) when source.tie(sink).
	Flow(source Agent) (err error)
	// destroy the link between sink and downstream sink.
	UnFlow(source Agent) (err error)
}
