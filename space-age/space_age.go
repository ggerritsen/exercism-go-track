// Package space provides functionality around calculating how old a person is on different planets
package space

// Planet is a type of planet
type Planet string

var earthYearsPerPlanet = map[Planet]float64 {
	"Mercury": 0.2408467,
	"Venus": 0.61519726,
	"Mars": 1.8808158,
	"Jupiter": 11.862615,
	"Saturn": 29.447498,
	"Uranus": 84.016846,
	"Neptune": 164.79132,
	"Earth": 1,
}

// Age will calculate how old a person is, given an age in seconds and a planet
func Age (seconds float64, planet Planet) float64 {
	return seconds / 31557600 / earthYearsPerPlanet[planet]
}
