package hamming

import "errors"

// Distance takes two base sequence as strings and returns the hamming distance.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("hamming distance is not defined for unequal length sequences")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}

	}
	return distance, nil
}
