package test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"github.com/koykov/inspector/testobj_ins"
	"github.com/stretchr/testify/assert"
)

func TestInspector(t *testing.T) {
	t.Run("compare", func(t *testing.T) {
		var buf bool
		testComparePtr(t, testobj_ins.TestObjectInspector{}, &buf)
	})
	t.Run("set", func(t *testing.T) {
		testSetterPtr(t, testobj_ins.TestObjectInspector{}, nil)
	})
	t.Run("set/buffer", func(t *testing.T) {
		ab := inspector.ByteBuffer{}
		testSetterPtr(t, testobj_ins.TestObjectInspector{}, &ab)
	})
	t.Run("deep equal", func(t *testing.T) {
		var ins testobj_ins.TestObjectInspector
		for i := range deqStages {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				stage := &deqStages[i]
				assert.Equal(t, ins.DeepEqualWithOptions(stage.l, stage.r, stage.opts), stage.eq)
			})
		}
	})
	t.Run("copy", func(t *testing.T) {
		var ins testobj_ins.TestObjectInspector
		obj := *testO
		obj.Name = []byte("foobar")
		cpy, _ := ins.Copy(obj)
		obj.Name[0] = 'F'
		assert.NotEqual(t, obj.Name, cpy.(*testobj.TestObject).Name)
	})
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

func TestInspectorReset(t *testing.T) {
	var ins testobj_ins.TestObjectInspector
	origin, _ := ins.Copy(testO)
	t.Run("full", func(t *testing.T) {
		cpy, _ := ins.Copy(origin)
		err := ins.Reset(cpy)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("nested single", func(t *testing.T) {
		cpy, _ := ins.Copy(origin)
		err := ins.Reset(cpy, "Finance", "Balance")
		if err != nil {
			t.Error(err)
		}
		to := cpy.(*testobj.TestObject)
		if to.Finance.Balance != 0 {
			t.FailNow()
		}
	})
	t.Run("nested map", func(t *testing.T) {
		t.Run("single key", func(t *testing.T) {
			cpy, _ := ins.Copy(origin)
			err := ins.Reset(cpy, "Flags", "Valid")
			if err != nil {
				t.Error(err)
			}
			to := cpy.(*testobj.TestObject)
			if to.Flags["Valid"] != 0 {
				t.FailNow()
			}
		})
		t.Run("full map", func(t *testing.T) {
			cpy, _ := ins.Copy(origin)
			err := ins.Reset(cpy, "Flags")
			if err != nil {
				t.Error(err)
			}
			to := cpy.(*testobj.TestObject)
			if len(to.Flags) != 0 {
				t.FailNow()
			}
		})
	})
	t.Run("nested slice", func(t *testing.T) {
		t.Run("single id", func(t *testing.T) {
			cpy, _ := ins.Copy(origin)
			err := ins.Reset(cpy, "Finance", "History", "1")
			if err != nil {
				t.Error(err)
			}
			to := cpy.(*testobj.TestObject)
			if x := to.Finance.History[1]; x.DateUnix != 0 || x.Cost != 0 || len(x.Comment) != 0 {
				t.FailNow()
			}
		})
		t.Run("full slice", func(t *testing.T) {
			cpy, _ := ins.Copy(origin)
			err := ins.Reset(cpy, "Finance", "History")
			if err != nil {
				t.Error(err)
			}
			to := cpy.(*testobj.TestObject)
			if len(to.Finance.History) != 0 {
				t.FailNow()
			}
		})
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

func TestInspectorEach(t *testing.T) {
	var ins testobj_ins.TestObjectInspector
	obj := *testO
	dst := make(map[string]any)
	_ = ins.Each(&obj, func(_ int, field string, value any) {
		dst[field] = value
	})
	t0, t1, t2, t3, t4 := "foo", []byte("bar"), 12.34, int32(78), uint64(0)
	expect := map[string]any{
		"id":         &t0,
		"name":       &t1,
		"cost":       &t2,
		"status":     &t3,
		"ustate":     &t4,
		"permission": &testobj.TestPermission{15: true, 23: false},
		"flags": testobj.TestFlag{
			"export": 17,
			"ro":     4,
			"rw":     7,
			"Valid":  1,
		},
		"history_tree": map[string]*testobj.TestHistory(nil),
		"finance": &testobj.TestFinance{
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
	if !reflect.DeepEqual(dst, expect) {
		t.Fail()
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

func BenchmarkInspectorEach(b *testing.B) {
	ins := testobj_ins.TestObjectInspector{}
	src := *testO
	dst := make(map[string]any)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		clear(dst)
		_ = ins.Each(&src, func(_ int, field string, value any) {
			dst[field] = value
		})
	}
}
