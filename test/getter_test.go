package test

import (
	"bytes"
	"testing"

	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"github.com/koykov/inspector/testobj_ins"
)

var (
	testO = &testobj.TestObject{
		Id:         "foo",
		Name:       []byte("bar"),
		Cost:       12.34,
		Status:     78,
		Permission: &testobj.TestPermission{15: true, 23: false},
		Flags: testobj.TestFlag{
			"export": 17,
			"ro":     4,
			"rw":     7,
			"Valid":  1,
		},
		Finance: &testobj.TestFinance{
			MoneyIn:  3200,
			MoneyOut: 1500.637657,
			Balance:  9000,
			AllowBuy: true,
			History: []testobj.TestHistory{
				{
					152354345634,
					14.345241,
					[]byte("pay for domain"),
				},
				{
					153465345246,
					-3.0000342543,
					[]byte("got refund"),
				},
				{
					156436535640,
					2325242534.35324523,
					[]byte("maintenance"),
				},
			},
		},
	}

	p0 = []string{"Id"}
	p1 = []string{"Name"}
	p2 = []string{"Permission", "23"}
	p3 = []string{"Flags", "export"}
	p4 = []string{"Finance", "Balance"}
	p5 = []string{"Finance", "History", "1", "DateUnix"}
	p6 = []string{"Finance", "History", "0", "Comment"}
	p7 = []string{"Status"}
	p8 = []string{"Finance", "MoneyIn"}

	expectFoo     = []byte("bar")
	expectComment = []byte("pay for domain")
)

func testGetter(t testing.TB, i inspector.Inspector) {
	id, _ := i.Get(testO, p0...)
	if id.(string) != "foo" {
		t.Error("object.Id: mismatch result and expectation")
	}

	name, _ := i.Get(testO, p1...)
	if !bytes.Equal(name.([]byte), expectFoo) {
		t.Error("object.Name: mismatch result and expectation")
	}

	perm, _ := i.Get(testO, p2...)
	if perm.(bool) != false {
		t.Error("object.Permission.23: mismatch result and expectation")
	}

	flag, _ := i.Get(testO, p3...)
	if flag.(int32) != 17 {
		t.Error("object.Flags.export: mismatch result and expectation")
	}

	bal, _ := i.Get(testO, p4...)
	if bal.(float64) != 9000 {
		t.Error("object.Finance.Balance: mismatch result and expectation")
	}

	date, _ := i.Get(testO, p5...)
	if date.(int64) != 153465345246 {
		t.Error("object.Finance.History.1.DateUnix: mismatch result and expectation")
	}

	comment, _ := i.Get(testO, p6...)
	if !bytes.Equal(comment.([]byte), expectComment) {
		t.Error("object.Finance.History.0.DateUnix: mismatch result and expectation")
	}
}

func testGetterPtr(t testing.TB, i inspector.Inspector, buf interface{}) {
	_ = i.GetTo(testO, &buf, p0...)
	if *buf.(*string) != "foo" {
		t.Error("object.Id: mismatch result and expectation")
	}

	_ = i.GetTo(testO, &buf, p1...)
	if !bytes.Equal(*buf.(*[]byte), expectFoo) {
		t.Error("object.Name: mismatch result and expectation")
	}

	_ = i.GetTo(testO, &buf, p2...)
	if *buf.(*bool) != false {
		t.Error("object.Permission.23: mismatch result and expectation")
	}

	_ = i.GetTo(testO, &buf, p3...)
	if *buf.(*int32) != 17 {
		t.Error("object.Flags.export: mismatch result and expectation")
	}

	_ = i.GetTo(testO, &buf, p4...)
	if *buf.(*float64) != 9000 {
		t.Error("object.Finance.Balance: mismatch result and expectation")
	}

	_ = i.GetTo(testO, &buf, p5...)
	if *buf.(*int64) != 153465345246 {
		t.Error("object.Finance.History.1.DateUnix: mismatch result and expectation")
	}

	_ = i.GetTo(testO, &buf, p6...)
	if !bytes.Equal(*buf.(*[]byte), expectComment) {
		t.Error("object.Finance.History.0.Comment: mismatch result and expectation")
	}
}

func testCmpPtr(t testing.TB, i inspector.Inspector, buf *bool) {
	_ = i.Cmp(testO, inspector.OpEq, "foo", buf, p0...)
	if !*buf {
		t.Error("object.Id: mismatch result and expectation")
	}

	_ = i.Cmp(testO, inspector.OpEq, "bar", buf, p1...)
	if !*buf {
		t.Error("object.Name: mismatch result and expectation")
	}

	_ = i.Cmp(testO, inspector.OpGtq, "60", buf, p7...)
	if !*buf {
		t.Error("object.Status: mismatch result and expectation")
	}

	_ = i.Cmp(testO, inspector.OpLtq, "5000", buf, p8...)
	if !*buf {
		t.Error("object.Finance.MoneyIn: mismatch result and expectation")
	}
}

func TestReflectInspector_Get(t *testing.T) {
	ins := &inspector.ReflectInspector{}
	testGetter(t, ins)
}

func BenchmarkReflectInspector_Get(b *testing.B) {
	ins := &inspector.ReflectInspector{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testGetter(b, ins)
	}
}

func TestCompiledInspector_Get(t *testing.T) {
	var buf interface{}
	ins := &testobj_ins.TestObjectInspector{}
	testGetterPtr(t, ins, buf)
}

func BenchmarkCompiledInspector_Get(b *testing.B) {
	ins := &testobj_ins.TestObjectInspector{}
	var buf interface{}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testGetterPtr(b, ins, buf)
	}
}

func TestCompiledInspector_Cmp(t *testing.T) {
	ins := &testobj_ins.TestObjectInspector{}
	var buf bool
	testCmpPtr(t, ins, &buf)
}

func BenchmarkCompiledInspector_Cmp(b *testing.B) {
	ins := &testobj_ins.TestObjectInspector{}
	var buf bool

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCmpPtr(b, ins, &buf)
	}
}
