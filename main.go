package main

import (
	"fmt"
	"github.com/eduardoraider/goclass-object-oriented/db"
	"github.com/eduardoraider/goclass-object-oriented/users"
)

func main() {
	app := users.NewApp("goApp")
	fmt.Println(app)

	u := users.NewUser("John Doe", "john.doe", "123456", 30)
	fmt.Println(u)

	fmt.Println(u.Name)

	err := db.Save(app)
	if err != nil {
		fmt.Println(err)
	}
	err2 := db.Save(u)
	if err2 != nil {
		fmt.Println(err2)
	}
}
