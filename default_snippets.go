package inspector

import (
	"math/rand"
	"strconv"
)

func strToIntSnippet(typ string) string {
	i := strconv.Itoa(rand.Intn(99))
	snippet := "t" + i + ", err" + i + " := strconv.ParseInt(!{arg}, 0, 0)\n"
	snippet += "if err" + i + " != nil { return err" + i + " }\n"
	snippet += "!{var} = " + typ + "(t" + i + ")"
	return snippet
}

func strToUintSnippet(typ string) string {
	i := strconv.Itoa(rand.Intn(99))
	snippet := "t" + i + ", err" + i + " := strconv.ParseUint(!{arg}, 0, 0)\n"
	snippet += "if err" + i + " != nil { return err" + i + " }\n"
	snippet += "!{var} = " + typ + "(t" + i + ")"
	return snippet
}

func strToFloatSnippet(typ string) string {
	i := strconv.Itoa(rand.Intn(99))
	snippet := "t" + i + ", err" + i + " := strconv.ParseFloat(!{arg}, 0)\n"
	snippet += "if err" + i + " != nil { return err" + i + " }\n"
	snippet += "!{var} = " + typ + "(t" + i + ")"
	return snippet
}
