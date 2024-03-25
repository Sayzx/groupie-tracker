

![App](https://media.discordapp.net/attachments/1012749489402023956/1221826582797226154/Sans_titre-3.png?ex=6613fdad&is=660188ad&hm=6360f9c6dafbeb80b4772fe8c398c8ac044af0770935e3eeda58e6cbcb6400e0&=&format=webp&quality=lossless&width=1440&height=240)


## Notre Projet

Groupie Tracker a duré un mois et demi, période pendant laquelle nous avons pu mettre nos compétences à l'épreuve en manipulant de multiples API et en faisant face à de nombreux problèmes de requêtes via JavaScript. Par exemple, nous avons découvert SQL à travers des fonctionnalités bonus et avons également appris à manier les API en Go."

## Languages

#### Front-End

HTML : Pour la structure des pages 
CSS : Pour le style des pages 
JS : Pour avoir des pages dynamique et des effects sur les pagres

#### Back-End

Goland: Pour gérer les données et faire des call API
SQL : Pour stocker des données comme les commentaires 


## Features

### API Multi-Parties
Intégration et gestion de quatre APIs distinctes offrant des détails sur les artistes, leurs lieux et dates de concerts, ainsi que les liens entre ces éléments.

### Affichage Intuitif des Données
Présentation des données des artistes à l'aide de visualisations variées, telles que blocs, cartes, tableaux, et graphiques.

### Filtrage Avancé
Utilisateurs peuvent filtrer les artistes selon plusieurs critères :
- Date de création
- Date du premier album
- Nombre de membres
- Lieux de concerts
Avec des filtres par plage de valeurs et par sélection multiple.

### Géolocalisation des Concerts
Affichage des emplacements de concerts sur une carte, grâce à la conversion des adresses en coordonnées géographiques.

### Barre de Recherche Dynamique
Fonction de recherche pour localiser facilement artistes, membres de groupe, et lieux, avec des suggestions de saisie automatique et indication du type de suggestion.

## Fonctionnalités Supplémentaires

- **Connexion Discord** : Connexion au site via Discord pour une expérience utilisateur améliorée.
- **Système de Commentaires** : Espace permettant aux utilisateurs de laisser des commentaires sur les pages des artistes, avec stockage des données en SQL.
- **Écoute des Artistes** : Intégration de fonctionnalités permettant l'écoute de musique directement depuis les pages des artistes.
- **Mode Clair/Sombre** : Possibilité pour les utilisateurs de basculer entre un thème clair et sombre.



## Comment Lancer le Projet ?

Pour démarrer le projet Groupie Tracker, assurez-vous d'avoir un environnement Go configuré sur votre machine. Suivez ces étapes pour installer et exécuter le projet :

### Prérequis

- Go installé sur votre machine.
- Accès à un terminal ou invite de commande.

### Installation et Exécution

1. **Clonage du dépôt :** Clonez le dépôt du projet avec la commande suivante :

```bash
git clone https://github.com/Sayzx/groupie-tracker/
cd groupie-tracker
go get
go run cmd/main.go
```

## Configuration de la Base de Données Personnalisée ( Facultatif )

Pour utiliser votre propre serveur SQL avec Groupie Tracker ( Ce n'est pas obligatoire c'est seulement si vous le voulez), veuillez suivre ces instructions :

1. **Préparation de la Base de Données :** Commencez par préparer votre base de données en exécutant le script SQL fourni dans `internal/db/sql.sql`. Ce script créera la structure nécessaire pour stocker les données du projet.

2. **Modification des Paramètres de Connexion :** Ensuite, ouvrez le fichier de configuration de la base de données situé à `internal/db/sqlgo`. À la ligne 13, vous trouverez les paramètres de connexion à la base de données (tels que le nom d'utilisateur, le mot de passe, l'adresse du serveur, le nom de la base de données, etc.).

   Remplacez les valeurs par défaut par celles correspondant à votre environnement de base de données SQL. Voici un exemple de ce à quoi pourrait ressembler cette modification :
   
   `sql.Open("mysql", "user:mdp@tcp(IP:3306)/database")`

   


👏 Bonne découverte  ici : `http://localhost:8080`




## Demo





![App](https://media.discordapp.net/attachments/1012749489402023956/1221833480799785060/image.png?ex=6614041a&is=66018f1a&hm=ccbaec249252d1e8c35092b38294b5a1a11881895aa69268bf5be4206ed860d3&=&format=webp&quality=lossless&width=1305&height=662)

