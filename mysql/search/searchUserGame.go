package SearchIntoTables

import (
	"database/sql"
	structures "fauxrome/server/Structures"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SearchByUserIntoGame(db *sql.DB, utilisateurRecherche string, nameTable string) {

	query := "SELECT ID, Utilisateur, Message, Image , Date FROM " + nameTable + " WHERE Utilisateur = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(utilisateurRecherche)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Creer un tableau de Games
	var conversations_game []structures.Conversation_Game_Search
	for rows.Next() {
		var u structures.Conversation_Game_Search
		if err := rows.Scan(&u.ID, &u.Utilisateur, &u.Message, &u.Image, &u.Date); err != nil {
			log.Fatal(err)
		}
		conversations_game = append(conversations_game, u)
	}

	// Afficher les résultats récupérés
	if len(conversations_game) > 0 {
		fmt.Println("Utilisateur trouvé :")
		for _, u := range conversations_game {
			// Affichage de l'utilisateur avec tous les champs
			fmt.Printf("ID: %d, Utilisateur: %s, Message: %s, Image: %s , Date: %s \n",
				u.ID, u.Utilisateur, u.Message, u.Image, u.Date)
		}
	} else {
		fmt.Println("Aucun utilisateur trouvé avec ce nom.")
	}

}

func AllIntoGame(db *sql.DB, nameTable string, u structures.Conversation_Game, conversations_game []structures.Conversation_Game) (structures.Conversation_Game, []structures.Conversation_Game) {
	query := "SELECT ID, Utilisateur, Message, Image, Date FROM " + nameTable + " ORDER BY ID DESC"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Vider la slice avant de la remplir (pour éviter la duplication des anciens messages)
	conversations_game = nil // Vider la slice des messages précédents

	for rows.Next() {
		// Scanner les données dans la struct
		if err := rows.Scan(&u.ID, &u.Utilisateur, &u.Message, &u.Image, &u.Date); err != nil {
			log.Fatal(err)
		}
		conversations_game = append(conversations_game, u)
	}

	return u, conversations_game
}

func DisplayIntoAllGame(u structures.Conversation_Game, conversations_game []structures.Conversation_Game) {
	if len(conversations_game) > 0 {
		fmt.Println("Conversations trouvées :")
		for _, u := range conversations_game {
			fmt.Printf("ID: %d\n Utilisateur: %s\n Message: %s\n Image: %s\n Date: %s \n\n",
				u.ID, u.Utilisateur, u.Message, u.Image, u.Date)
		}
	} else {
		fmt.Println("Aucune conversation trouvée dans la table.")
	}
}
