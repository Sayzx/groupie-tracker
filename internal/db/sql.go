package db

import (
	"database/sql"
	"log"
)

func initDB() *sql.DB {
	// Initialisation et connexion à la base de données MySQL.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/groupie")
	if err != nil {
		log.Fatalf("Erreur lors de la connexion à la base de données: %v\n", err)
	}
	// Vérifiez que la connexion à la base de données est réussie.
	if err := db.Ping(); err != nil {
		log.Fatalf("Échec de la connexion à la base de données: %v\n", err)
	}
	return db
}
