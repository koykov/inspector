package inspector

import (
	"errors"
	"strings"
)

type StrConv struct {
	Snippet string
	Import  []string
}

var (
	convSnippet = map[string]StrConv{}

	ErrNoConvFunc = errors.New("convert function doesn't exists")
)

func RegisterStrToFunc(typ, snippet string, imports []string) {
	convSnippet[typ] = StrConv{snippet, imports}
}

func StrConvSnippet(s, typ, _var string) (string, []string, error) {
	if sc, ok := convSnippet[typ]; ok {
		snippet := strings.Replace(sc.Snippet, "!{arg}", s, -1)
		snippet = strings.Replace(snippet, "!{var}", _var, -1)
		return snippet, sc.Import, nil
	}
	return "", nil, ErrNoConvFunc
}
