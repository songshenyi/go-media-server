package protocol

import "bytes"

type FLVHeader struct {
	Data [13]byte
}

func (h *FLVHeader)SetAudioFlag(flag bool){
	f := 0x04 * flag
	h.Data[4] = (( h.Data[4] & 0x01) |  f)
}

func (h *FLVHeader)SetVideoFlag(flag bool){
	f := 0x01 * flag
	h.Data[4] = (( h.Data[4] & 0x04) |  f)
}

func (h *FLVHeader)GetAudioFlag()(flag bool){
	return h.Data[4] & 0x04
}

func (h *FLVHeader)GetVideoFlag()(flag bool){
	return h.Data[4] & 0x01
}

type FLVTag struct{
	tagType uint8
	dataSize uint32
	timeStamp uint32
	streamId uint32
	data bytes.Buffer
}


