package main

import (
	"fmt"

	_ "github.com/Orwell-coder/ginf/bootstrap"
	"github.com/Orwell-coder/ginf/router"
)

func main() {
	fmt.Println("test")
	router.Router().Run(":8080")
}
