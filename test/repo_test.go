package test

import (
	"github.com/koykov/inspector"
	"github.com/koykov/inspector/testobj"
)

var (
	testO = &testobj.TestObject{
		Id:         "foo",
		Name:       []byte("bar"),
		Cost:       12.34,
		Status:     78,
		Permission: &testobj.TestPermission{15: true, 23: false},
		Flags: testobj.TestFlag{
			"export": 17,
			"ro":     4,
			"rw":     7,
			"Valid":  1,
		},
		Finance: &testobj.TestFinance{
			MoneyIn:  3200,
			MoneyOut: 1500.637657,
			Balance:  9000,
			AllowBuy: true,
			History: []testobj.TestHistory{
				{
					152354345634,
					14.345241,
					[]byte("pay for domain"),
				},
				{
					153465345246,
					-3.0000342543,
					[]byte("got refund"),
				},
				{
					156436535640,
					2325242534.35324523,
					[]byte("maintenance"),
				},
			},
		},
	}

	p0 = []string{"Id"}
	p1 = []string{"Name"}
	p2 = []string{"Permission", "23"}
	p3 = []string{"Flags", "export"}
	p4 = []string{"Finance", "Balance"}
	p5 = []string{"Finance", "History", "1", "DateUnix"}
	p6 = []string{"Finance", "History", "0", "Comment"}
	p7 = []string{"Status"}
	p8 = []string{"Finance", "MoneyIn"}
	p9 = []string{"Finance", "AllowBuy"}

	expectFoo      = []byte("bar")
	expectName     = []byte("2000")
	expectComment  = []byte("pay for domain")
	expectComment1 = []byte("lorem ipsum dolor sit amet")

	deqStages = []struct {
		l, r *testobj.TestObject
		opts *inspector.DEQOptions
		eq   bool
	}{
		{
			l:  nil,
			r:  nil,
			eq: true,
		},
		{
			l:  &testobj.TestObject{Id: "foo"},
			r:  &testobj.TestObject{Id: "bar"},
			eq: false,
		},
		{
			l:  &testobj.TestObject{Cost: 3.1415},
			r:  &testobj.TestObject{Cost: 3.1415},
			eq: true,
		},
		{
			l:  &testobj.TestObject{Name: []byte("qwe")},
			r:  &testobj.TestObject{Name: []byte("rty")},
			eq: false,
		},
		{
			l:  &testobj.TestObject{Flags: testobj.TestFlag{"xxx": 15}},
			r:  &testobj.TestObject{Flags: testobj.TestFlag{"xxx": 15}},
			eq: true,
		},
		{
			l:  &testobj.TestObject{Flags: testobj.TestFlag{"xxx": 15}},
			r:  &testobj.TestObject{Flags: testobj.TestFlag{"xxx": 54}},
			eq: false,
		},
		{
			l:  &testobj.TestObject{Flags: testobj.TestFlag{"xxx": 15}},
			r:  &testobj.TestObject{Flags: testobj.TestFlag{"yyy": 15}},
			eq: false,
		},
		{
			l:  &testobj.TestObject{Finance: &testobj.TestFinance{History: []testobj.TestHistory{{Cost: 22}}}},
			r:  &testobj.TestObject{Finance: &testobj.TestFinance{History: []testobj.TestHistory{{Cost: 22}}}},
			eq: true,
		},
		{
			l:  &testobj.TestObject{Finance: &testobj.TestFinance{History: []testobj.TestHistory{{Comment: []byte("aaa")}}}},
			r:  &testobj.TestObject{Finance: &testobj.TestFinance{History: []testobj.TestHistory{{}}}},
			eq: false,
		},
		{
			l:    &testobj.TestObject{Id: "foobar", Name: []byte("qwe")},
			r:    &testobj.TestObject{Id: "foobar", Name: []byte("rty")},
			opts: &inspector.DEQOptions{Exclude: map[string]struct{}{"Name": {}}},
			eq:   true,
		},
		{
			l:    &testobj.TestObject{Id: "foobar", Name: []byte("qwe")},
			r:    &testobj.TestObject{Id: "foobar", Name: []byte("rty")},
			opts: &inspector.DEQOptions{Filter: map[string]struct{}{"Id": {}}},
			eq:   true,
		},
	}
)
