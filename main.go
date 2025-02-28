package main

import (
	"fauxrome/mysql/ConnectAndDisconnect"
	"fauxrome/mysql/CreateAndDelete"
	server "fauxrome/server/Handlers"
	structures "fauxrome/server/Structures"

	"fmt"
	"log"
	"net/http"
)

const port = ":1678"

func main() {
	structures.DB.UserName = "AdminSupreme"       // Utilisateur mysql à creer
	structures.DB.PassWD = "AdminSupreme123!"     // Mot de passe mysql a definir pour l'utilisateur
	structures.DB.DatabaseName = "Database_Forum" // Base de donnée à créer au préalable

	db, err := ConnectAndDisconnect.ConnectToBDD_Mysql()
	if err != nil {
		log.Fatalf("Erreur: %v", err)
	}
	CreateAndDelete.CreateDataBase(db, structures.DB.DatabaseName)

	structures.Tbl.User = "Utilisateur"
	structures.Tbl.Profil = "Profil"
	structures.Tbl.Game = "Games"
	structures.Tbl.Forum = "GameLeagueOfLegends"
	structures.Tbl.Bans = "Bannissement"

	CreateAndDelete.CreateAllTables(db)
	// ------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes du serveur
	http.HandleFunc("/", server.AccueilHandler)
	http.HandleFunc("/games", server.GamesHandler)
	http.HandleFunc("/signin", server.SigninHandler)
	http.HandleFunc("/login", server.LoginHandler)
	http.HandleFunc("/forum", server.ForumHandler)
	http.HandleFunc("/profil", server.ProfilHandler)

	fmt.Println("Serveur démarré sur http://localhost:1678")
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
