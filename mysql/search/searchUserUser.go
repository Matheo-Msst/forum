package SearchIntoTables

import (
	"database/sql"
	structures "fauxrome/server/Structures"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SearchByUserIntoUser(db *sql.DB, utilisateurRecherche string, nameTable string, u structures.Utilisateur_Search, users []structures.Utilisateur_Search) (structures.Utilisateur_Search, []structures.Utilisateur_Search) {

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
		users = append(users, u)
	}
	return u, users
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
func IfNOtPassword(user string, password string, u structures.Utilisateur_Search, users []structures.Utilisateur_Search) bool {
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
	if u.MotDePasse != password && u.Utilisateur != user {
		return true
	}
	return false
}
