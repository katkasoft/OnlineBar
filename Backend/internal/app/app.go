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
	r.LoadHTMLGlob("Frontend/templates/*")

	r.Static("/static", "./Frontend")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Главная страница",
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"title": "Login page",
		})
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", gin.H{
			"title": "Register page",
		})
	})

	r.GET("/ping", rest.TestFunc)
	r.GET("/product", rest.GetProductList)
	r.GET("/balance/get", rest.GetBalance)
	r.POST("/login", rest.LoginHandler)
	r.POST("/register", rest.RegisterHandler)
	r.POST("/buy", rest.BuyHandler)
	r.POST("/balance/update", rest.UpdateBalance)

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
