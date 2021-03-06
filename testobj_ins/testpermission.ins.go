// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

type TestPermissionInspector struct {
	inspector.BaseInspector
}

func (i4 *TestPermissionInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i4.GetTo(src, &buf, path...)
	return buf, err
}

func (i4 *TestPermissionInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
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
		t36, err36 := strconv.ParseInt(path[0], 0, 0)
		if err36 != nil {
			return err36
		}
		k = int32(t36)
		x0 := (*x)[k]
		_ = x0
		*buf = &x0
		return
	}
	return
}

func (i4 *TestPermissionInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
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
		t37, err37 := strconv.ParseInt(path[0], 0, 0)
		if err37 != nil {
			return err37
		}
		k = int32(t37)
		x0 := (*x)[k]
		_ = x0
		var rightExact bool
		t38, err38 := strconv.ParseBool(right)
		if err38 != nil {
			return err38
		}
		rightExact = bool(t38)
		if cond == inspector.OpEq {
			*result = x0 == rightExact
		} else {
			*result = x0 != rightExact
		}
		return
	}
	return
}

func (i4 *TestPermissionInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
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

func (i4 *TestPermissionInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
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
		t39, err39 := strconv.ParseInt(path[0], 0, 0)
		if err39 != nil {
			return err39
		}
		k = int32(t39)
		x0 := (*x)[k]
		_ = x0
		inspector.AssignBuf(&x0, value, buf)
		(*x)[k] = x0
		return nil
	}
	return nil
}

func (i4 *TestPermissionInspector) Set(dst, value interface{}, path ...string) error {
	return i4.SetWB(dst, value, nil, path...)
}
