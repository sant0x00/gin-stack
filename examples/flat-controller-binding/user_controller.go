package main

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (c *UserController) Create(ctx *gin.Context) {
	ctx.JSON(201, gin.H{"created": true})
}

func (c *UserController) Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"user": "some user"})
}

func (c *UserController) Update(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"updated": true})
}

func (c *UserController) Delete(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"deleted": true})
}
