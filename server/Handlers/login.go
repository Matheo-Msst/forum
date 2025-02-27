package server

import (
	"database/sql"
	"fauxrome/mysql/ConnectAndDisconnect"
	SearchIntoTables "fauxrome/mysql/search"
	structures "fauxrome/server/Structures"
	Roles "fauxrome/server/roles"

	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Récupérer les valeurs du formulaire
		username := r.FormValue("username")
		password := r.FormValue("password")
		structures.User_Connected = username
		// Si l'utilisateur se connecte en tant qu'invité, le nom d'utilisateur et le mot de passe seront "guest"
		if username == "guest" && password == "guest" {
			structures.Role_ConnectedUser = "GUEST"
			role := structures.Role_ConnectedUser
			path := Roles.IfRole(role)
			structures.Role_ConnectedUser = path
			fmt.Println("Connexion en tant qu'invité, chemin :", role)

			// Rediriger vers le forum en tant qu'invité
			http.Redirect(w, r, "/forum", http.StatusSeeOther)
			return
		}

		// Connexion à la base de données
		db := MysqlConf()

		// Table Utilisateur
		nameTableUser := "Utilisateur"

		// Recherche de l'utilisateur dans la base de données via l'utilisateur
		structures.Simple_Utilisateurs_Search, structures.Slice_Utilisateurs_Search = SearchIntoTables.SearchByUserIntoUser(db, username, nameTableUser, structures.Simple_Utilisateurs_Search, structures.Slice_Utilisateurs_Search)

		// Affichage des informations de l'utilisateur pour le débogage
		fmt.Println("Nom d'utilisateur:", username)
		fmt.Println("Mot de passe:", password)
		SearchIntoTables.DisplaySearchUser(structures.Simple_Utilisateurs_Search, structures.Slice_Utilisateurs_Search)

		// Vérification du mot de passe de l'utilisateur
		if SearchIntoTables.IfNOtPassword(username, password, structures.Simple_Utilisateurs_Search, structures.Slice_Utilisateurs_Search) {
			// Si le mot de passe est incorrect, afficher le formulaire de connexion
			AfficherTemplate(w, "login", nil)
		} else {
			structures.User_Connected = structures.Simple_Utilisateurs_Search.Utilisateur
			test := structures.User_Connected
			fmt.Println("L'utilsateur (login) est : ", test)
			role := structures.Simple_Utilisateurs_Search.Role
			// Déterminer le chemin à suivre pour le rôle
			path := Roles.IfRole(role)
			structures.Role_ConnectedUser = path

			// Rediriger l'utilisateur vers le forum en fonction de son rôle
			http.Redirect(w, r, "/forum", http.StatusSeeOther)
		}
	} else {
		// Si la méthode est GET, afficher le formulaire de connexion
		AfficherTemplate(w, "login", nil)
	}
}

func MysqlConf() *sql.DB {
	userDB := "AdminSupreme"     // Utilisateur mysql à creer
	PassWD := "AdminSupreme123!" // Mot de passe mysql a definir pour l'utilisateur
	dbName := "Database_Forum"   // Base de donnée à créer au préalable

	// Connexion à la base de données
	db, err := ConnectAndDisconnect.ConnectToBDD_Mysql(userDB, PassWD, dbName)
	if err != nil {
		log.Fatalf("Erreur: %v", err)
	}
	return db
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
