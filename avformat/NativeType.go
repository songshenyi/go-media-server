package avformat

import (
	"encoding/binary"
	"io"
)

type NativeUint8 uint8

func (v *NativeUint8)MarshalBinary() (data []byte, err error) {
	return []byte{byte(*v)}, nil
}

func (v *NativeUint8) Size() int {
	return 1
}

func (v *NativeUint8) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = NativeUint8(data[0])
	return
}


type NativeUint16 uint16

func (v *NativeUint16) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Size())
	binary.BigEndian.PutUint16(data, uint16(*v))
	return data, nil
}

func (v *NativeUint16) Size() int {
	return 2
}

func (v *NativeUint16) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = NativeUint16(binary.BigEndian.Uint16(data))
	return
}

type NativeUint24 uint32

func (v *NativeUint24) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Size())
	data[0] = byte(*v >> 16)
	data[1] = byte(*v >> 8)
	data[2] = byte(*v)
	return data, nil
}

func (v *NativeUint24) Size() int {
	return 3
}

func (v *NativeUint24) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = NativeUint24(uint32(data[0])<<16 | uint32(data[1])<<8 | uint32(data[2]))
	return
}


type NativeUint32 uint32

func (v *NativeUint32) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Size())
	binary.BigEndian.PutUint32(data, uint32(*v) )
	return data, nil
}

func (v *NativeUint32) Size() int {
	return 4
}

func (v *NativeUint32) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = NativeUint32(binary.BigEndian.Uint32(data))
	return
}

type NativeUint64 uint64

func (v *NativeUint64) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Size())
	binary.BigEndian.PutUint64(data, uint64(*v) )
	return data, nil
}

func (v *NativeUint64) Size() int {
	return 8
}

func (v *NativeUint64) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = NativeUint64(binary.BigEndian.Uint64(data))
	return
}