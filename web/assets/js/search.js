console.log('search.js loaded');
function searchArtist() {
    var input = document.getElementById('searchInput').value;
    if (input.length > 0) {
        fetch('/api/search/artists?query=' + encodeURIComponent(input))
            .then(response => response.json())
            .then(data => {
                var suggestions = data.filter(artist => artist.name.toLowerCase().includes(input.toLowerCase()));
                showSuggestions(suggestions);
            })
            .catch(error => console.error('Error:', error));
    } else {
        document.getElementById('suggestions').style.display = 'none';
    }
}

function showSuggestions(suggestions) {
    var suggestionsContainer = document.getElementById('suggestions');
    suggestionsContainer.innerHTML = '';
    if (suggestions.length > 0) {
        suggestions.forEach(artist => {
            var suggestionElement = document.createElement('div');
            suggestionElement.innerHTML = `<a href="/artist_info?id=${artist.id}">${artist.name}</a>`; // Fixed template literal
            suggestionsContainer.appendChild(suggestionElement);
        });
        suggestionsContainer.style.display = 'block';
    } else {
        suggestionsContainer.style.display = 'none';
    }
}


function displayResults(data) {
    var resultsDiv = document.querySelector('.results');
    resultsDiv.innerHTML = ''; 

    data.forEach(artist => {
        var artistHtml = `<div class="artist-result">
            <img src="${artist.Image}" alt="${artist.Name}">
            <div class="info">
                <h2>${artist.Name}</h2>
                <p>Date de première sortie : ${artist.FirstAlbum}</p>
                <p>Nombre de membres : ${artist.Members.length}</p>
            </div>
        </div>`;
        resultsDiv.innerHTML += artistHtml;
    });
}

    document.getElementById('filterToggle').addEventListener('click', function() {
        var filtersPanel = document.getElementById('filtersPanel');
        filtersPanel.style.display = filtersPanel.style.display === 'block' ? 'none' : 'block';
    });

    document.addEventListener('DOMContentLoaded', function() {
        var yearSelect = document.getElementById('yearSelect');
        if (yearSelect.length === 1) { // S'assure qu'il n'y a pas déjà des options ajoutées
            for (var year = 1960; year <= 2024; year++) {
                var option = document.createElement('option');
                option.value = year;
                option.textContent = year;
                yearSelect.appendChild(option);
            }
        }
    });
