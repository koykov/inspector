// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"encoding/json"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
)

type TestStringFloatPtrMapInspector struct {
	inspector.BaseInspector
}

func (i9 TestStringFloatPtrMapInspector) TypeName() string {
	return "TestStringFloatPtrMap"
}

func (i9 TestStringFloatPtrMapInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i9.GetTo(src, &buf, path...)
	return buf, err
}

func (i9 TestStringFloatPtrMapInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestStringFloatPtrMap
	_ = x
	if p, ok := src.(**testobj.TestStringFloatPtrMap); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStringFloatPtrMap); ok {
		x = p
	} else if v, ok := src.(testobj.TestStringFloatPtrMap); ok {
		x = &v
	} else {
		return
	}
	if len(path) == 0 {
		*buf = &(*x)
		return
	}

	if len(path) > 0 {
		if x0, ok := (*x)[path[0]]; ok {
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

func (i9 TestStringFloatPtrMapInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestStringFloatPtrMap
	_ = x
	if p, ok := src.(**testobj.TestStringFloatPtrMap); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStringFloatPtrMap); ok {
		x = p
	} else if v, ok := src.(testobj.TestStringFloatPtrMap); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if x0, ok := (*x)[path[0]]; ok {
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

func (i9 TestStringFloatPtrMapInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestStringFloatPtrMap
	_ = x
	if p, ok := src.(**testobj.TestStringFloatPtrMap); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestStringFloatPtrMap); ok {
		x = p
	} else if v, ok := src.(testobj.TestStringFloatPtrMap); ok {
		x = &v
	} else {
		return
	}

	for k := range *x {
		if l.RequireKey() {
			*buf = append((*buf)[:0], k...)
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

func (i9 TestStringFloatPtrMapInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestStringFloatPtrMap
	_ = x
	if p, ok := dst.(**testobj.TestStringFloatPtrMap); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestStringFloatPtrMap); ok {
		x = p
	} else if v, ok := dst.(testobj.TestStringFloatPtrMap); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		x0 := (*x)[path[0]]
		_ = x0
		if x0 == nil {
			return nil
		}
		inspector.AssignBuf(x0, value, buf)
		(*x)[path[0]] = x0
		return nil
	}
	return nil
}

func (i9 TestStringFloatPtrMapInspector) Set(dst, value interface{}, path ...string) error {
	return i9.SetWB(dst, value, nil, path...)
}

func (i9 TestStringFloatPtrMapInspector) DeepEqual(l, r interface{}) bool {
	return i9.DeepEqualWithOptions(l, r, nil)
}

func (i9 TestStringFloatPtrMapInspector) DeepEqualWithOptions(l, r interface{}, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestStringFloatPtrMap
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestStringFloatPtrMap); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestStringFloatPtrMap); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestStringFloatPtrMap); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestStringFloatPtrMap); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestStringFloatPtrMap); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestStringFloatPtrMap); ok {
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

func (i9 TestStringFloatPtrMapInspector) Unmarshal(p []byte, typ inspector.Encoding) (interface{}, error) {
	var x testobj.TestStringFloatPtrMap
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i9 TestStringFloatPtrMapInspector) Copy(x interface{}) (interface{}, error) {
	var origin, cpy testobj.TestStringFloatPtrMap
	switch x.(type) {
	case testobj.TestStringFloatPtrMap:
		origin = x.(testobj.TestStringFloatPtrMap)
	case *testobj.TestStringFloatPtrMap:
		origin = *x.(*testobj.TestStringFloatPtrMap)
	case **testobj.TestStringFloatPtrMap:
		origin = **x.(**testobj.TestStringFloatPtrMap)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i9.calcBytes(&origin)
	buf1 := make([]byte, 0, bc)
	if err := i9.cpy(buf1, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i9 TestStringFloatPtrMapInspector) CopyWB(x interface{}, buf inspector.AccumulativeBuffer) (interface{}, error) {
	var origin, cpy testobj.TestStringFloatPtrMap
	switch x.(type) {
	case testobj.TestStringFloatPtrMap:
		origin = x.(testobj.TestStringFloatPtrMap)
	case *testobj.TestStringFloatPtrMap:
		origin = *x.(*testobj.TestStringFloatPtrMap)
	case **testobj.TestStringFloatPtrMap:
		origin = **x.(**testobj.TestStringFloatPtrMap)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	buf1 := buf.AcquireBytes()
	defer buf.ReleaseBytes(buf1)
	if err := i9.cpy(buf1, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i9 TestStringFloatPtrMapInspector) calcBytes(x *testobj.TestStringFloatPtrMap) (c int) {
	for k0, v0 := range *x {
		_, _ = k0, v0
		c += len(k0)
	}
	return c
}

func (i9 TestStringFloatPtrMapInspector) cpy(buf []byte, l, r *testobj.TestStringFloatPtrMap) error {
	if len(*r) > 0 {
		buf0 := make(testobj.TestStringFloatPtrMap, len(*r))
		_ = buf0
		for rk, rv := range *r {
			_, _ = rk, rv
			var lk string
			buf, lk = inspector.BufferizeString(buf, rk)
			var lv *float64
			lv = rv
			(*l)[lk] = lv
		}
	}
	return nil
}

func (i9 TestStringFloatPtrMapInspector) Reset(x interface{}) {
	var origin testobj.TestStringFloatPtrMap
	_ = origin
	switch x.(type) {
	case testobj.TestStringFloatPtrMap:
		origin = x.(testobj.TestStringFloatPtrMap)
	case *testobj.TestStringFloatPtrMap:
		origin = *x.(*testobj.TestStringFloatPtrMap)
	case **testobj.TestStringFloatPtrMap:
		origin = **x.(**testobj.TestStringFloatPtrMap)
	default:
		return
	}
	if l := len((origin)); l > 0 {
		for k, _ := range origin {
			delete((origin), k)
		}
	}
}
