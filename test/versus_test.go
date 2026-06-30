package test

import (
	"testing"

	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj_ins"
)

func TestVersus(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		t.Run("reflect", func(t *testing.T) {
			testGetter(t, inspector.ReflectInspector{})
		})
		t.Run("inspector", func(t *testing.T) {
			var buf any
			testGetterPtr(t, testobj_ins.TestObjectInspector{}, buf)
		})
	})
}

func BenchmarkVersus(b *testing.B) {
	b.Run("get", func(b *testing.B) {
		b.Run("reflect", func(b *testing.B) {
			var ins inspector.ReflectInspector
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				testGetter(b, ins)
			}
		})
		b.Run("inspector", func(b *testing.B) {
			var ins testobj_ins.TestObjectInspector
			var buf any
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				testGetterPtr(b, ins, buf)
			}
		})
	})
}
