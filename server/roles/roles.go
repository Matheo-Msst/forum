package Roles

import (
	"database/sql"
	SearchIntoTables "fauxrome/mysql/search"
	structures "fauxrome/server/Structures"
)

func IfRole(role string) string {
	path := ""
	if role == "ADMIN" {
		path = "/admins/"
		return path
	}
	if role == "USER" {
		path = "/users/"
		return path
	} else {
		role = "GUEST"
		path = "/guests/"
		return path
	}
}

func IfBanned(db *sql.DB, username string) bool {
	SearchIntoTables.SearchByUserIntoBans(db, username)
	if structures.Simple_Bans.Utilisateur == username {
		return true
	}
	return false
}
