package test

import (
	"testing"

	"github.com/koykov/inspector"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringStringMap(t *testing.T) {
	ins := inspector.StringStringMapInspector{}
	src := map[string]string{"a": "1", "b": "2"}

	t.Run("type name instance", func(t *testing.T) {
		assert.Equal(t, "map[string]string", ins.TypeName())
		assert.Equal(t, map[string]string{}, ins.Instance(false))
		assert.Equal(t, &map[string]string{}, ins.Instance(true))
	})

	t.Run("get", func(t *testing.T) {
		v, err := ins.Get(src, "a")
		require.NoError(t, err)
		assert.Equal(t, "1", *asPtr[string](t, v))

		// only path[0] is consulted: extra path elements are ignored
		nested, err := ins.Get(src, "a", "ignored")
		require.NoError(t, err)
		assert.Equal(t, *asPtr[string](t, v), *asPtr[string](t, nested))

		v, err = ins.Get(src, "missing")
		require.NoError(t, err)
		assert.Nil(t, v)

		// missing path[0] with extra elements is still a miss
		v, err = ins.Get(src, "missing", "ignored")
		require.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("get to", func(t *testing.T) {
		var buf any
		require.NoError(t, ins.GetTo(src, &buf, "b"))
		assert.Equal(t, "2", *asPtr[string](t, buf))

		buf = "sentinel"
		require.NoError(t, ins.GetTo(src, &buf, "missing"))
		assert.Equal(t, "sentinel", buf)

		// empty path copies the source as-is
		var raw any
		require.NoError(t, ins.GetTo(src, &raw))
		assert.Equal(t, src, raw)
	})

	t.Run("set", func(t *testing.T) {
		dst := map[string]string{}
		require.NoError(t, ins.Set(dst, "x", "k"))
		assert.Equal(t, "x", dst["k"])

		require.NoError(t, ins.Set(dst, []byte("y"), "k2"))
		assert.Equal(t, "y", dst["k2"])

		var buf inspector.ByteBuffer
		require.NoError(t, ins.SetWithBuffer(dst, "z", &buf, "k3"))
		assert.Equal(t, "z", dst["k3"])
	})

	t.Run("compare", func(t *testing.T) {
		var res bool
		require.NoError(t, ins.Compare(src, inspector.OpEq, "1", &res, "a"))
		assert.True(t, res)
		require.NoError(t, ins.Compare(src, inspector.OpNq, "1", &res, "a"))
		assert.False(t, res)

		res = true
		require.NoError(t, ins.Compare(src, inspector.OpEq, "1", &res, "missing"))
		assert.True(t, res) // result untouched
	})

	t.Run("loop", func(t *testing.T) {
		it := &collectingIterator{requireKey: true}
		require.NoError(t, ins.Loop(src, it, new([]byte)))
		assert.Len(t, it.keys, 2)
		assert.Len(t, it.vals, 2)

		// LoopCtlBrk stops after the first visited key
		it = &collectingIterator{requireKey: true, ctl: inspector.LoopCtlBrk}
		require.NoError(t, ins.Loop(src, it, new([]byte)))
		assert.Len(t, it.keys, 1)

		// LoopCtlCnt continues: every key is still visited
		it = &collectingIterator{requireKey: true, ctl: inspector.LoopCtlCnt}
		require.NoError(t, ins.Loop(src, it, new([]byte)))
		assert.Len(t, it.keys, 2)
	})

	t.Run("deep equal regression", func(t *testing.T) {
		// The comparison is asymmetric: indir2 requires the right-hand operand to be a
		// pointer (a bare value map is rejected with ErrMustPointerType before any
		// comparison happens). Pointer-right cases below exercise the real compare branch.
		l := map[string]string{"a": "1"}
		rEqual := map[string]string{"a": "1"}
		rDiff := map[string]string{"a": "2"}
		rExtra := map[string]string{"a": "1", "b": "2"}

		assert.True(t, ins.DeepEqual(l, &rEqual))
		assert.False(t, ins.DeepEqual(l, &rDiff))
		assert.False(t, ins.DeepEqual(l, &rExtra))
		assert.False(t, ins.DeepEqual(nil, nil))
		assert.True(t, ins.DeepEqualWithOptions(l, &rEqual, &inspector.DEQOptions{Exclude: map[string]struct{}{"a": {}}}))
		assert.False(t, ins.DeepEqualWithOptions(l, &rDiff, &inspector.DEQOptions{Exclude: map[string]struct{}{"a": {}}}))
		assert.False(t, ins.DeepEqual(l, rEqual))
	})

	t.Run("copy", func(t *testing.T) {
		cp, err := ins.Copy(src)
		require.NoError(t, err)
		assert.Equal(t, src, cp)

		dst := map[string]string{"old": "value"}
		require.NoError(t, ins.CopyTo(src, &dst, nil))
		assert.Equal(t, src, dst)
	})

	t.Run("length capacity each reset", func(t *testing.T) {
		var n = -1
		require.NoError(t, ins.Length(src, &n))
		assert.Equal(t, 2, n)

		b := []byte("abc")
		require.NoError(t, ins.Length(b, &n))
		assert.Equal(t, 3, n)

		n = 0
		require.NoError(t, ins.Capacity([]byte("abc"), &n))
		assert.Equal(t, 3, n) // []byte: cap

		var fields int
		require.NoError(t, ins.Each(src, func(i int, field string, value any) {
			fields++
			assert.NotEmpty(t, field)
		}))
		assert.Equal(t, 2, fields)

		m := map[string]string{"a": "1", "b": "2"}
		require.NoError(t, ins.Reset(m))
		assert.Empty(t, m)
	})

	t.Run("append not implemented", func(t *testing.T) {
		v, err := ins.Append(src, "x")
		require.ErrorIs(t, err, inspector.ErrNotImplement)
		assert.Equal(t, src, v)
	})

	t.Run("unmarshal", func(t *testing.T) {
		v, err := ins.Unmarshal([]byte(`{"a":"1"}`), inspector.EncodingJSON)
		require.NoError(t, err)
		assert.Equal(t, map[string]string{"a": "1"}, v)

		_, err = ins.Unmarshal(nil, inspector.Encoding(42))
		require.ErrorIs(t, err, inspector.ErrUnknownEncodingType)
	})
}

type collectingIterator struct {
	keys, vals []any
	requireKey bool
	ctl        inspector.LoopCtl
}

func (it *collectingIterator) RequireKey() bool                    { return it.requireKey }
func (it *collectingIterator) SetKey(v any, _ inspector.Inspector) { it.keys = append(it.keys, v) }
func (it *collectingIterator) SetVal(v any, _ inspector.Inspector) { it.vals = append(it.vals, v) }
func (it *collectingIterator) Iterate() inspector.LoopCtl          { return it.ctl }
