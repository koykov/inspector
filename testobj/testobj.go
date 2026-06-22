package testobj

type TestPermission map[int32]bool
type TestFlag map[string]int32

type TestHistory struct {
	DateUnix int64   `json:"date_unix"`
	Cost     float64 `json:"cost"`
	Comment  []byte  `json:"comment"`
}

type TestFinance struct {
	MoneyIn  float64       `json:"money_in"`
	MoneyOut float64       `json:"money_out"`
	Balance  float64       `json:"balance"`
	AllowBuy bool          `json:"allow_buy"`
	History  []TestHistory `json:"history"`
}

type TestObject struct {
	Id          string                  `json:"id"`
	Name        []byte                  `json:"name"`
	Status      int32                   `json:"status"`
	Ustate      uint64                  `json:"ustate"`
	Cost        float64                 `json:"cost"`
	Permission  *TestPermission         `json:"permission,omitempty"`
	HistoryTree map[string]*TestHistory `json:"history_tree"`
	Flags       TestFlag                `json:"flags"`
	Finance     *TestFinance            `json:"finance,omitempty"`
}

func (h *TestHistory) Clear() {
	if h == nil {
		return
	}
	h.DateUnix, h.Cost = 0, 0
	h.Comment = h.Comment[:0]
}

func (f *TestFinance) Clear() {
	if f == nil {
		return
	}
	f.MoneyIn, f.MoneyOut, f.Balance, f.AllowBuy = 0, 0, 0, false
	for i := range f.History {
		f.History[i].Clear()
	}
	f.History = f.History[:0]
}

func (o *TestObject) Clear() {
	if o == nil {
		return
	}
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
