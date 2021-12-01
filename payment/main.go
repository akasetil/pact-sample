package main

import (
	"github.com/gin-gonic/gin"
	"payment/controller"
)

func main() {
	engine := gin.Default()

	engine.POST("/payment", controller.Payment)

	engine.Run(":3000")
}
