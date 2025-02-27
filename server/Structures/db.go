package structures

var DB Database
var Tbl Tables

type Database struct {
	UserName     string
	PassWD       string
	DatabaseName string
}
type Tables struct {
	User   string
	Profil string
	Forum  string
	Game   string
	Bans   string
}
