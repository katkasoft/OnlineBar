package main

import (
	"OnlineBar/internal/app"
	_ "OnlineBar/internal/database/postgresql"
	"log"
)

func main() {
	log.Println("Server starting...")

	app.StartServer()
}
