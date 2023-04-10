package inspector

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/koykov/fastconv"
)

// StringsInspector is a built-in inspector for []string/[][]byte types.
type StringsInspector struct{}

func (i StringsInspector) TypeName() string {
	return "strings"
}

func (i StringsInspector) Get(src any, path ...string) (any, error) {
	var x any
	err := i.GetTo(src, &x, path...)
	return x, err
}

func (i StringsInspector) GetTo(src any, buf *any, path ...string) error {
	if len(path) != 1 {
		return nil
	}
	ss, pp, ok := i.sp(src)
	if !ok {
		return nil
	}
	idx, err := strconv.Atoi(path[0])
	if err != nil {
		return err
	}
	if idx < 0 {
		return nil
	}
	switch {
	case len(ss) > 0 && idx < len(ss):
		*buf = &ss[idx]
	case len(pp) > 0 && idx < len(pp):
		*buf = &pp[idx]
	}
	return nil
}

func (i StringsInspector) Set(dst, value any, path ...string) error {
	var buf ByteBuffer
	return i.SetWithBuffer(dst, value, &buf, path...)
}

func (i StringsInspector) SetWithBuffer(dst, value any, buf AccumulativeBuffer, path ...string) error {
	if len(path) != 1 {
		return nil
	}
	ss, pp, ok := i.sp(dst)
	if !ok {
		return nil
	}
	idx, err := strconv.Atoi(path[0])
	if err != nil {
		return err
	}
	if idx < 0 {
		return nil
	}
	var p []byte
	switch {
	case len(ss) > 0 && idx < len(ss):
		switch value.(type) {
		case string:
			p = fastconv.S2B(value.(string))
		case *string:
			p = fastconv.S2B(*value.(*string))
		}
		if len(p) > 0 {
			ss[idx] = fastconv.B2S(buf.Bufferize(p))
		}
	case len(pp) > 0 && idx < len(pp):
		switch value.(type) {
		case []byte:
			p = value.([]byte)
		case *[]byte:
			p = *value.(*[]byte)
		}
		if len(p) > 0 {
			pp[idx] = buf.Bufferize(p)
		}
	}
	return nil
}

func (i StringsInspector) Compare(src any, cond Op, right string, result *bool, path ...string) error {
	if len(path) != 1 {
		return nil
	}
	ss, pp, ok := i.sp(src)
	if !ok {
		return nil
	}
	idx, err := strconv.Atoi(path[0])
	if err != nil {
		return err
	}
	if idx < 0 {
		return nil
	}
	var s string
	switch {
	case len(ss) > 0 && idx < len(ss):
		s = ss[idx]
	case len(pp) > 0 && idx < len(pp):
		s = fastconv.B2S(pp[idx])
	}
	switch cond {
	case OpNq:
		*result = s != right
	case OpEq:
		*result = s == right
	case OpGt:
		*result = s > right
	case OpGtq:
		*result = s >= right
	case OpLt:
		*result = s < right
	case OpLtq:
		*result = s <= right
	}
	return nil
}

func (i StringsInspector) Loop(src any, l Iterator, buf *[]byte, path ...string) error {
	if len(path) > 0 {
		return nil
	}
	ss, pp, ok := i.sp(src)
	if !ok {
		return nil
	}

	switch {
	case len(ss) > 0:
		for j := 0; j < len(ss); j++ {
			if l.RequireKey() {
				*buf = strconv.AppendInt((*buf)[:0], int64(j), 10)
				l.SetKey(buf, StaticInspector{})
			}
			l.SetVal(&ss[j], StaticInspector{})
			ctl := l.Iterate()
			if ctl == LoopCtlBrk {
				break
			}
			if ctl == LoopCtlCnt {
				continue
			}
		}
	case len(pp) > 0:
		for j := 0; j < len(pp); j++ {
			if l.RequireKey() {
				*buf = strconv.AppendInt((*buf)[:0], int64(j), 10)
				l.SetKey(buf, StaticInspector{})
			}
			l.SetVal(&pp[j], StaticInspector{})
			ctl := l.Iterate()
			if ctl == LoopCtlBrk {
				break
			}
			if ctl == LoopCtlCnt {
				continue
			}
		}
	}
	return nil
}

func (i StringsInspector) DeepEqual(l, r any) bool {
	return i.DeepEqualWithOptions(l, r, nil)
}

func (i StringsInspector) DeepEqualWithOptions(l, r any, _ *DEQOptions) bool {
	ssL, ppL, okL := i.sp(l)
	if !okL {
		return false
	}
	ssR, ppR, okR := i.sp(r)
	if !okR {
		return false
	}
	ssLn, ssRn, ppLn, ppRn := len(ssL), len(ssR), len(ppL), len(ppR)
	switch {
	case ssLn > 0 && ssRn > 0 && ssLn == ssRn:
		for j := 0; j < ssLn; j++ {
			if ssL[j] != ssR[j] {
				return false
			}
		}
		return true
	case ssLn > 0 && ppRn > 0 && ssLn == ppRn:
		for j := 0; j < ssLn; j++ {
			if ssL[j] != fastconv.B2S(ppR[j]) {
				return false
			}
		}
		return true
	case ppLn > 0 && ssRn > 0 && ppLn == ssRn:
		for j := 0; j < ppLn; j++ {
			if fastconv.B2S(ppL[j]) != ssR[j] {
				return false
			}
		}
		return true
	case ppLn > 0 && ppRn > 0 && ppLn == ppRn:
		for j := 0; j < ppLn; j++ {
			if !bytes.Equal(ppL[j], ppR[j]) {
				return false
			}
		}
		return true
	}
	return false
}

func (i StringsInspector) Unmarshal(p []byte, typ Encoding) (any, error) {
	var x []string
	switch typ {
	case EncodingJSON:
		err := json.Unmarshal(p, &x)
		return &x, err
	default:
		return nil, ErrUnknownEncodingType
	}
}

func (i StringsInspector) Copy(x any) (any, error) {
	var buf ByteBuffer
	var dst []string
	err := i.CopyTo(x, &dst, &buf)
	return dst, err
}

func (i StringsInspector) CopyTo(src, dst any, buf AccumulativeBuffer) error {
	ssR, ppR, okR := i.sp(src)
	if !okR {
		return ErrUnsupportedType
	}

	var (
		ss *[]string
		pp *[][]byte
	)
	switch dst.(type) {
	case []string:
		return ErrMustPointerType
	case *[]string:
		ss = dst.(*[]string)
		switch {
		case len(ssR) > 0:
			for j := 0; j < len(ssR); j++ {
				cpy := buf.BufferizeString(ssR[j])
				*ss = append(*ss, cpy)
			}
		case len(ppR) > 0:
			for j := 0; j < len(ppR); j++ {
				cpy := buf.Bufferize(ppR[j])
				*ss = append(*ss, fastconv.B2S(cpy))
			}
		}
	case [][]byte:
		return ErrMustPointerType
	case *[][]byte:
		pp = dst.(*[][]byte)
		switch {
		case len(ssR) > 0:
			for j := 0; j < len(ssR); j++ {
				cpy := buf.BufferizeString(ssR[j])
				*pp = append(*pp, fastconv.S2B(cpy))
			}
		case len(ppR) > 0:
			for j := 0; j < len(ppR); j++ {
				cpy := buf.Bufferize(ppR[j])
				*pp = append(*pp, cpy)
			}
		}
	default:
		return ErrUnsupportedType
	}
	return nil
}

func (i StringsInspector) Length(src any, result *int, path ...string) error {
	ss, pp, ok := i.sp(src)
	if !ok {
		return nil
	}

	if len(path) == 1 {
		idx, err := strconv.Atoi(path[0])
		if err != nil {
			return err
		}
		switch {
		case len(ss) > 0 && idx >= 0 && idx < len(ss):
			*result = len(ss[idx])
		case len(pp) > 0 && idx >= 0 && idx < len(pp):
			*result = len(pp[idx])
		}
	} else {
		if len(ss) > 0 {
			*result = len(ss)
		}
		if len(pp) > 0 {
			*result = len(pp)
		}
	}
	return nil
}

func (i StringsInspector) Capacity(src any, result *int, path ...string) error {
	_, pp, ok := i.sp(src)
	if !ok {
		return nil
	}

	if len(path) == 1 {
		idx, err := strconv.Atoi(path[0])
		if err != nil {
			return err
		}
		if len(pp) > 0 && idx >= 0 && idx < len(pp) {
			*result = cap(pp[idx])
		}
	} else {
		if len(pp) > 0 {
			*result = cap(pp)
		}
	}
	return nil
}

func (i StringsInspector) Reset(x any) error {
	switch x.(type) {
	case []string:
		return ErrMustPointerType
	case *[]string:
		ss := x.(*[]string)
		*ss = (*ss)[:0]
	case [][]byte:
		return ErrMustPointerType
	case *[][]byte:
		pp := x.(*[][]byte)
		*pp = (*pp)[:0]
	}
	return nil
}

func (i StringsInspector) sp(x any) (ss []string, pp [][]byte, ok bool) {
	ok = true
	switch x.(type) {
	case []string:
		ss = x.([]string)
	case *[]string:
		ss = *(x.(*[]string))
	case [][]byte:
		pp = x.([][]byte)
	case *[][]byte:
		pp = *(x.(*[][]byte))
	default:
		ok = false
	}
	return
}
