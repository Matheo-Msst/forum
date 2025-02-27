package remove

import (
	"database/sql"
	"fmt"
	"log"
)

// Commande mysql exemple : DELETE FROM utilisateurs WHERE id = 5;

func RemoveToUser(db *sql.DB, userName, nameTable string) {
	query := fmt.Sprintf("DELETE FROM `%s` WHERE Utilisateur = ?;", nameTable)

	// Préparation de la requête
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Erreur lors de la préparation de la requête: %v", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(userName)
	if err != nil {
		log.Fatalf("Erreur lors de l'effacement %v", err)
	} else {
		fmt.Println("Effacé avec succès")
	}
}
