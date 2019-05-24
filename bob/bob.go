package bob

import "strings"

// Hey takes a remark and returns Bobs answer.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	answer := ""
	switch {
	case remark == "":
		answer = "Fine. Be that way!"
	case strings.ToUpper(remark) == remark && string(remark[len(remark)-1]) == "?" && strings.ToLower(remark) != remark:
		answer = "Calm down, I know what I'm doing!"
	case string(remark[len(remark)-1]) == "?":
		answer = "Sure."
	case strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark:
		answer = "Whoa, chill out!"
	default:
		answer = "Whatever."
	}
	return answer
}
