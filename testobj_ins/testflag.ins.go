// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

type TestFlagInspector struct {
	inspector.BaseInspector
}

func (i1 *TestFlagInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i1.GetTo(src, &buf, path...)
	return buf, err
}

func (i1 *TestFlagInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestFlag
	_ = x
	if p, ok := src.(**testobj.TestFlag); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFlag); ok {
		x = p
	} else if v, ok := src.(testobj.TestFlag); ok {
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
			*buf = &x0
			return
		}
	}
	return
}

func (i1 *TestFlagInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestFlag
	_ = x
	if p, ok := src.(**testobj.TestFlag); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFlag); ok {
		x = p
	} else if v, ok := src.(testobj.TestFlag); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if x0, ok := (*x)[path[0]]; ok {
			_ = x0
			var rightExact int32
			t10, err10 := strconv.ParseInt(right, 0, 0)
			if err10 != nil {
				return err10
			}
			rightExact = int32(t10)
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

func (i1 *TestFlagInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestFlag
	_ = x
	if p, ok := src.(**testobj.TestFlag); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFlag); ok {
		x = p
	} else if v, ok := src.(testobj.TestFlag); ok {
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

func (i1 *TestFlagInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestFlag
	_ = x
	if p, ok := dst.(**testobj.TestFlag); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestFlag); ok {
		x = p
	} else if v, ok := dst.(testobj.TestFlag); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		x0 := (*x)[path[0]]
		_ = x0
		inspector.AssignBuf(&x0, value, buf)
		(*x)[path[0]] = x0
		return nil
	}
	return nil
}

func (i1 *TestFlagInspector) Set(dst, value interface{}, path ...string) error {
	return i1.SetWB(dst, value, nil, path...)
}
