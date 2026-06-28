package inspector

import "strconv"

func (c *Compiler) writeNodeEach(node_ *node, v string, depth int) error {
	if node_.ptr {
		// Value may be nil on pointer types.
		c.wl("if ", v, " == nil { return nil }")
	}
	tagfb := []string{"json", "xml", "yaml", "bson"}
	for i := 0; i < len(node_.chld); i++ {
		ch := node_.chld[i]
		name := ch.name
		for j := 0; j < len(tagfb); j++ {
			for k := 0; k < len(ch.tags); k++ {
				if ch.tags[k].Key == tagfb[j] && len(ch.tags[k].Name) > 0 {
					name = ch.tags[k].Name
				}
			}
		}
		pfx_ := ""
		if !ch.ptr && (ch.slct == nil || ch.slct.typu == "byte") && ch.mapk == nil {
			pfx_ = "&"
		}
		c.wl("fn(", strconv.Itoa(i), `,"`, name, `",`, pfx_, `x.`, ch.name, ")")
	}

	return nil
}
