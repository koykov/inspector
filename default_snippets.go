package inspector

func strToBoolSnippet(typ string) string {
	snippet := "t!{tmp}, err!{tmp} := strconv.ParseBool(!{arg})\n"
	snippet += "if err!{tmp} != nil { return err!{tmp} }\n"
	snippet += "!{var} = " + typ + "(t!{tmp})"
	return snippet
}

func strToByteSnippet(_ string) string {
	snippet := "t!{tmp} := byteconv.S2B(!{arg})\n"
	snippet += "if len(t!{tmp}) > 0{ !{var} = t!{tmp}[0] }\n"
	return snippet
}

func strToIntSnippet(typ string) string {
	snippet := "t!{tmp}, err!{tmp} := strconv.ParseInt(!{arg}, 0, 0)\n"
	snippet += "if err!{tmp} != nil { return err!{tmp} }\n"
	snippet += "!{var} = " + typ + "(t!{tmp})"
	return snippet
}

func strToUintSnippet(typ string) string {
	snippet := "t!{tmp}, err!{tmp} := strconv.ParseUint(!{arg}, 0, 0)\n"
	snippet += "if err!{tmp} != nil { return err!{tmp} }\n"
	snippet += "!{var} = " + typ + "(t!{tmp})"
	return snippet
}

func strToFloatSnippet(typ string) string {
	snippet := "t!{tmp}, err!{tmp} := strconv.ParseFloat(!{arg}, 0)\n"
	snippet += "if err!{tmp} != nil { return err!{tmp} }\n"
	snippet += "!{var} = " + typ + "(t!{tmp})"
	return snippet
}

func strToBytesSnippet(_ string) string {
	snippet := "!{var} = byteconv.S2B(!{arg})\n"
	return snippet
}

func strToStrSnippet(_ string) string {
	snippet := "!{var} = !{arg}\n"
	return snippet
}
