package CreateAndDelete

import (
	"database/sql"
	SearchIntoTables "fauxrome/mysql/search"
	structures "fauxrome/server/Structures"
	"fmt"
	"log"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDataBase(db *sql.DB, nameDB string) {
	query := "CREATE DATABASE IF NOT EXISTS " + nameDB + ";"
	_, err := db.Exec(query)
	if err != nil {
		handleError(err)
	}
}

func CreateTable(db *sql.DB, query string, nameTable string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Erreur lors de la création de la table: %v", err)
	} else {
		fmt.Println("Table " + nameTable + " crée ou déjà existante.")
	}
}

func NameTableForum(nameTableGame string) string {
	nameTableGame = "Game_" + nameTableGame
	return nameTableGame
}

// ---------------------------------------------------------

func CreateAllTables(db *sql.DB) {
	fmt.Println("----------------------------------------------------")
	CreateTableUser(db, structures.Tbl.User)
	CreateTableProfil(db, structures.Tbl.Profil)
	CreateTableGame(db, structures.Tbl.Game)
	CreateTableBans(db, structures.Tbl.Bans)
	SearchIntoTables.AllIntoGames(db)
	// Boucle pour parcourir Slice_Games
	for i := 0; i < len(structures.Slice_Games); i++ {
		game := structures.Slice_Games[i]
		fmt.Println("Nom du jeu:", game.NomJeu)
		gameTable := NameTableForum(game.NomJeu)
		CreateTableForum(db, gameTable)
		structures.NamesTables = append(structures.NamesTables, gameTable)
	}
	fmt.Println("----------------------------------------------------")
	fmt.Println(structures.NamesTables)
	fmt.Println("----------------------------------------------------")
}
func CreateTableBans(db *sql.DB, nameTable string) {
	query := "CREATE TABLE IF NOT EXISTS `" + nameTable + "` (" +
		"`ID` INT NOT NULL AUTO_INCREMENT, " +
		"`Utilisateur` VARCHAR(100) NOT NULL, " +
		"`MotDePasse` VARCHAR(100) NOT NULL, " +
		"`Cause` TEXT NOT NULL, " +
		"`Date_Bannissement` VARCHAR(100) DEFAULT NULL, " +
		"`PhotoProfil` VARCHAR(100) NOT NULL, " +
		"PRIMARY KEY (`ID`));"
	CreateTable(db, query, nameTable)
}

func CreateTableForum(db *sql.DB, nameTable string) {
	query := "CREATE TABLE IF NOT EXISTS `" + nameTable + "` (" +
		"`ID` INT NOT NULL AUTO_INCREMENT, " +
		"`Utilisateur` VARCHAR(100) NOT NULL, " +
		"`Message` TEXT NOT NULL, " +
		"`Image` VARCHAR(100) DEFAULT NULL, " +
		"`Date` VARCHAR(100) NOT NULL, " +
		"PRIMARY KEY (`ID`));"
	CreateTable(db, query, nameTable)
}

func CreateTableProfil(db *sql.DB, nameTable string) {
	query := "CREATE TABLE IF NOT EXISTS `" + nameTable + "` (" +
		"`ID` INT NOT NULL AUTO_INCREMENT, " +
		"`Utilisateur` VARCHAR(100) NOT NULL, " +
		"`Prenom` VARCHAR(100) NOT NULL, " +
		"`Nom` VARCHAR(100) NOT NULL, " +
		"`Age` VARCHAR(3) NOT NULL, " +
		"`Email` VARCHAR(100) NOT NULL, " +
		"`PhotoProfil` VARCHAR(255) NOT NULL, " +
		"`Description` TEXT NOT NULL, " +
		"PRIMARY KEY (`ID`));"
	CreateTable(db, query, nameTable)
}

func CreateTableUser(db *sql.DB, nameTable string) {
	query := "CREATE TABLE IF NOT EXISTS `" + nameTable + "` (" +
		"`ID` INT NOT NULL AUTO_INCREMENT, " +
		"`Utilisateur` VARCHAR(100) NOT NULL, " +
		"`MotDePasse` VARCHAR(100) NOT NULL, " +
		"`Role` VARCHAR(5) NOT NULL DEFAULT 'USER', " +
		"PRIMARY KEY (`ID`));"
	CreateTable(db, query, nameTable)
}

func CreateTableGame(db *sql.DB, nameTable string) {
	query := "CREATE TABLE IF NOT EXISTS `" + nameTable + "` (" +
		"`ID` INT NOT NULL AUTO_INCREMENT, " +
		"`NomJeu` VARCHAR(100) NOT NULL, " +
		"`ImageJeu` VARCHAR(100) NOT NULL, " +
		"`Description` TEXT NOT NULL, " +
		"`Types` VARCHAR(200) NOT NULL, " +
		"PRIMARY KEY (`ID`));"
	CreateTable(db, query, nameTable)
}
