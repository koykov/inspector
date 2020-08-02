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

type TestFinanceInspector struct {
	inspector.BaseInspector
}

func (i0 *TestFinanceInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i0.GetTo(src, &buf, path...)
	return buf, err
}

func (i0 *TestFinanceInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
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
	*buf = &(*x)
	return
}

func (i0 *TestFinanceInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
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
							rightExact = fastconv.S2B(right)

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

func (i0 *TestFinanceInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
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

func (i0 *TestFinanceInspector) Set(dst, value interface{}, path ...string) error {
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
			inspector.Assign(&x.MoneyIn, value)
			return nil
		}
		if path[0] == "MoneyOut" {
			inspector.Assign(&x.MoneyOut, value)
			return nil
		}
		if path[0] == "Balance" {
			inspector.Assign(&x.Balance, value)
			return nil
		}
		if path[0] == "AllowBuy" {
			inspector.Assign(&x.AllowBuy, value)
			return nil
		}
		if path[0] == "History" {
			x0 := x.History
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
							inspector.Assign(&x1.DateUnix, value)
							return nil
						}
						if path[2] == "Cost" {
							inspector.Assign(&x1.Cost, value)
							return nil
						}
						if path[2] == "Comment" {
							inspector.Assign(&x1.Comment, value)
							return nil
						}
					}
					(x0)[i] = *x1
					return nil
				}
			}
		}
	}
	return nil
}
