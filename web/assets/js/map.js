document.addEventListener('DOMContentLoaded', function() {
    var map = L.map('map').setView([47.1, 2.4], 4); // Vue initiale centrée sur l'Europe, ajustez selon les besoins
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors'
    }).addTo(map);

    // Extraction des données des concerts à partir des éléments de la liste
    var concertItems = document.querySelectorAll('.relations ul li');
    concertItems.forEach(function(item) {
        var text = item.innerText;
        var parts = text.split(' - ');
        var date = parts[0].replace('Date: ', '');
        var city = parts[1].replace('City: ', '');

        // Convertir la ville en coordonnées (exemple fictif)
        var coords = getCoordinatesForCity(city); // Implémentez cette fonction selon vos besoins

        if (coords) {
            L.marker([coords.lat, coords.lon]).addTo(map)
                .bindPopup('Date: ' + date + '<br>City: ' + city);
        }
    });

    function getCoordinatesForCity(city) {
        var cities = {
            'budapest-hungary': {lat: 47.4979, lon: 19.0402},
            'paris-france': {lat: 48.8566, lon: 2.3522},
            'london-uk': {lat: 51.5074, lon: 0.1278},
            'bogota-colombia': {lat: 4.7110, lon: -74.0721},
            'georgia-usa': {lat: 32.1656, lon: -82.9001},
            'lisbon-portugal': {lat: 38.7223, lon: -9.1393},
            'sao_paulo-brazil': {lat: -23.5505, lon: -46.6333},
            'stockholm-sweden': {lat: 59.3293, lon: 18.0686},
            'werchter-belgium': {lat: 50.9716, lon: 4.7004},
            'bilbao-spain': {lat: 43.2630, lon: -2.9350},
            'dunedin-new_zealand': {lat: -45.8742, lon: 170.5036},
            'los_angeles-usa': {lat: 34.0522, lon: -118.2437},
            'nagoya-japan': {lat: 35.1815, lon: 136.9066},
            'north_carolina-usa': {lat: 35.6301, lon: -79.8064},
            'osaka-japan': {lat: 34.6937, lon: 135.5023},
            'penrose-new_zealand': {lat: -36.9055, lon: 174.8107},
            'saitama-japan': {lat: 35.8616, lon: 139.6455},
            'noumea-new_caledonia': {lat: -22.2558, lon: 166.4505},
            'papeete-french_polynesia': {lat: -17.5516, lon: -149.5585},
            'playa_del_carmen-mexico': {lat: 20.6296, lon: -87.0739},
            'california-usa': {lat: 36.7783, lon: -119.4179},
            'nevada-usa': {lat: 38.8026, lon: -116.4194},
            'yogyakarta-indonesia': {lat: -7.797, lon: 110.370},
            'auckland-new_zealand': {lat: -36.8485, lon: 174.7633},
            'bratislava-slovakia': {lat: 48.1486, lon: 17.1077},
            'minsk-belarus': {lat: 53.9045, lon: 27.5615},
            'new_south_wales-australia': {lat: -33.8688, lon: 151.2093},
            'queensland-australia': {lat: -20.9176, lon: 142.7028},
            'victoria-australia': {lat: -37.8136, lon: 144.9631},
            'new_york-usa': {lat: 40.7128, lon: -74.0060},
            'maine-usa': {lat: 45.2538, lon: -69.4455},
            'abu_dhabi-united_arab_emirates': {lat: 24.4539, lon: 54.3773},
            'pennsylvania-usa': {lat: 41.2033, lon: -77.1945},
            'manchester-uk': {lat: 53.4808, lon: -2.2426},
            'gothenburg-sweden': {lat: 57.7089, lon: 11.9746},
            'aarhus-denmark': {lat: 56.1629, lon: 10.2039},
            'berlin-germany': {lat: 52.5200, lon: 13.4050},
            'coppenhagen-denmark': {lat: 55.6761, lon: 12.5683},
            'west_melbourne-australia': {lat: -37.8136, lon: 144.9631},
            'amsterdam-netherlands': {lat: 52.3676, lon: 4.9041},
            'missouri-usa': {lat: 37.9643, lon: -91.8318},
            'birmingham-uk': {lat: 52.4862, lon: -1.8904},
            'chicago-usa': {lat: 41.8781, lon: -87.6298},
        };
        return cities[city.toLowerCase()];
    }
});