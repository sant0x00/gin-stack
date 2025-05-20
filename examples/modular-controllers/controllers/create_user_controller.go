package controllers

import (
	"github.com/gin-gonic/gin"
	router "github.com/sant0x00/gin-stack"
)

type CreateUserController struct{}

func (c CreateUserController) GetBind() router.ControllerBind {
	return router.ControllerBind{
		Method:       "post", // You can use too, for example: POST or http.MethodPost
		RelativePath: "/user",
	}
}

func (c CreateUserController) Execute(ctx *gin.Context) {
	var user struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"user": user})
}
