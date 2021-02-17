// Code generated by bebopc-go; DO NOT EDIT.

package generated

import (
	"bytes"
	"io"

	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
)

var _ bebop.Record = &D{}

type D struct {
	S string
}

func (bbp D) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp D) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.S)))
	at += 4
	copy(buf[at:at+len(bbp.S)], []byte(bbp.S))
	at += len(bbp.S)
}

func (bbp *D) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	bbp.S, err = iohelp.ReadStringBytes(buf[at:])
	if err != nil {
		 return err
	}
	at += 4 + len(bbp.S)
	return nil
}

func (bbp *D) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.S = iohelp.MustReadStringBytes(buf[at:])
	at += 4+len(bbp.S)
}

func (bbp D) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(len(bbp.S)))
	w.Write([]byte(bbp.S))
	return w.Err
}

func (bbp *D) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.S = iohelp.ReadString(r)
	return r.Err
}

func (bbp D) bodyLen() int {
	bodyLen := 0
	bodyLen += 4
	bodyLen += len(bbp.S)
	return bodyLen
}

func makeD(r iohelp.ErrorReader) (D, error) {
	v := D{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeDFromBytes(buf []byte) (D, error) {
	v := D{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeDFromBytes(buf []byte) D {
	v := D{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &X{}

type X struct {
	X bool
}

func (bbp X) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp X) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteBoolBytes(buf[at:], bbp.X)
	at += 1
}

func (bbp *X) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 1 {
		 return iohelp.ErrTooShort
	}
	bbp.X = iohelp.ReadBoolBytes(buf[at:])
	at += 1
	return nil
}

func (bbp *X) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.X = iohelp.ReadBoolBytes(buf[at:])
	at += 1
}

func (bbp X) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteBool(w, bbp.X)
	return w.Err
}

func (bbp *X) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.X = iohelp.ReadBool(r)
	return r.Err
}

func (bbp X) bodyLen() int {
	bodyLen := 0
	bodyLen += 1
	return bodyLen
}

func makeX(r iohelp.ErrorReader) (X, error) {
	v := X{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeXFromBytes(buf []byte) (X, error) {
	v := X{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeXFromBytes(buf []byte) X {
	v := X{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &W{}

type W struct {
	D *D
	X *X
}

func (bbp W) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp W) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.bodyLen()-4))
	at += 4
	if bbp.D != nil {
		buf[at] = 1
		at++
		(*bbp.D).MarshalBebopTo(buf[at:])
		at += (*bbp.D).bodyLen()
	}
	if bbp.X != nil {
		buf[at] = 2
		at++
		(*bbp.X).MarshalBebopTo(buf[at:])
		at += (*bbp.X).bodyLen()
	}
}

func (bbp *W) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.D = new(D)
			(*bbp.D), err = makeDFromBytes(buf[at:])
			if err != nil {
				 return err
			}
			at += ((*bbp.D)).bodyLen()
		case 2:
			at += 1
			bbp.X = new(X)
			(*bbp.X), err = makeXFromBytes(buf[at:])
			if err != nil {
				 return err
			}
			at += ((*bbp.X)).bodyLen()
		default:
			return nil
		}
	}
}

func (bbp *W) MustUnmarshalBebop(buf []byte) {
	at := 0
	for {
		switch buf[at] {
		case 1:
			bbp.D = new(D)
			(*bbp.D) = mustMakeDFromBytes(buf[at:])
			at += ((*bbp.D)).bodyLen()
		case 2:
			bbp.X = new(X)
			(*bbp.X) = mustMakeXFromBytes(buf[at:])
			at += ((*bbp.X)).bodyLen()
		default:
			return
		}
	}
}

func (bbp W) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.bodyLen()-4))
	if bbp.D != nil {
		w.Write([]byte{1})
		err = (*bbp.D).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	if bbp.X != nil {
		w.Write([]byte{2})
		err = (*bbp.X).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *W) DecodeBebop(ior io.Reader) (err error) {
	er := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(er)
	body := make([]byte, bodyLen)
	er.Read(body)
	r := iohelp.NewErrorReader(bytes.NewReader(body))
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.D = new(D)
			(*bbp.D), err = makeD(r)
			if err != nil {
				return err
			}
		case 2:
			bbp.X = new(X)
			(*bbp.X), err = makeX(r)
			if err != nil {
				return err
			}
		default:
			return er.Err
		}
	}
}

func (bbp W) bodyLen() int {
	bodyLen := 5
	if bbp.D != nil {
		bodyLen += 1
		bodyLen += (*bbp.D).bodyLen()
	}
	if bbp.X != nil {
		bodyLen += 1
		bodyLen += (*bbp.X).bodyLen()
	}
	return bodyLen
}

func makeW(r iohelp.ErrorReader) (W, error) {
	v := W{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeWFromBytes(buf []byte) (W, error) {
	v := W{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeWFromBytes(buf []byte) W {
	v := W{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &A{}

type A struct {
	A *uint32
}

func (bbp A) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp A) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.bodyLen()-4))
	at += 4
	if bbp.A != nil {
		buf[at] = 1
		at++
		iohelp.WriteUint32Bytes(buf[at:], *bbp.A)
		at += 4
	}
}

func (bbp *A) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.A = new(uint32)
			if len(buf[at:]) < 4 {
				 return iohelp.ErrTooShort
			}
			(*bbp.A) = iohelp.ReadUint32Bytes(buf[at:])
			at += 4
		default:
			return nil
		}
	}
}

func (bbp *A) MustUnmarshalBebop(buf []byte) {
	at := 0
	for {
		switch buf[at] {
		case 1:
			bbp.A = new(uint32)
			(*bbp.A) = iohelp.ReadUint32Bytes(buf[at:])
			at += 4
		default:
			return
		}
	}
}

func (bbp A) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.bodyLen()-4))
	if bbp.A != nil {
		w.Write([]byte{1})
		iohelp.WriteUint32(w, *bbp.A)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *A) DecodeBebop(ior io.Reader) (err error) {
	er := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(er)
	body := make([]byte, bodyLen)
	er.Read(body)
	r := iohelp.NewErrorReader(bytes.NewReader(body))
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.A = new(uint32)
			*bbp.A = iohelp.ReadUint32(r)
		default:
			return er.Err
		}
	}
}

func (bbp A) bodyLen() int {
	bodyLen := 5
	if bbp.A != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

func makeA(r iohelp.ErrorReader) (A, error) {
	v := A{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeAFromBytes(buf []byte) (A, error) {
	v := A{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeAFromBytes(buf []byte) A {
	v := A{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &B{}

//*
//     * This branch is, too!
//     
type B struct {
	B bool
}

func (bbp B) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp B) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteBoolBytes(buf[at:], bbp.B)
	at += 1
}

func (bbp *B) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 1 {
		 return iohelp.ErrTooShort
	}
	bbp.B = iohelp.ReadBoolBytes(buf[at:])
	at += 1
	return nil
}

func (bbp *B) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.B = iohelp.ReadBoolBytes(buf[at:])
	at += 1
}

func (bbp B) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteBool(w, bbp.B)
	return w.Err
}

func (bbp *B) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.B = iohelp.ReadBool(r)
	return r.Err
}

func (bbp B) bodyLen() int {
	bodyLen := 0
	bodyLen += 1
	return bodyLen
}

func makeB(r iohelp.ErrorReader) (B, error) {
	v := B{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeBFromBytes(buf []byte) (B, error) {
	v := B{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeBFromBytes(buf []byte) B {
	v := B{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &C{}

type C struct {
}

func (bbp C) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp C) MarshalBebopTo(buf []byte) {
}

func (bbp *C) UnmarshalBebop(buf []byte) (err error) {
	return nil
}

func (bbp *C) MustUnmarshalBebop(buf []byte) {
}

func (bbp C) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	return w.Err
}

func (bbp *C) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	return r.Err
}

func (bbp C) bodyLen() int {
	bodyLen := 0
	return bodyLen
}

func makeC(r iohelp.ErrorReader) (C, error) {
	v := C{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeCFromBytes(buf []byte) (C, error) {
	v := C{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeCFromBytes(buf []byte) C {
	v := C{}
	v.MustUnmarshalBebop(buf)
	return v
}

const UOpCode = 0x79656168

var _ bebop.Record = &U{}

type U struct {
	A *A
	B *B
	C *C
	W *W
}

func (bbp U) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp U) MarshalBebopTo(buf []byte) {
	iohelp.WriteUint32Bytes(buf, uint32(UOpCode))
	at := 4
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.bodyLen()-8))
	at += 4
	if bbp.A != nil {
		buf[at] = 1
		at++
		(*bbp.A).MarshalBebopTo(buf[at:])
		at += (*bbp.A).bodyLen()
	}
	if bbp.B != nil {
		buf[at] = 2
		at++
		(*bbp.B).MarshalBebopTo(buf[at:])
		at += (*bbp.B).bodyLen()
	}
	if bbp.C != nil {
		buf[at] = 3
		at++
		(*bbp.C).MarshalBebopTo(buf[at:])
		at += (*bbp.C).bodyLen()
	}
	if bbp.W != nil {
		buf[at] = 4
		at++
		(*bbp.W).MarshalBebopTo(buf[at:])
		at += (*bbp.W).bodyLen()
	}
}

func (bbp *U) UnmarshalBebop(buf []byte) (err error) {
	at := 4
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.A = new(A)
			(*bbp.A), err = makeAFromBytes(buf[at:])
			if err != nil {
				 return err
			}
			at += ((*bbp.A)).bodyLen()
		case 2:
			at += 1
			bbp.B = new(B)
			(*bbp.B), err = makeBFromBytes(buf[at:])
			if err != nil {
				 return err
			}
			at += ((*bbp.B)).bodyLen()
		case 3:
			at += 1
			bbp.C = new(C)
			(*bbp.C), err = makeCFromBytes(buf[at:])
			if err != nil {
				 return err
			}
			at += ((*bbp.C)).bodyLen()
		case 4:
			at += 1
			bbp.W = new(W)
			(*bbp.W), err = makeWFromBytes(buf[at:])
			if err != nil {
				 return err
			}
			at += ((*bbp.W)).bodyLen()
		default:
			return nil
		}
	}
}

func (bbp *U) MustUnmarshalBebop(buf []byte) {
	at := 4
	for {
		switch buf[at] {
		case 1:
			bbp.A = new(A)
			(*bbp.A) = mustMakeAFromBytes(buf[at:])
			at += ((*bbp.A)).bodyLen()
		case 2:
			bbp.B = new(B)
			(*bbp.B) = mustMakeBFromBytes(buf[at:])
			at += ((*bbp.B)).bodyLen()
		case 3:
			bbp.C = new(C)
			(*bbp.C) = mustMakeCFromBytes(buf[at:])
			at += ((*bbp.C)).bodyLen()
		case 4:
			bbp.W = new(W)
			(*bbp.W) = mustMakeWFromBytes(buf[at:])
			at += ((*bbp.W)).bodyLen()
		default:
			return
		}
	}
}

func (bbp U) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(UOpCode))
	iohelp.WriteUint32(w, uint32(bbp.bodyLen()-8))
	if bbp.A != nil {
		w.Write([]byte{1})
		err = (*bbp.A).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	if bbp.B != nil {
		w.Write([]byte{2})
		err = (*bbp.B).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	if bbp.C != nil {
		w.Write([]byte{3})
		err = (*bbp.C).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	if bbp.W != nil {
		w.Write([]byte{4})
		err = (*bbp.W).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *U) DecodeBebop(ior io.Reader) (err error) {
	er := iohelp.NewErrorReader(ior)
	iohelp.ReadUint32(er)
	bodyLen := iohelp.ReadUint32(er)
	body := make([]byte, bodyLen)
	er.Read(body)
	r := iohelp.NewErrorReader(bytes.NewReader(body))
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.A = new(A)
			(*bbp.A), err = makeA(r)
			if err != nil {
				return err
			}
		case 2:
			bbp.B = new(B)
			(*bbp.B), err = makeB(r)
			if err != nil {
				return err
			}
		case 3:
			bbp.C = new(C)
			(*bbp.C), err = makeC(r)
			if err != nil {
				return err
			}
		case 4:
			bbp.W = new(W)
			(*bbp.W), err = makeW(r)
			if err != nil {
				return err
			}
		default:
			return er.Err
		}
	}
}

func (bbp U) bodyLen() int {
	bodyLen := 5
	bodyLen += 4
	if bbp.A != nil {
		bodyLen += 1
		bodyLen += (*bbp.A).bodyLen()
	}
	if bbp.B != nil {
		bodyLen += 1
		bodyLen += (*bbp.B).bodyLen()
	}
	if bbp.C != nil {
		bodyLen += 1
		bodyLen += (*bbp.C).bodyLen()
	}
	if bbp.W != nil {
		bodyLen += 1
		bodyLen += (*bbp.W).bodyLen()
	}
	return bodyLen
}

func makeU(r iohelp.ErrorReader) (U, error) {
	v := U{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeUFromBytes(buf []byte) (U, error) {
	v := U{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeUFromBytes(buf []byte) U {
	v := U{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &Cons{}

type Cons struct {
	Head uint32
	Tail List
}

func (bbp Cons) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp Cons) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], bbp.Head)
	at += 4
	(bbp.Tail).MarshalBebopTo(buf[at:])
	at += (bbp.Tail).bodyLen()
}

func (bbp *Cons) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 4 {
		 return iohelp.ErrTooShort
	}
	bbp.Head = iohelp.ReadUint32Bytes(buf[at:])
	at += 4
	bbp.Tail, err = makeListFromBytes(buf[at:])
	if err != nil {
		 return err
	}
	at += (bbp.Tail).bodyLen()
	return nil
}

func (bbp *Cons) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.Head = iohelp.ReadUint32Bytes(buf[at:])
	at += 4
	bbp.Tail = mustMakeListFromBytes(buf[at:])
	at += (bbp.Tail).bodyLen()
}

func (bbp Cons) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, bbp.Head)
	err = (bbp.Tail).EncodeBebop(w)
	if err != nil {
		return err
	}
	return w.Err
}

func (bbp *Cons) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.Head = iohelp.ReadUint32(r)
	(bbp.Tail), err = makeList(r)
	if err != nil {
		return err
	}
	return r.Err
}

func (bbp Cons) bodyLen() int {
	bodyLen := 0
	bodyLen += 4
	bodyLen += (bbp.Tail).bodyLen()
	return bodyLen
}

func makeCons(r iohelp.ErrorReader) (Cons, error) {
	v := Cons{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeConsFromBytes(buf []byte) (Cons, error) {
	v := Cons{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeConsFromBytes(buf []byte) Cons {
	v := Cons{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &Nil{}

type Nil struct {
}

func (bbp Nil) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp Nil) MarshalBebopTo(buf []byte) {
}

func (bbp *Nil) UnmarshalBebop(buf []byte) (err error) {
	return nil
}

func (bbp *Nil) MustUnmarshalBebop(buf []byte) {
}

func (bbp Nil) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	return w.Err
}

func (bbp *Nil) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	return r.Err
}

func (bbp Nil) bodyLen() int {
	bodyLen := 0
	return bodyLen
}

func makeNil(r iohelp.ErrorReader) (Nil, error) {
	v := Nil{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeNilFromBytes(buf []byte) (Nil, error) {
	v := Nil{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeNilFromBytes(buf []byte) Nil {
	v := Nil{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &List{}

type List struct {
	Cons *Cons
	Nil *Nil
}

func (bbp List) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp List) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.bodyLen()-4))
	at += 4
	if bbp.Cons != nil {
		buf[at] = 1
		at++
		(*bbp.Cons).MarshalBebopTo(buf[at:])
		at += (*bbp.Cons).bodyLen()
	}
	if bbp.Nil != nil {
		buf[at] = 2
		at++
		(*bbp.Nil).MarshalBebopTo(buf[at:])
		at += (*bbp.Nil).bodyLen()
	}
}

func (bbp *List) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.Cons = new(Cons)
			(*bbp.Cons), err = makeConsFromBytes(buf[at:])
			if err != nil {
				 return err
			}
			at += ((*bbp.Cons)).bodyLen()
		case 2:
			at += 1
			bbp.Nil = new(Nil)
			(*bbp.Nil), err = makeNilFromBytes(buf[at:])
			if err != nil {
				 return err
			}
			at += ((*bbp.Nil)).bodyLen()
		default:
			return nil
		}
	}
}

func (bbp *List) MustUnmarshalBebop(buf []byte) {
	at := 0
	for {
		switch buf[at] {
		case 1:
			bbp.Cons = new(Cons)
			(*bbp.Cons) = mustMakeConsFromBytes(buf[at:])
			at += ((*bbp.Cons)).bodyLen()
		case 2:
			bbp.Nil = new(Nil)
			(*bbp.Nil) = mustMakeNilFromBytes(buf[at:])
			at += ((*bbp.Nil)).bodyLen()
		default:
			return
		}
	}
}

func (bbp List) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.bodyLen()-4))
	if bbp.Cons != nil {
		w.Write([]byte{1})
		err = (*bbp.Cons).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	if bbp.Nil != nil {
		w.Write([]byte{2})
		err = (*bbp.Nil).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *List) DecodeBebop(ior io.Reader) (err error) {
	er := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(er)
	body := make([]byte, bodyLen)
	er.Read(body)
	r := iohelp.NewErrorReader(bytes.NewReader(body))
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.Cons = new(Cons)
			(*bbp.Cons), err = makeCons(r)
			if err != nil {
				return err
			}
		case 2:
			bbp.Nil = new(Nil)
			(*bbp.Nil), err = makeNil(r)
			if err != nil {
				return err
			}
		default:
			return er.Err
		}
	}
}

func (bbp List) bodyLen() int {
	bodyLen := 5
	if bbp.Cons != nil {
		bodyLen += 1
		bodyLen += (*bbp.Cons).bodyLen()
	}
	if bbp.Nil != nil {
		bodyLen += 1
		bodyLen += (*bbp.Nil).bodyLen()
	}
	return bodyLen
}

func makeList(r iohelp.ErrorReader) (List, error) {
	v := List{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeListFromBytes(buf []byte) (List, error) {
	v := List{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeListFromBytes(buf []byte) List {
	v := List{}
	v.MustUnmarshalBebop(buf)
	return v
}

