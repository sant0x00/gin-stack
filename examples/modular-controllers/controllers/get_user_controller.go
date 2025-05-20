package controllers

import (
	"github.com/gin-gonic/gin"
	router "github.com/sant0x00/gin-stack"
)

type GetUserController struct{}

func (g GetUserController) GetBind() router.ControllerBind {
	return router.ControllerBind{
		Method:       "get", // You can use too, for example: GET or http.MethodGet
		RelativePath: "/user/:id",
	}
}

func (g GetUserController) Execute(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	// Simulate fetching user data from a database
	user := map[string]string{
		"id":   id,
		"name": "John Doe",
	}

	ctx.JSON(200, gin.H{"user": user})
}
