package test

import (
	"testing"

	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
	"github.com/koykov/inspector/testobj_ins"
	"github.com/stretchr/testify/assert"
)

func testGetter(t testing.TB, i inspector.Inspector) {
	id, _ := i.Get(testO, p0...)
	assert.Equal(t, id, "foo")

	name, _ := i.Get(testO, p1...)
	assert.Equal(t, name, expectFoo)

	perm, _ := i.Get(testO, p2...)
	assert.Equal(t, perm, false)

	flag, _ := i.Get(testO, p3...)
	assert.Equal(t, flag, int32(17))

	bal, _ := i.Get(testO, p4...)
	assert.Equal(t, bal, float64(9000))

	date, _ := i.Get(testO, p5...)
	assert.Equal(t, date, int64(153465345246))

	comment, _ := i.Get(testO, p6...)
	assert.Equal(t, comment, expectComment)
}

func testGetterPtr(t testing.TB, i inspector.Inspector, buf any) {
	_ = i.GetTo(testO, &buf, p0...)
	assert.Equal(t, *buf.(*string), "foo")

	_ = i.GetTo(testO, &buf, p1...)
	assert.Equal(t, *buf.(*[]byte), expectFoo)

	_ = i.GetTo(testO, &buf, p2...)
	assert.Equal(t, *buf.(*bool), false)

	_ = i.GetTo(testO, &buf, p3...)
	assert.Equal(t, *buf.(*int32), int32(17))

	_ = i.GetTo(testO, &buf, p4...)
	assert.Equal(t, *buf.(*float64), float64(9000))

	_ = i.GetTo(testO, &buf, p5...)
	assert.Equal(t, *buf.(*int64), int64(153465345246))

	_ = i.GetTo(testO, &buf, p6...)
	assert.Equal(t, *buf.(*[]byte), expectComment)
}

func testComparePtr(t testing.TB, i inspector.Inspector, buf *bool) {
	_ = i.Compare(testO, inspector.OpEq, "foo", buf, p0...)
	assert.False(t, !*buf)

	_ = i.Compare(testO, inspector.OpEq, "bar", buf, p1...)
	assert.False(t, !*buf)

	_ = i.Compare(testO, inspector.OpGtq, "60", buf, p7...)
	assert.False(t, !*buf)

	_ = i.Compare(testO, inspector.OpLtq, "5000", buf, p8...)
	assert.False(t, !*buf)

	_ = i.Compare(testO, inspector.OpEq, "true", buf, p9...)
	assert.False(t, !*buf)
}

func testSetterPtr(t testing.TB, i inspector.Inspector, ab inspector.AccumulativeBuffer) {
	var cins testobj_ins.TestObjectInspector
	obj1, _ := cins.Copy(testO)
	obj := obj1.(*testobj.TestObject)
	obj.Id = ""
	_ = i.SetWithBuffer(obj, 1984, ab, p0...)
	assert.Equal(t, obj.Id, "1984")

	_ = i.SetWithBuffer(obj, 2000, ab, p1...)
	assert.Equal(t, obj.Name, expectName)

	_ = i.SetWithBuffer(obj, false, ab, p2...)
	assert.False(t, (*obj.Permission)[23])

	_ = i.SetWithBuffer(obj, int32(23), ab, p3...)
	assert.Equal(t, obj.Flags["export"], int32(23))

	_ = i.SetWithBuffer(obj, float64(9000), ab, p4...)
	assert.Equal(t, obj.Finance.Balance, float64(9000))

	_ = i.SetWithBuffer(obj, int64(153465345246), ab, p5...)
	assert.Equal(t, obj.Finance.History[1].DateUnix, int64(153465345246))

	_ = i.SetWithBuffer(obj, &expectComment1, ab, p6...)
	assert.Equal(t, obj.Finance.History[0].Comment, expectComment1)
}
