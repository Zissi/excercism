package proverb

import "fmt"

// Proverb takes a list of pieces and creates a proverb from them.
func Proverb(rhyme []string) []string {
	templates := []string{"For want of a %s the %s was lost.",
		"And all for the want of a %s."}
	proverb := make([]string, len(rhyme))
	for position, piece := range rhyme {
		if position < len(rhyme)-1 {
			proverb[position] = fmt.Sprintf(templates[0], piece, rhyme[position+1])
		} else {
			proverb[position] = fmt.Sprintf(templates[1], rhyme[0])
		}
	}
	return proverb
}
