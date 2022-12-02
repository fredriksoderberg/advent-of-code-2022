package cast

import "strconv"

func ToInt(input string) int {
	var val int
	val, err := strconv.Atoi(input)
	if err != nil {
		panic("failed parse" + err.Error())
	}
	return val
}

func ToString(input int) string {
	return strconv.Itoa(input)
}
