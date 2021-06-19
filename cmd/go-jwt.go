package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jimmiepr/Go-JWT/internal/service"
)

func main() {
	r := gin.Default()
	r.POST("/login", service.Login)
	log.Fatal(r.Run(":3000"))
}
