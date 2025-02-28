package structures

var Role_ConnectedUser string
var User_Connected string

var Simple_Profil_Search Profil_Search
var Simple_Utilisateurs_Search Utilisateur_Search
var Simple_Convs_Search Conversation_Game_Search
var Simple_Bans_Search Bans_Search
var Simple_Game_Search Games_Search

var Slice_Profil_Search []Profil_Search
var Slice_Utilisateurs_Search []Utilisateur_Search
var Slice_Convs_Search []Conversation_Game_Search
var Slice_Bans_Search []Bans_Search
var Slice_Games_Search []Games_Search

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
type Bans_Search struct {
	ID                int
	Utilisateur       string
	MotDePasse        string
	Cause             string
	Date_Bannissement string
	PhotoProfil       string
}
type Games_Search struct {
	ID          int
	NomJeu      string
	ImageJeu    string
	Description string
	Types       string
}

var NamesTables []string

// ---------------------------------------------------

var Simple_Profil Profil
var Simple_Utilisateur Utilisateur
var Simple_Conv Conversation_Game
var Simple_Bans Bans
var Simple_Game Games

var Slice_Profils []Profil
var Slice_Utilisateurs []Utilisateur
var Slice_Convs []Conversation_Game
var Slice_Bans []Bans
var Slice_Games []Games

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
type Bans struct {
	ID                int
	Utilisateur       string
	MotDePasse        string
	Cause             string
	Date_Bannissement string
	PhotoProfil       string
}
type Games struct {
	ID          int
	NomJeu      string
	ImageJeu    string
	Description string
	Types       string
}
