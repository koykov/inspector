package test

import (
	"math"
	"reflect"
	"testing"

	"github.com/koykov/inspector"
)

type dummyIterator struct{}

func (dummyIterator) RequireKey() bool                        { return true }
func (dummyIterator) SetKey(val any, ins inspector.Inspector) { _, _ = val, ins }
func (dummyIterator) SetVal(val any, ins inspector.Inspector) { _, _ = val, ins }
func (dummyIterator) Iterate() inspector.LoopCtl              { return inspector.LoopCtlNone }

var (
	testM_ = &map[string]any{
		"nested": []byte("some bytes"),
	}
	testM = map[string]any{
		"foo": map[string]any{
			"noptr": map[string]any{
				"str":   "my string",
				"bytes": "my bytes",
				"int":   15,
			},
		},
		"bar": &map[string]any{
			"dptr": &testM_,
		},
		"str":   "some string",
		"int":   -123456,
		"uint":  123456,
		"float": -math.Pi,
	}
	testReset = map[string]any{
		"str":   "some string",
		"int":   -123456,
		"uint":  123456,
		"float": -math.Pi,
	}
)

func TestStringAnyMap(t *testing.T) {
	t.Run("wrong type", func(t *testing.T) {
		_, err := inspector.StringAnyMapInspector{}.Get("foo", "bar")
		if err != inspector.ErrUnsupportedType {
			t.FailNow()
		}
	})
	t.Run("get", func(t *testing.T) {
		var (
			buf any
			err error
			ins inspector.StringAnyMapInspector
		)
		if err = ins.GetTo(testM, &buf, "bar", "dptr", "nested"); err != nil {
			t.Error(err)
		}
		p, ok := buf.([]byte)
		if !ok {
			t.FailNow()
		}
		if string(p) != "some bytes" {
			t.FailNow()
		}
	})
	t.Run("set", func(t *testing.T) {
		var ins inspector.StringAnyMapInspector
		if err := ins.Set(testM, 20, "foo", "noptr", "int"); err != nil {
			t.Error(err)
		}
		v := testM["foo"].(map[string]any)["noptr"].(map[string]any)["int"]
		if v != 20 {
			t.FailNow()
		}
	})
	t.Run("compare", func(t *testing.T) {
		var (
			ins inspector.StringAnyMapInspector
			r   bool
		)
		if err := ins.Compare(testM, inspector.OpLt, "100", &r, "foo", "noptr", "int"); err != nil {
			t.Error(err)
		}
		if !r {
			t.FailNow()
		}
	})
	t.Run("loop", func(t *testing.T) {
		ins := inspector.StringAnyMapInspector{}
		var buf []byte
		if err := ins.Loop(testM, dummyIterator{}, &buf, "foo", "noptr"); err != nil {
			t.Error(err)
		}
	})
	t.Run("copy", func(t *testing.T) {
		var ins inspector.StringAnyMapInspector
		cpy, err := ins.Copy(&testM)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(cpy, testM) {
			t.FailNow()
		}
	})
	t.Run("len", func(t *testing.T) {
		var (
			ins inspector.StringAnyMapInspector
			r   int
		)
		if err := ins.Length(testM, &r, "foo", "noptr"); err != nil {
			t.Error(err)
		}
		if r != 3 {
			t.FailNow()
		}
	})
	t.Run("cap", func(t *testing.T) {
		var (
			ins inspector.StringAnyMapInspector
			r   int
		)
		if err := ins.Capacity(testM, &r, "bar", "dptr", "nested"); err != nil {
			t.Error(err)
		}
		if r != 10 {
			t.FailNow()
		}
	})
	t.Run("reset", func(t *testing.T) {
		t.Run("stranymap", func(t *testing.T) {
			t.Run("full", func(t *testing.T) {
				var ins inspector.StringAnyMapInspector
				if err := ins.Reset(&testReset); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(testReset, map[string]any{}) {
					t.FailNow()
				}
			})
			t.Run("nested single", func(t *testing.T) {
				var ins inspector.StringAnyMapInspector
				var m = map[string]any{"a": map[string]any{"b": map[string]any{"c": "d"}}}
				var e = map[string]any{"a": map[string]any{"b": map[string]any{}}}
				if err := ins.Reset(&m, "a", "b", "c"); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(m, e) {
					t.FailNow()
				}
			})
			t.Run("nested multi", func(t *testing.T) {
				var ins inspector.StringAnyMapInspector
				var m = map[string]any{"a": map[string]any{"b": map[string]any{"c": "d", "e": "f"}}}
				var e = map[string]any{"a": map[string]any{"b": map[string]any{}}}
				if err := ins.Reset(&m, "a", "b"); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(m, e) {
					t.FailNow()
				}
			})
		})
	})
}

func BenchmarkStringAnyMap(b *testing.B) {
	b.Run("get", func(b *testing.B) {
		b.ReportAllocs()
		var (
			buf  any
			path = []string{"bar", "dptr", "nested"}
			ins  inspector.StringAnyMapInspector
		)
		for i := 0; i < b.N; i++ {
			_ = ins.GetTo(testM, &buf, path...)
		}
		_ = buf
	})
	b.Run("set", func(b *testing.B) {
		b.ReportAllocs()
		var (
			ins  inspector.StringAnyMapInspector
			path = []string{"foo", "noptr", "int"}
			buf  inspector.ByteBuffer
		)
		for i := 0; i < b.N; i++ {
			_ = ins.SetWithBuffer(testM, 20, &buf, path...)
			buf.Reset()
		}
	})
	b.Run("set string", func(b *testing.B) {
		b.ReportAllocs()
		var (
			ins  inspector.StringAnyMapInspector
			path = []string{"foo", "noptr", "str"}
			s_   = "my string new"
			buf  inspector.ByteBuffer
		)
		for i := 0; i < b.N; i++ {
			_ = ins.SetWithBuffer(testM, &s_, &buf, path...)
			buf.Reset()
		}
	})
	b.Run("set bytes", func(b *testing.B) {
		b.ReportAllocs()
		var (
			ins  inspector.StringAnyMapInspector
			path = []string{"foo", "noptr", "str"}
			b_   = []byte("my bytes new")
			buf  inspector.ByteBuffer
		)
		for i := 0; i < b.N; i++ {
			_ = ins.SetWithBuffer(testM, &b_, &buf, path...)
			buf.Reset()
		}
	})
	b.Run("compare", func(b *testing.B) {
		b.ReportAllocs()
		var (
			ins  inspector.StringAnyMapInspector
			path = []string{"foo", "noptr", "str"}
			r    bool
		)
		for i := 0; i < b.N; i++ {
			_ = ins.Compare(testM, inspector.OpLt, "100", &r, path...)
		}
	})
	b.Run("loop", func(b *testing.B) {
		b.ReportAllocs()
		var (
			ins  inspector.StringAnyMapInspector
			path = []string{"foo", "noptr"}
			buf  []byte
		)
		for i := 0; i < b.N; i++ {
			_ = ins.Loop(testM, dummyIterator{}, &buf, path...)
			buf = buf[:0]
		}
	})
	b.Run("copy", func(b *testing.B) {
		b.ReportAllocs()
		var (
			ins inspector.StringAnyMapInspector
			cpy map[string]any
			buf inspector.ByteBuffer
		)
		for i := 0; i < b.N; i++ {
			_ = ins.CopyTo(testM, &cpy, &buf)
			buf.Reset()
		}
	})
	b.Run("len", func(b *testing.B) {
		b.ReportAllocs()
		var (
			ins  inspector.StringAnyMapInspector
			r    int
			path = []string{"foo", "noptr"}
		)
		for i := 0; i < b.N; i++ {
			_ = ins.Length(testM, &r, path...)
		}
	})
	b.Run("cap", func(b *testing.B) {
		b.ReportAllocs()
		var (
			ins  inspector.StringAnyMapInspector
			r    int
			path = []string{"bar", "dptr", "nested"}
		)
		for i := 0; i < b.N; i++ {
			_ = ins.Capacity(testM, &r, path...)
		}
	})
}
