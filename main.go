package main

import (
	"fmt"
	"github.com/eduardoraider/goclass-object-oriented/users"
)

func main() {
	app := users.NewApp("goApp")
	fmt.Println(app)
}
