package cast

import "strconv"

func StringToInt(input string) int {
	var val int
	val, err := strconv.Atoi(input)
	if err != nil {
		panic("failed parse" + err.Error())
	}
	return val
}

func RuneToInt(input rune) int {
	s := string(input)
	return StringToInt(s)
}

func ToString(input int) string {
	return strconv.Itoa(input)
}
