// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"encoding/json"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

type TestPermissionInspector struct {
	inspector.BaseInspector
}

func (i7 TestPermissionInspector) TypeName() string {
	return "TestPermission"
}

func (i7 TestPermissionInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i7.GetTo(src, &buf, path...)
	return buf, err
}

func (i7 TestPermissionInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestPermission
	_ = x
	if p, ok := src.(**testobj.TestPermission); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := src.(testobj.TestPermission); ok {
		x = &v
	} else {
		return
	}
	if len(path) == 0 {
		*buf = &(*x)
		return
	}

	if len(path) > 0 {
		var k int32
		t267, err267 := strconv.ParseInt(path[0], 0, 0)
		if err267 != nil {
			return err267
		}
		k = int32(t267)
		x0 := (*x)[k]
		_ = x0
		*buf = &x0
		return
	}
	return
}

func (i7 TestPermissionInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestPermission
	_ = x
	if p, ok := src.(**testobj.TestPermission); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := src.(testobj.TestPermission); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		var k int32
		t268, err268 := strconv.ParseInt(path[0], 0, 0)
		if err268 != nil {
			return err268
		}
		k = int32(t268)
		x0 := (*x)[k]
		_ = x0
		var rightExact bool
		t269, err269 := strconv.ParseBool(right)
		if err269 != nil {
			return err269
		}
		rightExact = bool(t269)
		if cond == inspector.OpEq {
			*result = x0 == rightExact
		} else {
			*result = x0 != rightExact
		}
		return
	}
	return
}

func (i7 TestPermissionInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestPermission
	_ = x
	if p, ok := src.(**testobj.TestPermission); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := src.(testobj.TestPermission); ok {
		x = &v
	} else {
		return
	}

	for k := range *x {
		if l.RequireKey() {
			*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)
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

func (i7 TestPermissionInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestPermission
	_ = x
	if p, ok := dst.(**testobj.TestPermission); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := dst.(testobj.TestPermission); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		var k int32
		t270, err270 := strconv.ParseInt(path[0], 0, 0)
		if err270 != nil {
			return err270
		}
		k = int32(t270)
		x0 := (*x)[k]
		_ = x0
		inspector.AssignBuf(&x0, value, buf)
		(*x)[k] = x0
		return nil
	}
	return nil
}

func (i7 TestPermissionInspector) Set(dst, value interface{}, path ...string) error {
	return i7.SetWB(dst, value, nil, path...)
}

func (i7 TestPermissionInspector) DeepEqual(l, r interface{}) bool {
	return i7.DeepEqualWithOptions(l, r, nil)
}

func (i7 TestPermissionInspector) DeepEqualWithOptions(l, r interface{}, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestPermission
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestPermission); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestPermission); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestPermission); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestPermission); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestPermission); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestPermission); ok {
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
		if lx1 != rx1 {
			return false
		}
	}
	return true
}

func (i7 TestPermissionInspector) Unmarshal(p []byte, typ inspector.Encoding) (interface{}, error) {
	var x testobj.TestPermission
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i7 TestPermissionInspector) Copy(x interface{}) (interface{}, error) {
	var origin, cpy testobj.TestPermission
	switch x.(type) {
	case testobj.TestPermission:
		origin = x.(testobj.TestPermission)
	case *testobj.TestPermission:
		origin = *x.(*testobj.TestPermission)
	case **testobj.TestPermission:
		origin = **x.(**testobj.TestPermission)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i7.calcBytes(&origin)
	buf1 := make([]byte, 0, bc)
	if err := i7.cpy(buf1, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i7 TestPermissionInspector) CopyWB(x interface{}, buf inspector.AccumulativeBuffer) (interface{}, error) {
	var origin, cpy testobj.TestPermission
	switch x.(type) {
	case testobj.TestPermission:
		origin = x.(testobj.TestPermission)
	case *testobj.TestPermission:
		origin = *x.(*testobj.TestPermission)
	case **testobj.TestPermission:
		origin = **x.(**testobj.TestPermission)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	buf1 := buf.AcquireBytes()
	defer buf.ReleaseBytes(buf1)
	if err := i7.cpy(buf1, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i7 TestPermissionInspector) calcBytes(x *testobj.TestPermission) (c int) {
	return c
}

func (i7 TestPermissionInspector) cpy(buf []byte, l, r *testobj.TestPermission) error {
	if len(*r) > 0 {
		buf0 := make(testobj.TestPermission, len(*r))
		_ = buf0
		for rk0, rv0 := range *r {
			_, _ = rk0, rv0
			var lk0 int32
			lk0 = rk0
			var lv0 bool
			lv0 = rv0
			(*l)[lk0] = lv0
		}
	}
	return nil
}

func (i7 TestPermissionInspector) Reset(x interface{}) {
	var origin testobj.TestPermission
	_ = origin
	switch x.(type) {
	case testobj.TestPermission:
		origin = x.(testobj.TestPermission)
	case *testobj.TestPermission:
		origin = *x.(*testobj.TestPermission)
	case **testobj.TestPermission:
		origin = **x.(**testobj.TestPermission)
	default:
		return
	}
	if l := len((origin)); l > 0 {
		for k, _ := range origin {
			delete((origin), k)
		}
	}
}
