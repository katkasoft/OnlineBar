package main

import (
	"OnlineBar/Backend/internal/app"
	_ "OnlineBar/Backend/internal/database/postgresql"
	"log"
)

func main() {
	log.Println("Server starting...")

	app.StartServer()
}
