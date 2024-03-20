console.log('search.js loaded');
let allArtists = [];
let debounceTimeout;

let cityToArtistsMap = {}; // Cette map va associer les villes aux artistes

function searchArtist() {
    var input = document.getElementById('searchInput').value.trim().toLowerCase();

    clearTimeout(debounceTimeout);
    debounceTimeout = setTimeout(() => {
        if (input.length > 0) {
            let filteredSuggestions = [];
            let suggestionsSet = new Set();

            allArtists.forEach(artist => {
                if (artist.name.toLowerCase().includes(input) && !suggestionsSet.has(artist.id)) {
                    filteredSuggestions.push({ label: `${artist.name} - Artiste`, data: artist });
                    suggestionsSet.add(artist.id);
                }

                if (artist.members) {
                    artist.members.forEach(member => {
                        if (member.toLowerCase().includes(input) && !suggestionsSet.has(artist.id)) {
                            filteredSuggestions.push({ label: `${member} - Membre`, data: artist });
                            suggestionsSet.add(artist.id);
                        }
                    });
                }

                if (artist.firstAlbum && artist.firstAlbum.toLowerCase().includes(input) && !suggestionsSet.has(artist.id)) {
                    filteredSuggestions.push({ label: `${artist.name} (${artist.firstAlbum}) - Première album`, data: artist });
                    suggestionsSet.add(artist.id);
                }

                if (artist.creationDate && artist.creationDate.toString().includes(input) && !suggestionsSet.has(artist.id)) {
                    filteredSuggestions.push({ label: `${artist.name} (${artist.creationDate}) - Date de création`, data: artist });
                    suggestionsSet.add(artist.id);
                }
            });

            if (input.length >= 3) {
                Object.keys(cityToArtistsMap).forEach(city => {
                    if (city.includes(input)) {
                        cityToArtistsMap[city].forEach(artistInCity => {
                            if (!suggestionsSet.has(artistInCity.id)) {
                                filteredSuggestions.push({ label: `${artistInCity.name} (${city}) - Ville de concert`, data: artistInCity });
                                suggestionsSet.add(artistInCity.id);
                            }
                        });
                    }
                });
            }

            showSuggestions(filteredSuggestions);
        } else {
            document.getElementById('suggestions').style.display = 'none';
        }
    }, 300);
}

function showSuggestions(suggestions) {
    var suggestionsContainer = document.getElementById('suggestions');
    suggestionsContainer.innerHTML = '';
    if (suggestions.length > 0) {
        suggestions.forEach(suggestion => {
            var suggestionElement = document.createElement('div');
            suggestionElement.className = 'suggestion-item';
            suggestionElement.innerHTML = `
                <img src="${suggestion.data.image}" alt="${suggestion.data.name}" class="suggestion-image">
                <a href="/artist_info?id=${suggestion.data.id}" class="suggestion-link">${suggestion.label}</a>
            `;
            suggestionsContainer.appendChild(suggestionElement);
        });
        suggestionsContainer.style.display = 'block';
    } else {
        suggestionsContainer.style.display = 'none';
    }
}

function loadLocationsAndMapToArtists() {
    fetch('/api/locations') // Utilisation du proxy pour contourner la politique CORS
        .then(response => response.json())
        .then(data => {
            // Supposons que data.index contient vos données de locations
            data.index.forEach(location => {
                location.locations.forEach(city => {
                    // Pour chaque ville, nous associons les artistes qui y ont joué
                    if (!cityToArtistsMap[city]) cityToArtistsMap[city] = [];
                    // Ici, au lieu d'ajouter juste l'id, on cherche l'artiste correspondant dans allArtists
                    const artistDetails = allArtists.find(artist => artist.id === location.id);
                    if (artistDetails) {
                        cityToArtistsMap[city].push(artistDetails);
                    }
                });
            });
        })
        .catch(error => console.error('Error loading locations:', error));
}

document.addEventListener('DOMContentLoaded', function() {
    // Chargement initial des artistes
    fetch('/api/search/artists')
        .then(response => response.json())
        .then(data => {
            allArtists = data;
            displayResults(allArtists); // Affiche tous les artistes par défaut
            populateCityToArtistsMap(allArtists);
            loadLocationsAndMapToArtists();
        })
        .catch(error => console.error('Error:', error));

    // Empêcher la soumission standard du formulaire et filtrer sur la base des entrées
    document.getElementById('searchForm').addEventListener('submit', function(event) {
        event.preventDefault(); // Empêche la soumission standard du formulaire
        filterAndDisplayArtists();
    });

    // Gestionnaire pour le filtrage dynamique à la saisie
    document.getElementById('searchInput').addEventListener('input', filterAndDisplayArtists);

    // Gestionnaire pour le filtrage dynamique des filtres
    document.getElementById('yearSelect').addEventListener('change', filterAndDisplayArtists);
    document.getElementById('creationYearSelect').addEventListener('change', filterAndDisplayArtists);
    document.getElementById('membersSelect').addEventListener('change', filterAndDisplayArtists);

    // Toggle pour les filtres supplémentaires
    document.getElementById('filterToggle').addEventListener('click', function() {
        var filtersPanel = document.getElementById('filtersPanel');
        filtersPanel.style.display = filtersPanel.style.display === 'block' ? 'none' : 'block';
    });

    document.getElementById('resetFilters').addEventListener('click', function() {
        document.getElementById('searchInput').value = '';
        document.getElementById('yearSelect').selectedIndex = 0;
        document.getElementById('creationYearSelect').selectedIndex = 0;
        document.getElementById('membersSelect').selectedIndex = 0;
        filterAndDisplayArtists(); // Affiche tous les artistes par défaut après la réinitialisation
    });
});

function populateCityToArtistsMap(artists) {
    cityToArtistsMap = artists.reduce((map, artist) => {
        let locations = artist.locations;
        // Si 'locations' n'est pas un tableau, mais une chaîne, convertissons-le en tableau
        if (typeof locations === 'string') {
            locations = [locations];
        }

        if (Array.isArray(locations)) {
            locations.forEach(city => {
                if (!map[city]) map[city] = [];
                map[city].push(artist);
            });
        }
        return map;
    }, {});
}

function displayResults(artists) {
    const resultsDiv = document.querySelector('.results');
    resultsDiv.innerHTML = '';

    if (artists.length > 0) {
        artists.forEach((artist, index) => {
            const artistHtml = `
            <a href="/artist_info?id=${artist.id}" class="artist-link">
                <div class="artist-result">
                    <img src="${artist.image}" alt="${artist.name}">
                    <div class="info">
                        <h2>${artist.name}</h2>
                        <p>Date de première sortie : ${artist.firstAlbum}</p>
                        <p>Nombre de membres : ${artist.members.length}</p>
                    </div>
                </div>
            </a>
        `;
        resultsDiv.innerHTML += artistHtml;
});

        // Applique l'animation après que les éléments sont ajoutés au DOM
        requestAnimationFrame(() => {
            const artistElements = resultsDiv.querySelectorAll('.artist-result');
            artistElements.forEach(el => el.classList.add('show'));
        });
    } else {
        resultsDiv.innerHTML = `<div class="no-results">Aucun Résultats...</div>`;
    }
}

function filterAndDisplayArtists() {
    const nameInput = document.getElementById('searchInput').value.trim().toLowerCase();
    const yearInput = document.getElementById('yearSelect').value;
    const creationYearInput = document.getElementById('creationYearSelect').value;
    const membersInput = document.getElementById('membersSelect').value;

    // Filtrez les artistes basé sur les inputs et les sélecteurs
    const filteredArtists = allArtists.filter(artist => {
        return (nameInput === "" || artist.name.toLowerCase().includes(nameInput)) &&
               (yearInput === "" || artist.firstAlbum.endsWith(yearInput)) &&
               (creationYearInput === "" || artist.creationDate === parseInt(creationYearInput)) &&
               (membersInput === "" || artist.members.length === parseInt(membersInput));
    });

    displayResults(filteredArtists);
}
