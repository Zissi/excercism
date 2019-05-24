package acronym

import "strings"

// Abbreviate takes a space separated string of words and returns the abbreviation of it.
func Abbreviate(s string) string {
	words := strings.Fields(s)
	abbreviation := ""
	for _, word := range words {
		abbreviation += string(word[0])
	}
	return abbreviation
}
