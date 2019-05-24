package space

// Planet type is needed for tests to pass
type Planet string

const earthYear = 31557600.0

var planetFactors = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age takes an age in seconds and returns how old someone would be on a different planet.
func Age(seconds float64, planet Planet) float64 {
	planetFactor := planetFactors[planet]
	return seconds / (earthYear * planetFactor)
}
