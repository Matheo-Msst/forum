package setupdefault

import structures "fauxrome/server/Structures"

func SetupDefaultProfil(p structures.Profil_Search) structures.Profil_Search {
	p.Utilisateur = "guest"
	p.Prenom = ""
	p.Nom = ""
	p.Age = ""
	p.Email = ""
	p.Description = ""
	p.PhotoProfil = "/static/images/icons/profil.png"
	return p
}

func SetupDefaultUser(u structures.Utilisateur_Search) structures.Utilisateur_Search {
	u.Utilisateur = "guest"
	u.MotDePasse = "guest"
	u.Role = "GUEST"
	return u
}
