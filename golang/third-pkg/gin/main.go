package main

import (
	"gin-example/routes"
	"log"
)

func main() {

	if err := routes.Router.Run(":9090"); err != nil {
		log.Fatalln(err)
	}
}
