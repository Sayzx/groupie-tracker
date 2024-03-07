let allArtists = [];
let debounceTimeout;

document.addEventListener('DOMContentLoaded', function() {
    // Supposez que vous chargez tous les artistes ici
    fetch('/api/search/artists')
        .then(response => response.json())
        .then(data => {
            allArtists = data; // Adaptez cette ligne selon la structure de votre réponse API
            displayResults(allArtists); // Affichez tous les artistes par défaut
        })
        .catch(error => console.error('Error:', error));

    const searchInput = document.getElementById('searchInput');
    searchInput.addEventListener('input', () => {
        const input = searchInput.value.trim().toLowerCase();
        const filteredArtists = input ? allArtists.filter(artist => artist.name.toLowerCase().includes(input)) : allArtists;
        displayResults(filteredArtists); // Affiche les artistes filtrés
    });
});

function displayResults(artists) {
    const resultsDiv = document.querySelector('.results');
    resultsDiv.innerHTML = ''; // Efface les résultats précédents

    artists.forEach(artist => {
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
        resultsDiv.innerHTML += artistHtml; // Ajoute la carte de l'artiste aux résultats
    });
}

function searchArtist() {
    var input = document.getElementById('searchInput').value.trim().toLowerCase();

    clearTimeout(debounceTimeout);
    debounceTimeout = setTimeout(() => {
        if (input.length > 0) {
            // Filtrez les artistes basé sur l'input
            let filteredArtists = allArtists.filter(artist => artist.name.toLowerCase().includes(input));
            showSuggestions(filteredArtists);
        } else {
            document.getElementById('suggestions').style.display = 'none';
        }
    }, 300);
}

function showSuggestions(suggestions) {
    var suggestionsContainer = document.getElementById('suggestions');
    suggestionsContainer.innerHTML = '';
    if (suggestions.length > 0) {
        suggestions.forEach(artist => {
            var suggestionElement = document.createElement('div');
            suggestionElement.innerHTML = `<a href="/artist_info?id=${artist.id}">${artist.name}</a>`;
            suggestionsContainer.appendChild(suggestionElement);
        });
        suggestionsContainer.style.display = 'block';
    } else {
        suggestionsContainer.style.display = 'none';
    }
}
