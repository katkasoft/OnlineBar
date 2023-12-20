package app

import (
	"OnlineBar/internal/transport/rest"
	"OnlineBar/pkg/cfg"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	config := cfg.ServerConfig()
	port := config.Server.Port
	host := config.Server.Host
	r := gin.Default()

	r.GET("/ping", rest.TestFunc)
	r.GET("/login", rest.LoginHandler)

	if err := r.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		log.Println("Failed to start server")
	}

	fmt.Println("Server started!")
}
