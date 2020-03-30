package test

import (
	"bytes"
	"testing"

	"github.com/koykov/inspector"
)

func TestCompiler(t *testing.T) {
	buf := &bytes.Buffer{}
	c := inspector.NewCompiler("github.com/koykov/inspector/testobj", "", buf)
	err := c.Compile()
	if err != nil {
		t.Error(err)
	}
}