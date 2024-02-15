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
	fmt.Printf("Email: %s\nPassword: %s\nPseudo: %s\n", email, password, pseudo)
	// si dans le dossier web/data le dossier account.json existe pas on le crée
	// sinon on ajoute les données dans le fichier account.json
	if _, err := os.Stat("./web/data/account.json"); os.IsNotExist(err) {
		// le fichier n'existe pas
		// on crée le fichier
		file, err := os.Create("./web/data/account.json")
		if err != nil {
			fmt.Printf("Erreur lors de la création du fichier: %v\n", err)
		}
		defer file.Close()
		// on écrit les données dans le fichier
		_, err = file.WriteString(fmt.Sprintf("{\"email\": \"%s\", \"password\": \"%s\", \"pseudo\": \"%s\"}", email, password, pseudo))
		if err != nil {
			fmt.Printf("Erreur lors de l'écriture dans le fichier: %v\n", err)
		}
	} else {
		// le fichier existe
		// on ouvre le fichier
		file, err := os.OpenFile("./web/data/account.json", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Erreur lors de l'ouverture du fichier: %v\n", err)
		}
		defer file.Close()
		// on écrit les données dans le fichier
		_, err = file.WriteString(fmt.Sprintf(",{\"email\": \"%s\", \"password\": \"%s\", \"pseudo\": \"%s\"}", email, password, pseudo))
		if err != nil {
			fmt.Printf("Erreur lors de l'écriture dans le fichier: %v\n", err)
		}
	}
	// on redirige l'utilisateur vers la page d'accueil
	http.ServeFile(w, r, "./web/templates/register.html")
}
