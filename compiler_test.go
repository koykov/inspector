package inspector

import (
	"bytes"
	"testing"
)

func TestCompiler(t *testing.T) {
	buf := &bytes.Buffer{}
	c := NewCompiler("github.com/koykov/inspector/testobj", "", buf)
	err := c.Compile()
	if err != nil {
		t.Error(err)
	}
}
