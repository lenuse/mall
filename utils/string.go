package utils

func StringToUpper(str string, sep rune) string {
	temp := []rune(str)
	var i int
	targetKey := -3
	var newStr []rune
	for k, v := range temp {
		if k == 0 || (k == targetKey+1) {
			newStr = append(newStr, v-32)
			goto add
		}
		if v == sep {
			targetKey = k
			continue
		}
		newStr = append(newStr, v)
	add:
		i = i + 1
	}
	return string(newStr)
}
