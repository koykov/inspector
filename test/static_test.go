package test

import (
	"testing"

	"github.com/koykov/inspector"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type countingIterator struct {
	keys, vals int
	ctl        inspector.LoopCtl
}

func (it *countingIterator) RequireKey() bool                { return true }
func (it *countingIterator) SetKey(any, inspector.Inspector) { it.keys++ }
func (it *countingIterator) SetVal(any, inspector.Inspector) { it.vals++ }
func (it *countingIterator) Iterate() inspector.LoopCtl      { return it.ctl }

func TestStaticInspector(t *testing.T) {
	ins := inspector.StaticInspector{}

	t.Run("type name", func(t *testing.T) {
		assert.Equal(t, "static", ins.TypeName())
	})

	t.Run("get passthrough", func(t *testing.T) {
		v, err := ins.Get(int64(42))
		require.NoError(t, err)
		assert.Equal(t, int64(42), v)

		var buf any
		require.NoError(t, ins.GetTo("x", &buf))
		assert.Equal(t, "x", buf)

		buf = "sentinel"
		require.NoError(t, ins.GetTo(nil, &buf))
		assert.Nil(t, buf)
	})

	t.Run("set pointers", func(t *testing.T) {
		var i32 int32
		require.NoError(t, ins.Set(&i32, int64(42)))
		assert.Equal(t, int32(42), i32)

		var u64 uint64
		require.NoError(t, ins.Set(&u64, uint64(7)))
		assert.Equal(t, uint64(7), u64)

		var f64 float64
		require.NoError(t, ins.Set(&f64, 3.5))
		assert.Equal(t, 3.5, f64)

		var b bool
		require.NoError(t, ins.Set(&b, true))
		assert.True(t, b)

		var s string
		require.NoError(t, ins.Set(&s, []byte("abc")))
		assert.Equal(t, "abc", s)

		var bs []byte
		require.NoError(t, ins.Set(&bs, "db"))
		assert.Equal(t, []byte("db"), bs)
	})

	t.Run("set with buffer reuse", func(t *testing.T) {
		var ab inspector.ByteBuffer
		var s1, s2 string
		require.NoError(t, ins.SetWithBuffer(&s1, "one", &ab))
		require.NoError(t, ins.SetWithBuffer(&s2, "two", &ab))
		assert.Equal(t, "one", s1)
		assert.Equal(t, "two", s2)

		var b1, b2 []byte
		require.NoError(t, ins.SetWithBuffer(&b1, "aa", &ab))
		require.NoError(t, ins.SetWithBuffer(&b2, "bb", &ab))
		assert.Equal(t, []byte("aa"), b1)
		assert.Equal(t, []byte("bb"), b2)
	})

	t.Run("set edge", func(t *testing.T) {
		require.ErrorIs(t, ins.Set(int32(0), 1), inspector.ErrMustPointerType)

		i := int64(7)
		require.NoError(t, ins.Set(&i, "42")) // indInt не принимает string → тихий no-op
		assert.Equal(t, int64(7), i)

		require.NoError(t, ins.Set(&i, float64(3.9))) // float→int truncation
		assert.Equal(t, int64(3), i)

		var st struct{}
		require.NoError(t, ins.Set(&st, 1))
	})

	t.Run("compare", func(t *testing.T) {
		var res bool
		require.NoError(t, ins.Compare(int64(10), inspector.OpGtq, "10", &res))
		assert.True(t, res)
		require.NoError(t, ins.Compare("abc", inspector.OpEq, "abc", &res))
		assert.True(t, res)
		require.NoError(t, ins.Compare([]byte("ab"), inspector.OpNq, "cd", &res))
		assert.True(t, res)
		require.NoError(t, ins.Compare(true, inspector.OpEq, "true", &res))
		assert.True(t, res)

		var u uint = 5
		require.NoError(t, ins.Compare(u, inspector.OpEq, "5", &res))
		assert.True(t, res)
		require.NoError(t, ins.Compare(u, inspector.OpNq, "5", &res))
		assert.False(t, res)
		require.NoError(t, ins.Compare(u, inspector.OpGtq, "5", &res))
		assert.True(t, res)

		require.NoError(t, ins.Compare(float64(1.5), inspector.OpLt, "2.5", &res))
		assert.True(t, res)

		i64 := int64(10)
		require.NoError(t, ins.Compare(&i64, inspector.OpEq, "10", &res))
		assert.True(t, res)
		require.NoError(t, ins.Compare(&i64, inspector.OpGtq, "11", &res))
		assert.False(t, res)

		res = true
		require.NoError(t, ins.Compare(10, inspector.OpGtq, "bad", &res))
		assert.True(t, res) // parse error: result untouched

		res = true
		require.NoError(t, ins.Compare(struct{}{}, inspector.OpEq, "1", &res))
		assert.False(t, res) // unsupported src

		// unsupported op: current behavior is a false result, not an untouched one
		// (cmp* returns false for OpInc and Compare assigns it).
		res = true
		require.NoError(t, ins.Compare(int64(10), inspector.OpInc, "10", &res))
		assert.False(t, res)
	})

	t.Run("deep equal", func(t *testing.T) {
		assert.True(t, ins.DeepEqual(1, int64(1)))
		assert.False(t, ins.DeepEqual(1, "1"))
		assert.True(t, ins.DeepEqual(0.1, 0.1005))
		assert.False(t, ins.DeepEqual(0.1, 0.12))
		assert.True(t, ins.DeepEqual("a", "a"))
		assert.True(t, ins.DeepEqual([]byte("a"), []byte("a")))
		assert.False(t, ins.DeepEqual(nil, nil))

		// StaticInspector ignores DEQOptions: it always uses FloatPrecision (1e-3),
		// not the requested 0.1. Therefore, 0.1 vs 0.12 (diff 0.02) is still unequal.
		assert.False(t, ins.DeepEqualWithOptions(0.1, 0.12, &inspector.DEQOptions{Precision: 0.1}))
		assert.True(t, ins.DeepEqualWithOptions(0.1, 0.1005, &inspector.DEQOptions{Precision: 0.1}))
	})

	t.Run("copy", func(t *testing.T) {
		cp, err := ins.Copy([]byte("ab"))
		require.NoError(t, err)
		assert.Equal(t, []byte("ab"), cp)
		_, err = ins.Copy(nil)
		require.ErrorIs(t, err, inspector.ErrUnsupportedType)

		var dst int64
		require.NoError(t, ins.CopyTo(int64(5), &dst, &inspector.ByteBuffer{}))
		assert.Equal(t, int64(5), dst)
	})

	t.Run("append loop each", func(t *testing.T) {
		v, err := ins.Append(5, 6)
		require.NoError(t, err)
		assert.Equal(t, 5, v)
		require.NoError(t, ins.Each(1, func(int, string, any) { t.Fatal("no fields expected") }))
		it := &countingIterator{}
		require.NoError(t, ins.Loop(1, it, new([]byte)))
		assert.Zero(t, it.keys)
		assert.Zero(t, it.vals)
	})

	t.Run("length capacity", func(t *testing.T) {
		var n int
		require.NoError(t, ins.Length([]byte("abc"), &n))
		assert.Equal(t, 3, n)
		require.NoError(t, ins.Length("abcd", &n))
		assert.Equal(t, 4, n)
		require.NoError(t, ins.Length(1, &n))
		assert.Equal(t, 0, n)

		b := make([]byte, 2, 10)
		require.NoError(t, ins.Capacity(b, &n))
		assert.Equal(t, 10, n)
	})

	t.Run("reset", func(t *testing.T) {
		i := 5
		require.NoError(t, ins.Reset(&i))
		assert.Equal(t, 0, i)

		b := []byte("ab")
		require.NoError(t, ins.Reset(b))
		assert.Equal(t, []byte("ab"), b) // текущее поведение: без мутации
		require.NoError(t, ins.Reset(nil))
	})

	t.Run("unmarshal", func(t *testing.T) {
		v, err := ins.Unmarshal([]byte(`{"a":1}`), inspector.EncodingJSON)
		require.NoError(t, err)
		assert.Equal(t, map[string]any{"a": float64(1)}, v)

		_, err = ins.Unmarshal(nil, inspector.Encoding(42))
		require.ErrorIs(t, err, inspector.ErrUnknownEncodingType)
	})
}
