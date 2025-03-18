import { MapContainer, TileLayer, useMapEvents, Marker } from "react-leaflet";
import "leaflet/dist/leaflet.css";
import { useState } from "react";

function MapComponent({ setGuess }) {
    const [marker, setMarker] = useState(null);

    const MapClickHandler = () => {
        useMapEvents({
            click(e) {
                const { lat, lng } = e.latlng;
                setMarker([lat, lng]);
                setGuess({ latitude: lat, longitude: lng });
            },
        });
        return null;
    };

    return (
        <MapContainer center={[20, 0]} zoom={2} style={{ height: "500px", width: "100%" }}>
            <TileLayer
                url="https://basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png"
                attribution='&copy; <a href="https://carto.com/">CARTO</a> contributors'
            />
            <MapClickHandler />
            {marker && <Marker position={marker} />}
        </MapContainer>
    );
}

export default MapComponent;
