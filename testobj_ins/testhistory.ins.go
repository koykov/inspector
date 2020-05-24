// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"bytes"
	"github.com/koykov/fastconv"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

type TestHistoryInspector struct {
	inspector.BaseInspector
}

func (i2 *TestHistoryInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i2.GetTo(src, &buf, path...)
	return buf, err
}

func (i2 *TestHistoryInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestHistory
	_ = x
	if p, ok := src.(*testobj.TestHistory); ok {
		x = p
	} else if v, ok := src.(testobj.TestHistory); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "DateUnix" {
			*buf = &x.DateUnix
			return
		}
		if path[0] == "Cost" {
			*buf = &x.Cost
			return
		}
		if path[0] == "Comment" {
			*buf = &x.Comment
			return
		}
	}
	return
}

func (i2 *TestHistoryInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestHistory
	_ = x
	if p, ok := src.(*testobj.TestHistory); ok {
		x = p
	} else if v, ok := src.(testobj.TestHistory); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "DateUnix" {
			var rightExact int64
			t10, err10 := strconv.ParseInt(right, 0, 0)
			if err10 != nil {
				return err10
			}
			rightExact = int64(t10)
			switch cond {
			case inspector.OpEq:
				*result = x.DateUnix == rightExact
			case inspector.OpNq:
				*result = x.DateUnix != rightExact
			case inspector.OpGt:
				*result = x.DateUnix > rightExact
			case inspector.OpGtq:
				*result = x.DateUnix >= rightExact
			case inspector.OpLt:
				*result = x.DateUnix < rightExact
			case inspector.OpLtq:
				*result = x.DateUnix <= rightExact
			}
			return
		}
		if path[0] == "Cost" {
			var rightExact float64
			t11, err11 := strconv.ParseFloat(right, 0)
			if err11 != nil {
				return err11
			}
			rightExact = float64(t11)
			switch cond {
			case inspector.OpEq:
				*result = x.Cost == rightExact
			case inspector.OpNq:
				*result = x.Cost != rightExact
			case inspector.OpGt:
				*result = x.Cost > rightExact
			case inspector.OpGtq:
				*result = x.Cost >= rightExact
			case inspector.OpLt:
				*result = x.Cost < rightExact
			case inspector.OpLtq:
				*result = x.Cost <= rightExact
			}
			return
		}
		if path[0] == "Comment" {
			var rightExact []byte
			rightExact = fastconv.S2B(right)

			if cond == inspector.OpEq {
				*result = bytes.Equal(x.Comment, rightExact)
			} else {
				*result = !bytes.Equal(x.Comment, rightExact)
			}
			return
		}
	}
	return
}

func (i2 *TestHistoryInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestHistory
	_ = x
	if p, ok := src.(*testobj.TestHistory); ok {
		x = p
	} else if v, ok := src.(testobj.TestHistory); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
	}
	return
}

func (i2 *TestHistoryInspector) Set(dst, value interface{}, path ...string) {
}