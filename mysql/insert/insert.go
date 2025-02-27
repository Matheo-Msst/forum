package insert

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertMessageToBans(db *sql.DB, userName, cause, Date_Ban, PhotoProfil, nameTable string) {
	query := fmt.Sprintf("INSERT INTO `%s` (`Utilisateur`,`Cause`, `Date_Bannissement`, `PhotoProfil`) VALUES (?,?,?,?)", nameTable)

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

func InsertMessageToGameForum(db *sql.DB, userName, message, image, Date, nameTable string) {
	query := fmt.Sprintf("INSERT INTO `%s` (`Utilisateur`,`Message`, `Image`, `Date`) VALUES (?,?,?,?)", nameTable)

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
		fmt.Println("la conv à été insérer avec succès dans le Forum")
	}
}

func InsertProfilToProfil(db *sql.DB, userName, firstname, lastname, age, Email, image, descriptions, nameTable string) {
	query := fmt.Sprintf("INSERT INTO `%s` (`Utilisateur`,`Prenom`, `Nom`, `Age`, `Email`,`PhotoProfil` ,`Description` ) VALUES (?,?,?,?,?,?,?)", nameTable)

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

func InsertUserToUser(db *sql.DB, userName, passUser, nameTable string) {
	query := fmt.Sprintf("INSERT INTO `%s` (`Utilisateur`, `MotDePasse`, `Role`) VALUES (?,?,?)", nameTable)

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

func InsertGameToGame(db *sql.DB, nameGame, imageGame, Description, Types, nameTable string) {
	query := fmt.Sprintf("INSERT INTO `%s` (`NomJeu`,`ImageJeu`, `Description`, `Types`) VALUES (?,?,?,?)", nameTable)

	// Préparation de la requête
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Erreur lors de la préparation de la requête: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(nameGame, imageGame, Description, Types)
	if err != nil {
		log.Fatalf("Erreur lors de l'insertion dans les jeux: %v", err)
	} else {
		fmt.Println("le jeu à été insérer avec succès dans les jeux")
	}
}
func InsertUserToBans(db *sql.DB, userName, Password, cause, DateBan, PhotoProfil, nameTable string) {
	query := fmt.Sprintf("INSERT INTO `%s` (`Utilisateur`,`MotDePasse`, `Cause`, `Date_Bannissement`, `PhotoProfil`) VALUES (?,?,?,?,?)", nameTable)

	// Préparation de la requête
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Erreur lors de la préparation de la requête: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userName, Password, cause, DateBan, PhotoProfil)
	if err != nil {
		log.Fatalf("Erreur lors de l'insertion dans les jeux: %v", err)
	} else {
		fmt.Println("le jeu à été insérer avec succès dans les jeux")
	}
}
