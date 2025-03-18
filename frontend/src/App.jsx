import { useState } from "react";
import MapComponent from "./MapComponent";
import "./styles.css";

function calculateDistance(lat1, lon1, lat2, lon2) {
    const R = 6371; // Radius of Earth in km
    const toRad = (degree) => (degree * Math.PI) / 180;

    const dLat = toRad(lat2 - lat1);
    const dLon = toRad(lon2 - lon1);

    const a =
        Math.sin(dLat / 2) * Math.sin(dLat / 2) +
        Math.cos(toRad(lat1)) * Math.cos(toRad(lat2)) *
        Math.sin(dLon / 2) * Math.sin(dLon / 2);

    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
    return R * c; // Distance in km
}

function calculateScore(distance) {
    const maxScore = 1000;
    const worldCircumference = 40075; // in km
    return Math.max(0, maxScore - (distance / worldCircumference) * maxScore);
}

function App() {
    const [guess, setGuess] = useState(null);
    const actual = { latitude: 20, longitude: 0 }; // Fixed actual location
    const [score, setScore] = useState(null);

    const handleGuess = () => {
        if (!guess) return alert("Please select a location on the map!");

        const distance = calculateDistance(
            guess.latitude, guess.longitude,
            actual.latitude, actual.longitude
        );

        setScore(calculateScore(distance).toFixed(2));
    };

    return (
        <div>
            <h1>Hello, GeoGuessr Clone!</h1>
            <MapComponent setGuess={setGuess} guess={guess} />
            <button onClick={handleGuess} style={{ marginTop: "10px", padding: "10px", cursor: "pointer" }}>
                Make a Guess
            </button>
            {score !== null && <h2>Score: {score}</h2>}
        </div>
    );
}

export default App;
