package inspector

import (
	"regexp"
	"strings"
)

type Tag struct {
	Key, Name string
	Options   []string
}

var reTagParse = regexp.MustCompile(`([^:]+):"([^"]*)"\s*`)

func parseTags(raw string) []Tag {
	if len(raw) == 0 {
		return nil
	}
	if m := reTagParse.FindAllStringSubmatch(raw, -1); len(m) > 0 {
		var dst []Tag
		for i := 0; i < len(m); i++ {
			mm := m[i]
			var t Tag
			t.Key = mm[1]
			if t.Options = strings.Split(mm[2], ","); len(t.Options) > 0 {
				t.Name = t.Options[0]
				t.Options = t.Options[1:]
			}
			if t.Name == "-" {
				continue
			}
			dst = append(dst, t)
		}
		return dst
	}
	return nil
}
