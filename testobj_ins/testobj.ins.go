package testobj_ins

import (
	"strconv"

	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
)

type TestFinanceInspector struct {
	inspector.BaseInspector
}

func (i0 *TestFinanceInspector) Get(src, buf interface{}, path ...string) interface{} {
	if len(path) == 0 {
		return nil
	}
	if src == nil {
		return nil
	}
	var x *testobj.TestFinance
	if p, ok := src.(*testobj.TestFinance); ok {
		x = p
	} else if v, ok := src.(testobj.TestFinance); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		if path[0] == "MoneyIn" {
			return x.MoneyIn
		}
		if path[0] == "MoneyOut" {
			return x.MoneyOut
		}
		if path[0] == "Balance" {
			return x.Balance
		}
		if path[0] == "History" {
			x0 := x.History
			_ = x0
			if len(path) > 1 {
				if k, err := i0.StrTo(path[1], "int"); err == nil {
					i := k.(int)
					if len(x0) > i {
						x1 := x0[i]
						_ = x1
						if len(path) > 2 {
							if path[2] == "DateUnix" {
								return x1.DateUnix
							}
							if path[2] == "Cost" {
								return x1.Cost
							}
							if path[2] == "Comment" {
								return x1.Comment
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func (i0 *TestFinanceInspector) Set(dst, value interface{}, path ...string) {
}

type TestFlagInspector struct {
	inspector.BaseInspector
}

func (i1 *TestFlagInspector) Get(src, buf interface{}, path ...string) interface{} {
	if len(path) == 0 {
		return nil
	}
	if src == nil {
		return nil
	}
	var x *testobj.TestFlag
	if p, ok := src.(*testobj.TestFlag); ok {
		x = p
	} else if v, ok := src.(testobj.TestFlag); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		if x0, ok := (*x)[path[0]]; ok {
			_ = x0
			return x0
		}
	}
	return nil
}

func (i1 *TestFlagInspector) Set(dst, value interface{}, path ...string) {
}

type TestHistoryInspector struct {
	inspector.BaseInspector
}

func (i2 *TestHistoryInspector) Get(src, buf interface{}, path ...string) interface{} {
	if len(path) == 0 {
		return nil
	}
	if src == nil {
		return nil
	}
	var x *testobj.TestHistory
	if p, ok := src.(*testobj.TestHistory); ok {
		x = p
	} else if v, ok := src.(testobj.TestHistory); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		if path[0] == "DateUnix" {
			return x.DateUnix
		}
		if path[0] == "Cost" {
			return x.Cost
		}
		if path[0] == "Comment" {
			return x.Comment
		}
	}
	return nil
}

func (i2 *TestHistoryInspector) Set(dst, value interface{}, path ...string) {
}

type TestObjectInspector struct {
	inspector.BaseInspector
}

func (i3 *TestObjectInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var buf interface{}
	err := i3.GetTo(src, &buf, path...)
	return buf, err
}

func (i3 *TestObjectInspector) GetTo(src interface{}, buf *interface{}, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	if src == nil {
		return
	}
	var x *testobj.TestObject
	if p, ok := src.(*testobj.TestObject); ok {
		x = p
	} else if v, ok := src.(testobj.TestObject); ok {
		x = &v
	} else {
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
				j, _ := strconv.Atoi(path[1])
				k := int32(j)
				x1 := (*x0)[k]
				_ = x1
				*buf = &x1
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
				}
			}
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
				if path[1] == "History" {
					x1 := x0.History
					_ = x1
					if len(path) > 2 {
						j, _ := strconv.Atoi(path[2])
						k := int32(j)
						i := k
						if len(x1) > int(i) {
							x2 := x1[i]
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
						}
					}
				}
			}
		}
	}
	return
}

func (i3 *TestObjectInspector) Set(dst, value interface{}, path ...string) {
}

type TestPermissionInspector struct {
	inspector.BaseInspector
}

func (i4 *TestPermissionInspector) Get(src, buf interface{}, path ...string) interface{} {
	if len(path) == 0 {
		return nil
	}
	if src == nil {
		return nil
	}
	var x *testobj.TestPermission
	if p, ok := src.(*testobj.TestPermission); ok {
		x = p
	} else if v, ok := src.(testobj.TestPermission); ok {
		x = &v
	} else {
		return nil
	}

	if len(path) > 0 {
		if k, err := i4.StrTo(path[0], "int32"); err == nil {
			x0 := (*x)[k.(int32)]
			_ = x0
			return x0
		}
	}
	return nil
}

func (i4 *TestPermissionInspector) Set(dst, value interface{}, path ...string) {
}
