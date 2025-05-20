package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sant0x00/gin-stack"
)

type HealthCheckController struct{}

func (c HealthCheckController) GetBind() router.ControllerBind {
	return router.ControllerBind{
		Method:       "get", // or GET, or http.MethodGet
		RelativePath: "/health",
	}
}

func (c HealthCheckController) Execute(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "ok"})
}

type ControllerModule struct{}

func (m ControllerModule) GetControllers() []router.Controller {
	return []router.Controller{
		HealthCheckController{},
	}
}

func main() {
	module := []router.ControllerModule{
		ControllerModule{},
	}

	r := router.GetRouters(module)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
