package main

import (
	"github.com/androidsr/sgin/controller"
	"github.com/androidsr/sgin/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	route.New(router.Group("/"), "controller", &controller.UserController{})
	router.Run(":8080")
}
