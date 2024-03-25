document.addEventListener('DOMContentLoaded', function() {
    const toggleIcon = document.getElementById('theme-toggle');
    toggleIcon.addEventListener('click', function() {
        document.body.classList.toggle('light-mode');

        if (document.body.classList.contains('light-mode')) {
            this.style.color = '#FFA500'; // Icône jaune en mode clair
            applyLightModeStyles();
        } else {
            this.style.color = '#fcfcfc'; // Icône blanche en mode sombre
            applyDarkModeStyles();
        }
    });
});

function applyLightModeStyles() {
    document.body.style.backgroundColor = '#f5f5f5';
    // Utilisation de #1DB954 pour les titres h1, provenant de lightmode2.js
    document.querySelectorAll('h1').forEach(el => el.style.color = '#1DB954');
    document.querySelectorAll('.card-container').forEach(el => el.style.backgroundColor = '#fff');
    document.querySelectorAll('.artist-card').forEach(el => el.style.backgroundColor = '#fff');
    document.querySelectorAll('p').forEach(el => el.style.color = '#555');
    // Ajout du style pour .rounded-input1 de lightmode3.js
    document.querySelectorAll('.rounded-input1').forEach(el => el.style.backgroundColor = '#fff');
}

function applyDarkModeStyles() {
    document.body.style.backgroundColor = '#333';
    document.querySelectorAll('h1').forEach(el => el.style.color = '#1DB954');
    document.querySelectorAll('.card-container').forEach(el => el.style.backgroundColor = '#333');
    document.querySelectorAll('.artist-card').forEach(el => el.style.backgroundColor = '#1f1f1f');
    document.querySelectorAll('p').forEach(el => el.style.color = '#ffffff');
    // Ajout du style pour .rounded-input1 de lightmode3.js
    document.querySelectorAll('.rounded-input1').forEach(el => el.style.backgroundColor = '#444');
}
