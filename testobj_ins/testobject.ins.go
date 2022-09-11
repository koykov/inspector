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

type TestObjectInspector struct {
	inspector.BaseInspector
}

func (i5 TestObjectInspector) TypeName() string {
	return "TestObject"
}

func (i5 TestObjectInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i5.GetTo(src, &buf, path...)
	return buf, err
}

func (i5 TestObjectInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if src == nil {
		return
	}
	var x *testobj.TestObject
	_ = x
	if p, ok := src.(**testobj.TestObject); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestObject); ok {
		x = p
	} else if v, ok := src.(testobj.TestObject); ok {
		x = &v
	} else {
		return
	}
	if len(path) == 0 {
		*buf = &(*x)
		return
	}

	if len(path) > 0 {
		if path[0] == "Id" {
			*buf = &x.Id
			return
		}
		if path[0] == "Name" {
			*buf = &x.Name
			return
		}
		if path[0] == "Status" {
			*buf = &x.Status
			return
		}
		if path[0] == "Ustate" {
			*buf = &x.Ustate
			return
		}
		if path[0] == "Cost" {
			*buf = &x.Cost
			return
		}
		if path[0] == "Permission" {
			x0 := x.Permission
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				var k int32
				t21, err21 := strconv.ParseInt(path[1], 0, 0)
				if err21 != nil {
					return err21
				}
				k = int32(t21)
				x1 := (*x0)[k]
				_ = x1
				*buf = &x1
				return
			}
			*buf = &x.Permission
			return
		}
		if path[0] == "HistoryTree" {
			x0 := x.HistoryTree
			_ = x0
			if len(path) > 1 {
				if x1, ok := (x0)[path[1]]; ok {
					_ = x1
					if len(path) > 2 {
						if x1 == nil {
							return
						}
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
					*buf = &x1
				}
			}
			*buf = &x.HistoryTree
			return
		}
		if path[0] == "Flags" {
			x0 := x.Flags
			_ = x0
			if len(path) > 1 {
				if x1, ok := (x0)[path[1]]; ok {
					_ = x1
					*buf = &x1
					return
				}
			}
			*buf = &x.Flags
			return
		}
		if path[0] == "Finance" {
			x0 := x.Finance
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				if path[1] == "MoneyIn" {
					*buf = &x0.MoneyIn
					return
				}
				if path[1] == "MoneyOut" {
					*buf = &x0.MoneyOut
					return
				}
				if path[1] == "Balance" {
					*buf = &x0.Balance
					return
				}
				if path[1] == "AllowBuy" {
					*buf = &x0.AllowBuy
					return
				}
				if path[1] == "History" {
					x1 := x0.History
					_ = x1
					if len(path) > 2 {
						var i int
						t22, err22 := strconv.ParseInt(path[2], 0, 0)
						if err22 != nil {
							return err22
						}
						i = int(t22)
						if len(x1) > i {
							x2 := &(x1)[i]
							_ = x2
							if len(path) > 3 {
								if path[3] == "DateUnix" {
									*buf = &x2.DateUnix
									return
								}
								if path[3] == "Cost" {
									*buf = &x2.Cost
									return
								}
								if path[3] == "Comment" {
									*buf = &x2.Comment
									return
								}
							}
							*buf = x2
						}
					}
					*buf = &x0.History
					return
				}
			}
			*buf = &x.Finance
			return
		}
	}
	return
}

func (i5 TestObjectInspector) Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestObject
	_ = x
	if p, ok := src.(**testobj.TestObject); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestObject); ok {
		x = p
	} else if v, ok := src.(testobj.TestObject); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "Id" {
			var rightExact string
			rightExact = right

			switch cond {
			case inspector.OpEq:
				*result = x.Id == rightExact
			case inspector.OpNq:
				*result = x.Id != rightExact
			case inspector.OpGt:
				*result = x.Id > rightExact
			case inspector.OpGtq:
				*result = x.Id >= rightExact
			case inspector.OpLt:
				*result = x.Id < rightExact
			case inspector.OpLtq:
				*result = x.Id <= rightExact
			}
			return
		}
		if path[0] == "Name" {
			var rightExact []byte
			rightExact = fastconv.S2B(right)

			if cond == inspector.OpEq {
				*result = bytes.Equal(x.Name, rightExact)
			} else {
				*result = !bytes.Equal(x.Name, rightExact)
			}
			return
		}
		if path[0] == "Status" {
			var rightExact int32
			t25, err25 := strconv.ParseInt(right, 0, 0)
			if err25 != nil {
				return err25
			}
			rightExact = int32(t25)
			switch cond {
			case inspector.OpEq:
				*result = x.Status == rightExact
			case inspector.OpNq:
				*result = x.Status != rightExact
			case inspector.OpGt:
				*result = x.Status > rightExact
			case inspector.OpGtq:
				*result = x.Status >= rightExact
			case inspector.OpLt:
				*result = x.Status < rightExact
			case inspector.OpLtq:
				*result = x.Status <= rightExact
			}
			return
		}
		if path[0] == "Ustate" {
			var rightExact uint64
			t26, err26 := strconv.ParseUint(right, 0, 0)
			if err26 != nil {
				return err26
			}
			rightExact = uint64(t26)
			switch cond {
			case inspector.OpEq:
				*result = x.Ustate == rightExact
			case inspector.OpNq:
				*result = x.Ustate != rightExact
			case inspector.OpGt:
				*result = x.Ustate > rightExact
			case inspector.OpGtq:
				*result = x.Ustate >= rightExact
			case inspector.OpLt:
				*result = x.Ustate < rightExact
			case inspector.OpLtq:
				*result = x.Ustate <= rightExact
			}
			return
		}
		if path[0] == "Cost" {
			var rightExact float64
			t27, err27 := strconv.ParseFloat(right, 0)
			if err27 != nil {
				return err27
			}
			rightExact = float64(t27)
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
		if path[0] == "Permission" {
			x0 := x.Permission
			_ = x0
			if right == inspector.Nil {
				if cond == inspector.OpEq {
					*result = x0 == nil
				} else {
					*result = x0 != nil
				}
				return
			}
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				var k int32
				t28, err28 := strconv.ParseInt(path[1], 0, 0)
				if err28 != nil {
					return err28
				}
				k = int32(t28)
				x1 := (*x0)[k]
				_ = x1
				var rightExact bool
				t29, err29 := strconv.ParseBool(right)
				if err29 != nil {
					return err29
				}
				rightExact = bool(t29)
				if cond == inspector.OpEq {
					*result = x1 == rightExact
				} else {
					*result = x1 != rightExact
				}
				return
			}
		}
		if path[0] == "HistoryTree" {
			x0 := x.HistoryTree
			_ = x0
			if len(path) > 1 {
				if x1, ok := (x0)[path[1]]; ok {
					_ = x1
					if len(path) > 2 {
						if x1 == nil {
							return
						}
						if path[2] == "DateUnix" {
							var rightExact int64
							t30, err30 := strconv.ParseInt(right, 0, 0)
							if err30 != nil {
								return err30
							}
							rightExact = int64(t30)
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
							t31, err31 := strconv.ParseFloat(right, 0)
							if err31 != nil {
								return err31
							}
							rightExact = float64(t31)
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
		if path[0] == "Flags" {
			x0 := x.Flags
			_ = x0
			if len(path) > 1 {
				if x1, ok := (x0)[path[1]]; ok {
					_ = x1
					var rightExact int32
					t33, err33 := strconv.ParseInt(right, 0, 0)
					if err33 != nil {
						return err33
					}
					rightExact = int32(t33)
					switch cond {
					case inspector.OpEq:
						*result = x1 == rightExact
					case inspector.OpNq:
						*result = x1 != rightExact
					case inspector.OpGt:
						*result = x1 > rightExact
					case inspector.OpGtq:
						*result = x1 >= rightExact
					case inspector.OpLt:
						*result = x1 < rightExact
					case inspector.OpLtq:
						*result = x1 <= rightExact
					}
					return
				}
			}
		}
		if path[0] == "Finance" {
			x0 := x.Finance
			_ = x0
			if right == inspector.Nil {
				if cond == inspector.OpEq {
					*result = x0 == nil
				} else {
					*result = x0 != nil
				}
				return
			}
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				if path[1] == "MoneyIn" {
					var rightExact float64
					t34, err34 := strconv.ParseFloat(right, 0)
					if err34 != nil {
						return err34
					}
					rightExact = float64(t34)
					switch cond {
					case inspector.OpEq:
						*result = x0.MoneyIn == rightExact
					case inspector.OpNq:
						*result = x0.MoneyIn != rightExact
					case inspector.OpGt:
						*result = x0.MoneyIn > rightExact
					case inspector.OpGtq:
						*result = x0.MoneyIn >= rightExact
					case inspector.OpLt:
						*result = x0.MoneyIn < rightExact
					case inspector.OpLtq:
						*result = x0.MoneyIn <= rightExact
					}
					return
				}
				if path[1] == "MoneyOut" {
					var rightExact float64
					t35, err35 := strconv.ParseFloat(right, 0)
					if err35 != nil {
						return err35
					}
					rightExact = float64(t35)
					switch cond {
					case inspector.OpEq:
						*result = x0.MoneyOut == rightExact
					case inspector.OpNq:
						*result = x0.MoneyOut != rightExact
					case inspector.OpGt:
						*result = x0.MoneyOut > rightExact
					case inspector.OpGtq:
						*result = x0.MoneyOut >= rightExact
					case inspector.OpLt:
						*result = x0.MoneyOut < rightExact
					case inspector.OpLtq:
						*result = x0.MoneyOut <= rightExact
					}
					return
				}
				if path[1] == "Balance" {
					var rightExact float64
					t36, err36 := strconv.ParseFloat(right, 0)
					if err36 != nil {
						return err36
					}
					rightExact = float64(t36)
					switch cond {
					case inspector.OpEq:
						*result = x0.Balance == rightExact
					case inspector.OpNq:
						*result = x0.Balance != rightExact
					case inspector.OpGt:
						*result = x0.Balance > rightExact
					case inspector.OpGtq:
						*result = x0.Balance >= rightExact
					case inspector.OpLt:
						*result = x0.Balance < rightExact
					case inspector.OpLtq:
						*result = x0.Balance <= rightExact
					}
					return
				}
				if path[1] == "AllowBuy" {
					var rightExact bool
					t37, err37 := strconv.ParseBool(right)
					if err37 != nil {
						return err37
					}
					rightExact = bool(t37)
					if cond == inspector.OpEq {
						*result = x0.AllowBuy == rightExact
					} else {
						*result = x0.AllowBuy != rightExact
					}
					return
				}
				if path[1] == "History" {
					x1 := x0.History
					_ = x1
					if len(path) > 2 {
						var i int
						t38, err38 := strconv.ParseInt(path[2], 0, 0)
						if err38 != nil {
							return err38
						}
						i = int(t38)
						if len(x1) > i {
							x2 := &(x1)[i]
							_ = x2
							if len(path) > 3 {
								if path[3] == "DateUnix" {
									var rightExact int64
									t39, err39 := strconv.ParseInt(right, 0, 0)
									if err39 != nil {
										return err39
									}
									rightExact = int64(t39)
									switch cond {
									case inspector.OpEq:
										*result = x2.DateUnix == rightExact
									case inspector.OpNq:
										*result = x2.DateUnix != rightExact
									case inspector.OpGt:
										*result = x2.DateUnix > rightExact
									case inspector.OpGtq:
										*result = x2.DateUnix >= rightExact
									case inspector.OpLt:
										*result = x2.DateUnix < rightExact
									case inspector.OpLtq:
										*result = x2.DateUnix <= rightExact
									}
									return
								}
								if path[3] == "Cost" {
									var rightExact float64
									t40, err40 := strconv.ParseFloat(right, 0)
									if err40 != nil {
										return err40
									}
									rightExact = float64(t40)
									switch cond {
									case inspector.OpEq:
										*result = x2.Cost == rightExact
									case inspector.OpNq:
										*result = x2.Cost != rightExact
									case inspector.OpGt:
										*result = x2.Cost > rightExact
									case inspector.OpGtq:
										*result = x2.Cost >= rightExact
									case inspector.OpLt:
										*result = x2.Cost < rightExact
									case inspector.OpLtq:
										*result = x2.Cost <= rightExact
									}
									return
								}
								if path[3] == "Comment" {
									var rightExact []byte
									rightExact = fastconv.S2B(right)

									if cond == inspector.OpEq {
										*result = bytes.Equal(x2.Comment, rightExact)
									} else {
										*result = !bytes.Equal(x2.Comment, rightExact)
									}
									return
								}
							}
						}
					}
				}
			}
		}
	}
	return
}

func (i5 TestObjectInspector) Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestObject
	_ = x
	if p, ok := src.(**testobj.TestObject); ok {
		x = *p
	} else if p, ok := src.(*testobj.TestObject); ok {
		x = p
	} else if v, ok := src.(testobj.TestObject); ok {
		x = &v
	} else {
		return
	}

	if len(path) > 0 {
		if path[0] == "Permission" {
			x0 := x.Permission
			_ = x0
			if x0 == nil {
				return
			}
			for k := range *x0 {
				if l.RequireKey() {
					*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)
					l.SetKey(buf, &inspector.StaticInspector{})
				}
				l.SetVal((*x0)[k], &inspector.StaticInspector{})
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
		if path[0] == "HistoryTree" {
			x0 := x.HistoryTree
			_ = x0
			for k := range x0 {
				if l.RequireKey() {
					*buf = append((*buf)[:0], k...)
					l.SetKey(buf, &inspector.StaticInspector{})
				}
				l.SetVal((x0)[k], &TestHistoryInspector{})
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
		if path[0] == "Flags" {
			x0 := x.Flags
			_ = x0
			for k := range x0 {
				if l.RequireKey() {
					*buf = append((*buf)[:0], k...)
					l.SetKey(buf, &inspector.StaticInspector{})
				}
				l.SetVal((x0)[k], &inspector.StaticInspector{})
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
		if path[0] == "Finance" {
			x0 := x.Finance
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return
				}
				if path[1] == "History" {
					x1 := x0.History
					_ = x1
					for k := range x1 {
						if l.RequireKey() {
							*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)
							l.SetKey(buf, &inspector.StaticInspector{})
						}
						l.SetVal(&(x1)[k], &TestHistoryInspector{})
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
		}
	}
	return
}

func (i5 TestObjectInspector) SetWB(dst, value interface{}, buf inspector.AccumulativeBuffer, path ...string) error {
	if len(path) == 0 {
		return nil
	}
	if dst == nil {
		return nil
	}
	var x *testobj.TestObject
	_ = x
	if p, ok := dst.(**testobj.TestObject); ok {
		x = *p
	} else if p, ok := dst.(*testobj.TestObject); ok {
		x = p
	} else if v, ok := dst.(testobj.TestObject); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		if path[0] == "Id" {
			inspector.AssignBuf(&x.Id, value, buf)
			return nil
		}
		if path[0] == "Name" {
			inspector.AssignBuf(&x.Name, value, buf)
			return nil
		}
		if path[0] == "Status" {
			inspector.AssignBuf(&x.Status, value, buf)
			return nil
		}
		if path[0] == "Ustate" {
			inspector.AssignBuf(&x.Ustate, value, buf)
			return nil
		}
		if path[0] == "Cost" {
			inspector.AssignBuf(&x.Cost, value, buf)
			return nil
		}
		if path[0] == "Permission" {
			x0 := x.Permission
			if uvalue, ok := value.(*testobj.TestPermission); ok {
				x0 = uvalue
			}
			if x0 == nil {
				z := make(testobj.TestPermission)
				x0 = &z
				x.Permission = x0
			}
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return nil
				}
				var k int32
				t42, err42 := strconv.ParseInt(path[1], 0, 0)
				if err42 != nil {
					return err42
				}
				k = int32(t42)
				x1 := (*x0)[k]
				_ = x1
				inspector.AssignBuf(&x1, value, buf)
				(*x0)[k] = x1
				return nil
			}
			x.Permission = x0
		}
		if path[0] == "HistoryTree" {
			x0 := x.HistoryTree
			if uvalue, ok := value.(*map[string]*testobj.TestHistory); ok {
				x0 = *uvalue
			}
			if x0 == nil {
				z := make(map[string]*testobj.TestHistory)
				x0 = z
				x.HistoryTree = x0
			}
			_ = x0
			if len(path) > 1 {
				x1 := (x0)[path[1]]
				_ = x1
				if len(path) > 2 {
					if x1 == nil {
						return nil
					}
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
				(x0)[path[1]] = x1
				return nil
			}
			x.HistoryTree = x0
		}
		if path[0] == "Flags" {
			x0 := x.Flags
			if uvalue, ok := value.(*testobj.TestFlag); ok {
				x0 = *uvalue
			}
			if x0 == nil {
				z := make(testobj.TestFlag)
				x0 = z
				x.Flags = x0
			}
			_ = x0
			if len(path) > 1 {
				x1 := (x0)[path[1]]
				_ = x1
				inspector.AssignBuf(&x1, value, buf)
				(x0)[path[1]] = x1
				return nil
			}
			x.Flags = x0
		}
		if path[0] == "Finance" {
			x0 := x.Finance
			if uvalue, ok := value.(*testobj.TestFinance); ok {
				x0 = uvalue
			}
			if x0 == nil {
				x0 = &testobj.TestFinance{}
				x.Finance = x0
			}
			_ = x0
			if len(path) > 1 {
				if x0 == nil {
					return nil
				}
				if path[1] == "MoneyIn" {
					inspector.AssignBuf(&x0.MoneyIn, value, buf)
					return nil
				}
				if path[1] == "MoneyOut" {
					inspector.AssignBuf(&x0.MoneyOut, value, buf)
					return nil
				}
				if path[1] == "Balance" {
					inspector.AssignBuf(&x0.Balance, value, buf)
					return nil
				}
				if path[1] == "AllowBuy" {
					inspector.AssignBuf(&x0.AllowBuy, value, buf)
					return nil
				}
				if path[1] == "History" {
					x1 := x0.History
					if uvalue, ok := value.(*[]testobj.TestHistory); ok {
						x1 = *uvalue
					}
					if x1 == nil {
						z := make([]testobj.TestHistory, 0)
						x1 = z
						x0.History = x1
					}
					_ = x1
					if len(path) > 2 {
						var i int
						t43, err43 := strconv.ParseInt(path[2], 0, 0)
						if err43 != nil {
							return err43
						}
						i = int(t43)
						if len(x1) > i {
							x2 := &(x1)[i]
							_ = x2
							if len(path) > 3 {
								if path[3] == "DateUnix" {
									inspector.AssignBuf(&x2.DateUnix, value, buf)
									return nil
								}
								if path[3] == "Cost" {
									inspector.AssignBuf(&x2.Cost, value, buf)
									return nil
								}
								if path[3] == "Comment" {
									inspector.AssignBuf(&x2.Comment, value, buf)
									return nil
								}
							}
							(x1)[i] = *x2
							return nil
						}
					}
					x0.History = x1
				}
			}
			x.Finance = x0
		}
	}
	return nil
}

func (i5 TestObjectInspector) Set(dst, value interface{}, path ...string) error {
	return i5.SetWB(dst, value, nil, path...)
}

func (i5 TestObjectInspector) DeepEqual(l, r interface{}) bool {
	return i5.DeepEqualWithOptions(l, r, nil)
}

func (i5 TestObjectInspector) DeepEqualWithOptions(l, r interface{}, opts *inspector.DEQOptions) bool {
	var (
		lx, rx   *testobj.TestObject
		leq, req bool
	)
	_, _, _, _ = lx, rx, leq, req
	if lp, ok := l.(**testobj.TestObject); ok {
		lx, leq = *lp, true
	} else if lp, ok := l.(*testobj.TestObject); ok {
		lx, leq = lp, true
	} else if lp, ok := l.(testobj.TestObject); ok {
		lx, leq = &lp, true
	}
	if rp, ok := r.(**testobj.TestObject); ok {
		rx, req = *rp, true
	} else if rp, ok := r.(*testobj.TestObject); ok {
		rx, req = rp, true
	} else if rp, ok := r.(testobj.TestObject); ok {
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

	if lx.Id != rx.Id && inspector.DEQMustCheck("Id", opts) {
		return false
	}
	if !bytes.Equal(lx.Name, rx.Name) && inspector.DEQMustCheck("Name", opts) {
		return false
	}
	if lx.Status != rx.Status && inspector.DEQMustCheck("Status", opts) {
		return false
	}
	if lx.Ustate != rx.Ustate && inspector.DEQMustCheck("Ustate", opts) {
		return false
	}
	if !inspector.EqualFloat64(lx.Cost, rx.Cost, opts) && inspector.DEQMustCheck("Cost", opts) {
		return false
	}
	lx1 := lx.Permission
	rx1 := rx.Permission
	_, _ = lx1, rx1
	if (lx1 == nil && rx1 != nil) || (lx1 != nil && rx1 == nil) {
		return false
	}
	if lx1 != nil && rx1 != nil {
		if inspector.DEQMustCheck("Permission", opts) {
			if len(*lx1) != len(*rx1) {
				return false
			}
			for k := range *lx1 {
				lx2 := (*lx1)[k]
				rx2, ok2 := (*rx1)[k]
				_, _, _ = lx2, rx2, ok2
				if !ok2 {
					return false
				}
				if lx2 != rx2 {
					return false
				}
			}
		}
	}
	lx3 := lx.HistoryTree
	rx3 := rx.HistoryTree
	_, _ = lx3, rx3
	if inspector.DEQMustCheck("HistoryTree", opts) {
		if len(lx3) != len(rx3) {
			return false
		}
		for k := range lx3 {
			lx4 := (lx3)[k]
			rx4, ok4 := (rx3)[k]
			_, _, _ = lx4, rx4, ok4
			if !ok4 {
				return false
			}
			if (lx4 == nil && rx4 != nil) || (lx4 != nil && rx4 == nil) {
				return false
			}
			if lx4 != nil && rx4 != nil {
				if lx4.DateUnix != rx4.DateUnix && inspector.DEQMustCheck("HistoryTree.DateUnix", opts) {
					return false
				}
				if !inspector.EqualFloat64(lx4.Cost, rx4.Cost, opts) && inspector.DEQMustCheck("HistoryTree.Cost", opts) {
					return false
				}
				if !bytes.Equal(lx4.Comment, rx4.Comment) && inspector.DEQMustCheck("HistoryTree.Comment", opts) {
					return false
				}
			}
		}
	}
	lx5 := lx.Flags
	rx5 := rx.Flags
	_, _ = lx5, rx5
	if inspector.DEQMustCheck("Flags", opts) {
		if len(lx5) != len(rx5) {
			return false
		}
		for k := range lx5 {
			lx6 := (lx5)[k]
			rx6, ok6 := (rx5)[k]
			_, _, _ = lx6, rx6, ok6
			if !ok6 {
				return false
			}
			if lx6 != rx6 {
				return false
			}
		}
	}
	lx7 := lx.Finance
	rx7 := rx.Finance
	_, _ = lx7, rx7
	if (lx7 == nil && rx7 != nil) || (lx7 != nil && rx7 == nil) {
		return false
	}
	if lx7 != nil && rx7 != nil {
		if inspector.DEQMustCheck("Finance", opts) {
			if !inspector.EqualFloat64(lx7.MoneyIn, rx7.MoneyIn, opts) && inspector.DEQMustCheck("Finance.MoneyIn", opts) {
				return false
			}
			if !inspector.EqualFloat64(lx7.MoneyOut, rx7.MoneyOut, opts) && inspector.DEQMustCheck("Finance.MoneyOut", opts) {
				return false
			}
			if !inspector.EqualFloat64(lx7.Balance, rx7.Balance, opts) && inspector.DEQMustCheck("Finance.Balance", opts) {
				return false
			}
			if lx7.AllowBuy != rx7.AllowBuy && inspector.DEQMustCheck("Finance.AllowBuy", opts) {
				return false
			}
			lx8 := lx7.History
			rx8 := rx7.History
			_, _ = lx8, rx8
			if inspector.DEQMustCheck("Finance.History", opts) {
				if len(lx8) != len(rx8) {
					return false
				}
				for i := 0; i < len(lx8); i++ {
					lx9 := (lx8)[i]
					rx9 := (rx8)[i]
					_, _ = lx9, rx9
					if lx9.DateUnix != rx9.DateUnix && inspector.DEQMustCheck("Finance.History.DateUnix", opts) {
						return false
					}
					if !inspector.EqualFloat64(lx9.Cost, rx9.Cost, opts) && inspector.DEQMustCheck("Finance.History.Cost", opts) {
						return false
					}
					if !bytes.Equal(lx9.Comment, rx9.Comment) && inspector.DEQMustCheck("Finance.History.Comment", opts) {
						return false
					}
				}
			}
		}
	}
	return true
}

func (i5 TestObjectInspector) Unmarshal(p []byte, typ inspector.Encoding) (interface{}, error) {
	var x testobj.TestObject
	switch typ {
	case inspector.EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, inspector.ErrUnknownEncodingType
	}
}

func (i5 TestObjectInspector) Copy(x interface{}) (interface{}, error) {
	var origin, cpy testobj.TestObject
	switch x.(type) {
	case testobj.TestObject:
		origin = x.(testobj.TestObject)
	case *testobj.TestObject:
		origin = *x.(*testobj.TestObject)
	case **testobj.TestObject:
		origin = **x.(**testobj.TestObject)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	bc := i5.calcBytes(&origin)
	buf1 := make([]byte, 0, bc)
	if err := i5.cpy(buf1, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i5 TestObjectInspector) CopyWB(x interface{}, buf inspector.AccumulativeBuffer) (interface{}, error) {
	var origin, cpy testobj.TestObject
	switch x.(type) {
	case testobj.TestObject:
		origin = x.(testobj.TestObject)
	case *testobj.TestObject:
		origin = *x.(*testobj.TestObject)
	case **testobj.TestObject:
		origin = **x.(**testobj.TestObject)
	default:
		return nil, inspector.ErrUnsupportedType
	}
	buf1 := buf.AcquireBytes()
	defer buf.ReleaseBytes(buf1)
	if err := i5.cpy(buf1, &cpy, &origin); err != nil {
		return nil, err
	}
	return cpy, nil
}

func (i5 TestObjectInspector) calcBytes(x *testobj.TestObject) (c int) {
	c += len(x.Id)
	c += len(x.Name)
	for k1, v1 := range x.HistoryTree {
		_, _ = k1, v1
		c += len(k1)
		c += len(v1.Comment)
	}
	for k1, v1 := range x.Flags {
		_, _ = k1, v1
		c += len(k1)
	}
	if x.Finance != nil {
		for i2 := 0; i2 < len(x.Finance.History); i2++ {
			x2 := &(x.Finance.History)[i2]
			c += len(x2.Comment)
		}
	}
	return c
}

func (i5 TestObjectInspector) cpy(buf []byte, l, r *testobj.TestObject) error {
	buf, l.Id = inspector.BufferizeString(buf, r.Id)
	buf, l.Name = inspector.Bufferize(buf, r.Name)
	l.Status = r.Status
	l.Ustate = r.Ustate
	l.Cost = r.Cost
	if r.Permission != nil {
		if len(*r.Permission) > 0 {
			buf1 := make(testobj.TestPermission, len(*r.Permission))
			_ = buf1
			for rk1, rv1 := range *r.Permission {
				_, _ = rk1, rv1
				var lk1 int32
				lk1 = rk1
				var lv1 bool
				lv1 = rv1
				(*l.Permission)[lk1] = lv1
			}
		}
	}
	if len(r.HistoryTree) > 0 {
		buf1 := make(map[string]*testobj.TestHistory, len(r.HistoryTree))
		_ = buf1
		for rk1, rv1 := range r.HistoryTree {
			_, _ = rk1, rv1
			var lk1 string
			buf, lk1 = inspector.BufferizeString(buf, rk1)
			var lv1 testobj.TestHistory
			lv1.DateUnix = rv1.DateUnix
			lv1.Cost = rv1.Cost
			buf, lv1.Comment = inspector.Bufferize(buf, rv1.Comment)
			(l.HistoryTree)[lk1] = &lv1
		}
	}
	if len(r.Flags) > 0 {
		buf1 := make(testobj.TestFlag, len(r.Flags))
		_ = buf1
		for rk1, rv1 := range r.Flags {
			_, _ = rk1, rv1
			var lk1 string
			buf, lk1 = inspector.BufferizeString(buf, rk1)
			var lv1 int32
			lv1 = rv1
			(l.Flags)[lk1] = lv1
		}
	}
	if r.Finance != nil {
		l.Finance.MoneyIn = r.Finance.MoneyIn
		l.Finance.MoneyOut = r.Finance.MoneyOut
		l.Finance.Balance = r.Finance.Balance
		l.Finance.AllowBuy = r.Finance.AllowBuy
		if len(r.Finance.History) > 0 {
			buf2 := make([]testobj.TestHistory, 0, len(r.Finance.History))
			for i2 := 0; i2 < len(r.Finance.History); i2++ {
				var b2 testobj.TestHistory
				x2 := &(l.Finance.History)[i2]
				b2.DateUnix = x2.DateUnix
				b2.Cost = x2.Cost
				buf, b2.Comment = inspector.Bufferize(buf, x2.Comment)
				buf2 = append(buf2, b2)
			}
			l.Finance.History = buf2
		}
	}
	return nil
}

func (i5 TestObjectInspector) Reset(x interface{}) {
	var origin testobj.TestObject
	_ = origin
	switch x.(type) {
	case testobj.TestObject:
		origin = x.(testobj.TestObject)
	case *testobj.TestObject:
		origin = *x.(*testobj.TestObject)
	case **testobj.TestObject:
		origin = **x.(**testobj.TestObject)
	default:
		return
	}
	origin.Id = ""
	if l := len((origin.Name)); l > 0 {
		(origin.Name) = (origin.Name)[:0]
	}
	origin.Status = 0
	origin.Ustate = 0
	origin.Cost = 0
	if origin.Permission != nil {
		if l := len((*origin.Permission)); l > 0 {
			for k, _ := range *origin.Permission {
				delete((*origin.Permission), k)
			}
		}
	}
	if l := len((origin.HistoryTree)); l > 0 {
		for k, _ := range origin.HistoryTree {
			delete((origin.HistoryTree), k)
		}
	}
	if l := len((origin.Flags)); l > 0 {
		for k, _ := range origin.Flags {
			delete((origin.Flags), k)
		}
	}
	if origin.Finance != nil {
		origin.Finance.MoneyIn = 0
		origin.Finance.MoneyOut = 0
		origin.Finance.Balance = 0
		origin.Finance.AllowBuy = false
		if l := len((origin.Finance.History)); l > 0 {
			_ = (origin.Finance.History)[l-1]
			for i := 0; i < l; i++ {
				x2 := &(origin.Finance.History)[i]
				x2.DateUnix = 0
				x2.Cost = 0
				if l := len((x2.Comment)); l > 0 {
					(x2.Comment) = (x2.Comment)[:0]
				}
			}
			(origin.Finance.History) = (origin.Finance.History)[:0]
		}
	}
}
