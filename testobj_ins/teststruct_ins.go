// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"bytes"
	"encoding/json"
	"github.com/koykov/byteconv"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

func init() {
	inspector.RegisterInspector("TestStruct", TestStructInspector{})
}

type TestStructInspector struct {
	inspector.BaseInspector
}

func (i11 TestStructInspector) TypeName() string {
	return "TestStruct"
}

func (i11 TestStructInspector) Get(src any, path ...string) (any, error) {
	var buf any
	err := i11.GetTo(src, &buf, path...)
	return buf, err
}

func (i11 TestStructInspector) GetTo(src any, buf *any, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestStruct
	_ = x
	if p, ok := src.(**testobj.TestStruct); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStruct); ok {
		x = p
	} else if v, ok := src.(testobj.TestStruct); ok {
		x = &v
	} else {
		return
	}
	if len(path) == 0 {
		*buf = &(*x)
		return
	}

	if len(path) > 0 {
		if path[0] == "A" {
			*buf = &x.A
			return
		}
		if path[0] == "S" {
			*buf = &x.S
			return
		}
		if path[0] == "B" {
			*buf = &x.B
			return
		}
		if path[0] == "I" {
			*buf = &x.I
			return
		}
		if path[0] == "I8" {
			*buf = &x.I8
			return
		}
		if path[0] == "I16" {
			*buf = &x.I16
			return
		}
		if path[0] == "I32" {
			*buf = &x.I32
			return
		}
		if path[0] == "I64" {
			*buf = &x.I64
			return
		}
		if path[0] == "U" {
			*buf = &x.U
			return
		}
		if path[0] == "U8" {
			*buf = &x.U8
			return
		}
		if path[0] == "U16" {
			*buf = &x.U16
			return
		}
		if path[0] == "U32" {
			*buf = &x.U32
			return
		}
		if path[0] == "U64" {
			*buf = &x.U64
			return
		}
		if path[0] == "F" {
			*buf = &x.F
			return
		}
		if path[0] == "D" {
			*buf = &x.D
			return
		}
	}
	return
}

func (i11 TestStructInspector) Compare(src any, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestStruct
	_ = x
	if p, ok := src.(**testobj.TestStruct); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStruct); ok {
		x = p
	} else if v, ok := src.(testobj.TestStruct); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "A" {
			var rightExact byte
			t356 := byteconv.S2B(right)
			if len(t356) > 0 {
				rightExact = t356[0]
			}

			switch cond {
			case inspector.OpEq:
				*result = x.A == rightExact
			case inspector.OpNq:
				*result = x.A != rightExact
			case inspector.OpGt:
				*result = x.A > rightExact
			case inspector.OpGtq:
				*result = x.A >= rightExact
			case inspector.OpLt:
				*result = x.A < rightExact
			case inspector.OpLtq:
				*result = x.A <= rightExact
			}
			return
		}
		if path[0] == "S" {
			var rightExact string
			rightExact = right

			switch cond {
			case inspector.OpEq:
				*result = x.S == rightExact
			case inspector.OpNq:
				*result = x.S != rightExact
			case inspector.OpGt:
				*result = x.S > rightExact
			case inspector.OpGtq:
				*result = x.S >= rightExact
			case inspector.OpLt:
				*result = x.S < rightExact
			case inspector.OpLtq:
				*result = x.S <= rightExact
			}
			return
		}
		if path[0] == "B" {
			var rightExact []byte
			rightExact = byteconv.S2B(right)

			if cond == inspector.OpEq {
				*result = bytes.Equal(x.B, rightExact)
			} else {
				*result = !bytes.Equal(x.B, rightExact)
			}
			return
		}
		if path[0] == "I" {
			var rightExact int
			t359, err359 := strconv.ParseInt(right, 0, 0)
			if err359 != nil {
				return err359
			}
			rightExact = int(t359)
			switch cond {
			case inspector.OpEq:
				*result = x.I == rightExact
			case inspector.OpNq:
				*result = x.I != rightExact
			case inspector.OpGt:
				*result = x.I > rightExact
			case inspector.OpGtq:
				*result = x.I >= rightExact
			case inspector.OpLt:
				*result = x.I < rightExact
			case inspector.OpLtq:
				*result = x.I <= rightExact
			}
			return
		}
		if path[0] == "I8" {
			var rightExact int8
			t360, err360 := strconv.ParseInt(right, 0, 0)
			if err360 != nil {
				return err360
			}
			rightExact = int8(t360)
			switch cond {
			case inspector.OpEq:
				*result = x.I8 == rightExact
			case inspector.OpNq:
				*result = x.I8 != rightExact
			case inspector.OpGt:
				*result = x.I8 > rightExact
			case inspector.OpGtq:
				*result = x.I8 >= rightExact
			case inspector.OpLt:
				*result = x.I8 < rightExact
			case inspector.OpLtq:
				*result = x.I8 <= rightExact
			}
			return
		}
		if path[0] == "I16" {
			var rightExact int16
			t361, err361 := strconv.ParseInt(right, 0, 0)
			if err361 != nil {
				return err361
			}
			rightExact = int16(t361)
			switch cond {
			case inspector.OpEq:
				*result = x.I16 == rightExact
			case inspector.OpNq:
				*result = x.I16 != rightExact
			case inspector.OpGt:
				*result = x.I16 > rightExact
			case inspector.OpGtq:
				*result = x.I16 >= rightExact
			case inspector.OpLt:
				*result = x.I16 < rightExact
			case inspector.OpLtq:
				*result = x.I16 <= rightExact
			}
			return
		}
		if path[0] == "I32" {
			var rightExact int32
			t362, err362 := strconv.ParseInt(right, 0, 0)
			if err362 != nil {
				return err362
			}
			rightExact = int32(t362)
			switch cond {
			case inspector.OpEq:
				*result = x.I32 == rightExact
			case inspector.OpNq:
				*result = x.I32 != rightExact
			case inspector.OpGt:
				*result = x.I32 > rightExact
			case inspector.OpGtq:
				*result = x.I32 >= rightExact
			case inspector.OpLt:
				*result = x.I32 < rightExact
			case inspector.OpLtq:
				*result = x.I32 <= rightExact
			}
			return
		}
		if path[0] == "I64" {
			var rightExact int64
			t363, err363 := strconv.ParseInt(right, 0, 0)
			if err363 != nil {
				return err363
			}
			rightExact = int64(t363)
			switch cond {
			case inspector.OpEq:
				*result = x.I64 == rightExact
			case inspector.OpNq:
				*result = x.I64 != rightExact
			case inspector.OpGt:
				*result = x.I64 > rightExact
			case inspector.OpGtq:
				*result = x.I64 >= rightExact
			case inspector.OpLt:
				*result = x.I64 < rightExact
			case inspector.OpLtq:
				*result = x.I64 <= rightExact
			}
			return
		}
		if path[0] == "U" {
			var rightExact uint
			t364, err364 := strconv.ParseUint(right, 0, 0)
			if err364 != nil {
				return err364
			}
			rightExact = uint(t364)
			switch cond {
			case inspector.OpEq:
				*result = x.U == rightExact
			case inspector.OpNq:
				*result = x.U != rightExact
			case inspector.OpGt:
				*result = x.U > rightExact
			case inspector.OpGtq:
				*result = x.U >= rightExact
			case inspector.OpLt:
				*result = x.U < rightExact
			case inspector.OpLtq:
				*result = x.U <= rightExact
			}
			return
		}
		if path[0] == "U8" {
			var rightExact uint8
			t365, err365 := strconv.ParseUint(right, 0, 0)
			if err365 != nil {
				return err365
			}
			rightExact = uint8(t365)
			switch cond {
			case inspector.OpEq:
				*result = x.U8 == rightExact
			case inspector.OpNq:
				*result = x.U8 != rightExact
			case inspector.OpGt:
				*result = x.U8 > rightExact
			case inspector.OpGtq:
				*result = x.U8 >= rightExact
			case inspector.OpLt:
				*result = x.U8 < rightExact
			case inspector.OpLtq:
				*result = x.U8 <= rightExact
			}
			return
		}
		if path[0] == "U16" {
			var rightExact uint16
			t366, err366 := strconv.ParseUint(right, 0, 0)
			if err366 != nil {
				return err366
			}
			rightExact = uint16(t366)
			switch cond {
			case inspector.OpEq:
				*result = x.U16 == rightExact
			case inspector.OpNq:
				*result = x.U16 != rightExact
			case inspector.OpGt:
				*result = x.U16 > rightExact
			case inspector.OpGtq:
				*result = x.U16 >= rightExact
			case inspector.OpLt:
				*result = x.U16 < rightExact
			case inspector.OpLtq:
				*result = x.U16 <= rightExact
			}
			return
		}
		if path[0] == "U32" {
			var rightExact uint32
			t367, err367 := strconv.ParseUint(right, 0, 0)
			if err367 != nil {
				return err367
			}
			rightExact = uint32(t367)
			switch cond {
			case inspector.OpEq:
				*result = x.U32 == rightExact
			case inspector.OpNq:
				*result = x.U32 != rightExact
			case inspector.OpGt:
				*result = x.U32 > rightExact
			case inspector.OpGtq:
				*result = x.U32 >= rightExact
			case inspector.OpLt:
				*result = x.U32 < rightExact
			case inspector.OpLtq:
				*result = x.U32 <= rightExact
			}
			return
		}
		if path[0] == "U64" {
			var rightExact uint64
			t368, err368 := strconv.ParseUint(right, 0, 0)
			if err368 != nil {
				return err368
			}
			rightExact = uint64(t368)
			switch cond {
			case inspector.OpEq:
				*result = x.U64 == rightExact
			case inspector.OpNq:
				*result = x.U64 != rightExact
			case inspector.OpGt:
				*result = x.U64 > rightExact
			case inspector.OpGtq:
				*result = x.U64 >= rightExact
			case inspector.OpLt:
				*result = x.U64 < rightExact
			case inspector.OpLtq:
				*result = x.U64 <= rightExact
			}
			return
		}
		if path[0] == "F" {
			var rightExact float32
			t369, err369 := strconv.ParseFloat(right, 0)
			if err369 != nil {
				return err369
			}
			rightExact = float32(t369)
			switch cond {
			case inspector.OpEq:
				*result = x.F == rightExact
			case inspector.OpNq:
				*result = x.F != rightExact
			case inspector.OpGt:
				*result = x.F > rightExact
			case inspector.OpGtq:
				*result = x.F >= rightExact
			case inspector.OpLt:
				*result = x.F < rightExact
			case inspector.OpLtq:
				*result = x.F <= rightExact
			}
			return
		}
		if path[0] == "D" {
			var rightExact float64
			t370, err370 := strconv.ParseFloat(right, 0)
			if err370 != nil {
				return err370
			}
			rightExact = float64(t370)
			switch cond {
			case inspector.OpEq:
				*result = x.D == rightExact
			case inspector.OpNq:
				*result = x.D != rightExact
			case inspector.OpGt:
				*result = x.D > rightExact
			case inspector.OpGtq:
				*result = x.D >= rightExact
			case inspector.OpLt:
				*result = x.D < rightExact
			case inspector.OpLtq:
				*result = x.D <= rightExact
			}
			return
		}
	}
	return
}

func (i11 TestStructInspector) Loop(src any, l inspector.Iterator, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestStruct
	_ = x
	if p, ok := src.(**testobj.TestStruct); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStruct); ok {
		x = p
	} else if v, ok := src.(testobj.TestStruct); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
	}
	return
}

func (i11 TestStructInspector) SetWithBuffer(dst, value any, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestStruct
	_ = x
	if p, ok := dst.(**testobj.TestStruct); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestStruct); ok {
		x = p
	} else if v, ok := dst.(testobj.TestStruct); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		if path[0] == "A" {
			inspector.AssignBuf(&x.A, value, buf)
			return nil
		}
		if path[0] == "S" {
			inspector.AssignBuf(&x.S, value, buf)
			return nil
		}
		if path[0] == "B" {
			inspector.AssignBuf(&x.B, value, buf)
			return nil
		}
		if path[0] == "I" {
			inspector.AssignBuf(&x.I, value, buf)
			return nil
		}
		if path[0] == "I8" {
			inspector.AssignBuf(&x.I8, value, buf)
			return nil
		}
		if path[0] == "I16" {
			inspector.AssignBuf(&x.I16, value, buf)
			return nil
		}
		if path[0] == "I32" {
			inspector.AssignBuf(&x.I32, value, buf)
			return nil
		}
		if path[0] == "I64" {
			inspector.AssignBuf(&x.I64, value, buf)
			return nil
		}
		if path[0] == "U" {
			inspector.AssignBuf(&x.U, value, buf)
			return nil
		}
		if path[0] == "U8" {
			inspector.AssignBuf(&x.U8, value, buf)
			return nil
		}
		if path[0] == "U16" {
			inspector.AssignBuf(&x.U16, value, buf)
			return nil
		}
		if path[0] == "U32" {
			inspector.AssignBuf(&x.U32, value, buf)
			return nil
		}
		if path[0] == "U64" {
			inspector.AssignBuf(&x.U64, value, buf)
			return nil
		}
		if path[0] == "F" {
			inspector.AssignBuf(&x.F, value, buf)
			return nil
		}
		if path[0] == "D" {
			inspector.AssignBuf(&x.D, value, buf)
			return nil
		}
	}
	return nil
}

func (i11 TestStructInspector) Set(dst, value any, path ...string) error {
	return i11.SetWithBuffer(dst, value, nil, path...)
}

func (i11 TestStructInspector) DeepEqual(l, r any) bool {
	return i11.DeepEqualWithOptions(l, r, nil)
}

func (i11 TestStructInspector) DeepEqualWithOptions(l, r any, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestStruct
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestStruct); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestStruct); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestStruct); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestStruct); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestStruct); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestStruct); ok {
		rx, req = &rp, true
	}
	if !leq || !req {
		return false
	}
	if lx == nil && rx == nil {
		return true
	}
	if (lx == nil && rx != nil) || (lx != nil && rx == nil) {
		return false
	}

	if lx.A != rx.A && inspector.DEQMustCheck("A", opts) {
		return false
	}
	if lx.S != rx.S && inspector.DEQMustCheck("S", opts) {
		return false
	}
	if !bytes.Equal(lx.B, rx.B) && inspector.DEQMustCheck("B", opts) {
		return false
	}
	if lx.I != rx.I && inspector.DEQMustCheck("I", opts) {
		return false
	}
	if lx.I8 != rx.I8 && inspector.DEQMustCheck("I8", opts) {
		return false
	}
	if lx.I16 != rx.I16 && inspector.DEQMustCheck("I16", opts) {
		return false
	}
	if lx.I32 != rx.I32 && inspector.DEQMustCheck("I32", opts) {
		return false
	}
	if lx.I64 != rx.I64 && inspector.DEQMustCheck("I64", opts) {
		return false
	}
	if lx.U != rx.U && inspector.DEQMustCheck("U", opts) {
		return false
	}
	if lx.U8 != rx.U8 && inspector.DEQMustCheck("U8", opts) {
		return false
	}
	if lx.U16 != rx.U16 && inspector.DEQMustCheck("U16", opts) {
		return false
	}
	if lx.U32 != rx.U32 && inspector.DEQMustCheck("U32", opts) {
		return false
	}
	if lx.U64 != rx.U64 && inspector.DEQMustCheck("U64", opts) {
		return false
	}
	if !inspector.EqualFloat32(lx.F, rx.F, opts) && inspector.DEQMustCheck("F", opts) {
		return false
	}
	if !inspector.EqualFloat64(lx.D, rx.D, opts) && inspector.DEQMustCheck("D", opts) {
		return false
	}
	return true
}

func (i11 TestStructInspector) Unmarshal(p []byte, typ inspector.Encoding) (any, error) {
	var x testobj.TestStruct
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i11 TestStructInspector) Copy(x any) (any, error) {
	var r testobj.TestStruct
	switch x.(type) {
	case testobj.TestStruct:
		r = x.(testobj.TestStruct)
	case *testobj.TestStruct:
		r = *x.(*testobj.TestStruct)
	case **testobj.TestStruct:
		r = **x.(**testobj.TestStruct)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i11.countBytes(&r)
	var l testobj.TestStruct
	err := i11.CopyTo(&r, &l, inspector.NewByteBuffer(bc))
	return &l, err
}

func (i11 TestStructInspector) CopyTo(src, dst any, buf inspector.AccumulativeBuffer) error {
	var r testobj.TestStruct
	switch src.(type) {
	case testobj.TestStruct:
		r = src.(testobj.TestStruct)
	case *testobj.TestStruct:
		r = *src.(*testobj.TestStruct)
	case **testobj.TestStruct:
		r = **src.(**testobj.TestStruct)
	default:
		return inspector.ErrUnsupportedType
	}
	var l *testobj.TestStruct
	switch dst.(type) {
	case testobj.TestStruct:
		return inspector.ErrMustPointerType
	case *testobj.TestStruct:
		l = dst.(*testobj.TestStruct)
	case **testobj.TestStruct:
		l = *dst.(**testobj.TestStruct)
	default:
		return inspector.ErrUnsupportedType
	}
	bb := buf.AcquireBytes()
	var err error
	if bb, err = i11.cpy(bb, l, &r); err != nil {
		return err
	}
	buf.ReleaseBytes(bb)
	return nil
}

func (i11 TestStructInspector) countBytes(x *testobj.TestStruct) (c int) {
	c += len(x.S)
	c += len(x.B)
	return c
}

func (i11 TestStructInspector) cpy(buf []byte, l, r *testobj.TestStruct) ([]byte, error) {
	l.A = r.A
	buf, l.S = inspector.BufferizeString(buf, r.S)
	buf, l.B = inspector.Bufferize(buf, r.B)
	l.I = r.I
	l.I8 = r.I8
	l.I16 = r.I16
	l.I32 = r.I32
	l.I64 = r.I64
	l.U = r.U
	l.U8 = r.U8
	l.U16 = r.U16
	l.U32 = r.U32
	l.U64 = r.U64
	l.F = r.F
	l.D = r.D
	return buf, nil
}

func (i11 TestStructInspector) Length(src any, result *int, path ...string) error {
	if src == nil {
		return nil
	}
	var x *testobj.TestStruct
	_ = x
	if p, ok := src.(**testobj.TestStruct); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStruct); ok {
		x = p
	} else if v, ok := src.(testobj.TestStruct); ok {
		x = &v
	} else {
		return inspector.ErrUnsupportedType
	}

	*result = 0
	if len(path) == 0 {
		return nil
	}
	if path[0] == "S" {
		*result = len(x.S)
		return nil
	}
	if path[0] == "B" {
		*result = len(x.B)
		return nil
	}
	return nil
}

func (i11 TestStructInspector) Capacity(src any, result *int, path ...string) error {
	if src == nil {
		return nil
	}
	var x *testobj.TestStruct
	_ = x
	if p, ok := src.(**testobj.TestStruct); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStruct); ok {
		x = p
	} else if v, ok := src.(testobj.TestStruct); ok {
		x = &v
	} else {
		return inspector.ErrUnsupportedType
	}

	*result = 0
	if len(path) == 0 {
		return nil
	}
	if path[0] == "S" {
	}
	if path[0] == "B" {
		*result = cap(x.B)
		return nil
	}
	return nil
}

func (i11 TestStructInspector) Reset(x any) error {
	var origin *testobj.TestStruct
	_ = origin
	switch x.(type) {
	case testobj.TestStruct:
		return inspector.ErrMustPointerType
	case *testobj.TestStruct:
		origin = x.(*testobj.TestStruct)
	case **testobj.TestStruct:
		origin = *x.(**testobj.TestStruct)
	default:
		return inspector.ErrUnsupportedType
	}
	origin.A = 0
	origin.S = ""
	if l := len((origin.B)); l > 0 {
		(origin.B) = (origin.B)[:0]
	}
	origin.I = 0
	origin.I8 = 0
	origin.I16 = 0
	origin.I32 = 0
	origin.I64 = 0
	origin.U = 0
	origin.U8 = 0
	origin.U16 = 0
	origin.U32 = 0
	origin.U64 = 0
	origin.F = 0
	origin.D = 0
	return nil
}
