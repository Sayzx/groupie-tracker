package pkg

import (
	"fmt"
	"net/http"
	"os"
)

func Run() {
	fmt.Println("Initialisation du serveur...")
	// Serveur de fichiers statiques pour les assets
	fs := http.FileServer(http.Dir("./web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("Route pour la page d'accueil")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/index.html")
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/register.html")
	})
	http.HandleFunc("/registerform", registerForm)
	fmt.Println("Server started at http://localhost:9999")
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur: %v\n", err)
	}

}

func registerForm(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	pseudo := r.FormValue("pseudo")
	formattedData := fmt.Sprintf("\n  {\"email\": \"%s\",\n  \"password\": \"%s\",\n  \"pseudo\": \"%s\"\n}", email, password, pseudo)
	fmt.Println(formattedData)

	filePath := "./web/data/account.json"

	// Lire le fichier actuel
	currentData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier: %v\n", err)
	}

	// Vérifier si le fichier est vide
	isEmpty := len(currentData) == 0

	// Ouvrir ou créer le fichier
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier: %v\n", err)
	}
	defer file.Close()

	// Ajouter la virgule si le fichier n'est pas vide
	if !isEmpty {
		_, err = file.WriteString(", ")
		if err != nil {
			fmt.Printf("Erreur lors de l'écriture dans le fichier: %v\n", err)
		}
	}

	// Ajouter les données formatées dans le fichier
	_, err = file.WriteString(fmt.Sprintf("[%s]", formattedData))
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture dans le fichier: %v\n", err)
	}

	// Rediriger l'utilisateur vers la page d'accueil
	http.ServeFile(w, r, "./web/templates/register.html")
}
