package server

import (
	"fauxrome/mysql/ConnectAndDisconnect"
	"fauxrome/mysql/insert"
	SearchIntoTables "fauxrome/mysql/search"
	structures "fauxrome/server/Structures"
	Roles "fauxrome/server/roles"
	"net/http"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		db, _ := ConnectAndDisconnect.ConnectToBDD_Mysql()

		nameTableUser := "Utilisateur"

		user := structures.Simple_Utilisateurs_Search
		users := structures.Slice_Utilisateurs_Search
		user, structures.Slice_Utilisateurs_Search = SearchIntoTables.SearchByUserIntoUser(db, username, user, users)
		structures.Simple_Utilisateurs_Search = user
		if user.Utilisateur == username {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			insert.InsertUserToUser(db, username, password, nameTableUser)
			imagePardefautProfil := "/static/images/icons/profil.png"
			insert.InsertProfilToProfil(db, username, " ", " ", " ", " ", imagePardefautProfil, " ", "Profil")
			structures.User_Connected = username
			role := "USER"
			role = Roles.IfRole(role)
			structures.Role_ConnectedUser = role
			structures.Conversation_game_var.Utilisateur = username
			http.Redirect(w, r, "/forum", http.StatusSeeOther)
		}
	} else {
		AfficherTemplate(w, "/signin", nil)
	}
}
