package inspector

import "testing"

func TestCompiler(t *testing.T) {
	c := NewCompiler("github.com/koykov/inspector/testobj", "")
	err := c.Compile()
	if err != nil {
		t.Error(err)
	}
}
