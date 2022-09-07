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

type TestHistoryInspector struct {
	inspector.BaseInspector
}

func (i4 TestHistoryInspector) TypeName() string {
	return "TestHistory"
}

func (i4 TestHistoryInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i4.GetTo(src, &buf, path...)
	return buf, err
}

func (i4 TestHistoryInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestHistory
	_ = x
	if p, ok := src.(**testobj.TestHistory); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestHistory); ok {
		x = p
	} else if v, ok := src.(testobj.TestHistory); ok {
		x = &v
	} else {
		return
	}
	if len(path) == 0 {
		*buf = &(*x)
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

func (i4 TestHistoryInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestHistory
	_ = x
	if p, ok := src.(**testobj.TestHistory); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestHistory); ok {
		x = p
	} else if v, ok := src.(testobj.TestHistory); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "DateUnix" {
			var rightExact int64
			t18, err18 := strconv.ParseInt(right, 0, 0)
			if err18 != nil {
				return err18
			}
			rightExact = int64(t18)
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
			t19, err19 := strconv.ParseFloat(right, 0)
			if err19 != nil {
				return err19
			}
			rightExact = float64(t19)
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

func (i4 TestHistoryInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestHistory
	_ = x
	if p, ok := src.(**testobj.TestHistory); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestHistory); ok {
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

func (i4 TestHistoryInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestHistory
	_ = x
	if p, ok := dst.(**testobj.TestHistory); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestHistory); ok {
		x = p
	} else if v, ok := dst.(testobj.TestHistory); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		if path[0] == "DateUnix" {
			inspector.AssignBuf(&x.DateUnix, value, buf)
			return nil
		}
		if path[0] == "Cost" {
			inspector.AssignBuf(&x.Cost, value, buf)
			return nil
		}
		if path[0] == "Comment" {
			inspector.AssignBuf(&x.Comment, value, buf)
			return nil
		}
	}
	return nil
}

func (i4 TestHistoryInspector) Set(dst, value interface{}, path ...string) error {
	return i4.SetWB(dst, value, nil, path...)
}

func (i4 TestHistoryInspector) DeepEqual(l, r interface{}) bool {
	return i4.DeepEqualWithOptions(l, r, nil)
}

func (i4 TestHistoryInspector) DeepEqualWithOptions(l, r interface{}, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestHistory
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestHistory); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestHistory); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestHistory); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestHistory); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestHistory); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestHistory); ok {
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

	if lx.DateUnix != rx.DateUnix && inspector.DEQMustCheck("DateUnix", opts) {
		return false
	}
	if !inspector.EqualFloat64(lx.Cost, rx.Cost, opts) && inspector.DEQMustCheck("Cost", opts) {
		return false
	}
	if !bytes.Equal(lx.Comment, rx.Comment) && inspector.DEQMustCheck("Comment", opts) {
		return false
	}
	return true
}

func (i4 TestHistoryInspector) Unmarshal(p []byte, typ inspector.Encoding) (interface{}, error) {
	var x testobj.TestHistory
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i4 TestHistoryInspector) Copy(x interface{}) (interface{}, error) {
	var origin, cpy testobj.TestHistory
	switch x.(type) {
	case testobj.TestHistory:
		origin = x.(testobj.TestHistory)
	case *testobj.TestHistory:
		origin = *x.(*testobj.TestHistory)
	case **testobj.TestHistory:
		origin = **x.(**testobj.TestHistory)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i4.calcBytes(&origin)
	buf := make([]byte, 0, bc)
	if err := i4.cpy(buf, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i4 TestHistoryInspector) calcBytes(x *testobj.TestHistory) (c int) {
	c += len(x.Comment)
	return c
}

func (i4 TestHistoryInspector) cpy(buf []byte, l, r *testobj.TestHistory) error {
	l.DateUnix = r.DateUnix
	l.Cost = r.Cost
	buf, l.Comment = inspector.Bufferize(buf, r.Comment)
	return nil
}
