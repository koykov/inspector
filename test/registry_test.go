package test

import (
	"testing"

	"github.com/koykov/inspector"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistry(t *testing.T) {
	ins, err := inspector.GetInspector("static")
	require.NoError(t, err)
	assert.Equal(t, "static", ins.TypeName())

	_, err = inspector.GetInspector("definitely-unknown")
	require.ErrorIs(t, err, inspector.ErrUnknownInspector)

	names := map[string]struct{}{}
	inspector.EachInspector(func(name string, _ inspector.Inspector) {
		names[name] = struct{}{}
	})
	for _, want := range []string{"static", "strings", "bytes", "stringAnyMap", "stringStringMap", "reflect"} {
		_, ok := names[want]
		assert.True(t, ok, "missing inspector %q", want)
	}
}
