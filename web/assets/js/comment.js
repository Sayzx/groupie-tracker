function displayComments(comments) {
    const container = document.getElementById('comments-container');
    comments.forEach(comment => {
        const commentDiv = document.createElement('div');
        commentDiv.classList.add('comment');
        
        // Création du conteneur pour l'avatar et le texte
        const contentDiv = document.createElement('div');
        contentDiv.classList.add('comment-content');
        
        const avatar = document.createElement('img');
        avatar.src = comment.discordAvatar ? comment.discordAvatar : '../assets/img/avatar-anonyme.png';
        avatar.alt = 'Avatar Discord';
        avatar.classList.add('discord-avatar');
        
        const text = document.createElement('p');
        text.innerText = comment.commentText;
        text.classList.add('comment-text');
        
        // Ajouter l'avatar et le texte au conteneur de contenu
        contentDiv.appendChild(avatar);
        contentDiv.appendChild(text);
        
        // Ajout du contenu et de l'auteur au div principal du commentaire
        commentDiv.appendChild(contentDiv);
        
        const author = document.createElement('p');
        author.innerText = `Par ${comment.discordName}`;
        author.classList.add('comment-author');
        commentDiv.appendChild(author);
        
        container.appendChild(commentDiv);
    });
}


// Requête à l'API pour obtenir les commentaires
fetch('/api/comments')
    .then(response => response.json())
    .then(comments => {
        displayComments(comments);
    })
    .catch(error => console.error('Erreur:', error));




// Requête à l'API pour obtenir les commentaires pour un artiste spécifique
function fetchCommentsForArtist() {
    const commentSection = document.getElementById('comment-section');
    const artistID = commentSection.getAttribute('data-artist-id');

    fetch(`/api/comments?id=${artistID}`) // Utilisez l'ID dans la requête
        .then(response => response.json())
        .then(comments => {
            displayComments(comments);
        })
        .catch(error => console.error('Erreur:', error));
}

// Assurez-vous que cette fonction est appelée lorsque la page est chargée
document.addEventListener('DOMContentLoaded', fetchCommentsForArtist);
