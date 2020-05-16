package inspector

import (
	"strconv"
)

var (
	tmpCntr int
)

func strToBoolSnippet(typ string) string {
	i := tmpIdx()
	snippet := "t" + i + ", err" + i + " := strconv.ParseBool(!{arg})\n"
	snippet += "if err" + i + " != nil { return err" + i + " }\n"
	snippet += "!{var} = " + typ + "(t" + i + ")"
	return snippet
}

func strToIntSnippet(typ string) string {
	i := tmpIdx()
	snippet := "t" + i + ", err" + i + " := strconv.ParseInt(!{arg}, 0, 0)\n"
	snippet += "if err" + i + " != nil { return err" + i + " }\n"
	snippet += "!{var} = " + typ + "(t" + i + ")"
	return snippet
}

func strToUintSnippet(typ string) string {
	i := tmpIdx()
	snippet := "t" + i + ", err" + i + " := strconv.ParseUint(!{arg}, 0, 0)\n"
	snippet += "if err" + i + " != nil { return err" + i + " }\n"
	snippet += "!{var} = " + typ + "(t" + i + ")"
	return snippet
}

func strToFloatSnippet(typ string) string {
	i := tmpIdx()
	snippet := "t" + i + ", err" + i + " := strconv.ParseFloat(!{arg}, 0)\n"
	snippet += "if err" + i + " != nil { return err" + i + " }\n"
	snippet += "!{var} = " + typ + "(t" + i + ")"
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

func tmpIdx() string {
	i := strconv.Itoa(tmpCntr)
	tmpCntr++
	return i
}
