package update

import (
	"database/sql"
	structures "fauxrome/server/Structures"
	"fmt"
)

// Fonction pour modifier les données d'une table "Profil"
func UpdateProfil(db *sql.DB, username, firstname, lastname, age, email, photoProfil, description string) error {
	// La requête SQL de mise à jour
	nameTable := structures.Tbl.Profil
	query := "UPDATE " + nameTable + " SET Prenom = ?, Nom = ?, Age = ?, Email = ?, PhotoProfil = ?, Description = ? WHERE Utilisateur = ?"

	// Exécution de la requête avec les paramètres
	_, err := db.Exec(query, firstname, lastname, age, email, photoProfil, description, username)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour du profil: %v", err)
	}

	return nil
}
