package structures

var Role_ConnectedUser string
var User_Connected string

var Simple_Profil_Search Profil_Search
var Simple_Utilisateurs_Search Utilisateur_Search
var Simple_Convs_Search Conversation_Game_Search

var Slice_Profil_Search []Profil_Search
var Slice_Utilisateurs_Search []Utilisateur_Search
var Slice_Convs_Search []Conversation_Game_Search

type Profil_Search struct {
	ID          int
	Utilisateur string
	Prenom      string
	Nom         string
	Age         string
	Email       string
	PhotoProfil string
	Description string
}
type Utilisateur_Search struct {
	ID          int
	Utilisateur string
	MotDePasse  string
	Role        string
}
type Conversation_Game_Search struct {
	ID          int
	Utilisateur string
	Message     string
	Image       string
	Date        string
}

var Profil_var Profil
var Utilisateur_var Utilisateur
var Conversation_game_var Conversation_Game

var Slice_Profils []Profil
var Slice_Utilisateurs []Utilisateur
var Slice_Convs []Conversation_Game

type Profil struct {
	ID          int
	Utilisateur string
	Prenom      string
	Nom         string
	Age         string
	Email       string
	PhotoProfil string
	Description string
}
type Utilisateur struct {
	ID          int
	Utilisateur string
	MotDePasse  string
	Role        string
}
type Conversation_Game struct {
	ID                    int
	Utilisateur_Connected string
	Utilisateur           string
	Message               string
	Image                 string
	Date                  string
}
