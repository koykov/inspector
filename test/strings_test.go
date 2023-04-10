package test

import (
	"testing"

	"github.com/koykov/inspector"
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
