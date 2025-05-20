package main

import (
	"github.com/sant0x00/gin-stack"
	"github.com/sant0x00/gin-stack/examples/modular-controllers/controllers"
)

type ControllerModule struct{}

func (m ControllerModule) GetControllers() []router.Controller {
	return []router.Controller{
		new(controllers.CreateUserController),
		new(controllers.GetUserController),
	}
}

func main() {
	module := []router.ControllerModule{
		new(ControllerModule),
	}

	routers, err := router.GetRouters(module)
	if err != nil {
		panic(err)
	}

	err = routers.Run(":8080")
	if err != nil {
		panic(err)
	}
}
