package test

import (
	"testing"

	"github.com/koykov/inspector"
)

func TestStrings(t *testing.T) {
	ss := []string{"foo", "bar"}
	t.Run("get", func(t *testing.T) {
		ins := inspector.StringsInspector{}
		x, _ := ins.Get(ss, "1")
		if *x.(*string) != "bar" {
			t.FailNow()
		}
	})
}

func BenchmarkStrings(b *testing.B) {
	ss := []string{"foo", "bar"}
	b.Run("get", func(b *testing.B) {
		b.ReportAllocs()
		ins := inspector.StringsInspector{}
		var x interface{}
		p := []string{"1"}
		for i := 0; i < b.N; i++ {
			_ = ins.GetTo(&ss, &x, p...)
			if *x.(*string) != "bar" {
				b.FailNow()
			}
		}
	})
}
