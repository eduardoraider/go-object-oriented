package main

import (
	"fmt"
	"github.com/eduardoraider/goclass-object-oriented/users"
)

func main() {
	app := users.NewApp("goApp")
	fmt.Println(app)

	u := users.NewUser("John Doe", "john.doe", "123456", 30)
	fmt.Println(u)

	fmt.Println(u.Name)
}
