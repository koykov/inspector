package test

import (
	"testing"

	"github.com/koykov/inspector"
	"github.com/stretchr/testify/assert"
)

func TestAssignBuiltins(t *testing.T) {
	t.Run("to int", func(t *testing.T) {
		var v int64
		assert.True(t, inspector.AssignToInt(&v, "42", nil))
		assert.Equal(t, int64(42), v)
		assert.False(t, inspector.AssignToInt(&v, "x", nil))
		assert.False(t, inspector.AssignToInt(&v, "999999999999999999999", nil))
		assert.False(t, inspector.AssignToInt(int64(0), "1", nil))
	})

	t.Run("to uint", func(t *testing.T) {
		var v uint64
		assert.True(t, inspector.AssignToUint(&v, "42", nil))
		assert.Equal(t, uint64(42), v)
		assert.False(t, inspector.AssignToUint(&v, "-1", nil))
	})

	t.Run("to float", func(t *testing.T) {
		var v float64
		assert.True(t, inspector.AssignToFloat(&v, "3.5", nil))
		assert.Equal(t, 3.5, v)
		assert.False(t, inspector.AssignToFloat(&v, "x", nil))
	})

	t.Run("to bool", func(t *testing.T) {
		var v bool
		assert.True(t, inspector.AssignToBool(&v, "true", nil))
		assert.True(t, v)
		assert.True(t, inspector.AssignToBool(&v, "1", nil))
		assert.False(t, v) // только литерал "true"
		assert.True(t, inspector.AssignToBool(&v, int64(2), nil))
		assert.True(t, v)
	})

	t.Run("to string", func(t *testing.T) {
		var s string
		assert.True(t, inspector.AssignToStr(&s, []byte("foobar"), nil))
		assert.Equal(t, "foobar", s)
	})

	t.Run("to bytes", func(t *testing.T) {
		var b []byte
		assert.True(t, inspector.AssignToBytes(&b, "foobar", nil))
		assert.Equal(t, []byte("foobar"), b)
	})

	t.Run("generic", func(t *testing.T) {
		var v int64
		assert.True(t, inspector.Assign(&v, "42"))
		assert.Equal(t, int64(42), v)
		assert.True(t, inspector.AssignBuf(&v, "43", nil))
		assert.Equal(t, int64(43), v)
	})
}
