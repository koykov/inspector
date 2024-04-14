// Code generated by inspc. DO NOT EDIT.
// source: github.com/koykov/inspector/testobj

package testobj_ins

import (
	"bytes"
	"encoding/json"
	"github.com/koykov/byteconv"
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"strconv"
)

func init() {
	inspector.RegisterInspector("TestFinance", TestFinanceInspector{})
}

type TestFinanceInspector struct {
	inspector.BaseInspector
}

func (i0 TestFinanceInspector) TypeName() string {
	return "TestFinance"
}

func (i0 TestFinanceInspector) Get(src any, path ...string) (any, error) {
	var buf any
	err := i0.GetTo(src, &buf, path...)
	return buf, err
}

func (i0 TestFinanceInspector) GetTo(src any, buf *any, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestFinance
	_ = x
	if p, ok := src.(**testobj.TestFinance); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFinance); ok {
		x = p
	} else if v, ok := src.(testobj.TestFinance); ok {
		x = &v
	} else {
		return
	}
	if len(path) == 0 {
		*buf = &(*x)
		return
	}

	if len(path) > 0 {
		if path[0] == "MoneyIn" {
			*buf = &x.MoneyIn
			return
		}
		if path[0] == "MoneyOut" {
			*buf = &x.MoneyOut
			return
		}
		if path[0] == "Balance" {
			*buf = &x.Balance
			return
		}
		if path[0] == "AllowBuy" {
			*buf = &x.AllowBuy
			return
		}
		if path[0] == "History" {
			x0 := x.History
			_ = x0
			if len(path) > 1 {
				var i int
				t0, err0 := strconv.ParseInt(path[1], 0, 0)
				if err0 != nil {
					return err0
				}
				i = int(t0)
				if len(x0) > i {
					x1 := &(x0)[i]
					_ = x1
					if len(path) > 2 {
						if path[2] == "DateUnix" {
							*buf = &x1.DateUnix
							return
						}
						if path[2] == "Cost" {
							*buf = &x1.Cost
							return
						}
						if path[2] == "Comment" {
							*buf = &x1.Comment
							return
						}
					}
					*buf = x1
				}
			}
			*buf = &x.History
			return
		}
	}
	return
}

func (i0 TestFinanceInspector) Compare(src any, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestFinance
	_ = x
	if p, ok := src.(**testobj.TestFinance); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFinance); ok {
		x = p
	} else if v, ok := src.(testobj.TestFinance); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "MoneyIn" {
			var rightExact float64
			t1, err1 := strconv.ParseFloat(right, 0)
			if err1 != nil {
				return err1
			}
			rightExact = float64(t1)
			switch cond {
			case inspector.OpEq:
				*result = x.MoneyIn == rightExact
			case inspector.OpNq:
				*result = x.MoneyIn != rightExact
			case inspector.OpGt:
				*result = x.MoneyIn > rightExact
			case inspector.OpGtq:
				*result = x.MoneyIn >= rightExact
			case inspector.OpLt:
				*result = x.MoneyIn < rightExact
			case inspector.OpLtq:
				*result = x.MoneyIn <= rightExact
			}
			return
		}
		if path[0] == "MoneyOut" {
			var rightExact float64
			t2, err2 := strconv.ParseFloat(right, 0)
			if err2 != nil {
				return err2
			}
			rightExact = float64(t2)
			switch cond {
			case inspector.OpEq:
				*result = x.MoneyOut == rightExact
			case inspector.OpNq:
				*result = x.MoneyOut != rightExact
			case inspector.OpGt:
				*result = x.MoneyOut > rightExact
			case inspector.OpGtq:
				*result = x.MoneyOut >= rightExact
			case inspector.OpLt:
				*result = x.MoneyOut < rightExact
			case inspector.OpLtq:
				*result = x.MoneyOut <= rightExact
			}
			return
		}
		if path[0] == "Balance" {
			var rightExact float64
			t3, err3 := strconv.ParseFloat(right, 0)
			if err3 != nil {
				return err3
			}
			rightExact = float64(t3)
			switch cond {
			case inspector.OpEq:
				*result = x.Balance == rightExact
			case inspector.OpNq:
				*result = x.Balance != rightExact
			case inspector.OpGt:
				*result = x.Balance > rightExact
			case inspector.OpGtq:
				*result = x.Balance >= rightExact
			case inspector.OpLt:
				*result = x.Balance < rightExact
			case inspector.OpLtq:
				*result = x.Balance <= rightExact
			}
			return
		}
		if path[0] == "AllowBuy" {
			var rightExact bool
			t4, err4 := strconv.ParseBool(right)
			if err4 != nil {
				return err4
			}
			rightExact = bool(t4)
			if cond == inspector.OpEq {
				*result = x.AllowBuy == rightExact
			} else {
				*result = x.AllowBuy != rightExact
			}
			return
		}
		if path[0] == "History" {
			x0 := x.History
			_ = x0
			if len(path) > 1 {
				var i int
				t5, err5 := strconv.ParseInt(path[1], 0, 0)
				if err5 != nil {
					return err5
				}
				i = int(t5)
				if len(x0) > i {
					x1 := &(x0)[i]
					_ = x1
					if len(path) > 2 {
						if path[2] == "DateUnix" {
							var rightExact int64
							t6, err6 := strconv.ParseInt(right, 0, 0)
							if err6 != nil {
								return err6
							}
							rightExact = int64(t6)
							switch cond {
							case inspector.OpEq:
								*result = x1.DateUnix == rightExact
							case inspector.OpNq:
								*result = x1.DateUnix != rightExact
							case inspector.OpGt:
								*result = x1.DateUnix > rightExact
							case inspector.OpGtq:
								*result = x1.DateUnix >= rightExact
							case inspector.OpLt:
								*result = x1.DateUnix < rightExact
							case inspector.OpLtq:
								*result = x1.DateUnix <= rightExact
							}
							return
						}
						if path[2] == "Cost" {
							var rightExact float64
							t7, err7 := strconv.ParseFloat(right, 0)
							if err7 != nil {
								return err7
							}
							rightExact = float64(t7)
							switch cond {
							case inspector.OpEq:
								*result = x1.Cost == rightExact
							case inspector.OpNq:
								*result = x1.Cost != rightExact
							case inspector.OpGt:
								*result = x1.Cost > rightExact
							case inspector.OpGtq:
								*result = x1.Cost >= rightExact
							case inspector.OpLt:
								*result = x1.Cost < rightExact
							case inspector.OpLtq:
								*result = x1.Cost <= rightExact
							}
							return
						}
						if path[2] == "Comment" {
							var rightExact []byte
							rightExact = byteconv.S2B(right)

							if cond == inspector.OpEq {
								*result = bytes.Equal(x1.Comment, rightExact)
							} else {
								*result = !bytes.Equal(x1.Comment, rightExact)
							}
							return
						}
					}
				}
			}
		}
	}
	return
}

func (i0 TestFinanceInspector) Loop(src any, l inspector.Iterator, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestFinance
	_ = x
	if p, ok := src.(**testobj.TestFinance); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFinance); ok {
		x = p
	} else if v, ok := src.(testobj.TestFinance); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "History" {
			x0 := x.History
			_ = x0
			for k := range x0 {
				if l.RequireKey() {
					*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)
					l.SetKey(buf, &inspector.StaticInspector{})
				}
				l.SetVal(&(x0)[k], &TestHistoryInspector{})
				ctl := l.Iterate()
				if ctl == inspector.LoopCtlBrk {
					break
				}
				if ctl == inspector.LoopCtlCnt {
					continue
				}
			}
			return
		}
	}
	return
}

func (i0 TestFinanceInspector) SetWithBuffer(dst, value any, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestFinance
	_ = x
	if p, ok := dst.(**testobj.TestFinance); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestFinance); ok {
		x = p
	} else if v, ok := dst.(testobj.TestFinance); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		if path[0] == "MoneyIn" {
			inspector.AssignBuf(&x.MoneyIn, value, buf)
			return nil
		}
		if path[0] == "MoneyOut" {
			inspector.AssignBuf(&x.MoneyOut, value, buf)
			return nil
		}
		if path[0] == "Balance" {
			inspector.AssignBuf(&x.Balance, value, buf)
			return nil
		}
		if path[0] == "AllowBuy" {
			inspector.AssignBuf(&x.AllowBuy, value, buf)
			return nil
		}
		if path[0] == "History" {
			x0 := x.History
			if uvalue, ok := value.(*[]testobj.TestHistory); ok {
				x0 = *uvalue
			}
			if x0 == nil {
				z := make([]testobj.TestHistory, 0)
				x0 = z
				x.History = x0
			}
			_ = x0
			if len(path) > 1 {
				var i int
				t9, err9 := strconv.ParseInt(path[1], 0, 0)
				if err9 != nil {
					return err9
				}
				i = int(t9)
				if len(x0) > i {
					x1 := &(x0)[i]
					_ = x1
					if len(path) > 2 {
						if path[2] == "DateUnix" {
							inspector.AssignBuf(&x1.DateUnix, value, buf)
							return nil
						}
						if path[2] == "Cost" {
							inspector.AssignBuf(&x1.Cost, value, buf)
							return nil
						}
						if path[2] == "Comment" {
							inspector.AssignBuf(&x1.Comment, value, buf)
							return nil
						}
					}
					(x0)[i] = *x1
					return nil
				}
			}
			x.History = x0
		}
	}
	return nil
}

func (i0 TestFinanceInspector) Set(dst, value any, path ...string) error {
	return i0.SetWithBuffer(dst, value, nil, path...)
}

func (i0 TestFinanceInspector) DeepEqual(l, r any) bool {
	return i0.DeepEqualWithOptions(l, r, nil)
}

func (i0 TestFinanceInspector) DeepEqualWithOptions(l, r any, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestFinance
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestFinance); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestFinance); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestFinance); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestFinance); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestFinance); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestFinance); ok {
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

	if !inspector.EqualFloat64(lx.MoneyIn, rx.MoneyIn, opts) && inspector.DEQMustCheck("MoneyIn", opts) {
		return false
	}
	if !inspector.EqualFloat64(lx.MoneyOut, rx.MoneyOut, opts) && inspector.DEQMustCheck("MoneyOut", opts) {
		return false
	}
	if !inspector.EqualFloat64(lx.Balance, rx.Balance, opts) && inspector.DEQMustCheck("Balance", opts) {
		return false
	}
	if lx.AllowBuy != rx.AllowBuy && inspector.DEQMustCheck("AllowBuy", opts) {
		return false
	}
	lx1 := lx.History
	rx1 := rx.History
	_, _ = lx1, rx1
	if inspector.DEQMustCheck("History", opts) {
		if len(lx1) != len(rx1) {
			return false
		}
		for i := 0; i < len(lx1); i++ {
			lx2 := (lx1)[i]
			rx2 := (rx1)[i]
			_, _ = lx2, rx2
			if lx2.DateUnix != rx2.DateUnix && inspector.DEQMustCheck("History.DateUnix", opts) {
				return false
			}
			if !inspector.EqualFloat64(lx2.Cost, rx2.Cost, opts) && inspector.DEQMustCheck("History.Cost", opts) {
				return false
			}
			if !bytes.Equal(lx2.Comment, rx2.Comment) && inspector.DEQMustCheck("History.Comment", opts) {
				return false
			}
		}
	}
	return true
}

func (i0 TestFinanceInspector) Unmarshal(p []byte, typ inspector.Encoding) (any, error) {
	var x testobj.TestFinance
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i0 TestFinanceInspector) Copy(x any) (any, error) {
	var r testobj.TestFinance
	switch x.(type) {
	case testobj.TestFinance:
		r = x.(testobj.TestFinance)
	case *testobj.TestFinance:
		r = *x.(*testobj.TestFinance)
	case **testobj.TestFinance:
		r = **x.(**testobj.TestFinance)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i0.countBytes(&r)
	var l testobj.TestFinance
	err := i0.CopyTo(&r, &l, inspector.NewByteBuffer(bc))
	return &l, err
}

func (i0 TestFinanceInspector) CopyTo(src, dst any, buf inspector.AccumulativeBuffer) error {
	var r testobj.TestFinance
	switch src.(type) {
	case testobj.TestFinance:
		r = src.(testobj.TestFinance)
	case *testobj.TestFinance:
		r = *src.(*testobj.TestFinance)
	case **testobj.TestFinance:
		r = **src.(**testobj.TestFinance)
	default:
		return inspector.ErrUnsupportedType
	}
	var l *testobj.TestFinance
	switch dst.(type) {
	case testobj.TestFinance:
		return inspector.ErrMustPointerType
	case *testobj.TestFinance:
		l = dst.(*testobj.TestFinance)
	case **testobj.TestFinance:
		l = *dst.(**testobj.TestFinance)
	default:
		return inspector.ErrUnsupportedType
	}
	bb := buf.AcquireBytes()
	var err error
	if bb, err = i0.cpy(bb, l, &r); err != nil {
		return err
	}
	buf.ReleaseBytes(bb)
	return nil
}

func (i0 TestFinanceInspector) countBytes(x *testobj.TestFinance) (c int) {
	for i1 := 0; i1 < len(x.History); i1++ {
		x1 := &(x.History)[i1]
		c += len(x1.Comment)
	}
	return c
}

func (i0 TestFinanceInspector) cpy(buf []byte, l, r *testobj.TestFinance) ([]byte, error) {
	l.MoneyIn = r.MoneyIn
	l.MoneyOut = r.MoneyOut
	l.Balance = r.Balance
	l.AllowBuy = r.AllowBuy
	if len(r.History) > 0 {
		buf1 := (l.History)
		if buf1 == nil {
			buf1 = make([]testobj.TestHistory, 0, len(r.History))
		}
		for i1 := 0; i1 < len(r.History); i1++ {
			var b1 testobj.TestHistory
			x1 := &(r.History)[i1]
			b1.DateUnix = x1.DateUnix
			b1.Cost = x1.Cost
			buf, b1.Comment = inspector.Bufferize(buf, x1.Comment)
			buf1 = append(buf1, b1)
		}
		l.History = buf1
	}
	return buf, nil
}

func (i0 TestFinanceInspector) Length(src any, result *int, path ...string) error {
	if src == nil {
		return nil
	}
	var x *testobj.TestFinance
	_ = x
	if p, ok := src.(**testobj.TestFinance); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFinance); ok {
		x = p
	} else if v, ok := src.(testobj.TestFinance); ok {
		x = &v
	} else {
		return inspector.ErrUnsupportedType
	}

	*result = 0
	if len(path) == 0 {
		return nil
	}
	if path[0] == "History" {
		if len(path) == 1 {
			*result = len(x.History)
			return nil
		}
		if len(path) < 2 {
			return nil
		}
		var i int
		t10, err10 := strconv.ParseInt(path[1], 0, 0)
		if err10 != nil {
			return err10
		}
		i = int(t10)
		if len(x.History) > i {
			x1 := &(x.History)[i]
			_ = x1
			if len(path) < 3 {
				return nil
			}
			if path[2] == "Comment" {
				*result = len(x1.Comment)
				return nil
			}
		}
	}
	return nil
}

func (i0 TestFinanceInspector) Capacity(src any, result *int, path ...string) error {
	if src == nil {
		return nil
	}
	var x *testobj.TestFinance
	_ = x
	if p, ok := src.(**testobj.TestFinance); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestFinance); ok {
		x = p
	} else if v, ok := src.(testobj.TestFinance); ok {
		x = &v
	} else {
		return inspector.ErrUnsupportedType
	}

	*result = 0
	if len(path) == 0 {
		return nil
	}
	if path[0] == "History" {
		if len(path) == 1 {
			*result = cap(x.History)
			return nil
		}
		if len(path) < 2 {
			return nil
		}
		var i int
		t11, err11 := strconv.ParseInt(path[1], 0, 0)
		if err11 != nil {
			return err11
		}
		i = int(t11)
		if len(x.History) > i {
			x1 := &(x.History)[i]
			_ = x1
			if len(path) < 3 {
				return nil
			}
			if path[2] == "Comment" {
				*result = cap(x1.Comment)
				return nil
			}
		}
	}
	return nil
}

func (i0 TestFinanceInspector) Reset(x any) error {
	var origin *testobj.TestFinance
	_ = origin
	switch x.(type) {
	case testobj.TestFinance:
		return inspector.ErrMustPointerType
	case *testobj.TestFinance:
		origin = x.(*testobj.TestFinance)
	case **testobj.TestFinance:
		origin = *x.(**testobj.TestFinance)
	default:
		return inspector.ErrUnsupportedType
	}
	origin.MoneyIn = 0
	origin.MoneyOut = 0
	origin.Balance = 0
	origin.AllowBuy = false
	if l := len((origin.History)); l > 0 {
		_ = (origin.History)[l-1]
		for i := 0; i < l; i++ {
			x1 := &(origin.History)[i]
			x1.DateUnix = 0
			x1.Cost = 0
			if l := len((x1.Comment)); l > 0 {
				(x1.Comment) = (x1.Comment)[:0]
			}
		}
		(origin.History) = (origin.History)[:0]
	}
	return nil
}
