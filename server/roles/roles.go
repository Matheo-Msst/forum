package Roles

func IfRole(role string) string {
	path := ""
	if role == "ADMIN" {
		path = "/admins/"
		return path
	}
	if role == "USER" {
		path = "/users/"
		return path
	} else {
		role = "GUEST"
		path = "/guests/"
		return path
	}
}
