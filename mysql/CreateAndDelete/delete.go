package CreateAndDelete

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteTable(db *sql.DB, nomTable string) error {

	query := fmt.Sprintf("DROP TABLE IF EXISTS %s;", nomTable)

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("Erreur lors de la suppression de la table %s : %v", nomTable, err)
	}
	log.Println("Table supprimée avec succès!")
	return nil
}
