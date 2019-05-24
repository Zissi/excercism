package strand

var transcriptionMap = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA takes a DNA sequence and returns the matching RNA sequence.
func ToRNA(dna string) string {
	rna := ""
	for _, base := range dna {
		rna += transcriptionMap[string(base)]
	}
	return rna

}
