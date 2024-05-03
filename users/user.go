package users

import "math/rand"

var letters = []rune("abcdefghijKlmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type user struct {
	Name     string
	Username string
	Secret   string
	Type     string
}

func NewApp(name string) user {
	u := user{Name: name, Type: "app"}
	u.Username = generateHash(5)
	u.Secret = generateHash(64)
	return u
}

func generateHash(n int) string {
	hash := make([]rune, n)
	for i := range hash {
		hash[i] = letters[rand.Intn(len(letters))]
	}
	return string(hash)
}
