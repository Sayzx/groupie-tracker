            // Fonction pour afficher les commentaires
            function displayComments(comments) {
                const container = document.getElementById('comments-container');
                comments.forEach(comment => {
                    const commentDiv = document.createElement('div');
                    commentDiv.classList.add('comment');
                    
                    const avatar = document.createElement('img');
                    avatar.src = comment.discordAvatar ? comment.discordAvatar : '../assets/img/avatar-anonyme.png';
                    avatar.alt = 'Avatar Discord';
                    avatar.classList.add('discord-avatar');
                    
                    const text = document.createElement('p');
                    text.innerText = comment.commentText;
                    
                    const author = document.createElement('p');
                    author.innerText = `Par ${comment.discordName}`;
                    
                    commentDiv.appendChild(avatar);
                    commentDiv.appendChild(text);
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