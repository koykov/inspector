// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"bytes"
	"encoding/json"
	"github.com/koykov/fastconv"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

type TestStructInspector struct {
	inspector.BaseInspector
}

func (i11 TestStructInspector) TypeName() string {
	return "TestStruct"
}

func (i11 TestStructInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i11.GetTo(src, &buf, path...)
	return buf, err
}

func (i11 TestStructInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
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

func (i11 TestStructInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
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
			t320 := fastconv.S2B(right)
			if len(t320) > 0 {
				rightExact = t320[0]
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
			rightExact = fastconv.S2B(right)

			if cond == inspector.OpEq {
				*result = bytes.Equal(x.B, rightExact)
			} else {
				*result = !bytes.Equal(x.B, rightExact)
			}
			return
		}
		if path[0] == "I" {
			var rightExact int
			t323, err323 := strconv.ParseInt(right, 0, 0)
			if err323 != nil {
				return err323
			}
			rightExact = int(t323)
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
			t324, err324 := strconv.ParseInt(right, 0, 0)
			if err324 != nil {
				return err324
			}
			rightExact = int8(t324)
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
			t325, err325 := strconv.ParseInt(right, 0, 0)
			if err325 != nil {
				return err325
			}
			rightExact = int16(t325)
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
			t326, err326 := strconv.ParseInt(right, 0, 0)
			if err326 != nil {
				return err326
			}
			rightExact = int32(t326)
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
			t327, err327 := strconv.ParseInt(right, 0, 0)
			if err327 != nil {
				return err327
			}
			rightExact = int64(t327)
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
			t328, err328 := strconv.ParseUint(right, 0, 0)
			if err328 != nil {
				return err328
			}
			rightExact = uint(t328)
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
			t329, err329 := strconv.ParseUint(right, 0, 0)
			if err329 != nil {
				return err329
			}
			rightExact = uint8(t329)
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
			t330, err330 := strconv.ParseUint(right, 0, 0)
			if err330 != nil {
				return err330
			}
			rightExact = uint16(t330)
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
			t331, err331 := strconv.ParseUint(right, 0, 0)
			if err331 != nil {
				return err331
			}
			rightExact = uint32(t331)
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
			t332, err332 := strconv.ParseUint(right, 0, 0)
			if err332 != nil {
				return err332
			}
			rightExact = uint64(t332)
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
			t333, err333 := strconv.ParseFloat(right, 0)
			if err333 != nil {
				return err333
			}
			rightExact = float32(t333)
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
			t334, err334 := strconv.ParseFloat(right, 0)
			if err334 != nil {
				return err334
			}
			rightExact = float64(t334)
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

func (i11 TestStructInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
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

func (i11 TestStructInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
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

func (i11 TestStructInspector) Set(dst, value interface{}, path ...string) error {
	return i11.SetWB(dst, value, nil, path...)
}

func (i11 TestStructInspector) DeepEqual(l, r interface{}) bool {
	return i11.DeepEqualWithOptions(l, r, nil)
}

func (i11 TestStructInspector) DeepEqualWithOptions(l, r interface{}, opts *inspector.DEQOptions) bool {
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

func (i11 TestStructInspector) Unmarshal(p []byte, typ inspector.Encoding) (interface{}, error) {
	var x testobj.TestStruct
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i11 TestStructInspector) Copy(x interface{}) (interface{}, error) {
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

func (i11 TestStructInspector) CopyTo(src, dst interface{}, buf inspector.AccumulativeBuffer) error {
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

func (i11 TestStructInspector) Reset(x interface{}) error {
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
