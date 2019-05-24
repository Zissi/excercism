package raindrops

import "strconv"

// Convert takes an integer and returns the rain sound of it.
func Convert(input int) string {
	sound := ""
	if input%3 == 0 {
		sound += "Pling"
	}
	if input%5 == 0 {
		sound += "Plang"
	}
	if input%7 == 0 {
		sound += "Plong"
	}
	if sound == "" {
		sound = strconv.Itoa(input)
	}
	return sound
}
