package counter

import "strings"

//Wcounter
func Wcounter(content string) map[string]int {
	if content == "" {
		return map[string]int{}
	}
	report := make(map[string]int)
	words := strings.FieldsFunc(content, spliters) //strings.Split(content, " ")
	for _, w := range words {
		report[w]++
	}
	return report
}

func spliters(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t' || r == ',' || r == '.'
}
