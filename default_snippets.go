package inspector

import (
	"math/rand"
	"strconv"
)

func strToIntSnippet(typ string) string {
	i := strconv.Itoa(rand.Int())
	return "t" + i + ", err" + i + " := strconv.ParseInt(!{arg}, 0, 0)\n!{var} = " + typ + "(t" + i + ")"
}

func strToUintSnippet(typ string) string {
	i := strconv.Itoa(rand.Int())
	return "t" + i + ", err" + i + " := strconv.ParseUint(!{arg}, 0, 0)\n!{var} = " + typ + "(t" + i + ")"
}

func strToFloatSnippet(typ string) string {
	i := strconv.Itoa(rand.Int())
	return "t" + i + ", err" + i + " := strconv.ParseFloat(!{arg}, 0)\n!{var} = " + typ + "(t" + i + ")"
}
