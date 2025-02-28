package SearchIntoTables

import (
	"database/sql"
	structures "fauxrome/server/Structures"
	"log"
)

func SearchByUserIntoBans(db *sql.DB, utilisateurRecherche string) {
	nameTable := structures.Tbl.Bans
	var u structures.Bans
	query := "SELECT ID, Utilisateur, MotDePasse, Cause, Date_Bannissement, PhotoProfil  FROM " + nameTable + " WHERE Utilisateur = ?"

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
		if err := rows.Scan(&u.ID, &u.Utilisateur, &u.MotDePasse, &u.Cause, &u.Date_Bannissement, &u.PhotoProfil); err != nil {
			log.Fatal(err)
		}
		// Remplissage de la struct
		structures.Simple_Bans.ID = u.ID
		structures.Simple_Bans.Utilisateur = u.Utilisateur
		structures.Simple_Bans.MotDePasse = u.MotDePasse
		structures.Simple_Bans.Cause = u.Cause
		structures.Simple_Bans.Date_Bannissement = u.Date_Bannissement
		structures.Simple_Bans.PhotoProfil = u.PhotoProfil
		// Remplissage de la slice
		structures.Slice_Utilisateurs_Search = append(structures.Slice_Utilisateurs_Search, structures.Simple_Utilisateurs_Search)
	}
}
