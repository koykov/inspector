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
