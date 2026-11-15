package test

import (
	"testing"

	"github.com/koykov/inspector"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testStringsIterator struct {
	k [][]byte
	v []string
}

func (i *testStringsIterator) RequireKey() bool { return true }
func (i *testStringsIterator) SetKey(val any, _ inspector.Inspector) {
	i.k = append(i.k, *val.(*[]byte))
}
func (i *testStringsIterator) SetVal(val any, _ inspector.Inspector) {
	i.v = append(i.v, *val.(*string))
}
func (i *testStringsIterator) Iterate() inspector.LoopCtl { return inspector.LoopCtlNone }
func (i *testStringsIterator) reset() {
	i.k = i.k[:0]
	i.v = i.v[:0]
}

func TestStrings(t *testing.T) {
	ss := []string{"foo", "bar"}
	pp := [][]byte{[]byte("foo"), []byte("bar")}
	ins := inspector.StringsInspector{}
	t.Run("get", func(t *testing.T) {
		x, _ := ins.Get(ss, "1")
		if *x.(*string) != "bar" {
			t.FailNow()
		}
	})
	t.Run("set", func(t *testing.T) {
		s := "asd"
		_ = ins.Set(&ss, s, "1")
		if ss[1] != "asd" {
			t.FailNow()
		}
	})
	t.Run("compare", func(t *testing.T) {
		var r bool
		_ = ins.Compare(&ss, inspector.OpEq, "asd", &r, "1")
		if !r {
			t.FailNow()
		}
		_ = ins.Compare(&ss, inspector.OpNq, "bar", &r, "1")
		if !r {
			t.FailNow()
		}
	})
	t.Run("loop", func(t *testing.T) {
		var buf []byte
		it := testStringsIterator{}
		_ = ins.Loop(ss, &it, &buf)
		if it.v[1] != "asd" {
			t.FailNow()
		}
	})
	t.Run("deep equal", func(t *testing.T) {
		pp := [][]byte{[]byte("foo"), []byte("asd")}
		r := ins.DeepEqual(ss, pp)
		if !r {
			t.FailNow()
		}
	})
	t.Run("copy", func(t *testing.T) {
		var ss1 []string
		_ = ins.CopyTo(ss, &ss1, &inspector.ByteBuffer{})
		if len(ss1) != len(ss) {
			t.FailNow()
		}
	})
	t.Run("len", func(t *testing.T) {
		var r int
		_ = ins.Length(ss, &r)
		if r != 2 {
			t.FailNow()
		}
	})
	t.Run("cap", func(t *testing.T) {
		var r int
		_ = ins.Capacity(pp, &r, "1")
		if r != 3 {
			t.FailNow()
		}
	})
	t.Run("reset", func(t *testing.T) {
		ss1 := []string{"foo", "bar", "str"}
		_ = ins.Reset(&ss1)
		if len(ss1) != 0 && cap(ss1) != 3 {
			t.FailNow()
		}
	})

	t.Run("copy to string slice", func(t *testing.T) {
		src := [][]byte{[]byte("a"), []byte("b")}
		var dst []string
		var ab inspector.ByteBuffer
		require.NoError(t, ins.CopyTo(src, &dst, &ab))
		assert.Equal(t, []string{"a", "b"}, dst)
	})

	t.Run("copy always strings", func(t *testing.T) {
		cp, err := ins.Copy([][]byte{[]byte("a")})
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, cp)
	})

	t.Run("copy to wrong type", func(t *testing.T) {
		require.ErrorIs(t, ins.CopyTo([]string{"a"}, []string{}, nil), inspector.ErrMustPointerType)
		require.ErrorIs(t, ins.CopyTo([]string{"a"}, &struct{}{}, nil), inspector.ErrUnsupportedType)
	})

	t.Run("each is noop", func(t *testing.T) {
		require.NoError(t, ins.Each([]string{"a"}, func(int, string, any) {
			t.Fatal("no fields expected")
		}))
	})

	t.Run("unmarshal", func(t *testing.T) {
		v, err := ins.Unmarshal([]byte(`["a","b"]`), inspector.EncodingJSON)
		require.NoError(t, err)
		assert.Equal(t, &[]string{"a", "b"}, v)

		_, err = ins.Unmarshal(nil, inspector.Encoding(42))
		require.ErrorIs(t, err, inspector.ErrUnknownEncodingType)
	})

	t.Run("capacity only bytes", func(t *testing.T) {
		var n int
		require.NoError(t, ins.Capacity([]string{"a", "b"}, &n))
		assert.Equal(t, 0, n)

		pp := [][]byte{make([]byte, 2, 8)}
		n = 0
		require.NoError(t, ins.Capacity(pp, &n, "0"))
		assert.Equal(t, 8, n) // cap(pp[0])
	})
}

func BenchmarkStrings(b *testing.B) {
	ss := []string{"foo", "bar"}
	pp := [][]byte{[]byte("foo"), []byte("bar")}
	ins := inspector.StringsInspector{}
	b.Run("get", func(b *testing.B) {
		b.ReportAllocs()
		var x interface{}
		p := []string{"1"}
		for i := 0; i < b.N; i++ {
			_ = ins.GetTo(&ss, &x, p...)
			if *x.(*string) != "bar" {
				b.FailNow()
			}
		}
	})
	b.Run("set", func(b *testing.B) {
		b.ReportAllocs()
		s := "asd"
		p := []string{"1"}
		buf := inspector.ByteBuffer{}
		for i := 0; i < b.N; i++ {
			_ = ins.SetWithBuffer(&ss, &s, &buf, p...)
			if ss[1] != "asd" {
				b.FailNow()
			}
		}
	})
	b.Run("compare", func(b *testing.B) {
		b.ReportAllocs()
		s := "asd"
		var r bool
		p := []string{"1"}
		for i := 0; i < b.N; i++ {
			_ = ins.Compare(&ss, inspector.OpEq, s, &r, p...)
			if !r {
				b.FailNow()
			}
		}
	})
	b.Run("loop", func(b *testing.B) {
		b.ReportAllocs()
		var buf []byte
		it := testStringsIterator{}
		for i := 0; i < b.N; i++ {
			it.reset()
			_ = ins.Loop(ss, &it, &buf)
			if it.v[1] != "asd" {
				b.FailNow()
			}
		}
	})
	b.Run("deep equal", func(b *testing.B) {
		b.ReportAllocs()
		pp := [][]byte{[]byte("foo"), []byte("asd")}
		for i := 0; i < b.N; i++ {
			r := ins.DeepEqual(&ss, &pp)
			if !r {
				b.FailNow()
			}
		}
	})
	b.Run("copy", func(b *testing.B) {
		b.ReportAllocs()
		var ss1 []string
		buf := inspector.ByteBuffer{}
		for i := 0; i < b.N; i++ {
			ss1 = ss1[:0]
			buf.Reset()
			_ = ins.CopyTo(ss, &ss1, &buf)
			if len(ss1) != len(ss) {
				b.FailNow()
			}
		}
	})
	b.Run("len", func(b *testing.B) {
		b.ReportAllocs()
		var r int
		for i := 0; i < b.N; i++ {
			_ = ins.Length(ss, &r)
			if r != 2 {
				b.FailNow()
			}
		}
	})
	b.Run("cap", func(b *testing.B) {
		b.ReportAllocs()
		var r int
		p := []string{"1"}
		for i := 0; i < b.N; i++ {
			_ = ins.Capacity(pp, &r, p...)
			if r != 3 {
				b.FailNow()
			}
		}
	})
	b.Run("reset", func(b *testing.B) {
		b.ReportAllocs()
		ss1 := []string{"foo", "bar", "str"}
		for i := 0; i < b.N; i++ {
			_ = ins.Reset(&ss1)
			if len(ss1) != 0 && cap(ss1) != 3 {
				b.FailNow()
			}
		}
	})
}
