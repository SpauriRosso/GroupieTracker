package main

import (
	Serveur "groupie-tracker/src/server"
	"log"
)

func main() {
	server := Serveur.NewServer()
	log.Fatal(server.Start(":5826"))
}
