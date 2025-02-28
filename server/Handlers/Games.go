package server

import (
	"fauxrome/mysql/ConnectAndDisconnect"
	"fauxrome/mysql/CreateAndDelete"
	SearchIntoTables "fauxrome/mysql/search"
	structures "fauxrome/server/Structures"
	"fmt"
	"net/http"
)

func GamesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler Games appelé") // Log pour vérifier si le handler est appelé
	db, _ := ConnectAndDisconnect.ConnectToBDD_Mysql()
	SearchIntoTables.AllIntoGames(db)
	structures.Slice_Games = reverseGames(structures.Slice_Games)
	if r.Method == http.MethodPost {
		// Log et traitement pour POST
		fmt.Println("Méthode POST reçue")
		for i := 0; i < len(structures.Slice_Games); i++ {
			game := structures.Slice_Games[i]
			NameGame := r.FormValue(game.NomJeu)

			if NameGame != "" {
				fmt.Println("Le jeu sélectionné est :", NameGame)
				NameGame = CreateAndDelete.NameTableForum(NameGame)
				fmt.Println("La table est :", NameGame)
				structures.Tbl.Forum = NameGame
				http.Redirect(w, r, "/forum", http.StatusSeeOther)
			}
		}

		// Affichage du template 'Games' après traitement
		fmt.Println("Affichage du template 'Games'")
		AfficherTemplate(w, "Games", structures.Slice_Games)
	}
	// Affichage du template 'Games' après traitement
	fmt.Println("Affichage du template 'Games'")
	AfficherTemplate(w, "Games", structures.Slice_Games)
}

func reverseGames(conversations_game []structures.Games) []structures.Games {
	for i, j := 0, len(conversations_game)-1; i < j; i, j = i+1, j-1 {
		conversations_game[i], conversations_game[j] = conversations_game[j], conversations_game[i]
	}
	return conversations_game
}
