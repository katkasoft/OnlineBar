package main

import (
	"fmt"

	"OnlineBar/internal/transport/rest"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", rest.TestFunc)
	r.GET("/login", rest.LoginHandler)

	r.Run() // listen and serve on 0.0.0.0:8080

	fmt.Printf("Server Started")
}
