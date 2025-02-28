package server

import (
	"database/sql"

	"fauxrome/mysql/ConnectAndDisconnect"
	"fauxrome/mysql/insert"
	SearchIntoTables "fauxrome/mysql/search"
	structures "fauxrome/server/Structures"
	"fmt"
	"net/http"
	"os"
	"time"
)

// ForumHandler gère la soumission du formulaire
func ForumHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le rôle de l'utilisateur connecté
	templatePath := structures.Role_ConnectedUser + "forum"
	fmt.Println("Le rôle templatepath forum:", structures.Role_ConnectedUser)

	db, _ := ConnectAndDisconnect.ConnectToBDD_Mysql()
	nameTableGame := "GameLeagueOfLegends"
	username := structures.User_Connected
	structures.Simple_Conv.Utilisateur_Connected = username
	fmt.Println("L'utilsateur (forum) est : ", username)
	fmt.Println("L'utilsateur (forum 2) est : ", structures.Simple_Conv.Utilisateur_Connected)

	if r.Method == http.MethodPost {

		message := r.FormValue("message")

		// Valider les champs obligatoires
		if message != "" {
			// Traiter l'image téléchargée (si présente)
			imagePath, err := handleImageUpload(r, username, nameTableGame)
			if err != nil {
				http.Error(w, "Erreur lors du téléchargement de l'image", http.StatusInternalServerError)
				return
			}

			// Obtenir la date et l'heure actuelles
			dateTime := getCurrentDateTime()
			insert.InsertMessageToGameForum(db, username, message, imagePath, dateTime, nameTableGame)
		}
		displayForumMessages(db)
		AfficherTemplate(w, templatePath, structures.Slice_Convs)

	} else {
		displayForumMessages(db)
		AfficherTemplate(w, templatePath, structures.Slice_Convs)
	}
}

// handleImageUpload traite le téléchargement de l'image
func handleImageUpload(r *http.Request, username string, nameTable string) (string, error) {
	CHEMIN_IMG := "static/images/forum/" + nameTable
	date := getCurrentDateTimeImage()
	// Créer le répertoire pour les forums si nécessaire
	if _, err := os.Stat(CHEMIN_IMG); os.IsNotExist(err) {
		err := os.MkdirAll(CHEMIN_IMG, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("Erreur lors de la création du répertoire: %v", err)
		}
	}
	// Vérifier si une image est téléchargée
	file, _, err := r.FormFile("image")
	if err == nil {
		// Si une image est téléchargée, la sauvegarder
		defer file.Close()
		imageName := username + "_" + date + ".png"
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

		return "", fmt.Errorf("Erreur lors de la récupération du fichier: %v", err)
	}

	return "", nil
}
func getCurrentDateTime() string {
	currentTime := time.Now()
	date := currentTime.Format("02/01/2006")
	timeStr := currentTime.Format("15:04")
	return date + " à " + timeStr
}
func getCurrentDateTimeImage() string {
	currentTime := time.Now()
	date := currentTime.Format("02/01/2006")
	timeStr := currentTime.Format("1504")
	return date + " à " + timeStr
}
func displayForumMessages(db *sql.DB) {
	SearchIntoTables.AllIntoForum(db)
	structures.Slice_Convs = reverseMessages(structures.Slice_Convs)

}
func reverseMessages(conversations_game []structures.Conversation_Game) []structures.Conversation_Game {
	for i, j := 0, len(conversations_game)-1; i < j; i, j = i+1, j-1 {
		conversations_game[i], conversations_game[j] = conversations_game[j], conversations_game[i]
	}
	return conversations_game
}
