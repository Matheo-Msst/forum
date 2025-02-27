package ConnectAndDisconnect

import (
	"database/sql"
	structures "fauxrome/server/Structures"
	"fmt"
	"log"
)

func ConnectToBDD_Mysql() (*sql.DB, error) {
	user_BDD := structures.DB.UserName
	password_BDD := structures.DB.PassWD
	dbName := structures.DB.DatabaseName

	dsn := user_BDD + ":" + password_BDD + "@tcp(127.0.0.1)/" + dbName

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("impossible de se connecter à la base de données: %v", err)
	}

	fmt.Println("Connexion réussie à la base de données!")
	return db, nil
}

func DisconnectFromDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("Erreur lors de la déconnexion de la base de données: %v", err)
	} else {
		fmt.Println("Déconnexion réussie de la base de données.")
	}
}
