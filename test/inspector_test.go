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
	p9 = []string{"Finance", "AllowBuy"}

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

func testComparePtr(t testing.TB, i inspector.Inspector, buf *bool) {
	_ = i.Compare(testO, inspector.OpEq, "foo", buf, p0...)
	if !*buf {
		t.Error("object.Id: mismatch result and expectation")
	}

	_ = i.Compare(testO, inspector.OpEq, "bar", buf, p1...)
	if !*buf {
		t.Error("object.Name: mismatch result and expectation")
	}

	_ = i.Compare(testO, inspector.OpGtq, "60", buf, p7...)
	if !*buf {
		t.Error("object.Status: mismatch result and expectation")
	}

	_ = i.Compare(testO, inspector.OpLtq, "5000", buf, p8...)
	if !*buf {
		t.Error("object.Finance.MoneyIn: mismatch result and expectation")
	}

	_ = i.Compare(testO, inspector.OpEq, "true", buf, p9...)
	if !*buf {
		t.Error("object.Finance.AllowBuy: mismatch result and expectation")
	}
}

func testSetterPtr(t testing.TB, i inspector.Inspector, ab inspector.AccumulativeBuffer) {
	var cins testobj_ins.TestObjectInspector
	obj1, _ := cins.Copy(testO)
	obj := obj1.(*testobj.TestObject)
	obj.Id = ""
	_ = i.SetWithBuffer(obj, 1984, ab, p0...)
	if obj.Id != "1984" {
		t.Error("object.Id: mismatch result and expectation")
	}

	_ = i.SetWithBuffer(obj, 2000, ab, p1...)
	if !bytes.Equal(obj.Name, expectName) {
		t.Error("object.Name: mismatch result and expectation")
	}

	_ = i.SetWithBuffer(obj, false, ab, p2...)
	if (*obj.Permission)[23] != false {
		t.Error("object.Permission.23: mismatch result and expectation")
	}

	_ = i.SetWithBuffer(obj, int32(23), ab, p3...)
	if obj.Flags["export"] != 23 {
		t.Error("object.Flags.export: mismatch result and expectation")
	}

	_ = i.SetWithBuffer(obj, float64(9000), ab, p4...)
	if obj.Finance.Balance != 9000 {
		t.Error("object.Finance.Balance: mismatch result and expectation")
	}

	_ = i.SetWithBuffer(obj, int64(153465345246), ab, p5...)
	if obj.Finance.History[1].DateUnix != 153465345246 {
		t.Error("object.Finance.History.1.DateUnix: mismatch result and expectation")
	}

	_ = i.SetWithBuffer(obj, &expectComment1, ab, p6...)
	if !bytes.Equal(obj.Finance.History[0].Comment, expectComment1) {
		t.Error("object.Finance.History.0.DateUnix: mismatch result and expectation")
	}
}

func TestInspector(t *testing.T) {
	t.Run("reflect/get", func(t *testing.T) {
		testGetter(t, inspector.ReflectInspector{})
	})
	t.Run("cg/get", func(t *testing.T) {
		var buf any
		testGetterPtr(t, testobj_ins.TestObjectInspector{}, buf)
	})
	t.Run("cg/cmp", func(t *testing.T) {
		var buf bool
		testComparePtr(t, testobj_ins.TestObjectInspector{}, &buf)
	})
	t.Run("cg/set", func(t *testing.T) {
		testSetterPtr(t, testobj_ins.TestObjectInspector{}, nil)
	})
	t.Run("cg/setBuf", func(t *testing.T) {
		ab := inspector.ByteBuffer{}
		testSetterPtr(t, testobj_ins.TestObjectInspector{}, &ab)
	})
}

func TestInspectorDeepEqual(t *testing.T) {
	var ins testobj_ins.TestObjectInspector
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
	var ins testobj_ins.TestObjectInspector
	obj := *testO
	obj.Name = []byte("foobar")
	cpy, _ := ins.Copy(obj)
	obj.Name[0] = 'F'
	if bytes.Equal(obj.Name, cpy.(*testobj.TestObject).Name) {
		t.FailNow()
	}
}

func TestInspectorAppend(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		var ins inspector.StringsInspector
		ss := []string{"foo", "bar"}
		raw, err := ins.Append(&ss, "qwe")
		if err != nil {
			t.Error(err)
		}
		ss1, ok := raw.(*[]string)
		if !ok {
			t.FailNow()
		}
		if len(*ss1) != 3 {
			t.FailNow()
		}
		if (*ss1)[2] != "qwe" {
			t.FailNow()
		}
	})
	t.Run("object", func(t *testing.T) {
		var ins testobj_ins.TestObjectInspector
		cpy, _ := ins.Copy(testO)
		raw, err := ins.Append(cpy, testobj.TestHistory{
			DateUnix: 111111,
			Cost:     3.1415,
			Comment:  nil,
		}, "Finance", "History")
		if err != nil {
			t.Error(err)
		}
		hist, ok := raw.(*[]testobj.TestHistory)
		if !ok {
			t.FailNow()
		}
		if len(*hist) != 4 {
			t.FailNow()
		}
		if (*hist)[3].DateUnix != 111111 {
			t.FailNow()
		}
	})
}

func TestInspectorLenCap(t *testing.T) {
	var ins testobj_ins.TestObjectInspector
	obj := *testO
	var l, c int

	_ = ins.Length(obj, &l, p0...)
	if l != 3 {
		t.FailNow()
	}

	_ = ins.Length(obj, &l, p1...)
	_ = ins.Capacity(obj, &c, p1...)
	if l != 3 || c != 3 {
		t.FailNow()
	}

	_ = ins.Length(obj, &l, "Permission")
	if l != 2 {
		t.FailNow()
	}

	_ = ins.Length(obj, &l, "Flags")
	if l != 4 {
		t.FailNow()
	}

	_ = ins.Length(obj, &l, "Finance", "History")
	_ = ins.Capacity(obj, &c, "Finance", "History")
	if l != 3 || c != 3 {
		t.FailNow()
	}

	_ = ins.Length(obj, &l, p6...)
	_ = ins.Capacity(obj, &c, p6...)
	if l != 14 || c != 14 {
		t.FailNow()
	}
}

func BenchmarkInspectorDeepEqual(b *testing.B) {
	var ins testobj_ins.TestObjectInspector
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
		var ins inspector.ReflectInspector
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			testGetter(b, ins)
		}
	})
	b.Run("cg/get", func(b *testing.B) {
		var ins testobj_ins.TestObjectInspector
		var buf any
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			testGetterPtr(b, ins, buf)
		}
	})
	b.Run("cg/cmp", func(b *testing.B) {
		var ins testobj_ins.TestObjectInspector
		var buf bool
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			testComparePtr(b, ins, &buf)
		}
	})
	b.Run("cg/set", func(b *testing.B) {
		var ins testobj_ins.TestObjectInspector
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			testSetterPtr(b, ins, nil)
		}
	})
	b.Run("cg/setBuf", func(b *testing.B) {
		var ins testobj_ins.TestObjectInspector
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

func BenchmarkInspectorAppend(b *testing.B) {
	b.Run("strings", func(b *testing.B) {
		var ins inspector.StringsInspector
		ss := []string{"foo", "bar"}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ss = ss[:2]
			_, err := ins.Append(&ss, "qwe")
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("object", func(b *testing.B) {
		var ins testobj_ins.TestObjectInspector
		raw, _ := ins.Copy(testO)
		cpy := raw.(*testobj.TestObject)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			cpy.Finance.History = cpy.Finance.History[:3]
			_, _ = ins.Append(cpy, testobj.TestHistory{
				DateUnix: 111111,
				Cost:     3.1415,
				Comment:  nil,
			}, "Finance", "History")
		}
	})
}

func BenchmarkInspectorLenCap(b *testing.B) {
	ins := testobj_ins.TestObjectInspector{}
	obj := *testO
	path := []string{"Finance", "History"}
	b.Run("len", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var l int
			_ = ins.Length(obj, &l, path...)
			if l != 3 {
				b.FailNow()
			}
		}
	})
	b.Run("cap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var c int
			_ = ins.Capacity(obj, &c, path...)
			if c != 3 {
				b.FailNow()
			}
		}
	})
}
