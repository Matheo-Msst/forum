package server

import (
	"fauxrome/mysql/ConnectAndDisconnect"
	SearchIntoTables "fauxrome/mysql/search"
	setupdefault "fauxrome/mysql/setup_default"
	structures "fauxrome/server/Structures"
	Roles "fauxrome/server/roles"

	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	u := setupdefault.SetupDefaultUser(structures.Simple_Utilisateurs_Search)
	if r.Method == http.MethodPost {
		// Récupérer les valeurs du formulaire
		username := r.FormValue("username")
		password := r.FormValue("password")
		structures.User_Connected = username
		// Si l'utilisateur se connecte en tant qu'invité, le nom d'utilisateur et le mot de passe seront "guest"
		if username == "guest" && password == "guest" {
			structures.Role_ConnectedUser = u.Role
			role := structures.Role_ConnectedUser
			path := Roles.IfRole(role)
			structures.Role_ConnectedUser = path
			fmt.Println("Connexion en tant qu'invité, chemin :", role)

			// Rediriger vers le forum en tant qu'invité
			http.Redirect(w, r, "/forum", http.StatusSeeOther)
			return
		}

		// Connexion à la base de données
		db, _ := ConnectAndDisconnect.ConnectToBDD_Mysql()

		SearchIntoTables.SearchByUserIntoUser(db, structures.User_Connected)
		if Roles.IfBanned(db, username) == true {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			var ban structures.Bans
			ban = structures.Simple_Bans
			fmt.Println("L'Utilisateur ", ban.Utilisateur, " à été banni pour : ", ban.Cause, " le ", ban.Date_Bannissement)
			return
		}
		fmt.Println("Nom d'utilisateur:", structures.User_Connected)
		fmt.Println("Mot de passe:", password)
		SearchIntoTables.DisplaySearchUser(structures.Simple_Utilisateurs_Search, structures.Slice_Utilisateurs_Search)

		if IfNOtPassword(structures.User_Connected, password, structures.Simple_Utilisateurs_Search, structures.Slice_Utilisateurs_Search) {
			AfficherTemplate(w, "login", nil)
		} else {
			structures.User_Connected = structures.Simple_Utilisateurs_Search.Utilisateur
			role := structures.Simple_Utilisateurs_Search.Role
			structures.Role_ConnectedUser = Roles.IfRole(role)

			http.Redirect(w, r, "/forum", http.StatusSeeOther)
		}
	} else {
		AfficherTemplate(w, "login", nil)
	}
}

func IfNOtPassword(user string, password string, u structures.Utilisateur_Search, users []structures.Utilisateur_Search) bool {
	// Afficher les résultats récupérés
	if len(users) > 0 {
		fmt.Println("Utilisateur trouvé :")
		for _, u := range users {
			// Affichage de l'utilisateur avec tous les champs
			fmt.Printf("ID: %d, Utilisateur: %s, MotDePasse: %s, Role: %s\n",
				u.ID, u.Utilisateur, u.MotDePasse, u.Role)
		}
	} else {
		fmt.Println("Aucun utilisateur trouvé avec ce nom.")
	}
	if u.MotDePasse != password && u.Utilisateur != user {
		return true
	}
	return false
}

func AfficherTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Obtenir le chemin absolu du répertoire de travail actuel
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Erreur lors de la récupération du répertoire de travail: %v\n", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Construire le chemin absolu vers le fichier du template
	tmplPath := filepath.Join(cwd, "static", "templates", fmt.Sprintf("%s.html", tmpl))
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Printf("Erreur de chargement du template: %v\n", err)
		http.Error(w, fmt.Sprintf("Erreur lors du chargement du template: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Println(tmplPath)
	err = t.Execute(w, data)
	if err != nil {
		fmt.Printf("Erreur d'exécution du template: %v\n", err)
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %s", err), http.StatusInternalServerError)
	}
}
