package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initialise et retourne une connexion à la base de données MySQL.
func InitDB() *sql.DB {
	// Tentative d'initialisation et de connexion à la base de données MySQL.
	db, err := sql.Open("mysql", "sayzx:MonNouveauMot2P@sse@tcp(pro.sayzx.fr:3306)/groupie")
	if err != nil {
		// Enregistre l'erreur sans interrompre l'exécution du programme.
		log.Printf("Erreur lors de la connexion à la base de données : %v", err)
		return nil // Retourne nil pour indiquer l'échec de la connexion.
	}

	// Vérifie que la connexion à la base de données est réussie.
	if err := db.Ping(); err != nil {
		// Enregistre l'erreur sans interrompre l'exécution du programme.
		log.Printf("Échec de la connexion à la base de données : %v", err)
		return nil // Retourne nil pour indiquer l'échec de la connexion.
	}
	log.Println("Connexion à la base de données réussie.")
	return db
}
