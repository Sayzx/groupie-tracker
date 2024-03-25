-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Hôte : localhost:3306
-- Généré le : lun. 25 mars 2024 à 16:06
-- Version du serveur :  8.0.36-0ubuntu0.20.04.1
-- Version de PHP : 7.4.3-4ubuntu2.20

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `groupie`
--

-- --------------------------------------------------------

--
-- Structure de la table `comments`
--

CREATE TABLE `comments` (
  `id` int UNSIGNED NOT NULL,
  `discord_name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `discord_avatar` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `comment` text COLLATE utf8mb4_general_ci,
  `artist_id` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `comments`
--

INSERT INTO `comments` (`id`, `discord_name`, `discord_avatar`, `comment`, `artist_id`) VALUES
(9, 'sayzx', 'https://cdn.discordapp.com/avatars/826826070899949601/a_8861167ea70612ca8256716976822f2a.png', 'Vraiment un artiste incroyable', 4),
(10, 'sayzx', 'https://cdn.discordapp.com/avatars/826826070899949601/a_8861167ea70612ca8256716976822f2a.png', 'Je t\'aime', 5),
(11, 'sayzx', 'https://cdn.discordapp.com/avatars/826826070899949601/a_8861167ea70612ca8256716976822f2a.png', 'Je suis payday', 5),
(12, 'sayzx', 'https://cdn.discordapp.com/avatars/826826070899949601/a_8861167ea70612ca8256716976822f2a.png', 'J\'aime ce mmec', 2),
(13, 'sayzx', 'https://cdn.discordapp.com/avatars/826826070899949601/a_8861167ea70612ca8256716976822f2a.png', 'Je pense qu\'enzo chante aussi bien que lui', 4),
(14, 'topwin', 'https://cdn.discordapp.com/avatars/366514410635657216/6a1034f11c127f7ce4da4f35fc4680df.png', 'e', 3),
(15, 'topwin', 'https://cdn.discordapp.com/avatars/366514410635657216/6a1034f11c127f7ce4da4f35fc4680df.png', 'pipi', 5),
(16, 'sayzx', 'https://cdn.discordapp.com/avatars/826826070899949601/a_8861167ea70612ca8256716976822f2a.png', 'Pink Floyd est vraimennt trop gentil', 3),
(17, 'sayzx', 'https://cdn.discordapp.com/avatars/826826070899949601/a_8861167ea70612ca8256716976822f2a.png', 'dev', 3),
(18, 'sayzx', 'https://cdn.discordapp.com/avatars/826826070899949601/a_8861167ea70612ca8256716976822f2a.png', 'Je suis un developpeur de énis', 3),
(19, 'sayzx', 'https://cdn.discordapp.com/avatars/826826070899949601/a_8861167ea70612ca8256716976822f2a.png', 'Dommage il est mort il faisais bien le taff', 6),
(20, 'sayzx', 'https://cdn.discordapp.com/avatars/826826070899949601/a_8861167ea70612ca8256716976822f2a.png', 'Je reste coentre les artiste c vrm le feu', 7),
(21, 'topwin', 'https://cdn.discordapp.com/avatars/366514410635657216/6a1034f11c127f7ce4da4f35fc4680df.png', 'pipi', 5),
(22, 'topwin', 'https://cdn.discordapp.com/avatars/366514410635657216/6a1034f11c127f7ce4da4f35fc4680df.png', 'momo', 5),
(23, 'topwin', 'https://cdn.discordapp.com/avatars/366514410635657216/6a1034f11c127f7ce4da4f35fc4680df.png', 'salut', 4),
(24, 'topwin', 'https://cdn.discordapp.com/avatars/366514410635657216/6a1034f11c127f7ce4da4f35fc4680df.png', 'caca', 47),
(25, 'topwin', 'https://cdn.discordapp.com/avatars/366514410635657216/6a1034f11c127f7ce4da4f35fc4680df.png', 'salut chef', 5),
(26, 'topwin', 'https://cdn.discordapp.com/avatars/366514410635657216/6a1034f11c127f7ce4da4f35fc4680df.png', 'sauce soja ou quoi la team xDDDDDD', 2),
(27, 'topwin', 'https://cdn.discordapp.com/avatars/366514410635657216/6a1034f11c127f7ce4da4f35fc4680df.png', 'Vive la france ! <3', 18);

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `comments`
--
ALTER TABLE `comments`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `comments`
--
ALTER TABLE `comments`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
