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

func (i2 TestFloatPtrSliceInspector) Get(src any, path ...string) (any, error) {
	var buf any
	err := i2.GetTo(src, &buf, path...)
	return buf, err
}

func (i2 TestFloatPtrSliceInspector) GetTo(src any, buf *any, path ...string) (err error) {
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
		t13, err13 := strconv.ParseInt(path[0], 0, 0)
		if err13 != nil {
			return err13
		}
		i = int(t13)
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

func (i2 TestFloatPtrSliceInspector) Compare(src any, cond inspector.Op, right string, result *bool, path ...string) (err error) {
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
		t14, err14 := strconv.ParseInt(path[0], 0, 0)
		if err14 != nil {
			return err14
		}
		i = int(t14)
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

func (i2 TestFloatPtrSliceInspector) Loop(src any, l inspector.Iterator, buf *[]byte, path ...string) (err error) {
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

func (i2 TestFloatPtrSliceInspector) SetWithBuffer(dst, value any, buf inspector.AccumulativeBuffer, path ...string) error {
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
		t15, err15 := strconv.ParseInt(path[0], 0, 0)
		if err15 != nil {
			return err15
		}
		i = int(t15)
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

func (i2 TestFloatPtrSliceInspector) Set(dst, value any, path ...string) error {
	return i2.SetWithBuffer(dst, value, nil, path...)
}

func (i2 TestFloatPtrSliceInspector) DeepEqual(l, r any) bool {
	return i2.DeepEqualWithOptions(l, r, nil)
}

func (i2 TestFloatPtrSliceInspector) DeepEqualWithOptions(l, r any, opts *inspector.DEQOptions) bool {
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

func (i2 TestFloatPtrSliceInspector) Unmarshal(p []byte, typ inspector.Encoding) (any, error) {
	var x testobj.TestFloatPtrSlice
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i2 TestFloatPtrSliceInspector) Copy(x any) (any, error) {
	var r testobj.TestFloatPtrSlice
	switch x.(type) {
	case testobj.TestFloatPtrSlice:
		r = x.(testobj.TestFloatPtrSlice)
	case *testobj.TestFloatPtrSlice:
		r = *x.(*testobj.TestFloatPtrSlice)
	case **testobj.TestFloatPtrSlice:
		r = **x.(**testobj.TestFloatPtrSlice)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i2.countBytes(&r)
	var l testobj.TestFloatPtrSlice
	err := i2.CopyTo(&r, &l, inspector.NewByteBuffer(bc))
	return &l, err
}

func (i2 TestFloatPtrSliceInspector) CopyTo(src, dst any, buf inspector.AccumulativeBuffer) error {
	var r testobj.TestFloatPtrSlice
	switch src.(type) {
	case testobj.TestFloatPtrSlice:
		r = src.(testobj.TestFloatPtrSlice)
	case *testobj.TestFloatPtrSlice:
		r = *src.(*testobj.TestFloatPtrSlice)
	case **testobj.TestFloatPtrSlice:
		r = **src.(**testobj.TestFloatPtrSlice)
	default:
		return inspector.ErrUnsupportedType
	}
	var l *testobj.TestFloatPtrSlice
	switch dst.(type) {
	case testobj.TestFloatPtrSlice:
		return inspector.ErrMustPointerType
	case *testobj.TestFloatPtrSlice:
		l = dst.(*testobj.TestFloatPtrSlice)
	case **testobj.TestFloatPtrSlice:
		l = *dst.(**testobj.TestFloatPtrSlice)
	default:
		return inspector.ErrUnsupportedType
	}
	bb := buf.AcquireBytes()
	var err error
	if bb, err = i2.cpy(bb, l, &r); err != nil {
		return err
	}
	buf.ReleaseBytes(bb)
	return nil
}

func (i2 TestFloatPtrSliceInspector) countBytes(x *testobj.TestFloatPtrSlice) (c int) {
	return c
}

func (i2 TestFloatPtrSliceInspector) cpy(buf []byte, l, r *testobj.TestFloatPtrSlice) ([]byte, error) {
	if len(*r) > 0 {
		buf0 := (*l)
		if buf0 == nil {
			buf0 = make(testobj.TestFloatPtrSlice, 0, len(*r))
		}
		for i0 := 0; i0 < len(*r); i0++ {
			var b0 *float32
			x0 := (*r)[i0]
			b0 = x0
			buf0 = append(buf0, b0)
		}
		l = &buf0
	}
	return buf, nil
}

func (i2 TestFloatPtrSliceInspector) Length(src any, result *int, path ...string) error {
	if src == nil {
		return nil
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
		return inspector.ErrUnsupportedType
	}

	*result = 0
	if len(path) == 0 {
		return nil
	}
	if len(path) == 0 {
		*result = len(*x)
		return nil
	}
	if len(path) < 1 {
		return nil
	}
	var i int
	t16, err16 := strconv.ParseInt(path[0], 0, 0)
	if err16 != nil {
		return err16
	}
	i = int(t16)
	if len(*x) > i {
		x0 := (*x)[i]
		_ = x0
		if x0 == nil {
			return nil
		}
	}
	return nil
}

func (i2 TestFloatPtrSliceInspector) Capacity(src any, result *int, path ...string) error {
	if src == nil {
		return nil
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
		return inspector.ErrUnsupportedType
	}

	*result = 0
	if len(path) == 0 {
		return nil
	}
	if len(path) == 0 {
		*result = cap(*x)
		return nil
	}
	if len(path) < 1 {
		return nil
	}
	var i int
	t17, err17 := strconv.ParseInt(path[0], 0, 0)
	if err17 != nil {
		return err17
	}
	i = int(t17)
	if len(*x) > i {
		x0 := (*x)[i]
		_ = x0
		if x0 == nil {
			return nil
		}
	}
	return nil
}

func (i2 TestFloatPtrSliceInspector) Reset(x any) error {
	var origin *testobj.TestFloatPtrSlice
	_ = origin
	switch x.(type) {
	case testobj.TestFloatPtrSlice:
		return inspector.ErrMustPointerType
	case *testobj.TestFloatPtrSlice:
		origin = x.(*testobj.TestFloatPtrSlice)
	case **testobj.TestFloatPtrSlice:
		origin = *x.(**testobj.TestFloatPtrSlice)
	default:
		return inspector.ErrUnsupportedType
	}
	if l := len((*origin)); l > 0 {
		(*origin) = (*origin)[:0]
	}
	return nil
}
