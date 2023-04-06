package inspector

import (
	"errors"
	"strconv"
	"strings"
)

// StrConv describes string to X conversion object.
type StrConv struct {
	// Code snippet.
	Snippet string
	// Imports list.
	Import []string
}

var (
	// Conversion snippets registry.
	convSnippetRegistry = map[string]StrConv{}
	// Counter for variables suffixes to avoid variable names duplicates.
	tmpCntr int

	ErrNoConvFunc = errors.New("convert function doesn't exists")
)

// RegisterStrToXFn registers new snippet.
func RegisterStrToXFn(typ, snippet string, imports []string) {
	convSnippetRegistry[typ] = StrConv{snippet, imports}
}

// StrConvSnippet builds a conversion snippet according arguments.
func StrConvSnippet(s, typn, typu, _var string) (string, []string, error) {
	var (
		sc StrConv
		ok bool
	)
	sc, ok = convSnippetRegistry[typn]
	if !ok {
		sc, ok = convSnippetRegistry[typu]
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
