package main

import (
	"fauxrome/mysql/ConnectAndDisconnect"
	"fauxrome/mysql/CreateAndDelete"
	server "fauxrome/server/Handlers"

	"fmt"
	"log"
	"net/http"
)

const port = ":1678"

func main() {
	userDB := "AdminSupreme"     // Utilisateur mysql à creer
	PassWD := "AdminSupreme123!" // Mot de passe mysql a definir pour l'utilisateur
	dbName := "Database_Forum"   // Base de donnée à créer au préalable

	db, err := ConnectAndDisconnect.ConnectToBDD_Mysql(userDB, PassWD, dbName)
	if err != nil {
		log.Fatalf("Erreur: %v", err)
	}
	CreateAndDelete.CreateDataBase(db, dbName)

	nameTableUser := "Utilisateur"
	nameTableProfil := "Profil"
	nameTableGame := "GameLeagueOfLegends"
	nameTableBans := "Bannissement"
	CreateAndDelete.CreateAllTables(db, nameTableUser, nameTableProfil, nameTableGame, nameTableBans)
	// ------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes du serveur
	http.HandleFunc("/", server.AccueilHandler)
	http.HandleFunc("/signin", server.SigninHandler)
	http.HandleFunc("/login", server.LoginHandler)
	http.HandleFunc("/forum", server.ForumHandler)
	http.HandleFunc("/profil", server.ProfilHandler)

	fmt.Println("Serveur démarré sur http://localhost:1678")
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
