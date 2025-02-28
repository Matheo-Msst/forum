package server

import (
	"fauxrome/mysql/ConnectAndDisconnect"
	SearchIntoTables "fauxrome/mysql/search"
	structures "fauxrome/server/Structures"
	"net/http"
)

func GamesHundler(w http.ResponseWriter, r *http.Request) {
	db, _ := ConnectAndDisconnect.ConnectToBDD_Mysql()
	SearchIntoTables.AllIntoGames(db)
	structures.Slice_Games = reverseGames(structures.Slice_Games)
	SearchIntoTables.DisplayIntoGames(structures.Simple_Game, structures.Slice_Games)
	AfficherTemplate(w, "forum_Games", structures.Slice_Games)
}
func reverseGames(conversations_game []structures.Games) []structures.Games {
	for i, j := 0, len(conversations_game)-1; i < j; i, j = i+1, j-1 {
		conversations_game[i], conversations_game[j] = conversations_game[j], conversations_game[i]
	}
	return conversations_game
}
