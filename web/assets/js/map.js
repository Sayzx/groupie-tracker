document.addEventListener('DOMContentLoaded', function() {
    var map = L.map('map').setView([0, 0], 2); // Initialiser la carte avec une vue globale

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
        maxZoom: 18,
    }).addTo(map);

    function getLatLng(location) {
        const locations = {
            // Liste des lieux prédéfinie
        };
        return locations[location] || [0, 0];
    }

    var artistId = document.getElementById('id').value;
    var apiUrl = `https://groupietrackers.herokuapp.com/api/relation/${artistId}`;
    console.log(apiUrl);

    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            const locations = data.datesLocations;
            for (let location in locations) {
                const latLng = getLatLng(location);
                const dates = locations[location].join(", ");
                L.marker(latLng).addTo(map)
                    .bindPopup(`<b>${location.replace(/-/g, " ")}</b><br>${dates}`);
            }
            map.fitBounds(L.featureGroup(Object.keys(locations).map(location => L.marker(getLatLng(location)))).getBounds());
        })
        .catch(error => console.error('Erreur :', error));
});
