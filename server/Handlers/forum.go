package server

import (
	"database/sql"
	"mime/multipart"

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
	fmt.Println("Name Table forum : ", structures.Tbl.Forum)
	db, _ := ConnectAndDisconnect.ConnectToBDD_Mysql()
	username := structures.User_Connected
	role := structures.Role_ConnectedUser
	structures.Simple_Conv.Utilisateur_Connected = username
	fmt.Println("L'utilsateur (forum) est : ", username)
	fmt.Println("Le role (forum) est : ", role)

	if r.Method == http.MethodPost {

		message := r.FormValue("message")

		// Valider les champs obligatoires
		if message != "" {
			// Traiter l'image téléchargée (si présente)
			imagePath, err := handleImageUpload(r, username, structures.Tbl.Forum)
			if err != nil {
				http.Error(w, "Erreur lors du téléchargement de l'image", http.StatusInternalServerError)
				return
			}

			// Obtenir la date et l'heure actuelles
			dateTime := getCurrentDateTime()
			insert.InsertMessageToGameForum(db, username, message, imagePath, dateTime, structures.Tbl.Forum)
		}
		displayForumMessages(db)
		AfficherTemplate(w, templatePath, structures.Slice_Convs)

	} else {
		displayForumMessages(db)
		AfficherTemplate(w, templatePath, structures.Slice_Convs)
	}
}

func handleImageUpload(r *http.Request, username string, nameTable string) (string, error) {
	// Définir le chemin du répertoire pour l'image
	CHEMIN_IMG := "static/images/forum/" + nameTable
	fmt.Println("Répertoire d'image prévu:", CHEMIN_IMG)

	// Récupérer la date et l'heure pour le nom de fichier
	date := getCurrentDateTimeImage()
	fmt.Println("Date et heure pour l'image:", date)

	// Vérifier si le répertoire existe, sinon, le créer
	if _, err := os.Stat(CHEMIN_IMG); os.IsNotExist(err) {
		fmt.Println("Répertoire n'existe pas, création en cours...")
		err := os.MkdirAll(CHEMIN_IMG, os.ModePerm)
		if err != nil {
			fmt.Println("Erreur lors de la création du répertoire:", err)
			return "", fmt.Errorf("Erreur lors de la création du répertoire: %v", err)
		}
		fmt.Println("Répertoire créé avec succès.")
	}

	// Vérifier si une image est téléchargée
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Erreur lors de la récupération du fichier:", err)
		return "", fmt.Errorf("Erreur lors de la récupération du fichier: %v", err)
	}
	defer file.Close()

	// Afficher le type MIME du fichier téléchargé
	fmt.Println("Type MIME du fichier téléchargé:", fileHeader.Header.Get("Content-Type"))

	// Vérifier si le fichier est une image valide (vous pouvez ajuster la vérification ici si nécessaire)
	if !isValidImage(fileHeader) {
		fmt.Println("Le fichier téléchargé n'est pas une image valide")
		return "", fmt.Errorf("Le fichier téléchargé n'est pas une image valide")
	}

	// Construire le chemin du fichier image
	imageName := username + "_" + date + ".png"
	imagePath := CHEMIN_IMG + "/" + imageName
	fmt.Println("Nom du fichier image:", imageName)
	fmt.Println("Chemin complet de l'image:", imagePath)

	// Créer le fichier image sur le disque
	outFile, err := os.Create(imagePath)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier image:", err)
		return "", fmt.Errorf("Erreur lors de la création du fichier image: %v", err)
	}
	defer outFile.Close()

	// Copier le contenu du fichier téléchargé dans le fichier local
	_, err = outFile.ReadFrom(file)
	if err != nil {
		fmt.Println("Erreur lors de la copie du fichier:", err)
		return "", fmt.Errorf("Erreur lors de la copie du fichier: %v", err)
	}

	fmt.Println("Image téléchargée avec succès à l'emplacement:", imagePath)
	return imagePath, nil
}

func isValidImage(fileHeader *multipart.FileHeader) bool {
	mimeType := fileHeader.Header.Get("Content-Type")
	valid := mimeType == "image/png" || mimeType == "image/jpeg" || mimeType == "image/jpg"
	fmt.Println("Validation de l'image:", valid)
	return valid
}

func getCurrentDateTime() string {
	currentTime := time.Now()
	date := currentTime.Format("02/01/2006")
	timeStr := currentTime.Format("15:04")
	return date + " à " + timeStr
}
func getCurrentDateTimeImage() string {
	currentTime := time.Now()
	date := currentTime.Format("02012006")
	timeStr := currentTime.Format("1504")
	return date + "_" + timeStr
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
