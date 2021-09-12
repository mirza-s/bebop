// Code generated by bebopc-go; DO NOT EDIT.

package generated

import (
	"io"
	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
)

var _ bebop.Record = &A{}

type A struct {
	B *uint32
}

func (bbp A) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.B != nil {
		buf[at] = 1
		at++
		iohelp.WriteUint32Bytes(buf[at:], *bbp.B)
		at += 4
	}
	return at
}

func (bbp *A) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.B = new(uint32)
			if len(buf[at:]) < 4 {
				 return io.ErrUnexpectedEOF
			}
			(*bbp.B) = iohelp.ReadUint32Bytes(buf[at:])
			at += 4
		default:
			return nil
		}
	}
}

func (bbp *A) MustUnmarshalBebop(buf []byte) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.B = new(uint32)
			(*bbp.B) = iohelp.ReadUint32Bytes(buf[at:])
			at += 4
		default:
			return
		}
	}
}

func (bbp A) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.B != nil {
		w.Write([]byte{1})
		iohelp.WriteUint32(w, *bbp.B)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *A) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.B = new(uint32)
			*bbp.B = iohelp.ReadUint32(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp A) Size() int {
	bodyLen := 5
	if bbp.B != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

func (bbp A) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeA(r iohelp.ErrorReader) (A, error) {
	v := A{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeAFromBytes(buf []byte) (A, error) {
	v := A{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeAFromBytes(buf []byte) A {
	v := A{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &B{}

//*
//     * This branch is, too!
//     
type B struct {
	C bool
}

func (bbp B) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteBoolBytes(buf[at:], bbp.C)
	at += 1
	return at
}

func (bbp *B) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 1 {
		 return io.ErrUnexpectedEOF
	}
	bbp.C = iohelp.ReadBoolBytes(buf[at:])
	at += 1
	return nil
}

func (bbp *B) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.C = iohelp.ReadBoolBytes(buf[at:])
	at += 1
}

func (bbp B) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteBool(w, bbp.C)
	return w.Err
}

func (bbp *B) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.C = iohelp.ReadBool(r)
	return r.Err
}

func (bbp B) Size() int {
	bodyLen := 0
	bodyLen += 1
	return bodyLen
}

func (bbp B) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeB(r iohelp.ErrorReader) (B, error) {
	v := B{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeBFromBytes(buf []byte) (B, error) {
	v := B{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeBFromBytes(buf []byte) B {
	v := B{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &C{}

type C struct {
}

func (bbp C) MarshalBebopTo(buf []byte) int {
	return 0
}

func (bbp *C) UnmarshalBebop(buf []byte) (err error) {
	return nil
}

func (bbp *C) MustUnmarshalBebop(buf []byte) {
}

func (bbp C) EncodeBebop(iow io.Writer) (err error) {
	return nil
}

func (bbp *C) DecodeBebop(ior io.Reader) (err error) {
	return nil
}

func (bbp C) Size() int {
	return 0
}

func (bbp C) MarshalBebop() []byte {
	return []byte{}
}

func MakeC(r iohelp.ErrorReader) (C, error) {
	return C{}, nil
}

func MakeCFromBytes(buf []byte) (C, error) {
	return C{}, nil
}

func MustMakeCFromBytes(buf []byte) C {
	return C{}
}

const UOpCode = 0x79656168

var _ bebop.Record = &U{}

//*
// * This union is so documented!
// 
type U struct {
	A *A
	B *B
	C *C
}

func (bbp U) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-5))
	at += 4
	if bbp.A != nil {
		buf[at] = 1
		at++
		(*bbp.A).MarshalBebopTo(buf[at:])
		at += (*bbp.A).Size()
		return at
	}
	if bbp.B != nil {
		buf[at] = 2
		at++
		(*bbp.B).MarshalBebopTo(buf[at:])
		at += (*bbp.B).Size()
		return at
	}
	if bbp.C != nil {
		buf[at] = 3
		at++
		(*bbp.C).MarshalBebopTo(buf[at:])
		at += (*bbp.C).Size()
		return at
	}
	return at
}

func (bbp *U) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	if len(buf) == 0 {
		return iohelp.ErrUnpopulatedUnion
	}
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.A = new(A)
			(*bbp.A), err = MakeAFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.A)).Size()
			return nil
		case 2:
			at += 1
			bbp.B = new(B)
			(*bbp.B), err = MakeBFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.B)).Size()
			return nil
		case 3:
			at += 1
			bbp.C = new(C)
			(*bbp.C), err = MakeCFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.C)).Size()
			return nil
		default:
			return nil
		}
	}
}

func (bbp *U) MustUnmarshalBebop(buf []byte) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.A = new(A)
			(*bbp.A) = MustMakeAFromBytes(buf[at:])
			at += ((*bbp.A)).Size()
			return
		case 2:
			at += 1
			bbp.B = new(B)
			(*bbp.B) = MustMakeBFromBytes(buf[at:])
			at += ((*bbp.B)).Size()
			return
		case 3:
			at += 1
			bbp.C = new(C)
			(*bbp.C) = MustMakeCFromBytes(buf[at:])
			at += ((*bbp.C)).Size()
			return
		default:
			return
		}
	}
}

func (bbp U) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-5))
	if bbp.A != nil {
		w.Write([]byte{1})
		err = (*bbp.A).EncodeBebop(w)
		if err != nil{
			return err
		}
		return w.Err
	}
	if bbp.B != nil {
		w.Write([]byte{2})
		err = (*bbp.B).EncodeBebop(w)
		if err != nil{
			return err
		}
		return w.Err
	}
	if bbp.C != nil {
		w.Write([]byte{3})
		err = (*bbp.C).EncodeBebop(w)
		if err != nil{
			return err
		}
		return w.Err
	}
	return w.Err
}

func (bbp *U) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)+1}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.A = new(A)
			(*bbp.A), err = MakeA(r)
			if err != nil{
				return err
			}
			io.ReadAll(r)
			return r.Err
		case 2:
			bbp.B = new(B)
			(*bbp.B), err = MakeB(r)
			if err != nil{
				return err
			}
			io.ReadAll(r)
			return r.Err
		case 3:
			bbp.C = new(C)
			(*bbp.C), err = MakeC(r)
			if err != nil{
				return err
			}
			io.ReadAll(r)
			return r.Err
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp U) Size() int {
	bodyLen := 4
	if bbp.A != nil {
		bodyLen += 1
		bodyLen += (*bbp.A).Size()
		return bodyLen
	}
	if bbp.B != nil {
		bodyLen += 1
		bodyLen += (*bbp.B).Size()
		return bodyLen
	}
	if bbp.C != nil {
		bodyLen += 1
		bodyLen += (*bbp.C).Size()
		return bodyLen
	}
	return bodyLen
}

func (bbp U) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeU(r iohelp.ErrorReader) (U, error) {
	v := U{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeUFromBytes(buf []byte) (U, error) {
	v := U{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeUFromBytes(buf []byte) U {
	v := U{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &Cons{}

type Cons struct {
	Head uint32
	Tail List
}

func (bbp Cons) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], bbp.Head)
	at += 4
	(bbp.Tail).MarshalBebopTo(buf[at:])
	at += (bbp.Tail).Size()
	return at
}

func (bbp *Cons) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 4 {
		 return io.ErrUnexpectedEOF
	}
	bbp.Head = iohelp.ReadUint32Bytes(buf[at:])
	at += 4
	bbp.Tail, err = MakeListFromBytes(buf[at:])
	if err != nil{
		return err
	}
	at += (bbp.Tail).Size()
	return nil
}

func (bbp *Cons) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.Head = iohelp.ReadUint32Bytes(buf[at:])
	at += 4
	bbp.Tail = MustMakeListFromBytes(buf[at:])
	at += (bbp.Tail).Size()
}

func (bbp Cons) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, bbp.Head)
	err = (bbp.Tail).EncodeBebop(w)
	if err != nil{
		return err
	}
	return w.Err
}

func (bbp *Cons) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.Head = iohelp.ReadUint32(r)
	(bbp.Tail), err = MakeList(r)
	if err != nil{
		return err
	}
	return r.Err
}

func (bbp Cons) Size() int {
	bodyLen := 0
	bodyLen += 4
	bodyLen += (bbp.Tail).Size()
	return bodyLen
}

func (bbp Cons) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeCons(r iohelp.ErrorReader) (Cons, error) {
	v := Cons{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeConsFromBytes(buf []byte) (Cons, error) {
	v := Cons{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeConsFromBytes(buf []byte) Cons {
	v := Cons{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &Nil{}

// nil is empty
type Nil struct {
}

func (bbp Nil) MarshalBebopTo(buf []byte) int {
	return 0
}

func (bbp *Nil) UnmarshalBebop(buf []byte) (err error) {
	return nil
}

func (bbp *Nil) MustUnmarshalBebop(buf []byte) {
}

func (bbp Nil) EncodeBebop(iow io.Writer) (err error) {
	return nil
}

func (bbp *Nil) DecodeBebop(ior io.Reader) (err error) {
	return nil
}

func (bbp Nil) Size() int {
	return 0
}

func (bbp Nil) MarshalBebop() []byte {
	return []byte{}
}

func MakeNil(r iohelp.ErrorReader) (Nil, error) {
	return Nil{}, nil
}

func MakeNilFromBytes(buf []byte) (Nil, error) {
	return Nil{}, nil
}

func MustMakeNilFromBytes(buf []byte) Nil {
	return Nil{}
}

var _ bebop.Record = &List{}

type List struct {
	Cons *Cons
	Nil *Nil
}

func (bbp List) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-5))
	at += 4
	if bbp.Cons != nil {
		buf[at] = 1
		at++
		(*bbp.Cons).MarshalBebopTo(buf[at:])
		at += (*bbp.Cons).Size()
		return at
	}
	if bbp.Nil != nil {
		buf[at] = 2
		at++
		(*bbp.Nil).MarshalBebopTo(buf[at:])
		at += (*bbp.Nil).Size()
		return at
	}
	return at
}

func (bbp *List) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	if len(buf) == 0 {
		return iohelp.ErrUnpopulatedUnion
	}
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.Cons = new(Cons)
			(*bbp.Cons), err = MakeConsFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.Cons)).Size()
			return nil
		case 2:
			at += 1
			bbp.Nil = new(Nil)
			(*bbp.Nil), err = MakeNilFromBytes(buf[at:])
			if err != nil{
				return err
			}
			at += ((*bbp.Nil)).Size()
			return nil
		default:
			return nil
		}
	}
}

func (bbp *List) MustUnmarshalBebop(buf []byte) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.Cons = new(Cons)
			(*bbp.Cons) = MustMakeConsFromBytes(buf[at:])
			at += ((*bbp.Cons)).Size()
			return
		case 2:
			at += 1
			bbp.Nil = new(Nil)
			(*bbp.Nil) = MustMakeNilFromBytes(buf[at:])
			at += ((*bbp.Nil)).Size()
			return
		default:
			return
		}
	}
}

func (bbp List) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-5))
	if bbp.Cons != nil {
		w.Write([]byte{1})
		err = (*bbp.Cons).EncodeBebop(w)
		if err != nil{
			return err
		}
		return w.Err
	}
	if bbp.Nil != nil {
		w.Write([]byte{2})
		err = (*bbp.Nil).EncodeBebop(w)
		if err != nil{
			return err
		}
		return w.Err
	}
	return w.Err
}

func (bbp *List) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)+1}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.Cons = new(Cons)
			(*bbp.Cons), err = MakeCons(r)
			if err != nil{
				return err
			}
			io.ReadAll(r)
			return r.Err
		case 2:
			bbp.Nil = new(Nil)
			(*bbp.Nil), err = MakeNil(r)
			if err != nil{
				return err
			}
			io.ReadAll(r)
			return r.Err
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp List) Size() int {
	bodyLen := 4
	if bbp.Cons != nil {
		bodyLen += 1
		bodyLen += (*bbp.Cons).Size()
		return bodyLen
	}
	if bbp.Nil != nil {
		bodyLen += 1
		bodyLen += (*bbp.Nil).Size()
		return bodyLen
	}
	return bodyLen
}

func (bbp List) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeList(r iohelp.ErrorReader) (List, error) {
	v := List{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeListFromBytes(buf []byte) (List, error) {
	v := List{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeListFromBytes(buf []byte) List {
	v := List{}
	v.MustUnmarshalBebop(buf)
	return v
}

