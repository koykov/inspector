package inspector

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTags(t *testing.T) {
	require.Nil(t, parseTags(""))
	require.Nil(t, parseTags("garbage"))

	got := parseTags(`json:"id" xml:"id,attr" yaml:"-"`)
	require.Len(t, got, 2)
	require.Equal(t, "json", got[0].Key)
	require.Equal(t, "id", got[0].Name)
	require.Empty(t, got[0].Options)
	require.Equal(t, "xml", got[1].Key)
	require.Equal(t, "id", got[1].Name)
	require.Equal(t, []string{"attr"}, got[1].Options)
}

func TestComposeAstTypeName(t *testing.T) {
	c := &Compiler{}
	require.Equal(t, "map[string]int", c.composeAstTypeName(&node{
		typ: typeMap, mapk: &node{typn: "string"}, mapv: &node{typn: "int"},
	}))
	require.Equal(t, "map[*string]*int", c.composeAstTypeName(&node{
		typ: typeMap, mapk: &node{typn: "string", ptr: true}, mapv: &node{typn: "int", ptr: true},
	}))
	require.Equal(t, "[]int", c.composeAstTypeName(&node{typ: typeSlice, slct: &node{typn: "int"}}))
	require.Equal(t, "[]*int", c.composeAstTypeName(&node{typ: typeSlice, slct: &node{typn: "int", ptr: true}}))
	require.Empty(t, c.composeAstTypeName(&node{typ: typeStruct}))
}

func TestFormatHelpers(t *testing.T) {
	c := &Compiler{}
	ptrNode, valNode := &node{ptr: true}, &node{}
	require.Equal(t, "(*v)", c.fmtV(ptrNode, "v"))
	require.Equal(t, "(v)", c.fmtV(valNode, "v"))
	require.Equal(t, "(*v)", c.fmtVd(valNode, "v", 0))
	require.Equal(t, "(v)", c.fmtVd(valNode, "v", 1))
	require.Equal(t, "*v", c.fmtVnb(valNode, "v", 0))
	require.Equal(t, "v", c.fmtVnb(valNode, "v", 1))
	require.Equal(t, "&v", c.fmtP(valNode, "v", 0))
	require.Equal(t, "v", c.fmtP(valNode, "v", 1))
	require.Equal(t, "Foo", c.fmtTfs(&node{typ: typeStruct, typn: "Foo"}, false))
	require.Equal(t, "testobj.Foo", c.fmtTfs(&node{typ: typeStruct, typn: "Foo", pkg: "testobj"}, false))
	require.Equal(t, "*testobj.Foo", c.fmtTfs(&node{typ: typeStruct, typn: "Foo", pkg: "testobj", ptr: true}, true))
}

func TestStrConvSnippet(t *testing.T) {
	old := make(map[string]StrConv, len(convSnippetRegistry))
	for k, v := range convSnippetRegistry {
		old[k] = v
	}
	oldCntr := tmpCntr
	t.Cleanup(func() {
		convSnippetRegistry = old
		tmpCntr = oldCntr
	})

	RegisterStrToXFn("testType", "parse(!{arg}, !{var}, !{tmp})", []string{"example.com/conv"})
	got, imports, err := StrConvSnippet("src", "testType", "", "dst")
	require.NoError(t, err)
	require.Equal(t, []string{"example.com/conv"}, imports)
	require.Contains(t, got, "parse(src, dst, ")

	_, _, err = StrConvSnippet("src", "nope", "nope2", "dst")
	require.ErrorIs(t, err, ErrNoConvFunc)
}

func TestConfigCopy(t *testing.T) {
	bl := map[string]struct{}{"X": {}}
	orig := &Config{
		Target: TargetFile, File: "f.go", Import: "i", Destination: "d",
		BlackList: bl, NoClean: true, XML: "x", Force: true,
	}
	cp := orig.Copy()
	require.Equal(t, orig.Target, cp.Target)
	require.Equal(t, orig.File, cp.File)
	require.True(t, cp.NoClean)
	cp.BlackList["Y"] = struct{}{}
	require.Contains(t, orig.BlackList, "Y") // Copy shallow: карта общая
}

func TestNodeWriteXML(t *testing.T) {
	var buf bytes.Buffer
	c := newFixtureCompiler(t, &buf)
	parseFixture(t, c, fixtureSrc)
	require.NotEmpty(t, c.nodes)

	var x1, x2 bytes.Buffer
	require.NoError(t, c.nodes[0].write(&x1))
	require.NoError(t, c.nodes[0].write(&x2))
	require.Equal(t, x1.String(), x2.String())
	require.Contains(t, x1.String(), `<?xml version="1.0" encoding="UTF-8"?>`)
	require.Contains(t, x1.String(), "<node ")
}
