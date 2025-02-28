package update

import (
	"database/sql"
	structures "fauxrome/server/Structures"
	"fmt"
)

func UpdateProfil(db *sql.DB, username, firstname, lastname, age, email, photoProfil, description string) error {
	nameTable := structures.Tbl.Profil

	query := "UPDATE " + nameTable + " SET Prenom = ?, Nom = ?, Age = ?, Email = ?, PhotoProfil = ?, Description = ? WHERE Utilisateur = ?"

	_, err := db.Exec(query, firstname, lastname, age, email, photoProfil, description, username)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour du profil: %v", err)
	}

	return nil
}
