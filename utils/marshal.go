package utils

import (
	"bytes"
	"encoding"
	"reflect"
)


// unmarshaler
type Marshaler interface {
	encoding.BinaryMarshaler
}

// marshal the object o to b
func Marshal(o Marshaler, b *bytes.Buffer) (err error) {
	if b == nil {
		panic("should not be nil.")
	}

	if o == nil {
		panic("should not be nil.")
	}

	if vb, err := o.MarshalBinary(); err != nil {
		return err
	} else if _, err := b.Write(vb); err != nil {
		return err
	}

	return
}

// marshal multiple o, which can be nil.
func Marshals(o ...Marshaler) (data []byte, err error) {
	var b bytes.Buffer

	for _, e := range o {
		if e == nil {
			continue
		}

		if rv := reflect.ValueOf(e); rv.IsNil() {
			continue
		}

		if err = Marshal(e, &b); err != nil {
			return
		}
	}

	return b.Bytes(), nil
}

// unmarshaler and sizer.
type UnmarshalSizer interface {
	encoding.BinaryUnmarshaler

	// the total size of bytes for this amf0 instance.
	Size() int
}

// unmarshal the object from b
func Unmarshal(o UnmarshalSizer, b *bytes.Buffer) (err error) {
	if b == nil {
		panic("should not be nil")
	}

	if o == nil {
		panic("should not be nil")
	}

	if err = o.UnmarshalBinary(b.Bytes()); err != nil {
		return
	}
	b.Next(o.Size())

	return err
}

// unmarshal multiple o pointers, which can be nil.
func Unmarshals(b *bytes.Buffer, o ...UnmarshalSizer) (err error) {
	for _, e := range o {
		if b.Len() == 0 {
			break
		}

		if e == nil {
			continue
		}

		if rv := reflect.ValueOf(e); rv.IsNil() {
			continue
		}

		if err = e.UnmarshalBinary(b.Bytes()); err != nil {
			return err
		}
		b.Next(e.Size())
	}

	return err
}
