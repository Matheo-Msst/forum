package SearchIntoTables

import (
	"database/sql"
	structures "fauxrome/server/Structures"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SearchByUserIntoUser(db *sql.DB, utilisateurRecherche string) {
	nameTable := structures.Tbl.User
	var u structures.Utilisateur_Search
	query := "SELECT ID, Utilisateur, MotDePasse, Role FROM " + nameTable + " WHERE Utilisateur = ?"

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
		// Integre les valeurs dans la struct de Utilisateur
		if err := rows.Scan(&u.ID, &u.Utilisateur, &u.MotDePasse, &u.Role); err != nil {
			log.Fatal(err)
		}
		// Remplissage de la struct
		structures.Simple_Utilisateurs_Search.ID = u.ID
		structures.Simple_Utilisateurs_Search.Utilisateur = u.Utilisateur
		structures.Simple_Utilisateurs_Search.MotDePasse = u.MotDePasse
		structures.Simple_Utilisateurs_Search.Role = u.Role
		// Remplissage de la slice
		structures.Slice_Utilisateurs_Search = append(structures.Slice_Utilisateurs_Search, structures.Simple_Utilisateurs_Search)
	}
}

func DisplaySearchUser(u structures.Utilisateur_Search, users []structures.Utilisateur_Search) {
	// Afficher les résultats récupérés
	if len(users) > 0 {
		fmt.Println("Utilisateur trouvé :")
		for _, u := range users {
			// Affichage de l'utilisateur avec tous les champs
			fmt.Printf("ID: %d, Utilisateur: %s, MotDePasse: %s, Role: %s\n",
				u.ID, u.Utilisateur, u.MotDePasse, u.Role)
		}
	} else {
		fmt.Println("Aucun utilisateur trouvé avec ce nom.")
	}
}
