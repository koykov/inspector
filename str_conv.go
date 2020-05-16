package inspector

import (
	"errors"
	"strconv"
	"strings"
)

type StrConv struct {
	Snippet string
	Import  []string
}

var (
	convSnippet = map[string]StrConv{}
	tmpCntr     int

	ErrNoConvFunc = errors.New("convert function doesn't exists")
)

func RegisterStrToFunc(typ, snippet string, imports []string) {
	convSnippet[typ] = StrConv{snippet, imports}
}

func StrConvSnippet(s, typn, typu, _var string) (string, []string, error) {
	var (
		sc StrConv
		ok bool
	)
	sc, ok = convSnippet[typn]
	if !ok {
		sc, ok = convSnippet[typu]
	}
	if ok {
		i := tmpIdx()
		snippet := strings.Replace(sc.Snippet, "!{arg}", s, -1)
		snippet = strings.Replace(snippet, "!{var}", _var, -1)
		snippet = strings.Replace(snippet, "!{tmp}", i, -1)
		return snippet, sc.Import, nil
	}
	return "", nil, ErrNoConvFunc
}

func tmpIdx() string {
	i := strconv.Itoa(tmpCntr)
	tmpCntr++
	return i
}
