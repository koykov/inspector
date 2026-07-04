package test

import (
	"math"
	"reflect"
	"testing"

	"github.com/koykov/inspector"
	"github.com/stretchr/testify/assert"
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
		assert.ErrorIs(t, err, inspector.ErrUnsupportedType)
	})

	t.Run("get", func(t *testing.T) {
		var (
			buf any
			ins inspector.StringAnyMapInspector
		)
		err := ins.GetTo(testM, &buf, "bar", "dptr", "nested")
		assert.NoError(t, err)
		p, ok := buf.([]byte)
		assert.True(t, ok)
		assert.Equal(t, string(p), "some bytes")
	})

	t.Run("set", func(t *testing.T) {
		var ins inspector.StringAnyMapInspector
		err := ins.Set(testM, 20, "foo", "noptr", "int")
		assert.NoError(t, err)
		v := testM["foo"].(map[string]any)["noptr"].(map[string]any)["int"]
		assert.Equal(t, v, 20)
	})

	t.Run("compare", func(t *testing.T) {
		var (
			ins inspector.StringAnyMapInspector
			r   bool
		)
		err := ins.Compare(testM, inspector.OpLt, "100", &r, "foo", "noptr", "int")
		assert.NoError(t, err)
		assert.True(t, r)
	})

	t.Run("loop", func(t *testing.T) {
		ins := inspector.StringAnyMapInspector{}
		var buf []byte
		err := ins.Loop(testM, dummyIterator{}, &buf, "foo", "noptr")
		assert.NoError(t, err)
	})

	t.Run("copy", func(t *testing.T) {
		var ins inspector.StringAnyMapInspector
		cpy, err := ins.Copy(&testM)
		assert.NoError(t, err)
		assert.Equal(t, cpy, testM)
	})

	t.Run("len", func(t *testing.T) {
		var (
			ins inspector.StringAnyMapInspector
			r   int
		)
		err := ins.Length(testM, &r, "foo", "noptr")
		assert.NoError(t, err)
		assert.Equal(t, r, 3)
	})

	t.Run("cap", func(t *testing.T) {
		var (
			ins inspector.StringAnyMapInspector
			r   int
		)
		err := ins.Capacity(testM, &r, "bar", "dptr", "nested")
		assert.NoError(t, err)
		assert.Equal(t, r, 10)
	})

	t.Run("reset", func(t *testing.T) {
		t.Run("full", func(t *testing.T) {
			var ins inspector.StringAnyMapInspector
			err := ins.Reset(&testReset)
			assert.NoError(t, err)
			assert.Equal(t, testReset, map[string]any{})
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
			err := ins.Reset(&m, "a", "b")
			assert.NoError(t, err)
			assert.Equal(t, m, e)
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
