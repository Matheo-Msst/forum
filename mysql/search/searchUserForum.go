package SearchIntoTables

import (
	"database/sql"
	structures "fauxrome/server/Structures"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SearchByUserIntoForum(db *sql.DB, utilisateurRecherche string) {
	nameTable := structures.Tbl.Forum
	var u structures.Conversation_Game_Search

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

	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Utilisateur, &u.Message, &u.Image, &u.Date); err != nil {
			log.Fatal(err)
		}
		// Remplissage de la struct
		structures.Simple_Convs_Search.ID = u.ID
		structures.Simple_Convs_Search.Utilisateur = u.Utilisateur
		structures.Simple_Convs_Search.Message = u.Message
		structures.Simple_Convs_Search.Date = u.Date
		// Remplissage de la slice
		structures.Slice_Convs_Search = append(structures.Slice_Convs_Search, structures.Simple_Convs_Search)
	}

}
func DisplayIntoForum(u structures.Conversation_Game, conversations_game []structures.Conversation_Game) {
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

func AllIntoForum(db *sql.DB) {
	nameTable := structures.Tbl.Forum
	var u structures.Conversation_Game
	query := "SELECT ID, Utilisateur, Message, Image, Date FROM " + nameTable + " ORDER BY ID DESC"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Vider la slice avant de la remplir (pour éviter la duplication des anciens messages)
	structures.Slice_Convs = nil // Vider la slice des messages précédents

	for rows.Next() {
		// Scanner les données dans la struct
		if err := rows.Scan(&u.ID, &u.Utilisateur, &u.Message, &u.Image, &u.Date); err != nil {
			log.Fatal(err)
		}
		// Remplissage de la struct
		structures.Simple_Conv.ID = u.ID
		structures.Simple_Conv.Utilisateur = u.Utilisateur
		structures.Simple_Conv.Message = u.Message
		structures.Simple_Conv.Image = u.Image
		structures.Simple_Conv.Date = u.Date
		// Remplissage de la slice
		structures.Slice_Convs = append(structures.Slice_Convs, structures.Simple_Conv)
	}
}
