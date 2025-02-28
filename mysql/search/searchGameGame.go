package SearchIntoTables

import (
	"database/sql"
	structures "fauxrome/server/Structures"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
CREATE TABLE `Games` (
    `ID` int NOT NULL AUTO_INCREMENT,
    `NomJeu` varchar(100) NOT NULL,
    `ImageJeu` varchar(100) NOT NULL,
    `Description` text NOT NULL,
    `Types` varchar(200) NOT NULL,
    PRIMARY KEY (`ID`)*/

func SearchByGamesIntoGames(db *sql.DB, jeurecherche string) {
	nameTable := structures.Tbl.Game
	var u structures.Games_Search

	query := "SELECT ID, NomJeu, ImageJeu, Description , Types " + nameTable + " WHERE NomJeu = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(jeurecherche)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.NomJeu, &u.ImageJeu, &u.Description, &u.Types); err != nil {
			log.Fatal(err)
		}
		// Remplissage de la struct
		structures.Simple_Game_Search.ID = u.ID
		structures.Simple_Game_Search.ImageJeu = u.ImageJeu
		structures.Simple_Game_Search.Description = u.Description
		structures.Simple_Game_Search.Types = u.Types

		// Remplissage de la slice
		structures.Slice_Games_Search = append(structures.Slice_Games_Search, structures.Simple_Game_Search)
	}

}
func DisplayIntoGames(u structures.Games, slice []structures.Games) {
	if len(slice) > 0 {
		fmt.Println("Jeux trouvés :")
		for _, u := range slice {
			fmt.Printf("ID: %d\n Utilisateur: %s\n Message: %s\n Image: %s\n Date: %s \n\n",
				u.ID, u.NomJeu, u.ImageJeu, u.Description, u.Types)
		}
	} else {
		fmt.Println("Aucune conversation trouvée dans la table.")
	}
}

func AllIntoGames(db *sql.DB) {
	nameTable := structures.Tbl.Game
	var u structures.Games
	query := "SELECT ID, NomJeu, ImageJeu, Description, Types FROM " + nameTable + " ORDER BY ID DESC;"
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	structures.Slice_Games = nil

	for rows.Next() {
		// Scanner les données dans la struct
		if err := rows.Scan(&u.ID, &u.NomJeu, &u.ImageJeu, &u.Description, &u.Types); err != nil {
			log.Fatal(err)
		}
		// Remplissage de la struct
		structures.Simple_Game.ID = u.ID
		structures.Simple_Game.NomJeu = u.NomJeu
		structures.Simple_Game.ImageJeu = u.ImageJeu
		structures.Simple_Game.Description = u.Description
		structures.Simple_Game.Types = u.Types
		// Remplissage de la slice
		structures.Slice_Games = append(structures.Slice_Games, structures.Simple_Game)
	}
}
