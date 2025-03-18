package game

import (
	"math/rand"
)

// RandomLocation generates a random latitude and longitude.
func RandomLocation() (float64, float64) {
	latitude := rand.Float64()*180 - 90   // Generate latitude between -90 and 90
	longitude := rand.Float64()*360 - 180 // Generate longitude between -180 and 180

	return latitude, longitude
}
