package server

import (
	"fauxrome/mysql/ConnectAndDisconnect"
	SearchIntoTables "fauxrome/mysql/search"
	"fauxrome/mysql/update"
	structures "fauxrome/server/Structures"
	"fmt"
	"net/http"
	"os"
)

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	db, _ := ConnectAndDisconnect.ConnectToBDD_Mysql()
	username := structures.User_Connected
	fmt.Println(username)
	var profil structures.Profil_Search
	var profils []structures.Profil_Search
	profil, _ = SearchIntoTables.SearchByUserIntoProfil(db, username, profil, profils)

	if r.Method == http.MethodPost {
		modif := r.FormValue("modif")
		if modif == "modif" {
			AfficherTemplate(w, "/modif_profil", profil)
			return
		}

		// Si ce n'est pas la page de modification
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		age := r.FormValue("age")
		email := r.FormValue("email")
		description := r.FormValue("description")

		// Gérer le téléchargement de l'image
		photoProfil, err := handleImageUploadProfil(r, username)
		if err != nil {
			http.Error(w, "Erreur lors de l'upload de l'image", http.StatusInternalServerError)
			return
		}

		err = update.UpdateProfil(db, username, firstname, lastname, age, email, photoProfil, description)
		if err != nil {
			http.Error(w, "Erreur lors de la mise à jour du profil", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/profil", http.StatusSeeOther)
		return
	}

	AfficherTemplate(w, "/profil", profil)
}

func handleImageUploadProfil(r *http.Request, username string) (string, error) {
	CHEMIN_IMG := "static/images/profils/" + username

	// Créer le répertoire pour les uploads si nécessaire
	if _, err := os.Stat(CHEMIN_IMG); os.IsNotExist(err) {
		err := os.MkdirAll(CHEMIN_IMG, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("Erreur lors de la création du répertoire: %v", err)
		}
	}

	// Vérifier si une image est téléchargée
	file, _, err := r.FormFile("photoProfil")
	if err == nil {
		// Si une image est téléchargée, la sauvegarder
		defer file.Close()
		imageName := username + "_profil.png"
		imagePath := CHEMIN_IMG + "/" + imageName
		outFile, err := os.Create(imagePath)
		if err != nil {
			return "", fmt.Errorf("Erreur lors de la création du fichier image: %v", err)
		}
		defer outFile.Close()

		// Copier le contenu du fichier téléchargé dans le fichier local
		_, err = outFile.ReadFrom(file)
		if err != nil {
			return "", fmt.Errorf("Erreur lors de la copie du fichier: %v", err)
		}
		return imagePath, nil
	} else if err != http.ErrMissingFile {
		// Si l'erreur est autre que "fichier manquant", renvoyer une erreur
		return "", fmt.Errorf("Erreur lors de la récupération du fichier: %v", err)
	}

	// Retourner une chaîne vide si aucune image n'est téléchargée
	return "", nil
}
