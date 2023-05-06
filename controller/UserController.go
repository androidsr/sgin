package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// @GET:/test
func (UserController) Get(c *gin.Context) {
	fmt.Println("Hello word")
}
