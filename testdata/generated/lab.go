// Code generated by bebopc-go; DO NOT EDIT.

package generated

import (
	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
	"io"
)

type VideoCodec uint32

const (
	VideoCodec_H264 VideoCodec = 0
	VideoCodec_H265 VideoCodec = 1
)

var _ bebop.Record = &Int32s{}

type Int32s struct {
	A []int32
}

func (bbp Int32s) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.A)))
	at += 4
	for _, v1 := range bbp.A {
		iohelp.WriteInt32Bytes(buf[at:], v1)
		at += 4
	}
	return at
}

func (bbp *Int32s) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 4 {
		return io.ErrUnexpectedEOF
	}
	bbp.A = make([]int32, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	if len(buf[at:]) < len(bbp.A)*4 {
		return io.ErrUnexpectedEOF
	}
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadInt32Bytes(buf[at:])
		at += 4
	}
	return nil
}

func (bbp *Int32s) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.A = make([]int32, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadInt32Bytes(buf[at:])
		at += 4
	}
}

func (bbp Int32s) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(len(bbp.A)))
	for _, elem := range bbp.A {
		iohelp.WriteInt32(w, elem)
	}
	return w.Err
}

func (bbp *Int32s) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.A = make([]int32, iohelp.ReadUint32(r))
	for i1 := range bbp.A {
		(bbp.A[i1]) = iohelp.ReadInt32(r)
	}
	return r.Err
}

func (bbp Int32s) Size() int {
	bodyLen := 0
	bodyLen += 4
	bodyLen += len(bbp.A) * 4
	return bodyLen
}

func (bbp Int32s) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeInt32s(r iohelp.ErrorReader) (Int32s, error) {
	v := Int32s{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeInt32sFromBytes(buf []byte) (Int32s, error) {
	v := Int32s{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeInt32sFromBytes(buf []byte) Int32s {
	v := Int32s{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &Uint32s{}

type Uint32s struct {
	A []uint32
}

func (bbp Uint32s) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.A)))
	at += 4
	for _, v1 := range bbp.A {
		iohelp.WriteUint32Bytes(buf[at:], v1)
		at += 4
	}
	return at
}

func (bbp *Uint32s) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 4 {
		return io.ErrUnexpectedEOF
	}
	bbp.A = make([]uint32, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	if len(buf[at:]) < len(bbp.A)*4 {
		return io.ErrUnexpectedEOF
	}
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadUint32Bytes(buf[at:])
		at += 4
	}
	return nil
}

func (bbp *Uint32s) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.A = make([]uint32, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadUint32Bytes(buf[at:])
		at += 4
	}
}

func (bbp Uint32s) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(len(bbp.A)))
	for _, elem := range bbp.A {
		iohelp.WriteUint32(w, elem)
	}
	return w.Err
}

func (bbp *Uint32s) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.A = make([]uint32, iohelp.ReadUint32(r))
	for i1 := range bbp.A {
		(bbp.A[i1]) = iohelp.ReadUint32(r)
	}
	return r.Err
}

func (bbp Uint32s) Size() int {
	bodyLen := 0
	bodyLen += 4
	bodyLen += len(bbp.A) * 4
	return bodyLen
}

func (bbp Uint32s) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeUint32s(r iohelp.ErrorReader) (Uint32s, error) {
	v := Uint32s{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeUint32sFromBytes(buf []byte) (Uint32s, error) {
	v := Uint32s{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeUint32sFromBytes(buf []byte) Uint32s {
	v := Uint32s{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &Float32s{}

type Float32s struct {
	A []float32
}

func (bbp Float32s) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.A)))
	at += 4
	for _, v1 := range bbp.A {
		iohelp.WriteFloat32Bytes(buf[at:], v1)
		at += 4
	}
	return at
}

func (bbp *Float32s) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 4 {
		return io.ErrUnexpectedEOF
	}
	bbp.A = make([]float32, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	if len(buf[at:]) < len(bbp.A)*4 {
		return io.ErrUnexpectedEOF
	}
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadFloat32Bytes(buf[at:])
		at += 4
	}
	return nil
}

func (bbp *Float32s) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.A = make([]float32, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadFloat32Bytes(buf[at:])
		at += 4
	}
}

func (bbp Float32s) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(len(bbp.A)))
	for _, elem := range bbp.A {
		iohelp.WriteFloat32(w, elem)
	}
	return w.Err
}

func (bbp *Float32s) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.A = make([]float32, iohelp.ReadUint32(r))
	for i1 := range bbp.A {
		(bbp.A[i1]) = iohelp.ReadFloat32(r)
	}
	return r.Err
}

func (bbp Float32s) Size() int {
	bodyLen := 0
	bodyLen += 4
	bodyLen += len(bbp.A) * 4
	return bodyLen
}

func (bbp Float32s) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeFloat32s(r iohelp.ErrorReader) (Float32s, error) {
	v := Float32s{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeFloat32sFromBytes(buf []byte) (Float32s, error) {
	v := Float32s{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeFloat32sFromBytes(buf []byte) Float32s {
	v := Float32s{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &Int64s{}

type Int64s struct {
	A []int64
}

func (bbp Int64s) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.A)))
	at += 4
	for _, v1 := range bbp.A {
		iohelp.WriteInt64Bytes(buf[at:], v1)
		at += 8
	}
	return at
}

func (bbp *Int64s) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 4 {
		return io.ErrUnexpectedEOF
	}
	bbp.A = make([]int64, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	if len(buf[at:]) < len(bbp.A)*8 {
		return io.ErrUnexpectedEOF
	}
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadInt64Bytes(buf[at:])
		at += 8
	}
	return nil
}

func (bbp *Int64s) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.A = make([]int64, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadInt64Bytes(buf[at:])
		at += 8
	}
}

func (bbp Int64s) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(len(bbp.A)))
	for _, elem := range bbp.A {
		iohelp.WriteInt64(w, elem)
	}
	return w.Err
}

func (bbp *Int64s) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.A = make([]int64, iohelp.ReadUint32(r))
	for i1 := range bbp.A {
		(bbp.A[i1]) = iohelp.ReadInt64(r)
	}
	return r.Err
}

func (bbp Int64s) Size() int {
	bodyLen := 0
	bodyLen += 4
	bodyLen += len(bbp.A) * 8
	return bodyLen
}

func (bbp Int64s) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeInt64s(r iohelp.ErrorReader) (Int64s, error) {
	v := Int64s{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeInt64sFromBytes(buf []byte) (Int64s, error) {
	v := Int64s{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeInt64sFromBytes(buf []byte) Int64s {
	v := Int64s{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &Uint64s{}

type Uint64s struct {
	A []uint64
}

func (bbp Uint64s) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.A)))
	at += 4
	for _, v1 := range bbp.A {
		iohelp.WriteUint64Bytes(buf[at:], v1)
		at += 8
	}
	return at
}

func (bbp *Uint64s) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 4 {
		return io.ErrUnexpectedEOF
	}
	bbp.A = make([]uint64, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	if len(buf[at:]) < len(bbp.A)*8 {
		return io.ErrUnexpectedEOF
	}
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadUint64Bytes(buf[at:])
		at += 8
	}
	return nil
}

func (bbp *Uint64s) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.A = make([]uint64, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadUint64Bytes(buf[at:])
		at += 8
	}
}

func (bbp Uint64s) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(len(bbp.A)))
	for _, elem := range bbp.A {
		iohelp.WriteUint64(w, elem)
	}
	return w.Err
}

func (bbp *Uint64s) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.A = make([]uint64, iohelp.ReadUint32(r))
	for i1 := range bbp.A {
		(bbp.A[i1]) = iohelp.ReadUint64(r)
	}
	return r.Err
}

func (bbp Uint64s) Size() int {
	bodyLen := 0
	bodyLen += 4
	bodyLen += len(bbp.A) * 8
	return bodyLen
}

func (bbp Uint64s) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeUint64s(r iohelp.ErrorReader) (Uint64s, error) {
	v := Uint64s{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeUint64sFromBytes(buf []byte) (Uint64s, error) {
	v := Uint64s{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeUint64sFromBytes(buf []byte) Uint64s {
	v := Uint64s{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &Float64s{}

type Float64s struct {
	A []float64
}

func (bbp Float64s) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.A)))
	at += 4
	for _, v1 := range bbp.A {
		iohelp.WriteFloat64Bytes(buf[at:], v1)
		at += 8
	}
	return at
}

func (bbp *Float64s) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 4 {
		return io.ErrUnexpectedEOF
	}
	bbp.A = make([]float64, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	if len(buf[at:]) < len(bbp.A)*8 {
		return io.ErrUnexpectedEOF
	}
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadFloat64Bytes(buf[at:])
		at += 8
	}
	return nil
}

func (bbp *Float64s) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.A = make([]float64, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	for i1 := range bbp.A {
		(bbp.A)[i1] = iohelp.ReadFloat64Bytes(buf[at:])
		at += 8
	}
}

func (bbp Float64s) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(len(bbp.A)))
	for _, elem := range bbp.A {
		iohelp.WriteFloat64(w, elem)
	}
	return w.Err
}

func (bbp *Float64s) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.A = make([]float64, iohelp.ReadUint32(r))
	for i1 := range bbp.A {
		(bbp.A[i1]) = iohelp.ReadFloat64(r)
	}
	return r.Err
}

func (bbp Float64s) Size() int {
	bodyLen := 0
	bodyLen += 4
	bodyLen += len(bbp.A) * 8
	return bodyLen
}

func (bbp Float64s) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeFloat64s(r iohelp.ErrorReader) (Float64s, error) {
	v := Float64s{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeFloat64sFromBytes(buf []byte) (Float64s, error) {
	v := Float64s{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeFloat64sFromBytes(buf []byte) Float64s {
	v := Float64s{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &VideoData{}

type VideoData struct {
	Time float64
	Width uint32
	Height uint32
	Fragment []byte
}

func (bbp VideoData) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteFloat64Bytes(buf[at:], bbp.Time)
	at += 8
	iohelp.WriteUint32Bytes(buf[at:], bbp.Width)
	at += 4
	iohelp.WriteUint32Bytes(buf[at:], bbp.Height)
	at += 4
	iohelp.WriteUint32Bytes(buf[at:], uint32(len(bbp.Fragment)))
	at += 4
	copy(buf[at:at+len(bbp.Fragment)], bbp.Fragment)
	at += len(bbp.Fragment)
	return at
}

func (bbp *VideoData) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 8 {
		return io.ErrUnexpectedEOF
	}
	bbp.Time = iohelp.ReadFloat64Bytes(buf[at:])
	at += 8
	if len(buf[at:]) < 4 {
		return io.ErrUnexpectedEOF
	}
	bbp.Width = iohelp.ReadUint32Bytes(buf[at:])
	at += 4
	if len(buf[at:]) < 4 {
		return io.ErrUnexpectedEOF
	}
	bbp.Height = iohelp.ReadUint32Bytes(buf[at:])
	at += 4
	if len(buf[at:]) < 4 {
		return io.ErrUnexpectedEOF
	}
	bbp.Fragment = make([]byte, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	if len(buf[at:]) < len(bbp.Fragment)*1 {
		return io.ErrUnexpectedEOF
	}
	copy(bbp.Fragment, buf[at:at+len(bbp.Fragment)])
	at += len(bbp.Fragment)
	return nil
}

func (bbp *VideoData) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.Time = iohelp.ReadFloat64Bytes(buf[at:])
	at += 8
	bbp.Width = iohelp.ReadUint32Bytes(buf[at:])
	at += 4
	bbp.Height = iohelp.ReadUint32Bytes(buf[at:])
	at += 4
	bbp.Fragment = make([]byte, iohelp.ReadUint32Bytes(buf[at:]))
	at += 4
	copy(bbp.Fragment, buf[at:at+len(bbp.Fragment)])
	at += len(bbp.Fragment)
}

func (bbp VideoData) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteFloat64(w, bbp.Time)
	iohelp.WriteUint32(w, bbp.Width)
	iohelp.WriteUint32(w, bbp.Height)
	iohelp.WriteUint32(w, uint32(len(bbp.Fragment)))
	for _, elem := range bbp.Fragment {
		iohelp.WriteByte(w, elem)
	}
	return w.Err
}

func (bbp *VideoData) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.Time = iohelp.ReadFloat64(r)
	bbp.Width = iohelp.ReadUint32(r)
	bbp.Height = iohelp.ReadUint32(r)
	bbp.Fragment = make([]byte, iohelp.ReadUint32(r))
	for i1 := range bbp.Fragment {
		(bbp.Fragment[i1]) = iohelp.ReadByte(r)
	}
	return r.Err
}

func (bbp VideoData) Size() int {
	bodyLen := 0
	bodyLen += 8
	bodyLen += 4
	bodyLen += 4
	bodyLen += 4
	bodyLen += len(bbp.Fragment) * 1
	return bodyLen
}

func (bbp VideoData) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeVideoData(r iohelp.ErrorReader) (VideoData, error) {
	v := VideoData{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeVideoDataFromBytes(buf []byte) (VideoData, error) {
	v := VideoData{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeVideoDataFromBytes(buf []byte) VideoData {
	v := VideoData{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &MediaMessage{}

type MediaMessage struct {
	Codec *VideoCodec
	Data *VideoData
}

func (bbp MediaMessage) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.Codec != nil {
		buf[at] = 1
		at++
		iohelp.WriteUint32Bytes(buf[at:], uint32(*bbp.Codec))
		at += 4
	}
	if bbp.Data != nil {
		buf[at] = 2
		at++
		(*bbp.Data).MarshalBebopTo(buf[at:])
		at += (*bbp.Data).Size()
	}
	return at
}

func (bbp *MediaMessage) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.Codec = new(VideoCodec)
			(*bbp.Codec) = VideoCodec(iohelp.ReadUint32Bytes(buf[at:]))
			at += 4
		case 2:
			at += 1
			bbp.Data = new(VideoData)
			(*bbp.Data) = MustMakeVideoDataFromBytes(buf[at:])
			if err != nil {
				return err
			}
			at += ((*bbp.Data)).Size()
		default:
			return nil
		}
	}
}

func (bbp *MediaMessage) MustUnmarshalBebop(buf []byte) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.Codec = new(VideoCodec)
			(*bbp.Codec) = VideoCodec(iohelp.ReadUint32Bytes(buf[at:]))
			at += 4
		case 2:
			at += 1
			bbp.Data = new(VideoData)
			(*bbp.Data) = MustMakeVideoDataFromBytes(buf[at:])
			at += ((*bbp.Data)).Size()
		default:
			return
		}
	}
}

func (bbp MediaMessage) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.Codec != nil {
		w.Write([]byte{1})
		iohelp.WriteUint32(w, uint32(*bbp.Codec))
	}
	if bbp.Data != nil {
		w.Write([]byte{2})
		err = (*bbp.Data).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *MediaMessage) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.Codec = new(VideoCodec)
			*bbp.Codec = VideoCodec(iohelp.ReadUint32(r))
		case 2:
			bbp.Data = new(VideoData)
			(*bbp.Data), err = MakeVideoData(r)
			if err != nil {
				return err
			}
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp MediaMessage) Size() int {
	bodyLen := 5
	if bbp.Codec != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Data != nil {
		bodyLen += 1
		bodyLen += (*bbp.Data).Size()
	}
	return bodyLen
}

func (bbp MediaMessage) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeMediaMessage(r iohelp.ErrorReader) (MediaMessage, error) {
	v := MediaMessage{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeMediaMessageFromBytes(buf []byte) (MediaMessage, error) {
	v := MediaMessage{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeMediaMessageFromBytes(buf []byte) MediaMessage {
	v := MediaMessage{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &SkipTestOld{}

// Should be able to decode a "SkipTestNewContainer" as a "SkipTestOldContainer".
type SkipTestOld struct {
	X *int32
	Y *int32
}

func (bbp SkipTestOld) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.X != nil {
		buf[at] = 1
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.X)
		at += 4
	}
	if bbp.Y != nil {
		buf[at] = 2
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.Y)
		at += 4
	}
	return at
}

func (bbp *SkipTestOld) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.X = new(int32)
			if len(buf[at:]) < 4 {
				return io.ErrUnexpectedEOF
			}
			(*bbp.X) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 2:
			at += 1
			bbp.Y = new(int32)
			if len(buf[at:]) < 4 {
				return io.ErrUnexpectedEOF
			}
			(*bbp.Y) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return nil
		}
	}
}

func (bbp *SkipTestOld) MustUnmarshalBebop(buf []byte) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.X = new(int32)
			(*bbp.X) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 2:
			at += 1
			bbp.Y = new(int32)
			(*bbp.Y) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return
		}
	}
}

func (bbp SkipTestOld) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.X != nil {
		w.Write([]byte{1})
		iohelp.WriteInt32(w, *bbp.X)
	}
	if bbp.Y != nil {
		w.Write([]byte{2})
		iohelp.WriteInt32(w, *bbp.Y)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *SkipTestOld) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.X = new(int32)
			*bbp.X = iohelp.ReadInt32(r)
		case 2:
			bbp.Y = new(int32)
			*bbp.Y = iohelp.ReadInt32(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp SkipTestOld) Size() int {
	bodyLen := 5
	if bbp.X != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Y != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

func (bbp SkipTestOld) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeSkipTestOld(r iohelp.ErrorReader) (SkipTestOld, error) {
	v := SkipTestOld{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeSkipTestOldFromBytes(buf []byte) (SkipTestOld, error) {
	v := SkipTestOld{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeSkipTestOldFromBytes(buf []byte) SkipTestOld {
	v := SkipTestOld{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &SkipTestNew{}

type SkipTestNew struct {
	X *int32
	Y *int32
	Z *int32
}

func (bbp SkipTestNew) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.X != nil {
		buf[at] = 1
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.X)
		at += 4
	}
	if bbp.Y != nil {
		buf[at] = 2
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.Y)
		at += 4
	}
	if bbp.Z != nil {
		buf[at] = 3
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.Z)
		at += 4
	}
	return at
}

func (bbp *SkipTestNew) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.X = new(int32)
			if len(buf[at:]) < 4 {
				return io.ErrUnexpectedEOF
			}
			(*bbp.X) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 2:
			at += 1
			bbp.Y = new(int32)
			if len(buf[at:]) < 4 {
				return io.ErrUnexpectedEOF
			}
			(*bbp.Y) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 3:
			at += 1
			bbp.Z = new(int32)
			if len(buf[at:]) < 4 {
				return io.ErrUnexpectedEOF
			}
			(*bbp.Z) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return nil
		}
	}
}

func (bbp *SkipTestNew) MustUnmarshalBebop(buf []byte) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.X = new(int32)
			(*bbp.X) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 2:
			at += 1
			bbp.Y = new(int32)
			(*bbp.Y) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 3:
			at += 1
			bbp.Z = new(int32)
			(*bbp.Z) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return
		}
	}
}

func (bbp SkipTestNew) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.X != nil {
		w.Write([]byte{1})
		iohelp.WriteInt32(w, *bbp.X)
	}
	if bbp.Y != nil {
		w.Write([]byte{2})
		iohelp.WriteInt32(w, *bbp.Y)
	}
	if bbp.Z != nil {
		w.Write([]byte{3})
		iohelp.WriteInt32(w, *bbp.Z)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *SkipTestNew) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.X = new(int32)
			*bbp.X = iohelp.ReadInt32(r)
		case 2:
			bbp.Y = new(int32)
			*bbp.Y = iohelp.ReadInt32(r)
		case 3:
			bbp.Z = new(int32)
			*bbp.Z = iohelp.ReadInt32(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp SkipTestNew) Size() int {
	bodyLen := 5
	if bbp.X != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Y != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Z != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

func (bbp SkipTestNew) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeSkipTestNew(r iohelp.ErrorReader) (SkipTestNew, error) {
	v := SkipTestNew{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeSkipTestNewFromBytes(buf []byte) (SkipTestNew, error) {
	v := SkipTestNew{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeSkipTestNewFromBytes(buf []byte) SkipTestNew {
	v := SkipTestNew{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &SkipTestOldContainer{}

type SkipTestOldContainer struct {
	S *SkipTestOld
	After *int32
}

func (bbp SkipTestOldContainer) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.S != nil {
		buf[at] = 1
		at++
		(*bbp.S).MarshalBebopTo(buf[at:])
		at += (*bbp.S).Size()
	}
	if bbp.After != nil {
		buf[at] = 2
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.After)
		at += 4
	}
	return at
}

func (bbp *SkipTestOldContainer) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.S = new(SkipTestOld)
			(*bbp.S) = MustMakeSkipTestOldFromBytes(buf[at:])
			if err != nil {
				return err
			}
			at += ((*bbp.S)).Size()
		case 2:
			at += 1
			bbp.After = new(int32)
			if len(buf[at:]) < 4 {
				return io.ErrUnexpectedEOF
			}
			(*bbp.After) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return nil
		}
	}
}

func (bbp *SkipTestOldContainer) MustUnmarshalBebop(buf []byte) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.S = new(SkipTestOld)
			(*bbp.S) = MustMakeSkipTestOldFromBytes(buf[at:])
			at += ((*bbp.S)).Size()
		case 2:
			at += 1
			bbp.After = new(int32)
			(*bbp.After) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return
		}
	}
}

func (bbp SkipTestOldContainer) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.S != nil {
		w.Write([]byte{1})
		err = (*bbp.S).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	if bbp.After != nil {
		w.Write([]byte{2})
		iohelp.WriteInt32(w, *bbp.After)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *SkipTestOldContainer) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.S = new(SkipTestOld)
			(*bbp.S), err = MakeSkipTestOld(r)
			if err != nil {
				return err
			}
		case 2:
			bbp.After = new(int32)
			*bbp.After = iohelp.ReadInt32(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp SkipTestOldContainer) Size() int {
	bodyLen := 5
	if bbp.S != nil {
		bodyLen += 1
		bodyLen += (*bbp.S).Size()
	}
	if bbp.After != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

func (bbp SkipTestOldContainer) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeSkipTestOldContainer(r iohelp.ErrorReader) (SkipTestOldContainer, error) {
	v := SkipTestOldContainer{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeSkipTestOldContainerFromBytes(buf []byte) (SkipTestOldContainer, error) {
	v := SkipTestOldContainer{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeSkipTestOldContainerFromBytes(buf []byte) SkipTestOldContainer {
	v := SkipTestOldContainer{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &SkipTestNewContainer{}

type SkipTestNewContainer struct {
	S *SkipTestNew
	After *int32
}

func (bbp SkipTestNewContainer) MarshalBebopTo(buf []byte) int {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.Size()-4))
	at += 4
	if bbp.S != nil {
		buf[at] = 1
		at++
		(*bbp.S).MarshalBebopTo(buf[at:])
		at += (*bbp.S).Size()
	}
	if bbp.After != nil {
		buf[at] = 2
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.After)
		at += 4
	}
	return at
}

func (bbp *SkipTestNewContainer) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.S = new(SkipTestNew)
			(*bbp.S) = MustMakeSkipTestNewFromBytes(buf[at:])
			if err != nil {
				return err
			}
			at += ((*bbp.S)).Size()
		case 2:
			at += 1
			bbp.After = new(int32)
			if len(buf[at:]) < 4 {
				return io.ErrUnexpectedEOF
			}
			(*bbp.After) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return nil
		}
	}
}

func (bbp *SkipTestNewContainer) MustUnmarshalBebop(buf []byte) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.S = new(SkipTestNew)
			(*bbp.S) = MustMakeSkipTestNewFromBytes(buf[at:])
			at += ((*bbp.S)).Size()
		case 2:
			at += 1
			bbp.After = new(int32)
			(*bbp.After) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return
		}
	}
}

func (bbp SkipTestNewContainer) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.Size()-4))
	if bbp.S != nil {
		w.Write([]byte{1})
		err = (*bbp.S).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	if bbp.After != nil {
		w.Write([]byte{2})
		iohelp.WriteInt32(w, *bbp.After)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *SkipTestNewContainer) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(r)
	r.Reader = &io.LimitedReader{R:r.Reader, N:int64(bodyLen)}
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.S = new(SkipTestNew)
			(*bbp.S), err = MakeSkipTestNew(r)
			if err != nil {
				return err
			}
		case 2:
			bbp.After = new(int32)
			*bbp.After = iohelp.ReadInt32(r)
		default:
			io.ReadAll(r)
			return r.Err
		}
	}
}

func (bbp SkipTestNewContainer) Size() int {
	bodyLen := 5
	if bbp.S != nil {
		bodyLen += 1
		bodyLen += (*bbp.S).Size()
	}
	if bbp.After != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

func (bbp SkipTestNewContainer) MarshalBebop() []byte {
	buf := make([]byte, bbp.Size())
	bbp.MarshalBebopTo(buf)
	return buf
}

func MakeSkipTestNewContainer(r iohelp.ErrorReader) (SkipTestNewContainer, error) {
	v := SkipTestNewContainer{}
	err := v.DecodeBebop(r)
	return v, err
}

func MakeSkipTestNewContainerFromBytes(buf []byte) (SkipTestNewContainer, error) {
	v := SkipTestNewContainer{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func MustMakeSkipTestNewContainerFromBytes(buf []byte) SkipTestNewContainer {
	v := SkipTestNewContainer{}
	v.MustUnmarshalBebop(buf)
	return v
}

