package test

import (
	"testing"

	"github.com/koykov/inspector"
	"github.com/stretchr/testify/assert"
)

func TestEqualFloat(t *testing.T) {
	assert.True(t, inspector.EqualFloat64(1.0, 1.0005, nil))
	assert.False(t, inspector.EqualFloat64(1.0, 1.002, nil))
	assert.True(t, inspector.EqualFloat64(1.0, 1.05, &inspector.DEQOptions{Precision: 0.1}))
	assert.True(t, inspector.EqualFloat32(1.0, 1.0005, nil))
	assert.False(t, inspector.EqualFloat32(1.0, 1.002, nil))
}

func TestDEQMustCheck(t *testing.T) {
	assert.True(t, inspector.DEQMustCheck("X", nil))
	assert.False(t, inspector.DEQMustCheck("X", &inspector.DEQOptions{
		Exclude: map[string]struct{}{"X": {}},
	}))
	assert.True(t, inspector.DEQMustCheck("X", &inspector.DEQOptions{
		Filter: map[string]struct{}{"X": {}},
	}))
	assert.False(t, inspector.DEQMustCheck("X", &inspector.DEQOptions{
		Filter: map[string]struct{}{"Y": {}},
	}))
}
