package SearchIntoTables

import (
	"database/sql"
	structures "fauxrome/server/Structures"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SearchByUserIntoProfil(db *sql.DB, utilisateurRecherche string, u structures.Profil_Search, Profils []structures.Profil_Search) (structures.Profil_Search, []structures.Profil_Search) {
	nameTable := structures.Tbl.Profil
	query := "SELECT ID, Utilisateur, Prenom, Nom, Age, Email, PhotoProfil, Description FROM " + nameTable + " WHERE Utilisateur = ?"

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

		// Integre les valeurs dans la struct de Profil
		if err := rows.Scan(&u.ID, &u.Utilisateur, &u.Prenom, &u.Nom, &u.Age, &u.Email, &u.PhotoProfil, &u.Description); err != nil {
			log.Fatal(err)
		}
		Profils = append(Profils, u)
	}
	return u, Profils
}

func AllIntoProfil(db *sql.DB) {
	nameTable := structures.Tbl.Profil
	var u structures.Profil
	query := "SELECT ID, Utilisateur, Prenom, Nom, Age, Email, PhotoProfil, Description FROM " + nameTable

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&u.ID, &u.Utilisateur); err != nil {
			log.Fatal(err)
		}
		// Remplissage de la struct
		structures.Simple_Profil.ID = u.ID
		structures.Simple_Profil.Utilisateur = u.Utilisateur
		// Remplissage de la slice
		structures.Slice_Profils = append(structures.Slice_Profils, structures.Simple_Profil)
	}

}

func DisplaySearchProfil(u structures.Profil_Search, Profils []structures.Profil_Search) {
	// Vérifier si des données ont été récupérées et les afficher
	if len(Profils) > 0 {
		fmt.Println("Conversations trouvées :")
		for _, u := range Profils {
			// Affichage de chaque conversation avec tous les champs
			fmt.Printf("ID: %d\n Utilisateur: %s\n Prenom: %s\n Nom: %s\n Age: %d\n Email: %s\n PhotoProfil: %s\n Description: %s\n\n", u.ID, u.Utilisateur, u.Prenom, u.Nom, u.Age, u.Email, u.PhotoProfil, u.Description)
		}
	} else {
		fmt.Println("Aucune conversation trouvée dans la table.")
	}
}
