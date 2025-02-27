package server

import (
	"database/sql"

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

	db := MysqlConf()
	nameTableGame := "GameLeagueOfLegends"
	username := structures.User_Connected
	structures.Conversation_game_var.Utilisateur_Connected = username
	fmt.Println("L'utilsateur (forum) est : ", username)
	fmt.Println("L'utilsateur (forum 2) est : ", structures.Conversation_game_var.Utilisateur_Connected)

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
		db := MysqlConf()
		displayForumMessages(db)
		AfficherTemplate(w, templatePath, structures.Slice_Convs)
	}
}

// handleImageUpload traite le téléchargement de l'image
func handleImageUpload(r *http.Request, username string, nameTable string) (string, error) {
	CHEMIN_IMG := "static/images/forum/" + nameTable
	date := time.Now()
	currenttime := date.Format("02-01-2006")
	timeStr := date.Format("15:04")

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
		imageName := username + "_" + currenttime + "_" + timeStr + ".png"
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

// getCurrentDateTime retourne la date et l'heure actuelles formatées
func getCurrentDateTime() string {
	currentTime := time.Now()
	date := currentTime.Format("02/01/2006")
	timeStr := currentTime.Format("15:04")
	return date + " à " + timeStr
}

// displayForumMessages récupère et affiche tous les messages du forum
func displayForumMessages(db *sql.DB) {
	structures.Conversation_game_var, structures.Slice_Convs = SearchIntoTables.AllIntoGame(db, "GameLeagueOfLegends", structures.Conversation_game_var, structures.Slice_Convs)
	// Inverser l'ordre des messages avant de les afficher
	structures.Slice_Convs = reverseMessages(structures.Slice_Convs)
	// Afficher le template avec les messages du forum

}

// reverseMessages inverse l'ordre des messages dans la conversation
func reverseMessages(conversations_game []structures.Conversation_Game) []structures.Conversation_Game {
	for i, j := 0, len(conversations_game)-1; i < j; i, j = i+1, j-1 {
		conversations_game[i], conversations_game[j] = conversations_game[j], conversations_game[i]
	}
	return conversations_game
}
