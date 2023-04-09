package test

import (
	"testing"

	"github.com/koykov/inspector"
)

func TestStrings(t *testing.T) {
	ss := []string{"foo", "bar"}
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
}

func BenchmarkStrings(b *testing.B) {
	ss := []string{"foo", "bar"}
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
}
