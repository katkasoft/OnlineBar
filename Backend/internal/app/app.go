package app

import (
	"OnlineBar/Backend/internal/transport/rest"
	"OnlineBar/Backend/pkg/cfg"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", rest.TestFunc)
	r.POST("/login", rest.LoginHandler)
	r.POST("/register", rest.RegisterHandler)

	return r
}

func StartServer() {
	config := cfg.ServerConfig()
	port := config.Server.Port
	host := config.Server.Host

	r := setupRouter()

	if err := r.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		log.Println("Failed to start server:", err)
		return
	}

	fmt.Println("Server started on", fmt.Sprintf("%s:%s", host, port))
}