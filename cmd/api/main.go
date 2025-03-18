package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sync"

	"guezzer/internal/game" //our game logic
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var (
	currentLocation Location
	mu              sync.Mutex // Mutex to protect the currentLocation variable
)

func haversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in kilo meters
	lat1Rad, lon1Rad := lat1*math.Pi/180, lon1*math.Pi/180
	lat2Rad, lon2Rad := lat2*math.Pi/180, lon2*math.Pi/180

	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c // Distance in km
}

// getRandomLocation handles API requests and returns a random location.
func getRandomLocation(w http.ResponseWriter, r *http.Request) {
	latitude, longitude := game.RandomLocation() // Call the function from game.go
	// response := Location{
	// 	Latitude:  latitude,
	// 	Longitude: longitude,
	// }
	mu.Lock()
	currentLocation = Location{Latitude: latitude, Longitude: longitude}
	mu.Unlock()

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentLocation)
}

// calculateScore calculates the score based on the distance
func calculateScore(distance float64) int {
	score := int(1000 * math.Exp(-distance/5000)) //exponential decay formula for dynamic scoring
	if score < 1 {
		score = 1
	}
	return score
}

func guessLocation(w http.ResponseWriter, r *http.Request) {
	var userGuess Location
	err := json.NewDecoder(r.Body).Decode(&userGuess)
	if err != nil {
		http.Error(w, "invalid json input", http.StatusBadRequest)
		return
	}
	mu.Lock()
	distance := haversineDistance(userGuess.Latitude, userGuess.Longitude, currentLocation.Latitude, currentLocation.Longitude)
	score := calculateScore(distance)
	mu.Unlock()

	//send response with distance
	response := map[string]interface{}{
		"distance": distance,
		"actual":   currentLocation,
		"guess":    userGuess,
		"score":    score,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	// 	if err := json.NewEncoder(w).Encode(response); err != nil {
	// 	http.Error(w, "Failed to generate response", http.StatusInternalServerError)
	// Respond with the random location in JSON format
	// fmt.Fprintf(w, `{"latitude": %f, "longitude": %f}`, latitude, longitude)
}

func main() {
	// Define the route for the API
	http.HandleFunc("/random-location", getRandomLocation)
	http.HandleFunc("/guess", guessLocation)

	// Start the HTTP server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
