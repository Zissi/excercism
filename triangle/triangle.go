package triangle

import "math"

// Kind gives the type of triangle
type Kind int

// This defines wich types of triangles exist
const (
	NaT = iota
	Equ
	Iso
	Sca
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	var k Kind
	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || a == math.Inf(0) || b == math.Inf(0) || c == math.Inf(0):
		k = NaT
	case a <= 0 || b <= 0 || c <= 0:
		k = NaT
	case a+b < c || b+c < a || a+c < b:
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || b == c || a == c:
		k = Iso
	default:
		k = Sca
	}
	return k
}
