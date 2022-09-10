// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"encoding/json"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

type TestFloatPtrSliceInspector struct {
	inspector.BaseInspector
}

func (i2 TestFloatPtrSliceInspector) TypeName() string {
	return "TestFloatPtrSlice"
}

func (i2 TestFloatPtrSliceInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i2.GetTo(src, &buf, path...)
	return buf, err
}

func (i2 TestFloatPtrSliceInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestFloatPtrSlice
	_ = x
	if p, ok := src.(**testobj.TestFloatPtrSlice); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFloatPtrSlice); ok {
		x = p
	} else if v, ok := src.(testobj.TestFloatPtrSlice); ok {
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
		t11, err11 := strconv.ParseInt(path[0], 0, 0)
		if err11 != nil {
			return err11
		}
		i = int(t11)
		if len(*x) > i {
			x0 := (*x)[i]
			_ = x0
			if x0 == nil {
				return
			}
			*buf = &x0
			return
		}
	}
	return
}

func (i2 TestFloatPtrSliceInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestFloatPtrSlice
	_ = x
	if p, ok := src.(**testobj.TestFloatPtrSlice); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFloatPtrSlice); ok {
		x = p
	} else if v, ok := src.(testobj.TestFloatPtrSlice); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		var i int
		t12, err12 := strconv.ParseInt(path[0], 0, 0)
		if err12 != nil {
			return err12
		}
		i = int(t12)
		if len(*x) > i {
			x0 := (*x)[i]
			_ = x0
			if x0 == nil {
				return
			}
			if right == inspector.Nil {
				if cond == inspector.OpEq {
					*result = x0 == nil
				} else {
					*result = x0 != nil
				}
				return
			}
			return
		}
	}
	return
}

func (i2 TestFloatPtrSliceInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestFloatPtrSlice
	_ = x
	if p, ok := src.(**testobj.TestFloatPtrSlice); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFloatPtrSlice); ok {
		x = p
	} else if v, ok := src.(testobj.TestFloatPtrSlice); ok {
		x = &v
	} else {
		return
	}

	for k := range *x {
		if l.RequireKey() {
			*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)
			l.SetKey(buf, &inspector.StaticInspector{})
		}
		l.SetVal(&(*x)[k], &inspector.StaticInspector{})
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

func (i2 TestFloatPtrSliceInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestFloatPtrSlice
	_ = x
	if p, ok := dst.(**testobj.TestFloatPtrSlice); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestFloatPtrSlice); ok {
		x = p
	} else if v, ok := dst.(testobj.TestFloatPtrSlice); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		var i int
		t13, err13 := strconv.ParseInt(path[0], 0, 0)
		if err13 != nil {
			return err13
		}
		i = int(t13)
		if len(*x) > i {
			x0 := (*x)[i]
			_ = x0
			if x0 == nil {
				return nil
			}
			inspector.AssignBuf(x0, value, buf)
			return nil
			(*x)[i] = x0
			return nil
		}
	}
	return nil
}

func (i2 TestFloatPtrSliceInspector) Set(dst, value interface{}, path ...string) error {
	return i2.SetWB(dst, value, nil, path...)
}

func (i2 TestFloatPtrSliceInspector) DeepEqual(l, r interface{}) bool {
	return i2.DeepEqualWithOptions(l, r, nil)
}

func (i2 TestFloatPtrSliceInspector) DeepEqualWithOptions(l, r interface{}, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestFloatPtrSlice
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestFloatPtrSlice); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestFloatPtrSlice); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestFloatPtrSlice); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestFloatPtrSlice); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestFloatPtrSlice); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestFloatPtrSlice); ok {
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
			if *lx1 != *rx1 {
				return false
			}
		}
	}
	return true
}

func (i2 TestFloatPtrSliceInspector) Unmarshal(p []byte, typ inspector.Encoding) (interface{}, error) {
	var x testobj.TestFloatPtrSlice
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i2 TestFloatPtrSliceInspector) Copy(x interface{}) (interface{}, error) {
	var origin, cpy testobj.TestFloatPtrSlice
	switch x.(type) {
	case testobj.TestFloatPtrSlice:
		origin = x.(testobj.TestFloatPtrSlice)
	case *testobj.TestFloatPtrSlice:
		origin = *x.(*testobj.TestFloatPtrSlice)
	case **testobj.TestFloatPtrSlice:
		origin = **x.(**testobj.TestFloatPtrSlice)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i2.calcBytes(&origin)
	buf1 := make([]byte, 0, bc)
	if err := i2.cpy(buf1, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i2 TestFloatPtrSliceInspector) CopyWB(x interface{}, buf inspector.AccumulativeBuffer) (interface{}, error) {
	var origin, cpy testobj.TestFloatPtrSlice
	switch x.(type) {
	case testobj.TestFloatPtrSlice:
		origin = x.(testobj.TestFloatPtrSlice)
	case *testobj.TestFloatPtrSlice:
		origin = *x.(*testobj.TestFloatPtrSlice)
	case **testobj.TestFloatPtrSlice:
		origin = **x.(**testobj.TestFloatPtrSlice)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	buf1 := buf.AcquireBytes()
	defer buf.ReleaseBytes(buf1)
	if err := i2.cpy(buf1, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i2 TestFloatPtrSliceInspector) calcBytes(x *testobj.TestFloatPtrSlice) (c int) {
	return c
}

func (i2 TestFloatPtrSliceInspector) cpy(buf []byte, l, r *testobj.TestFloatPtrSlice) error {
	if len(*r) > 0 {
		buf0 := make(testobj.TestFloatPtrSlice, 0, len(*r))
		for i0 := 0; i0 < len(*r); i0++ {
			var b0 *float32
			x0 := (*l)[i0]
			b0 = x0
			buf0 = append(buf0, b0)
		}
		l = &buf0
	}
	return nil
}
