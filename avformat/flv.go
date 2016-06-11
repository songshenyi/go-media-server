package avformat

import (
	"io"
	"github.com/songshenyi/go-media-server/logger"
	"encoding"
	"bytes"
	"github.com/songshenyi/go-media-server/utils"
)


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

type Marshaler interface {
	encoding.BinaryUnmarshaler
	encoding.BinaryMarshaler
	Size() int
}

type RtmpMessageType GMSUint8
func (v *RtmpMessageType)MarshalBinary() (data []byte, err error) {
	return []byte{byte(*v)}, nil
}

func (v *RtmpMessageType) Size() int {
	return 1
}

func (v *RtmpMessageType) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = RtmpMessageType(data[0])
	return
}


const (
	RtmpMsgAmf0DataMessage 	RtmpMessageType = 18 // 0x12
	RtmpMsgAmf3DataMessage 	RtmpMessageType = 15 // 0x0F

	RtmpMsgAudioMessage 	RtmpMessageType = 8 // 0x08
	RtmpMsgVideoMessage 	RtmpMessageType = 9 // 0x09
)

type FlvMessage struct {
	Tag *FlvTag
	Header *FlvHeader

	MetaData bool
	VideoSequenceHeader bool
	AudioSequenceHeader bool
}

func NewFlvMessage()(m *FlvMessage, e error){
	return &FlvMessage{}, nil
}

func (v* FlvMessage)Copy() (m* FlvMessage){
	m, _ = NewFlvMessage()
	if(v.Header != nil){
		header := *v.Header
		m.Header = &header
	}

	if(v.Tag != nil){
		tag := *v.Tag
		m.Tag = &tag
	}

	return m
}

const (
	FlvHeaderSize int = 13
	FlvTagHeaderSize int = 11
	FlvPreTagLenSize int = 4
)

type FlvHeaderSignature [3]byte

func (v *FlvHeaderSignature) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Size())
	copy(data, v[:])
	return data, nil
}

func (v *FlvHeaderSignature) Size() int {
	return 3
}

func (v *FlvHeaderSignature) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	copy(v[:], data)
	return
}

type FlvTagTimestamp uint64

func (v *FlvTagTimestamp) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Size())
	data[3] = byte(*v >> 24)
	data[2] = byte(*v)
	data[1] = byte(*v >> 8)
	data[0] = byte(*v >> 16)
	return data, nil
}

func (v *FlvTagTimestamp) Size() int {
	return 4
}

func (v *FlvTagTimestamp) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = FlvTagTimestamp(uint32(data[3]) << 24 | uint32(data[0]) << 16 | uint32(data[1]) << 8 | uint32(data[2]))
	return
}

type FlvHeader struct {
	Signature FlvHeaderSignature
	Version GMSUint8
	EnableAudio bool
	EnableVideo bool
	Offset GMSUint32
	PreTagSize0 GMSUint32
}

func (h *FlvHeader)ToMessage()(m *FlvMessage, err error){
	m, err = NewFlvMessage()
	m.Header = h

	return m, err
}

func ReadFlvHeader(r io.Reader)(h *FlvHeader, err error){
	h = &FlvHeader{}
	var buf bytes.Buffer
	if _, err := io.CopyN(&buf, r, int64(FlvHeaderSize)) ; err != nil {
		logger.Warn("read flv headerfailed")
		return h, err
	}

	data := buf.Bytes()
	var AVFlag GMSUint8
	if err = utils.Unmarshals(bytes.NewBuffer(data), &h.Signature, &h.Version, &AVFlag, &h.Offset, &h.PreTagSize0); err != nil{
		logger.Warn("Unmarshals flv header failed")
		return h, err
	}

	//copy(h.Signature[:], data)
	if string(h.Signature[:]) != "FLV" {
		logger.Warnf("flv header Signature is wrong, %s", string(h.Signature[:]))
		return h, err
	}

	if h.Version != 1 {
		logger.Warn("flv header Version invalid")
		return h, err
	}

	//AVFlag := data[4]
	h.EnableAudio = (AVFlag & 0x04) != 0
	h.EnableVideo = (AVFlag & 0x01) != 0

	logger.Info("read flv header")
	return
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (h *FlvHeader)ToData()(data []byte, err error){
	var AVFlag GMSUint8
	audioFlag := GMSUint8(0x04 * Btoi(h.EnableAudio))
	videoFlag := GMSUint8(0x01 * Btoi(h.EnableVideo))
	AVFlag = AVFlag | audioFlag | videoFlag

	return utils.Marshals(&h.Signature, &h.Version, &AVFlag, &h.Offset, &h.PreTagSize0)
}

type FlvTag struct{
	TagType   RtmpMessageType
	DataSize  GMSUint24
	TimeStamp FlvTagTimestamp
	StreamId  GMSUint24
	Payload   []byte
}

func (tag *FlvTag)ToMessage()(m *FlvMessage, err error){
	m, err = NewFlvMessage()
	m.Tag = tag

	switch tag.TagType {
	case RtmpMsgAmf0DataMessage:
		m.MetaData = true
	case RtmpMsgAudioMessage:
		m.AudioSequenceHeader = tag.isAudioSequenceHeader()
	case RtmpMsgVideoMessage:
		m.VideoSequenceHeader = tag.isVideoSequenceHeader()
	}

	return m, err
}

func (tag *FlvTag)isVideoSequenceHeader() bool {
	// TODO: FIXME: support other codecs.
	if len(tag.Payload) < 2 {
		return false
	}

	b := tag.Payload

	// sequence header only for h264
	codec := RtmpCodecVideo(b[0] & 0x0f)
	if codec != RtmpAVC {
		return false
	}

	frameType := RtmpAVCFrame((b[0] >> 4) & 0x0f)
	avcPacketType := RtmpVideoAVCType(b[1])
	return frameType == RtmpKeyFrame && avcPacketType == RtmpSequenceHeader
}

func (tag *FlvTag) isAudioSequenceHeader() bool {
	// TODO: FIXME: support other codecs.
	if len(tag.Payload) < 2 {
		return false
	}

	b := tag.Payload

	soundFormat := RtmpCodecAudio((b[0] >> 4) & 0x0f)
	if soundFormat != RtmpAAC {
		return false
	}

	aacPacketType := RtmpAacType(b[1])
	return aacPacketType == RtmpAacSequenceHeader
}


func FlvGetDataSize(data []byte)(size uint32, err error){
	size = uint32(data[0])<<16 | uint32(data[1])<<8 | uint32(data[2])
	return
}



func FlvGetTimestamp(data []byte)(ts uint32, err error){
	ts = uint32(data[3] << 24) | uint32(data[0] << 16) | uint32(data[1] << 8) | uint32(data[2]);
	return
}

func FlvGetStreamId(data []byte)(streamid uint32, err error){
	streamid = uint32(data[0])<<16 | uint32(data[1])<<8 | uint32(data[2])
	return
}

func FlvGetPreTagSize(data []byte)(size uint32, err error){
	size = uint32(data[0]) << 24 | uint32(data[1]) << 16 | uint32(data[2]) << 8 | uint32(data[3]);
	return
}


func ReadFlvTag(r io.Reader)(tag *FlvTag, err error){
	tag = &FlvTag{}
	var buf bytes.Buffer
	if _, err := io.CopyN(&buf, r, int64(FlvTagHeaderSize)) ; err != nil {
		logger.Warn("read flv headerfailed")
		return tag, err
	}

	data := buf.Bytes()

	if err = utils.Unmarshals(bytes.NewBuffer(data), &tag.TagType, &tag.DataSize, &tag.TimeStamp, &tag.StreamId); err != nil{
		logger.Warn("Unmarshals flv header failed")
		return tag, err
	}

	buf.Reset()

	written, err := io.CopyN(&buf, r, int64(tag.DataSize))
	if (written != int64(tag.DataSize) || err != nil) {
		logger.Warn("read flv tag Data failed")
		return tag, err
	}

	io.CopyN(&buf, r, 4)
	tag.Payload = make([]byte, tag.DataSize)
	copy(tag.Payload, buf.Bytes())

	logger.Trace("read flv tag")
	return
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
type RtmpCodecAudio uint8

const (
	RtmpLinearPCMPlatformEndian RtmpCodecAudio = iota
	RtmpADPCM
	RtmpMP3
	RtmpLinearPCMLittleEndian
	RtmpNellymoser16kHzMono
	RtmpNellymoser8kHzMono
	RtmpNellymoser
	RtmpReservedG711AlawLogarithmicPCM
	RtmpReservedG711MuLawLogarithmicPCM
	RtmpReserved
	RtmpAAC
	RtmpSpeex
	RtmpReserved1CodecAudio
	RtmpReserved2CodecAudio
	RtmpReservedMP3_8kHz
	RtmpReservedDeviceSpecificSound
	RtmpReserved3CodecAudio
	RtmpDisabledCodecAudio
)

// AACPacketType IF SoundFormat == 10 UI8
// The following values are defined:
//     0 = AAC sequence header
//     1 = AAC raw
type RtmpAacType uint8

const (
	RtmpAacSequenceHeader RtmpAacType = iota
	RtmpAacRawData
	RtmpAacReserved
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
type RtmpCodecVideo uint8

const (
	RtmpReservedCodecVideo RtmpCodecVideo = iota
	RtmpReserved1CodecVideo
	RtmpSorensonH263
	RtmpScreenVideo
	RtmpOn2VP6
	RtmpOn2VP6WithAlphaChannel
	RtmpScreenVideoVersion2
	RtmpAVC
	RtmpDisabledCodecVideo
	RtmpReserved2CodecVideo
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
	RtmpReservedAVCFrame RtmpAVCFrame = iota
	RtmpKeyFrame
	RtmpInterFrame
	RtmpDisposableInterFrame
	RtmpGeneratedKeyFrame
	RtmpVideoInfoFrame
	RtmpReserved1AVCFrame
)

// AVCPacketType IF CodecID == 7 UI8
// The following values are defined:
//     0 = AVC sequence header
//     1 = AVC NALU
//     2 = AVC end of sequence (lower level NALU sequence ender is
//         not required or supported)
type RtmpVideoAVCType uint8

const (
	RtmpSequenceHeader RtmpVideoAVCType = iota
	RtmpNALU
	RtmpSequenceHeaderEOF
	RtmpReservedAVCType
)

func (v RtmpMessageType) isAudio() bool {
	return v == RtmpMsgAudioMessage
}

func (v RtmpMessageType) isVideo() bool {
	return v == RtmpMsgVideoMessage
}

func (v RtmpMessageType) isData() bool {
	return v.isAmf0Data() || v.isAmf3Data()
}

func (v RtmpMessageType) isAmf0Data() bool {
	return v == RtmpMsgAmf0DataMessage
}

func (v RtmpMessageType) isAmf3Data() bool {
	return v == RtmpMsgAmf3DataMessage
}
