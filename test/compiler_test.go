package test

import (
	"bytes"
	"testing"

	"github.com/koykov/inspector"
)

func TestCompiler(t *testing.T) {
	t.Run("package", func(t *testing.T) {
		conf := inspector.Config{
			Target:      inspector.TargetPackage,
			Package:     "github.com/koykov/inspector/testobj",
			Destination: "github.com/koykov/inspector/testobj_ins",
			NoClean:     true,
			Buf:         &bytes.Buffer{},
		}
		c, err := inspector.NewCompiler(&conf)
		if err != nil {
			t.Error(err)
		}
		if err = c.Compile(); err != nil {
			t.Error(err)
		}
	})
}
