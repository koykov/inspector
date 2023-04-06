// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"encoding/json"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
)

type TestStringPtrFloatPtrMapInspector struct {
	inspector.BaseInspector
}

func (i10 TestStringPtrFloatPtrMapInspector) TypeName() string {
	return "TestStringPtrFloatPtrMap"
}

func (i10 TestStringPtrFloatPtrMapInspector) Get(src any, path ...string) (any, error) {
	var buf any
	err := i10.GetTo(src, &buf, path...)
	return buf, err
}

func (i10 TestStringPtrFloatPtrMapInspector) GetTo(src any, buf *any, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestStringPtrFloatPtrMap
	_ = x
	if p, ok := src.(**testobj.TestStringPtrFloatPtrMap); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStringPtrFloatPtrMap); ok {
		x = p
	} else if v, ok := src.(testobj.TestStringPtrFloatPtrMap); ok {
		x = &v
	} else {
		return
	}
	if len(path) == 0 {
		*buf = &(*x)
		return
	}

	if len(path) > 0 {
		if x0, ok := (*x)[&path[0]]; ok {
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

func (i10 TestStringPtrFloatPtrMapInspector) Cmp(src any, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestStringPtrFloatPtrMap
	_ = x
	if p, ok := src.(**testobj.TestStringPtrFloatPtrMap); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStringPtrFloatPtrMap); ok {
		x = p
	} else if v, ok := src.(testobj.TestStringPtrFloatPtrMap); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if x0, ok := (*x)[&path[0]]; ok {
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

func (i10 TestStringPtrFloatPtrMapInspector) Loop(src any, l inspector.Iterator, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestStringPtrFloatPtrMap
	_ = x
	if p, ok := src.(**testobj.TestStringPtrFloatPtrMap); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStringPtrFloatPtrMap); ok {
		x = p
	} else if v, ok := src.(testobj.TestStringPtrFloatPtrMap); ok {
		x = &v
	} else {
		return
	}

	for k := range *x {
		if l.RequireKey() {
			*buf = append((*buf)[:0], *k...)
			l.SetKey(buf, &inspector.StaticInspector{})
		}
		l.SetVal((*x)[k], &inspector.StaticInspector{})
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

func (i10 TestStringPtrFloatPtrMapInspector) SetWB(dst, value any, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestStringPtrFloatPtrMap
	_ = x
	if p, ok := dst.(**testobj.TestStringPtrFloatPtrMap); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestStringPtrFloatPtrMap); ok {
		x = p
	} else if v, ok := dst.(testobj.TestStringPtrFloatPtrMap); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		x0 := (*x)[&path[0]]
		_ = x0
		if x0 == nil {
			return nil
		}
		inspector.AssignBuf(x0, value, buf)
		(*x)[&path[0]] = x0
		return nil
	}
	return nil
}

func (i10 TestStringPtrFloatPtrMapInspector) Set(dst, value any, path ...string) error {
	return i10.SetWB(dst, value, nil, path...)
}

func (i10 TestStringPtrFloatPtrMapInspector) DeepEqual(l, r any) bool {
	return i10.DeepEqualWithOptions(l, r, nil)
}

func (i10 TestStringPtrFloatPtrMapInspector) DeepEqualWithOptions(l, r any, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestStringPtrFloatPtrMap
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestStringPtrFloatPtrMap); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestStringPtrFloatPtrMap); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestStringPtrFloatPtrMap); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestStringPtrFloatPtrMap); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestStringPtrFloatPtrMap); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestStringPtrFloatPtrMap); ok {
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
	for k := range *lx {
		lx1 := (*lx)[k]
		rx1, ok1 := (*rx)[k]
		_, _, _ = lx1, rx1, ok1
		if !ok1 {
			return false
		}
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

func (i10 TestStringPtrFloatPtrMapInspector) Unmarshal(p []byte, typ inspector.Encoding) (any, error) {
	var x testobj.TestStringPtrFloatPtrMap
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i10 TestStringPtrFloatPtrMapInspector) Copy(x any) (any, error) {
	var r testobj.TestStringPtrFloatPtrMap
	switch x.(type) {
	case testobj.TestStringPtrFloatPtrMap:
		r = x.(testobj.TestStringPtrFloatPtrMap)
	case *testobj.TestStringPtrFloatPtrMap:
		r = *x.(*testobj.TestStringPtrFloatPtrMap)
	case **testobj.TestStringPtrFloatPtrMap:
		r = **x.(**testobj.TestStringPtrFloatPtrMap)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i10.countBytes(&r)
	var l testobj.TestStringPtrFloatPtrMap
	err := i10.CopyTo(&r, &l, inspector.NewByteBuffer(bc))
	return &l, err
}

func (i10 TestStringPtrFloatPtrMapInspector) CopyTo(src, dst any, buf inspector.AccumulativeBuffer) error {
	var r testobj.TestStringPtrFloatPtrMap
	switch src.(type) {
	case testobj.TestStringPtrFloatPtrMap:
		r = src.(testobj.TestStringPtrFloatPtrMap)
	case *testobj.TestStringPtrFloatPtrMap:
		r = *src.(*testobj.TestStringPtrFloatPtrMap)
	case **testobj.TestStringPtrFloatPtrMap:
		r = **src.(**testobj.TestStringPtrFloatPtrMap)
	default:
		return inspector.ErrUnsupportedType
	}
	var l *testobj.TestStringPtrFloatPtrMap
	switch dst.(type) {
	case testobj.TestStringPtrFloatPtrMap:
		return inspector.ErrMustPointerType
	case *testobj.TestStringPtrFloatPtrMap:
		l = dst.(*testobj.TestStringPtrFloatPtrMap)
	case **testobj.TestStringPtrFloatPtrMap:
		l = *dst.(**testobj.TestStringPtrFloatPtrMap)
	default:
		return inspector.ErrUnsupportedType
	}
	bb := buf.AcquireBytes()
	var err error
	if bb, err = i10.cpy(bb, l, &r); err != nil {
		return err
	}
	buf.ReleaseBytes(bb)
	return nil
}

func (i10 TestStringPtrFloatPtrMapInspector) countBytes(x *testobj.TestStringPtrFloatPtrMap) (c int) {
	for k0, v0 := range *x {
		_, _ = k0, v0
		c += len(*k0)
	}
	return c
}

func (i10 TestStringPtrFloatPtrMapInspector) cpy(buf []byte, l, r *testobj.TestStringPtrFloatPtrMap) ([]byte, error) {
	if len(*r) > 0 {
		if l == nil {
			buf0 := make(testobj.TestStringPtrFloatPtrMap, len(*r))
			l = &buf0
		}
		for rk0, rv0 := range *r {
			_, _ = rk0, rv0
			var lk0 *string
			buf, *lk0 = inspector.BufferizeString(buf, *rk0)
			var lv0 *float64
			lv0 = rv0
			(*l)[lk0] = lv0
		}
	}
	return buf, nil
}

func (i10 TestStringPtrFloatPtrMapInspector) Reset(x any) error {
	var origin *testobj.TestStringPtrFloatPtrMap
	_ = origin
	switch x.(type) {
	case testobj.TestStringPtrFloatPtrMap:
		return inspector.ErrMustPointerType
	case *testobj.TestStringPtrFloatPtrMap:
		origin = x.(*testobj.TestStringPtrFloatPtrMap)
	case **testobj.TestStringPtrFloatPtrMap:
		origin = *x.(**testobj.TestStringPtrFloatPtrMap)
	default:
		return inspector.ErrUnsupportedType
	}
	if l := len((*origin)); l > 0 {
		for k, _ := range *origin {
			delete((*origin), k)
		}
	}
	return nil
}
