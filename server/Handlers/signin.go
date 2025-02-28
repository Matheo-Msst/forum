package server

import (
	"fauxrome/mysql/ConnectAndDisconnect"
	"fauxrome/mysql/insert"
	SearchIntoTables "fauxrome/mysql/search"
	setupdefault "fauxrome/mysql/setup_default"
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

		SearchIntoTables.SearchByUserIntoUser(db, username)

		if structures.Simple_Utilisateurs_Search.Utilisateur == username /*deja inscrit*/ || Roles.IfBanned(db, username) == true /*banni*/ {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			insert.InsertUserToUser(db, username, password, nameTableUser)
			profil := setupdefault.SetupDefaultProfil(structures.Simple_Profil_Search)
			insert.InsertProfilToProfil(db, username, profil.Prenom, profil.Nom, profil.Age, profil.Email, profil.PhotoProfil, profil.Description, structures.Tbl.Profil)
			structures.Simple_Profil_Search = profil
			structures.User_Connected = username
			structures.Simple_Conv.Utilisateur = username
			role := "USER"
			structures.Role_ConnectedUser = Roles.IfRole(role)
			http.Redirect(w, r, "/forum", http.StatusSeeOther)
		}
	} else {
		AfficherTemplate(w, "/signin", nil)
	}
}
