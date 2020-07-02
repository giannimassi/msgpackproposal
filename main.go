package main

import (
	"github.com/tinylib/msgp/msgp"
)

//go:generate msgp

//msgp:tuple StructV1

// StructV1 is the older implementation
type StructV1 struct {
	A uint32
	B uint32
}

// DecodeMsg1 implements a method for decoding a StructV1 from a msgpack array
// of arbitrary length
func (z *StructV1) DecodeMsg1(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 == 0 {
		return
	}
	z.A, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "A")
		return
	}
	if zb0001--; zb0001 == 0 {
		return
	}
	z.B, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "B")
		return
	}
	for ; zb0001 > 1; zb0001-- {
		if err = dc.Skip(); err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	return
}

// UnmarshalMsg1 implements a method for unmarshalling a byte slice into a
// StructV1 from a msgpack array of arbitrary length
func (z *StructV1) UnmarshalMsg1(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 == 0 {
		return
	}
	z.A, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "A")
		return
	}
	if zb0001--; zb0001 == 0 {
		return
	}
	z.B, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "B")
		return
	}
	for ; zb0001 > 1; zb0001-- {
		if bts, err = msgp.Skip(bts); err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	o = bts
	return
}

//msgp:tuple StructV2

// StructV2 is the newer implementation
type StructV2 struct {
	A uint32
	B uint32
	C uint32
}

// DecodeMsg1 implements a method for decoding a StructV2 from a msgpack array
// of arbitrary length
func (z *StructV2) DecodeMsg1(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001--; zb0001 == 0 {
		return
	}
	z.A, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "A")
		return
	}
	if zb0001--; zb0001 == 0 {
		return
	}
	z.B, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "B")
		return
	}
	if zb0001--; zb0001 == 0 {
		return
	}
	z.C, err = dc.ReadUint32()
	if err != nil {
		err = msgp.WrapError(err, "C")
		return
	}
	for ; zb0001 > 0; zb0001-- {
		if err = dc.Skip(); err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	return
}

// UnmarshalMsg1 implements a method for unmarshalling a byte slice into a
// StructV2 from a msgpack array of arbitrary length
func (z *StructV2) UnmarshalMsg1(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001--; zb0001 == 0 {
		return
	}
	z.A, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "A")
		return
	}
	if zb0001--; zb0001 == 0 {
		return
	}
	z.B, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "B")
		return
	}
	if zb0001--; zb0001 == 0 {
		return
	}
	z.C, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "C")
		return
	}
	for ; zb0001 > 0; zb0001-- {
		if bts, err = msgp.Skip(bts); err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
	}
	o = bts
	return
}
