package avformat

import (
	"encoding/binary"
	"io"
)

type GMSUint8 uint8

func (v *GMSUint8)MarshalBinary() (data []byte, err error) {
	return []byte{byte(*v)}, nil
}

func (v *GMSUint8) Size() int {
	return 1
}

func (v *GMSUint8) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = GMSUint8(data[0])
	return
}


type GMSUint16 uint16

func (v *GMSUint16) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Size())
	binary.BigEndian.PutUint16(data, uint16(*v))
	return data, nil
}

func (v *GMSUint16) Size() int {
	return 2
}

func (v *GMSUint16) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = GMSUint16(binary.BigEndian.Uint16(data))
	return
}

type GMSUint24 uint32

func (v *GMSUint24) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Size())
	data[0] = byte(*v >> 16)
	data[1] = byte(*v >> 8)
	data[2] = byte(*v)
	return data, nil
}

func (v *GMSUint24) Size() int {
	return 3
}

func (v *GMSUint24) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = GMSUint24(uint32(data[0])<<16 | uint32(data[1])<<8 | uint32(data[2]))
	return
}


type GMSUint32 uint32

func (v *GMSUint32) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Size())
	binary.BigEndian.PutUint32(data, uint32(*v) )
	return data, nil
}

func (v *GMSUint32) Size() int {
	return 4
}

func (v *GMSUint32) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = GMSUint32(binary.BigEndian.Uint32(data))
	return
}

type GMSUint64 uint64

func (v *GMSUint64) MarshalBinary() (data []byte, err error) {
	data = make([]byte, v.Size())
	binary.BigEndian.PutUint64(data, uint64(*v) )
	return data, nil
}

func (v *GMSUint64) Size() int {
	return 8
}

func (v *GMSUint64) UnmarshalBinary(data []byte) (err error) {
	if len(data) < v.Size() {
		return io.EOF
	}
	*v = GMSUint64(binary.BigEndian.Uint64(data))
	return
}