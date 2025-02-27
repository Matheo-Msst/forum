package insert

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertMessageToBans(db *sql.DB, userName string, cause string, Date_Ban string, PhotoProfil string, nomTable string) {
	query := fmt.Sprintf("INSERT INTO `%s` (`Utilisateur`,`Cause`, `Date_Bannissement`, `PhotoProfil`) VALUES (?,?,?,?)", nomTable)

	// Préparation de la requête
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Erreur lors de la préparation de la requête: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userName, cause, Date_Ban, PhotoProfil)
	if err != nil {
		log.Fatalf("Erreur lors de l'insertion du Banni: %v", err)
	} else {
		fmt.Println("le banni à été insérer avec succès dans les bannissements")
	}
}

func InsertMessageToGameForum(db *sql.DB, userName string, message string, image string, Date string, nomTable string) {
	query := fmt.Sprintf("INSERT INTO `%s` (`Utilisateur`,`Message`, `Image`, `Date`) VALUES (?,?,?,?)", nomTable)

	// Préparation de la requête
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Erreur lors de la préparation de la requête: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userName, message, image, Date)
	if err != nil {
		log.Fatalf("Erreur lors de l'insertion de la conv: %v", err)
	} else {
		fmt.Println("la conv à été insérer avec succès dans Game")
	}
}

func InsertProfilToProfil(db *sql.DB, userName string, firstname string, lastname string, age string, Email string, image string, descriptions string, nomTable string) {
	query := fmt.Sprintf("INSERT INTO `%s` (`Utilisateur`,`Prenom`, `Nom`, `Age`, `Email`,`PhotoProfil` ,`Description` ) VALUES (?,?,?,?,?,?,?)", nomTable)

	// Préparation de la requête
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Erreur lors de la préparation de la requête: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userName, firstname, lastname, age, Email, image, descriptions)
	if err != nil {
		log.Fatalf("Erreur lors de l'insertion du profil: %v", err)
	} else {
		fmt.Println("le profil à été insérer avec succès.")
	}
}

func InsertUserToUser(db *sql.DB, userName string, passUser string, nomTable string) {
	query := fmt.Sprintf("INSERT INTO `%s` (`Utilisateur`, `MotDePasse`, `Role`) VALUES (?,?,?)", nomTable)

	// Préparation de la requête
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Erreur lors de la préparation de la requête: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userName, passUser, "USER")
	if err != nil {
		log.Fatalf("Erreur lors de l'insertion de l'utilisateur: %v", err)
	} else {
		fmt.Println("Utilisateur inséré avec succès.")
	}
}
