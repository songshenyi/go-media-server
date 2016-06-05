package avformat

import "bytes"

type FLVHeader struct {
	Signature [3]byte
	Version uint8
	EnableAudio bool
	EnableVideo bool
	Offset uint32
}

//func (h *FLVHeader)SetAudioFlag(flag bool){
//	f := 0x04 * flag
//	h.Data[4] = (( h.Data[4] & 0x01) |  f)
//}
//
//func (h *FLVHeader)SetVideoFlag(flag bool){
//	f := 0x01 * flag
//	h.Data[4] = (( h.Data[4] & 0x04) |  f)
//}
//
//func (h *FLVHeader)GetAudioFlag()(flag bool){
//	return h.Data[4] & 0x04
//}
//
//func (h *FLVHeader)GetVideoFlag()(flag bool){
//	return h.Data[4] & 0x01
//}

type RTMPMessageType uint8

const (
	RTMPMsgAMF0DataMessage RTMPMessageType = 18 // 0x12
	RTMPMsgAMF3DataMessage RTMPMessageType = 15 // 0x0F

	RTMPMsgAudioMessage RTMPMessageType = 8 // 0x08
	RTMPMsgVideoMessage RTMPMessageType = 9 // 0x09
)

type FLVTagTimestamp uint64

type FLVTagUint24 uint32

type FLVMessage struct {
	tag *FLVTag

	MetaData bool
	VideoSequenceHeader bool
	AudioSequenceHeader bool
}

type FLVTag struct{
	TagType   RTMPMessageType
	DataSize  FLVTagUint24
	TimeStamp FLVTagTimestamp
	StreamId  FLVTagUint24
	Payload      []byte
}


// SoundFormat UB [4]
// Format of SoundData. The following values are defined:
//     0 = Linear PCM, platform endian
//     1 = ADPCM
//     2 = MP3
//     3 = Linear PCM, little endian
//     4 = Nellymoser 16 kHz mono
//     5 = Nellymoser 8 kHz mono
//     6 = Nellymoser
//     7 = G.711 A-law logarithmic PCM
//     8 = G.711 mu-law logarithmic PCM
//     9 = reserved
//     10 = AAC
//     11 = Speex
//     14 = MP3 8 kHz
//     15 = Device-specific sound
// Formats 7, 8, 14, and 15 are reserved.
// AAC is supported in Flash Player 9,0,115,0 and higher.
// Speex is supported in Flash Player 10 and higher.
type RTMPCodecAudio uint8

const (
	RTMPLinearPCMPlatformEndian RTMPCodecAudio = iota
	RTMPADPCM
	RTMPMP3
	RTMPLinearPCMLittleEndian
	RTMPNellymoser16kHzMono
	RTMPNellymoser8kHzMono
	RTMPNellymoser
	RTMPReservedG711AlawLogarithmicPCM
	RTMPReservedG711MuLawLogarithmicPCM
	RTMPReserved
	RTMPAAC
	RTMPSpeex
	RTMPReserved1CodecAudio
	RTMPReserved2CodecAudio
	RTMPReservedMP3_8kHz
	RTMPReservedDeviceSpecificSound
	RTMPReserved3CodecAudio
	RTMPDisabledCodecAudio
)

// AACPacketType IF SoundFormat == 10 UI8
// The following values are defined:
//     0 = AAC sequence header
//     1 = AAC raw
type RTMPAacType uint8

const (
	RTMPAacSequenceHeader RTMPAacType = iota
	RTMPAacRawData
	RTMPAacReserved
)

// E.4.3.1 VIDEODATA
// CodecID UB [4]
// Codec Identifier. The following values are defined:
//     2 = Sorenson H.263
//     3 = Screen video
//     4 = On2 VP6
//     5 = On2 VP6 with alpha channel
//     6 = Screen video version 2
//     7 = AVC
type RTMPCodecVideo uint8

const (
	RTMPReservedCodecVideo RTMPCodecVideo = iota
	RTMPReserved1CodecVideo
	RTMPSorensonH263
	RTMPScreenVideo
	RTMPOn2VP6
	RTMPOn2VP6WithAlphaChannel
	RTMPScreenVideoVersion2
	RTMPAVC
	RTMPDisabledCodecVideo
	RTMPReserved2CodecVideo
)

// E.4.3.1 VIDEODATA
// Frame Type UB [4]
// Type of video frame. The following values are defined:
//     1 = key frame (for AVC, a seekable frame)
//     2 = inter frame (for AVC, a non-seekable frame)
//     3 = disposable inter frame (H.263 only)
//     4 = generated key frame (reserved for server use only)
//     5 = video info/command frame
type RtmpAVCFrame uint8

const (
	RTMPReservedAVCFrame RtmpAVCFrame = iota
	RTMPKeyFrame
	RTMPInterFrame
	RTMPDisposableInterFrame
	RTMPGeneratedKeyFrame
	RTMPVideoInfoFrame
	RTMPReserved1AVCFrame
)

// AVCPacketType IF CodecID == 7 UI8
// The following values are defined:
//     0 = AVC sequence header
//     1 = AVC NALU
//     2 = AVC end of sequence (lower level NALU sequence ender is
//         not required or supported)
type RtmpVideoAVCType uint8

const (
	RTMPSequenceHeader RtmpVideoAVCType = iota
	RTMPNALU
	RTMPSequenceHeaderEOF
	RTMPReservedAVCType
)
type RTMPCodecAudio uint8


func (v *FLVMessage) isVideoSequenceHeader() bool {
	// TODO: FIXME: support other codecs.
	if len(v.tag.Payload) < 2 {
		return false
	}

	b := v.tag.Payload

	// sequence header only for h264
	codec := RTMPCodecVideo(b[0] & 0x0f)
	if codec != RTMPAVC {
		return false
	}

	frameType := RtmpAVCFrame((b[0] >> 4) & 0x0f)
	avcPacketType := RtmpVideoAVCType(b[1])
	return frameType == RTMPKeyFrame && avcPacketType == RTMPSequenceHeader
}

func (v *FLVMessage) isAudioSequenceHeader() bool {
	// TODO: FIXME: support other codecs.
	if len(v.tag.Payload) < 2 {
		return false
	}

	b := v.tag.Payload

	soundFormat := RTMPCodecAudio((b[0] >> 4) & 0x0f)
	if soundFormat != RTMPAAC {
		return false
	}

	aacPacketType := RTMPAacType(b[1])
	return aacPacketType == RTMPAacSequenceHeader
}

func (v RTMPMessageType) isAudio() bool {
	return v == RTMPMsgAudioMessage
}

func (v RTMPMessageType) isVideo() bool {
	return v == RTMPMsgVideoMessage
}

func (v RTMPMessageType) isData() bool {
	return v.isAmf0Data() || v.isAmf3Data()
}

func (v RTMPMessageType) isAmf0Data() bool {
	return v == RTMPMsgAMF0DataMessage
}

func (v RTMPMessageType) isAmf3Data() bool {
	return v == RTMPMsgAMF3DataMessage
}
