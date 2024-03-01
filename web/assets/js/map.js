document.addEventListener('DOMContentLoaded', function() {
    var map = L.map('map').setView([47.1, 2.4], 4);
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors'
    }).addTo(map);


    // Ici je fait un call sur cette url pour récupérer les coordonnées des villes
    fetch('https://groupie.sayzx.fr/api/sayzx.json')
    .then(response => response.json())
    .then(cities => {
        var concertItems = document.querySelectorAll('.relations ul li');
        concertItems.forEach(function(item) {
            var text = item.innerText;
            var parts = text.split(' - ');
            var date = parts[0].replace('Date: ', '');
            var city = parts[1].replace('City: ', '');
            var cityKey = city.toLowerCase().replace(/ /g, '_');

            // Je récupère les coordonnées de la ville
            var coords = cities[cityKey];

            // Si les coordonnées existent, j'ajoute un marqueur sur la carte
            if (coords) {
                L.marker([coords.lat, coords.lon]).addTo(map)
                    .bindPopup('Date: ' + date + '<br>City: ' + city);
            }
        });
    })
    .catch(error => console.error('Erreur lors de la récupération des données:', error));
});
