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

type TestStructSliceLiteralInspector struct {
	inspector.BaseInspector
}

func (i12 TestStructSliceLiteralInspector) TypeName() string {
	return "TestStructSliceLiteral"
}

func (i12 TestStructSliceLiteralInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i12.GetTo(src, &buf, path...)
	return buf, err
}

func (i12 TestStructSliceLiteralInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestStructSliceLiteral
	_ = x
	if p, ok := src.(**testobj.TestStructSliceLiteral); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStructSliceLiteral); ok {
		x = p
	} else if v, ok := src.(testobj.TestStructSliceLiteral); ok {
		x = &v
	} else {
		return
	}
	if len(path) == 0 {
		*buf = &(*x)
		return
	}

	if len(path) > 0 {
		var i int
		t305, err305 := strconv.ParseInt(path[0], 0, 0)
		if err305 != nil {
			return err305
		}
		i = int(t305)
		if len(*x) > i {
			x0 := (*x)[i]
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				if path[1] == "A" {
					*buf = &x0.A
					return
				}
				if path[1] == "S" {
					*buf = &x0.S
					return
				}
				if path[1] == "B" {
					*buf = &x0.B
					return
				}
				if path[1] == "I" {
					*buf = &x0.I
					return
				}
				if path[1] == "I8" {
					*buf = &x0.I8
					return
				}
				if path[1] == "I16" {
					*buf = &x0.I16
					return
				}
				if path[1] == "I32" {
					*buf = &x0.I32
					return
				}
				if path[1] == "I64" {
					*buf = &x0.I64
					return
				}
				if path[1] == "U" {
					*buf = &x0.U
					return
				}
				if path[1] == "U8" {
					*buf = &x0.U8
					return
				}
				if path[1] == "U16" {
					*buf = &x0.U16
					return
				}
				if path[1] == "U32" {
					*buf = &x0.U32
					return
				}
				if path[1] == "U64" {
					*buf = &x0.U64
					return
				}
				if path[1] == "F" {
					*buf = &x0.F
					return
				}
				if path[1] == "D" {
					*buf = &x0.D
					return
				}
			}
			*buf = x0
		}
	}
	return
}

func (i12 TestStructSliceLiteralInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestStructSliceLiteral
	_ = x
	if p, ok := src.(**testobj.TestStructSliceLiteral); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStructSliceLiteral); ok {
		x = p
	} else if v, ok := src.(testobj.TestStructSliceLiteral); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		var i int
		t306, err306 := strconv.ParseInt(path[0], 0, 0)
		if err306 != nil {
			return err306
		}
		i = int(t306)
		if len(*x) > i {
			x0 := (*x)[i]
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				if path[1] == "A" {
					var rightExact byte
					t307 := fastconv.S2B(right)
					if len(t307) > 0 {
						rightExact = t307[0]
					}

					switch cond {
					case inspector.OpEq:
						*result = x0.A == rightExact
					case inspector.OpNq:
						*result = x0.A != rightExact
					case inspector.OpGt:
						*result = x0.A > rightExact
					case inspector.OpGtq:
						*result = x0.A >= rightExact
					case inspector.OpLt:
						*result = x0.A < rightExact
					case inspector.OpLtq:
						*result = x0.A <= rightExact
					}
					return
				}
				if path[1] == "S" {
					var rightExact string
					rightExact = right

					switch cond {
					case inspector.OpEq:
						*result = x0.S == rightExact
					case inspector.OpNq:
						*result = x0.S != rightExact
					case inspector.OpGt:
						*result = x0.S > rightExact
					case inspector.OpGtq:
						*result = x0.S >= rightExact
					case inspector.OpLt:
						*result = x0.S < rightExact
					case inspector.OpLtq:
						*result = x0.S <= rightExact
					}
					return
				}
				if path[1] == "B" {
					var rightExact []byte
					rightExact = fastconv.S2B(right)

					if cond == inspector.OpEq {
						*result = bytes.Equal(x0.B, rightExact)
					} else {
						*result = !bytes.Equal(x0.B, rightExact)
					}
					return
				}
				if path[1] == "I" {
					var rightExact int
					t310, err310 := strconv.ParseInt(right, 0, 0)
					if err310 != nil {
						return err310
					}
					rightExact = int(t310)
					switch cond {
					case inspector.OpEq:
						*result = x0.I == rightExact
					case inspector.OpNq:
						*result = x0.I != rightExact
					case inspector.OpGt:
						*result = x0.I > rightExact
					case inspector.OpGtq:
						*result = x0.I >= rightExact
					case inspector.OpLt:
						*result = x0.I < rightExact
					case inspector.OpLtq:
						*result = x0.I <= rightExact
					}
					return
				}
				if path[1] == "I8" {
					var rightExact int8
					t311, err311 := strconv.ParseInt(right, 0, 0)
					if err311 != nil {
						return err311
					}
					rightExact = int8(t311)
					switch cond {
					case inspector.OpEq:
						*result = x0.I8 == rightExact
					case inspector.OpNq:
						*result = x0.I8 != rightExact
					case inspector.OpGt:
						*result = x0.I8 > rightExact
					case inspector.OpGtq:
						*result = x0.I8 >= rightExact
					case inspector.OpLt:
						*result = x0.I8 < rightExact
					case inspector.OpLtq:
						*result = x0.I8 <= rightExact
					}
					return
				}
				if path[1] == "I16" {
					var rightExact int16
					t312, err312 := strconv.ParseInt(right, 0, 0)
					if err312 != nil {
						return err312
					}
					rightExact = int16(t312)
					switch cond {
					case inspector.OpEq:
						*result = x0.I16 == rightExact
					case inspector.OpNq:
						*result = x0.I16 != rightExact
					case inspector.OpGt:
						*result = x0.I16 > rightExact
					case inspector.OpGtq:
						*result = x0.I16 >= rightExact
					case inspector.OpLt:
						*result = x0.I16 < rightExact
					case inspector.OpLtq:
						*result = x0.I16 <= rightExact
					}
					return
				}
				if path[1] == "I32" {
					var rightExact int32
					t313, err313 := strconv.ParseInt(right, 0, 0)
					if err313 != nil {
						return err313
					}
					rightExact = int32(t313)
					switch cond {
					case inspector.OpEq:
						*result = x0.I32 == rightExact
					case inspector.OpNq:
						*result = x0.I32 != rightExact
					case inspector.OpGt:
						*result = x0.I32 > rightExact
					case inspector.OpGtq:
						*result = x0.I32 >= rightExact
					case inspector.OpLt:
						*result = x0.I32 < rightExact
					case inspector.OpLtq:
						*result = x0.I32 <= rightExact
					}
					return
				}
				if path[1] == "I64" {
					var rightExact int64
					t314, err314 := strconv.ParseInt(right, 0, 0)
					if err314 != nil {
						return err314
					}
					rightExact = int64(t314)
					switch cond {
					case inspector.OpEq:
						*result = x0.I64 == rightExact
					case inspector.OpNq:
						*result = x0.I64 != rightExact
					case inspector.OpGt:
						*result = x0.I64 > rightExact
					case inspector.OpGtq:
						*result = x0.I64 >= rightExact
					case inspector.OpLt:
						*result = x0.I64 < rightExact
					case inspector.OpLtq:
						*result = x0.I64 <= rightExact
					}
					return
				}
				if path[1] == "U" {
					var rightExact uint
					t315, err315 := strconv.ParseUint(right, 0, 0)
					if err315 != nil {
						return err315
					}
					rightExact = uint(t315)
					switch cond {
					case inspector.OpEq:
						*result = x0.U == rightExact
					case inspector.OpNq:
						*result = x0.U != rightExact
					case inspector.OpGt:
						*result = x0.U > rightExact
					case inspector.OpGtq:
						*result = x0.U >= rightExact
					case inspector.OpLt:
						*result = x0.U < rightExact
					case inspector.OpLtq:
						*result = x0.U <= rightExact
					}
					return
				}
				if path[1] == "U8" {
					var rightExact uint8
					t316, err316 := strconv.ParseUint(right, 0, 0)
					if err316 != nil {
						return err316
					}
					rightExact = uint8(t316)
					switch cond {
					case inspector.OpEq:
						*result = x0.U8 == rightExact
					case inspector.OpNq:
						*result = x0.U8 != rightExact
					case inspector.OpGt:
						*result = x0.U8 > rightExact
					case inspector.OpGtq:
						*result = x0.U8 >= rightExact
					case inspector.OpLt:
						*result = x0.U8 < rightExact
					case inspector.OpLtq:
						*result = x0.U8 <= rightExact
					}
					return
				}
				if path[1] == "U16" {
					var rightExact uint16
					t317, err317 := strconv.ParseUint(right, 0, 0)
					if err317 != nil {
						return err317
					}
					rightExact = uint16(t317)
					switch cond {
					case inspector.OpEq:
						*result = x0.U16 == rightExact
					case inspector.OpNq:
						*result = x0.U16 != rightExact
					case inspector.OpGt:
						*result = x0.U16 > rightExact
					case inspector.OpGtq:
						*result = x0.U16 >= rightExact
					case inspector.OpLt:
						*result = x0.U16 < rightExact
					case inspector.OpLtq:
						*result = x0.U16 <= rightExact
					}
					return
				}
				if path[1] == "U32" {
					var rightExact uint32
					t318, err318 := strconv.ParseUint(right, 0, 0)
					if err318 != nil {
						return err318
					}
					rightExact = uint32(t318)
					switch cond {
					case inspector.OpEq:
						*result = x0.U32 == rightExact
					case inspector.OpNq:
						*result = x0.U32 != rightExact
					case inspector.OpGt:
						*result = x0.U32 > rightExact
					case inspector.OpGtq:
						*result = x0.U32 >= rightExact
					case inspector.OpLt:
						*result = x0.U32 < rightExact
					case inspector.OpLtq:
						*result = x0.U32 <= rightExact
					}
					return
				}
				if path[1] == "U64" {
					var rightExact uint64
					t319, err319 := strconv.ParseUint(right, 0, 0)
					if err319 != nil {
						return err319
					}
					rightExact = uint64(t319)
					switch cond {
					case inspector.OpEq:
						*result = x0.U64 == rightExact
					case inspector.OpNq:
						*result = x0.U64 != rightExact
					case inspector.OpGt:
						*result = x0.U64 > rightExact
					case inspector.OpGtq:
						*result = x0.U64 >= rightExact
					case inspector.OpLt:
						*result = x0.U64 < rightExact
					case inspector.OpLtq:
						*result = x0.U64 <= rightExact
					}
					return
				}
				if path[1] == "F" {
					var rightExact float32
					t320, err320 := strconv.ParseFloat(right, 0)
					if err320 != nil {
						return err320
					}
					rightExact = float32(t320)
					switch cond {
					case inspector.OpEq:
						*result = x0.F == rightExact
					case inspector.OpNq:
						*result = x0.F != rightExact
					case inspector.OpGt:
						*result = x0.F > rightExact
					case inspector.OpGtq:
						*result = x0.F >= rightExact
					case inspector.OpLt:
						*result = x0.F < rightExact
					case inspector.OpLtq:
						*result = x0.F <= rightExact
					}
					return
				}
				if path[1] == "D" {
					var rightExact float64
					t321, err321 := strconv.ParseFloat(right, 0)
					if err321 != nil {
						return err321
					}
					rightExact = float64(t321)
					switch cond {
					case inspector.OpEq:
						*result = x0.D == rightExact
					case inspector.OpNq:
						*result = x0.D != rightExact
					case inspector.OpGt:
						*result = x0.D > rightExact
					case inspector.OpGtq:
						*result = x0.D >= rightExact
					case inspector.OpLt:
						*result = x0.D < rightExact
					case inspector.OpLtq:
						*result = x0.D <= rightExact
					}
					return
				}
			}
		}
	}
	return
}

func (i12 TestStructSliceLiteralInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestStructSliceLiteral
	_ = x
	if p, ok := src.(**testobj.TestStructSliceLiteral); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStructSliceLiteral); ok {
		x = p
	} else if v, ok := src.(testobj.TestStructSliceLiteral); ok {
		x = &v
	} else {
		return
	}

	for k := range *x {
		if l.RequireKey() {
			*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)
			l.SetKey(buf, &inspector.StaticInspector{})
		}
		l.SetVal(&(*x)[k], &TestStructInspector{})
		ctl := l.Iterate()
		if ctl == inspector.LoopCtlBrk {
			break
		}
		if ctl == inspector.LoopCtlCnt {
			continue
		}
	}
	return
	return
}

func (i12 TestStructSliceLiteralInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestStructSliceLiteral
	_ = x
	if p, ok := dst.(**testobj.TestStructSliceLiteral); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestStructSliceLiteral); ok {
		x = p
	} else if v, ok := dst.(testobj.TestStructSliceLiteral); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		var i int
		t322, err322 := strconv.ParseInt(path[0], 0, 0)
		if err322 != nil {
			return err322
		}
		i = int(t322)
		if len(*x) > i {
			x0 := (*x)[i]
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return nil
				}
				if path[1] == "A" {
					inspector.AssignBuf(&x0.A, value, buf)
					return nil
				}
				if path[1] == "S" {
					inspector.AssignBuf(&x0.S, value, buf)
					return nil
				}
				if path[1] == "B" {
					inspector.AssignBuf(&x0.B, value, buf)
					return nil
				}
				if path[1] == "I" {
					inspector.AssignBuf(&x0.I, value, buf)
					return nil
				}
				if path[1] == "I8" {
					inspector.AssignBuf(&x0.I8, value, buf)
					return nil
				}
				if path[1] == "I16" {
					inspector.AssignBuf(&x0.I16, value, buf)
					return nil
				}
				if path[1] == "I32" {
					inspector.AssignBuf(&x0.I32, value, buf)
					return nil
				}
				if path[1] == "I64" {
					inspector.AssignBuf(&x0.I64, value, buf)
					return nil
				}
				if path[1] == "U" {
					inspector.AssignBuf(&x0.U, value, buf)
					return nil
				}
				if path[1] == "U8" {
					inspector.AssignBuf(&x0.U8, value, buf)
					return nil
				}
				if path[1] == "U16" {
					inspector.AssignBuf(&x0.U16, value, buf)
					return nil
				}
				if path[1] == "U32" {
					inspector.AssignBuf(&x0.U32, value, buf)
					return nil
				}
				if path[1] == "U64" {
					inspector.AssignBuf(&x0.U64, value, buf)
					return nil
				}
				if path[1] == "F" {
					inspector.AssignBuf(&x0.F, value, buf)
					return nil
				}
				if path[1] == "D" {
					inspector.AssignBuf(&x0.D, value, buf)
					return nil
				}
			}
			(*x)[i] = x0
			return nil
		}
	}
	return nil
}

func (i12 TestStructSliceLiteralInspector) Set(dst, value interface{}, path ...string) error {
	return i12.SetWB(dst, value, nil, path...)
}

func (i12 TestStructSliceLiteralInspector) DeepEqual(l, r interface{}) bool {
	return i12.DeepEqualWithOptions(l, r, nil)
}

func (i12 TestStructSliceLiteralInspector) DeepEqualWithOptions(l, r interface{}, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestStructSliceLiteral
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestStructSliceLiteral); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestStructSliceLiteral); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestStructSliceLiteral); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestStructSliceLiteral); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestStructSliceLiteral); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestStructSliceLiteral); ok {
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

	if len(*lx) != len(*rx) {
		return false
	}
	for i := 0; i < len(*lx); i++ {
		lx1 := (*lx)[i]
		rx1 := (*rx)[i]
		_, _ = lx1, rx1
		if (lx1 == nil && rx1 != nil) || (lx1 != nil && rx1 == nil) {
			return false
		}
		if lx1 != nil && rx1 != nil {
			if lx1.A != rx1.A && inspector.DEQMustCheck("A", opts) {
				return false
			}
			if lx1.S != rx1.S && inspector.DEQMustCheck("S", opts) {
				return false
			}
			if !bytes.Equal(lx1.B, rx1.B) && inspector.DEQMustCheck("B", opts) {
				return false
			}
			if lx1.I != rx1.I && inspector.DEQMustCheck("I", opts) {
				return false
			}
			if lx1.I8 != rx1.I8 && inspector.DEQMustCheck("I8", opts) {
				return false
			}
			if lx1.I16 != rx1.I16 && inspector.DEQMustCheck("I16", opts) {
				return false
			}
			if lx1.I32 != rx1.I32 && inspector.DEQMustCheck("I32", opts) {
				return false
			}
			if lx1.I64 != rx1.I64 && inspector.DEQMustCheck("I64", opts) {
				return false
			}
			if lx1.U != rx1.U && inspector.DEQMustCheck("U", opts) {
				return false
			}
			if lx1.U8 != rx1.U8 && inspector.DEQMustCheck("U8", opts) {
				return false
			}
			if lx1.U16 != rx1.U16 && inspector.DEQMustCheck("U16", opts) {
				return false
			}
			if lx1.U32 != rx1.U32 && inspector.DEQMustCheck("U32", opts) {
				return false
			}
			if lx1.U64 != rx1.U64 && inspector.DEQMustCheck("U64", opts) {
				return false
			}
			if !inspector.EqualFloat32(lx1.F, rx1.F, opts) && inspector.DEQMustCheck("F", opts) {
				return false
			}
			if !inspector.EqualFloat64(lx1.D, rx1.D, opts) && inspector.DEQMustCheck("D", opts) {
				return false
			}
		}
	}
	return true
}

func (i12 TestStructSliceLiteralInspector) Unmarshal(p []byte, typ inspector.Encoding) (interface{}, error) {
	var x testobj.TestStructSliceLiteral
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i12 TestStructSliceLiteralInspector) Copy(x interface{}) (interface{}, error) {
	var origin, cpy testobj.TestStructSliceLiteral
	switch x.(type) {
	case testobj.TestStructSliceLiteral:
		origin = x.(testobj.TestStructSliceLiteral)
	case *testobj.TestStructSliceLiteral:
		origin = *x.(*testobj.TestStructSliceLiteral)
	case **testobj.TestStructSliceLiteral:
		origin = **x.(**testobj.TestStructSliceLiteral)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i12.calcBytes(&origin)
	buf1 := make([]byte, 0, bc)
	if err := i12.cpy(buf1, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i12 TestStructSliceLiteralInspector) CopyWB(x interface{}, buf inspector.AccumulativeBuffer) (interface{}, error) {
	var origin, cpy testobj.TestStructSliceLiteral
	switch x.(type) {
	case testobj.TestStructSliceLiteral:
		origin = x.(testobj.TestStructSliceLiteral)
	case *testobj.TestStructSliceLiteral:
		origin = *x.(*testobj.TestStructSliceLiteral)
	case **testobj.TestStructSliceLiteral:
		origin = **x.(**testobj.TestStructSliceLiteral)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	buf1 := buf.AcquireBytes()
	defer buf.ReleaseBytes(buf1)
	if err := i12.cpy(buf1, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i12 TestStructSliceLiteralInspector) calcBytes(x *testobj.TestStructSliceLiteral) (c int) {
	for i0 := 0; i0 < len(*x); i0++ {
		x0 := (*x)[i0]
		c += len(x0.S)
		c += len(x0.B)
	}
	return c
}

func (i12 TestStructSliceLiteralInspector) cpy(buf []byte, l, r *testobj.TestStructSliceLiteral) error {
	if len(*r) > 0 {
		buf0 := make(testobj.TestStructSliceLiteral, 0, len(*r))
		for i0 := 0; i0 < len(*r); i0++ {
			var b0 testobj.TestStruct
			x0 := (*l)[i0]
			b0.A = x0.A
			buf, b0.S = inspector.BufferizeString(buf, x0.S)
			buf, b0.B = inspector.Bufferize(buf, x0.B)
			b0.I = x0.I
			b0.I8 = x0.I8
			b0.I16 = x0.I16
			b0.I32 = x0.I32
			b0.I64 = x0.I64
			b0.U = x0.U
			b0.U8 = x0.U8
			b0.U16 = x0.U16
			b0.U32 = x0.U32
			b0.U64 = x0.U64
			b0.F = x0.F
			b0.D = x0.D
			buf0 = append(buf0, &b0)
		}
		l = &buf0
	}
	return nil
}

func (i12 TestStructSliceLiteralInspector) Reset(x interface{}) {
	var origin testobj.TestStructSliceLiteral
	_ = origin
	switch x.(type) {
	case testobj.TestStructSliceLiteral:
		origin = x.(testobj.TestStructSliceLiteral)
	case *testobj.TestStructSliceLiteral:
		origin = *x.(*testobj.TestStructSliceLiteral)
	case **testobj.TestStructSliceLiteral:
		origin = **x.(**testobj.TestStructSliceLiteral)
	default:
		return
	}
	if l := len((origin)); l > 0 {
		_ = (origin)[l-1]
		for i := 0; i < l; i++ {
			x0 := (origin)[i]
			x0.A = 0
			x0.S = ""
			if l := len((x0.B)); l > 0 {
				(x0.B) = (x0.B)[:0]
			}
			x0.I = 0
			x0.I8 = 0
			x0.I16 = 0
			x0.I32 = 0
			x0.I64 = 0
			x0.U = 0
			x0.U8 = 0
			x0.U16 = 0
			x0.U32 = 0
			x0.U64 = 0
			x0.F = 0
			x0.D = 0
		}
		(origin) = (origin)[:0]
	}
}