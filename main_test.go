package main

import (
	"bytes"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestMarshalStructV2UnmarshalStructV1(t *testing.T) {
	v2 := StructV2{}
	bts, err := v2.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}

	v1 := StructV1{}
	left, err := v1.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestMarshalStructV2UnmarshalStructV1Variable(t *testing.T) {
	v2 := StructV2{}
	bts, err := v2.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}

	v1 := StructV1{}
	left, err := v1.UnmarshalMsg1(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestMarshalStructV1UnmarshalStructV2(t *testing.T) {
	v1 := StructV1{}
	bts, err := v1.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}

	v2 := StructV2{}
	left, err := v2.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestMarshalStructV1UnmarshalStructV2Variable(t *testing.T) {
	v1 := StructV1{}
	bts, err := v1.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}

	v2 := StructV2{}
	left, err := v2.UnmarshalMsg1(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestEncodeStructV2DecodeStructV1(t *testing.T) {
	v := StructV2{}
	var buf bytes.Buffer
	msgp.Encode(&buf, &v)

	m := v.Msgsize()
	if buf.Len() > m {
		t.Logf("WARNING: Msgsize() for %v is inaccurate", v)
	}

	vn := StructV1{}
	err := vn.DecodeMsg(msgp.NewReader(&buf))
	if err != nil {
		t.Error(err)
	}

	buf.Reset()
	msgp.Encode(&buf, &v)
	err = msgp.NewReader(&buf).Skip()
	if err != nil {
		t.Error(err)
	}
}

func TestEncodeStructV1DecodeStructV2(t *testing.T) {
	v := StructV1{}
	var buf bytes.Buffer
	msgp.Encode(&buf, &v)

	m := v.Msgsize()
	if buf.Len() > m {
		t.Logf("WARNING: Msgsize() for %v is inaccurate", v)
	}

	vn := StructV2{}
	err := vn.DecodeMsg(msgp.NewReader(&buf))
	if err != nil {
		t.Error(err)
	}

	buf.Reset()
	msgp.Encode(&buf, &v)
	err = msgp.NewReader(&buf).Skip()
	if err != nil {
		t.Error(err)
	}
}

func TestEncodeStructV2DecodeStructV1Variable(t *testing.T) {
	v := StructV2{}
	var buf bytes.Buffer
	msgp.Encode(&buf, &v)

	m := v.Msgsize()
	if buf.Len() > m {
		t.Logf("WARNING: Msgsize() for %v is inaccurate", v)
	}

	vn := StructV1{}
	err := vn.DecodeMsg1(msgp.NewReader(&buf))
	if err != nil {
		t.Error(err)
	}

	buf.Reset()
	msgp.Encode(&buf, &v)
	err = msgp.NewReader(&buf).Skip()
	if err != nil {
		t.Error(err)
	}
}

func TestEncodeStructV1DecodeStructV2Variable(t *testing.T) {
	v := StructV1{}
	var buf bytes.Buffer
	msgp.Encode(&buf, &v)

	m := v.Msgsize()
	if buf.Len() > m {
		t.Logf("WARNING: Msgsize() for %v is inaccurate", v)
	}

	vn := StructV2{}
	err := vn.DecodeMsg1(msgp.NewReader(&buf))
	if err != nil {
		t.Error(err)
	}

	buf.Reset()
	msgp.Encode(&buf, &v)
	err = msgp.NewReader(&buf).Skip()
	if err != nil {
		t.Error(err)
	}
}
