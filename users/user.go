package users

type profile struct {
	user
	Age uint8
}

func NewUser(name, username, pass string, age uint8) profile {
	u := profile{Age: age}
	u.user = user{
		Name:     name,
		Username: username,
		Secret:   pass,
		Type:     "user",
	}
	return u
}
