package server

import (
	setupdefault "fauxrome/mysql/setup_default"
	structures "fauxrome/server/Structures"
	Roles "fauxrome/server/roles"
	"fmt"
	"net/http"
	"text/template"
)

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	u := setupdefault.SetupDefaultUser(structures.Simple_Utilisateurs_Search)
	structures.Role_ConnectedUser = u.Role
	structures.Role_ConnectedUser = Roles.IfRole(structures.Role_ConnectedUser)
	// Charger le template
	templateAcc := "./static/templates/home.html"
	t, err := template.ParseFiles(templateAcc)
	if err != nil {
		fmt.Printf("Erreur de chargement du template: %v\n", err)
		http.Error(w, fmt.Sprintf("Erreur lors du chargement du template: %s", err), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Printf("Erreur d'exécution du template: %v\n", err)
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %s", err), http.StatusInternalServerError)
	}
}
