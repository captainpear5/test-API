package main

import (
	"github.com/captainpear5/test-API/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.Run()
}
