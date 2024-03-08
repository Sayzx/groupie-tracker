console.log('search.js loaded');
let allArtists = [];
let debounceTimeout;

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
            suggestionElement.className = 'suggestion-item';
            suggestionElement.innerHTML = `
                <img src="${artist.image}" alt="${artist.name}" class="suggestion-image">
                <a href="/artist_info?id=${artist.id}" class="suggestion-link">${artist.name}</a>
            `;
            suggestionsContainer.appendChild(suggestionElement);
        });
        suggestionsContainer.style.display = 'block';
    } else {
        suggestionsContainer.style.display = 'none';
    }
}


function displayResults(artists) {
    const resultsDiv = document.querySelector('.results');
    resultsDiv.innerHTML = '';

    if (artists.length > 0) {
        artists.forEach((artist, index) => {
            const artistHtml = `
                <a href="/artist_info?id=${artist.id}" class="artist-link">
                    <div class="artist-result" style="transition-delay: ${index * 50}ms">
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

document.addEventListener('DOMContentLoaded', function() {
    // Chargement initial des artistes
    fetch('/api/search/artists')
        .then(response => response.json())
        .then(data => {
            allArtists = data;
            displayResults(allArtists); // Affiche tous les artistes par défaut
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