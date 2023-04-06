package test

import (
	"bytes"
	"strconv"
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

	expectFoo      = []byte("bar")
	expectName     = []byte("2000")
	expectComment  = []byte("pay for domain")
	expectComment1 = []byte("lorem ipsum dolor sit amet")

	deqStages = []struct {
		l, r *testobj.TestObject
		opts *inspector.DEQOptions
		eq   bool
	}{
		{
			l:  nil,
			r:  nil,
			eq: true,
		},
		{
			l:  &testobj.TestObject{Id: "foo"},
			r:  &testobj.TestObject{Id: "bar"},
			eq: false,
		},
		{
			l:  &testobj.TestObject{Cost: 3.1415},
			r:  &testobj.TestObject{Cost: 3.1415},
			eq: true,
		},
		{
			l:  &testobj.TestObject{Name: []byte("qwe")},
			r:  &testobj.TestObject{Name: []byte("rty")},
			eq: false,
		},
		{
			l:  &testobj.TestObject{Flags: testobj.TestFlag{"xxx": 15}},
			r:  &testobj.TestObject{Flags: testobj.TestFlag{"xxx": 15}},
			eq: true,
		},
		{
			l:  &testobj.TestObject{Flags: testobj.TestFlag{"xxx": 15}},
			r:  &testobj.TestObject{Flags: testobj.TestFlag{"xxx": 54}},
			eq: false,
		},
		{
			l:  &testobj.TestObject{Flags: testobj.TestFlag{"xxx": 15}},
			r:  &testobj.TestObject{Flags: testobj.TestFlag{"yyy": 15}},
			eq: false,
		},
		{
			l:  &testobj.TestObject{Finance: &testobj.TestFinance{History: []testobj.TestHistory{{Cost: 22}}}},
			r:  &testobj.TestObject{Finance: &testobj.TestFinance{History: []testobj.TestHistory{{Cost: 22}}}},
			eq: true,
		},
		{
			l:  &testobj.TestObject{Finance: &testobj.TestFinance{History: []testobj.TestHistory{{Comment: []byte("aaa")}}}},
			r:  &testobj.TestObject{Finance: &testobj.TestFinance{History: []testobj.TestHistory{{}}}},
			eq: false,
		},
		{
			l:    &testobj.TestObject{Id: "foobar", Name: []byte("qwe")},
			r:    &testobj.TestObject{Id: "foobar", Name: []byte("rty")},
			opts: &inspector.DEQOptions{Exclude: map[string]struct{}{"Name": {}}},
			eq:   true,
		},
		{
			l:    &testobj.TestObject{Id: "foobar", Name: []byte("qwe")},
			r:    &testobj.TestObject{Id: "foobar", Name: []byte("rty")},
			opts: &inspector.DEQOptions{Filter: map[string]struct{}{"Id": {}}},
			eq:   true,
		},
	}
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

func testGetterPtr(t testing.TB, i inspector.Inspector, buf any) {
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

func testSetterPtr(t testing.TB, i inspector.Inspector, ab inspector.AccumulativeBuffer) {
	testO.Id = ""
	_ = i.SetWB(testO, 1984, ab, p0...)
	if testO.Id != "1984" {
		t.Error("object.Id: mismatch result and expectation")
	}

	_ = i.SetWB(testO, 2000, ab, p1...)
	if !bytes.Equal(testO.Name, expectName) {
		t.Error("object.Name: mismatch result and expectation")
	}

	_ = i.SetWB(testO, false, ab, p2...)
	if (*testO.Permission)[23] != false {
		t.Error("object.Permission.23: mismatch result and expectation")
	}

	_ = i.SetWB(testO, int32(23), ab, p3...)
	if testO.Flags["export"] != 23 {
		t.Error("object.Flags.export: mismatch result and expectation")
	}

	_ = i.SetWB(testO, float64(9000), ab, p4...)
	if testO.Finance.Balance != 9000 {
		t.Error("object.Finance.Balance: mismatch result and expectation")
	}

	_ = i.SetWB(testO, int64(153465345246), ab, p5...)
	if testO.Finance.History[1].DateUnix != 153465345246 {
		t.Error("object.Finance.History.1.DateUnix: mismatch result and expectation")
	}

	_ = i.SetWB(testO, &expectComment1, ab, p6...)
	if !bytes.Equal(testO.Finance.History[0].Comment, expectComment1) {
		t.Error("object.Finance.History.0.DateUnix: mismatch result and expectation")
	}
}

func TestInspector(t *testing.T) {
	t.Run("reflect/get", func(t *testing.T) {
		testGetter(t, &inspector.ReflectInspector{})
	})
	t.Run("cg/get", func(t *testing.T) {
		var buf any
		testGetterPtr(t, &testobj_ins.TestObjectInspector{}, buf)
	})
	t.Run("cg/cmp", func(t *testing.T) {
		var buf bool
		testCmpPtr(t, &testobj_ins.TestObjectInspector{}, &buf)
	})
	t.Run("cg/set", func(t *testing.T) {
		testSetterPtr(t, &testobj_ins.TestObjectInspector{}, nil)
	})
	t.Run("cg/setBuf", func(t *testing.T) {
		ab := inspector.ByteBuffer{}
		testSetterPtr(t, &testobj_ins.TestObjectInspector{}, &ab)
	})
}

func TestInspectorDeepEqual(t *testing.T) {
	ins := &testobj_ins.TestObjectInspector{}
	for i := range deqStages {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			stage := &deqStages[i]
			if ins.DeepEqualWithOptions(stage.l, stage.r, stage.opts) != stage.eq {
				t.FailNow()
			}
		})
	}
}

func TestInspectorCopy(t *testing.T) {
	ins := &testobj_ins.TestObjectInspector{}
	obj := *testO
	obj.Name = []byte("foobar")
	cpy, _ := ins.Copy(obj)
	obj.Name[0] = 'F'
	if bytes.Equal(obj.Name, cpy.(*testobj.TestObject).Name) {
		t.FailNow()
	}
}

func BenchmarkInspectorDeepEqual(b *testing.B) {
	ins := &testobj_ins.TestObjectInspector{}
	for i := range deqStages {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			b.ReportAllocs()
			stage := &deqStages[i]
			for j := 0; j < b.N; j++ {
				if ins.DeepEqualWithOptions(stage.l, stage.r, stage.opts) != stage.eq {
					b.FailNow()
				}
			}
		})
	}
}

func BenchmarkInspector(b *testing.B) {
	b.Run("reflect/get", func(b *testing.B) {
		ins := &inspector.ReflectInspector{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			testGetter(b, ins)
		}
	})
	b.Run("cg/get", func(b *testing.B) {
		ins := &testobj_ins.TestObjectInspector{}
		var buf any
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			testGetterPtr(b, ins, buf)
		}
	})
	b.Run("cg/cmp", func(b *testing.B) {
		ins := &testobj_ins.TestObjectInspector{}
		var buf bool
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			testCmpPtr(b, ins, &buf)
		}
	})
	b.Run("cg/set", func(b *testing.B) {
		ins := &testobj_ins.TestObjectInspector{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			testSetterPtr(b, ins, nil)
		}
	})
	b.Run("cg/setBuf", func(b *testing.B) {
		ins := &testobj_ins.TestObjectInspector{}
		ab := inspector.ByteBuffer{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			testSetterPtr(b, ins, &ab)
			ab.Reset()
		}
	})
}

func BenchmarkInspectorCopy(b *testing.B) {
	b.Run("testobj", func(b *testing.B) {
		var (
			cpy testobj.TestObject
			ins testobj_ins.TestObjectInspector
			buf inspector.ByteBuffer
		)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ins.Reset(&cpy)
			buf.Reset()
			if err := ins.CopyTo(testO, &cpy, &buf); err != nil {
				b.Fatal(err)
			}
		}
	})
	fn1 := func(b *testing.B, origin *testobj.TestObject1) {
		var (
			cpy testobj.TestObject1
			ins testobj_ins.TestObject1Inspector
			buf inspector.ByteBuffer
		)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if err := ins.Reset(&cpy); err != nil {
				b.Fatal(err)
			}
			buf.Reset()
			if err := ins.CopyTo(origin, &cpy, &buf); err != nil {
				b.Fatal(err)
			}
		}
	}
	b.Run("testobj1+alloc-free-cases", func(b *testing.B) {
		origin := testobj.TestObject1{
			IntSlice:        []int32{0, 1, 2, 3, 4, 5},
			ByteSlice:       []byte("lorem ipsum ..."),
			FloatSlice:      testobj.TestFloatSlice{0, 1, 2, 3, 4, 5},
			StructSlice:     []testobj.TestStruct{{A: 1, S: "foobar", B: []byte("foobar"), I: 12, I8: 8, I16: 16, I32: 32, I64: 64, U: 256, U8: 8, U16: 16, U32: 32, U64: 64, F: 3.1415, D: 3.141561}},
			IntStringMap:    map[int]string{0: "foo", 1: "bar", 2: "qwe"},
			StringFloatMap:  testobj.TestStringFloatMap{"foo": 1, "bar": 2, "qwe": 3},
			FloatStructMap:  map[float64]testobj.TestStruct{1: {A: 1, S: "foobar", B: []byte("foobar"), I: 12, I8: 8, I16: 16, I32: 32, I64: 64, U: 256, U8: 8, U16: 16, U32: 32, U64: 64, F: 3.1415, D: 3.141561}},
			NestedStruct:    testobj.TestStruct{A: 1, S: "foobar", B: []byte("foobar"), I: 12, I8: 8, I16: 16, I32: 32, I64: 64, U: 256, U8: 8, U16: 16, U32: 32, U64: 64, F: 3.1415, D: 3.141561},
			NestedStructPtr: &testobj.TestStruct{A: 1, S: "foobar", B: []byte("foobar"), I: 12, I8: 8, I16: 16, I32: 32, I64: 64, U: 256, U8: 8, U16: 16, U32: 32, U64: 64, F: 3.1415, D: 3.141561},
		}
		fn1(b, &origin)
	})
	b.Run("testobj1+string maps", func(b *testing.B) {
		origin := testobj.TestObject1{
			IntSlice:       []int32{0, 1, 2, 3, 4, 5},
			ByteSlice:      []byte("lorem ipsum ..."),
			FloatSlice:     testobj.TestFloatSlice{0, 1, 2, 3, 4, 5},
			IntStringMap:   map[int]string{0: "foo", 1: "bar", 2: "qwe"},
			StringFloatMap: testobj.TestStringFloatMap{"foo": 1, "bar": 2, "qwe": 3},
		}
		fn1(b, &origin)
	})
	b.Run("testobj1+struct pointer slice", func(b *testing.B) {
		origin := testobj.TestObject1{
			IntSlice:           []int32{0, 1, 2, 3, 4, 5},
			ByteSlice:          []byte("lorem ipsum ..."),
			FloatSlice:         testobj.TestFloatSlice{0, 1, 2, 3, 4, 5},
			StructPtrSlice:     []*testobj.TestStruct{{A: 1, S: "foobar", B: []byte("foobar"), I: 12, I8: 8, I16: 16, I32: 32, I64: 64, U: 256, U8: 8, U16: 16, U32: 32, U64: 64, F: 3.1415, D: 3.141561}},
			StructSliceLiteral: []*testobj.TestStruct{{A: 1, S: "foobar", B: []byte("foobar"), I: 12, I8: 8, I16: 16, I32: 32, I64: 64, U: 256, U8: 8, U16: 16, U32: 32, U64: 64, F: 3.1415, D: 3.141561}},
		}
		fn1(b, &origin)
	})
	b.Run("testobj1+nested maps", func(b *testing.B) {
		origin := testobj.TestObject1{
			IntSlice:   []int32{0, 1, 2, 3, 4, 5},
			ByteSlice:  []byte("lorem ipsum ..."),
			FloatSlice: testobj.TestFloatSlice{0, 1, 2, 3, 4, 5},
			IntIntMapMap: map[int32]map[int32]int32{
				0: {0: 1, 1: 2, 2: 3},
			},
		}
		fn1(b, &origin)
	})
	b.Run("testobj1+struct ptr map", func(b *testing.B) {
		origin := testobj.TestObject1{
			IntSlice:          []int32{0, 1, 2, 3, 4, 5},
			ByteSlice:         []byte("lorem ipsum ..."),
			FloatSlice:        testobj.TestFloatSlice{0, 1, 2, 3, 4, 5},
			FloatStructPtrMap: map[float64]*testobj.TestStruct{1: {A: 1, S: "foobar", B: []byte("foobar"), I: 12, I8: 8, I16: 16, I32: 32, I64: 64, U: 256, U8: 8, U16: 16, U32: 32, U64: 64, F: 3.1415, D: 3.141561}},
		}
		fn1(b, &origin)
	})
}
