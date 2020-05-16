package inspector

func strToBoolSnippet(typ string) string {
	snippet := "t!{tmp}, err!{tmp} := strconv.ParseBool(!{arg})\n"
	snippet += "if err!{tmp} != nil { return err!{tmp} }\n"
	snippet += "!{var} = " + typ + "(t!{tmp})"
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

func strToBytesSnippet(typ string) string {
	snippet := "!{var} = fastconv.S2B(!{arg})\n"
	return snippet
}

func strToStrSnippet(_ string) string {
	snippet := "!{var} = !{arg}\n"
	return snippet
}
