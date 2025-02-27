package server

import (
	"fmt"
	"net/http"
	"text/template"
)

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
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
