package testobj

type TestPermission map[int32]bool
type TestFlag map[string]int32

type TestHistory struct {
	DateUnix int64
	Cost     float64
	Comment  []byte
}

type TestFinance struct {
	MoneyIn  float64
	MoneyOut float64
	Balance  float64
	AllowBuy bool
	History  []TestHistory
}

type TestObject struct {
	Id          string
	Name        []byte
	Status      int32
	Cost        float64
	Permission  *TestPermission
	HistoryTree map[string]*TestHistory
	Flags       TestFlag
	Finance     *TestFinance
}

func (h *TestHistory) Clear() {
	h.DateUnix, h.Cost = 0, 0
	h.Comment = h.Comment[:0]
}

func (f *TestFinance) Clear() {
	f.MoneyIn, f.MoneyOut, f.Balance, f.AllowBuy = 0, 0, 0, false
	for i := range f.History {
		f.History[i].Clear()
	}
}

func (o *TestObject) Clear() {
	o.Id = ""
	o.Name = o.Name[:0]
	o.Status, o.Cost = 0, 0
	if o.Permission != nil {
		for k := range *o.Permission {
			(*o.Permission)[k] = false
		}
	}
	for k, v := range o.HistoryTree {
		v.Clear()
		o.HistoryTree[k] = v
	}
	for k := range o.Flags {
		o.Flags[k] = 0
	}
	o.Finance.Clear()
}
