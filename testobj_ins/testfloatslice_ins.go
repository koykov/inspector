// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"encoding/json"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

func init() {
	inspector.RegisterInspector("TestFloatSlice", TestFloatSliceInspector{})
}

type TestFloatSliceInspector struct {
	inspector.BaseInspector
}

func (i3 TestFloatSliceInspector) TypeName() string {
	return "TestFloatSlice"
}

func (i3 TestFloatSliceInspector) Get(src any, path ...string) (any, error) {
	var buf any
	err := i3.GetTo(src, &buf, path...)
	return buf, err
}

func (i3 TestFloatSliceInspector) GetTo(src any, buf *any, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestFloatSlice
	_ = x
	if p, ok := src.(**testobj.TestFloatSlice); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFloatSlice); ok {
		x = p
	} else if v, ok := src.(testobj.TestFloatSlice); ok {
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
		t16, err16 := strconv.ParseInt(path[0], 0, 0)
		if err16 != nil {
			return err16
		}
		i = int(t16)
		if len(*x) > i {
			x0 := (*x)[i]
			_ = x0
			*buf = &x0
			return
		}
	}
	return
}

func (i3 TestFloatSliceInspector) Compare(src any, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestFloatSlice
	_ = x
	if p, ok := src.(**testobj.TestFloatSlice); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFloatSlice); ok {
		x = p
	} else if v, ok := src.(testobj.TestFloatSlice); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		var i int
		t17, err17 := strconv.ParseInt(path[0], 0, 0)
		if err17 != nil {
			return err17
		}
		i = int(t17)
		if len(*x) > i {
			x0 := (*x)[i]
			_ = x0
			var rightExact float32
			t18, err18 := strconv.ParseFloat(right, 0)
			if err18 != nil {
				return err18
			}
			rightExact = float32(t18)
			switch cond {
			case inspector.OpEq:
				*result = x0 == rightExact
			case inspector.OpNq:
				*result = x0 != rightExact
			case inspector.OpGt:
				*result = x0 > rightExact
			case inspector.OpGtq:
				*result = x0 >= rightExact
			case inspector.OpLt:
				*result = x0 < rightExact
			case inspector.OpLtq:
				*result = x0 <= rightExact
			}
			return
		}
	}
	return
}

func (i3 TestFloatSliceInspector) Loop(src any, l inspector.Iterator, buf *[]byte, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestFloatSlice
	_ = x
	if p, ok := src.(**testobj.TestFloatSlice); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFloatSlice); ok {
		x = p
	} else if v, ok := src.(testobj.TestFloatSlice); ok {
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

func (i3 TestFloatSliceInspector) SetWithBuffer(dst, value any, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestFloatSlice
	_ = x
	if p, ok := dst.(**testobj.TestFloatSlice); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestFloatSlice); ok {
		x = p
	} else if v, ok := dst.(testobj.TestFloatSlice); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		var i int
		t19, err19 := strconv.ParseInt(path[0], 0, 0)
		if err19 != nil {
			return err19
		}
		i = int(t19)
		if len(*x) > i {
			x0 := (*x)[i]
			_ = x0
			inspector.AssignBuf(&x0, value, buf)
			return nil
			(*x)[i] = x0
			return nil
		}
	}
	return nil
}

func (i3 TestFloatSliceInspector) Set(dst, value any, path ...string) error {
	return i3.SetWithBuffer(dst, value, nil, path...)
}

func (i3 TestFloatSliceInspector) DeepEqual(l, r any) bool {
	return i3.DeepEqualWithOptions(l, r, nil)
}

func (i3 TestFloatSliceInspector) DeepEqualWithOptions(l, r any, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestFloatSlice
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestFloatSlice); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestFloatSlice); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestFloatSlice); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestFloatSlice); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestFloatSlice); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestFloatSlice); ok {
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
		if lx1 != rx1 {
			return false
		}
	}
	return true
}

func (i3 TestFloatSliceInspector) Unmarshal(p []byte, typ inspector.Encoding) (any, error) {
	var x testobj.TestFloatSlice
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i3 TestFloatSliceInspector) Copy(x any) (any, error) {
	var r testobj.TestFloatSlice
	switch x.(type) {
	case testobj.TestFloatSlice:
		r = x.(testobj.TestFloatSlice)
	case *testobj.TestFloatSlice:
		r = *x.(*testobj.TestFloatSlice)
	case **testobj.TestFloatSlice:
		r = **x.(**testobj.TestFloatSlice)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i3.countBytes(&r)
	var l testobj.TestFloatSlice
	err := i3.CopyTo(&r, &l, inspector.NewByteBuffer(bc))
	return &l, err
}

func (i3 TestFloatSliceInspector) CopyTo(src, dst any, buf inspector.AccumulativeBuffer) error {
	var r testobj.TestFloatSlice
	switch src.(type) {
	case testobj.TestFloatSlice:
		r = src.(testobj.TestFloatSlice)
	case *testobj.TestFloatSlice:
		r = *src.(*testobj.TestFloatSlice)
	case **testobj.TestFloatSlice:
		r = **src.(**testobj.TestFloatSlice)
	default:
		return inspector.ErrUnsupportedType
	}
	var l *testobj.TestFloatSlice
	switch dst.(type) {
	case testobj.TestFloatSlice:
		return inspector.ErrMustPointerType
	case *testobj.TestFloatSlice:
		l = dst.(*testobj.TestFloatSlice)
	case **testobj.TestFloatSlice:
		l = *dst.(**testobj.TestFloatSlice)
	default:
		return inspector.ErrUnsupportedType
	}
	bb := buf.AcquireBytes()
	var err error
	if bb, err = i3.cpy(bb, l, &r); err != nil {
		return err
	}
	buf.ReleaseBytes(bb)
	return nil
}

func (i3 TestFloatSliceInspector) countBytes(x *testobj.TestFloatSlice) (c int) {
	return c
}

func (i3 TestFloatSliceInspector) cpy(buf []byte, l, r *testobj.TestFloatSlice) ([]byte, error) {
	if len(*r) > 0 {
		buf0 := (*l)
		if buf0 == nil {
			buf0 = make(testobj.TestFloatSlice, 0, len(*r))
		}
		for i0 := 0; i0 < len(*r); i0++ {
			var b0 float32
			x0 := (*r)[i0]
			b0 = x0
			buf0 = append(buf0, b0)
		}
		l = &buf0
	}
	return buf, nil
}

func (i3 TestFloatSliceInspector) Length(src any, result *int, path ...string) error {
	if src == nil {
		return nil
	}
	var x *testobj.TestFloatSlice
	_ = x
	if p, ok := src.(**testobj.TestFloatSlice); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFloatSlice); ok {
		x = p
	} else if v, ok := src.(testobj.TestFloatSlice); ok {
		x = &v
	} else {
		return inspector.ErrUnsupportedType
	}

	*result = 0
	if len(path) == 0 {
		return nil
	}
	return nil
}

func (i3 TestFloatSliceInspector) Capacity(src any, result *int, path ...string) error {
	if src == nil {
		return nil
	}
	var x *testobj.TestFloatSlice
	_ = x
	if p, ok := src.(**testobj.TestFloatSlice); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFloatSlice); ok {
		x = p
	} else if v, ok := src.(testobj.TestFloatSlice); ok {
		x = &v
	} else {
		return inspector.ErrUnsupportedType
	}

	*result = 0
	if len(path) == 0 {
		return nil
	}
	return nil
}

func (i3 TestFloatSliceInspector) Reset(x any) error {
	var origin *testobj.TestFloatSlice
	_ = origin
	switch x.(type) {
	case testobj.TestFloatSlice:
		return inspector.ErrMustPointerType
	case *testobj.TestFloatSlice:
		origin = x.(*testobj.TestFloatSlice)
	case **testobj.TestFloatSlice:
		origin = *x.(**testobj.TestFloatSlice)
	default:
		return inspector.ErrUnsupportedType
	}
	if l := len((*origin)); l > 0 {
		(*origin) = (*origin)[:0]
	}
	return nil
}