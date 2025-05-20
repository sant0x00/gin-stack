package main

import (
	router "github.com/sant0x00/gin-stack"
)

type UserModule struct {
	controller *UserController
}

func NewUserModule() router.ControllerModule {
	controller := new(UserController)
	return &UserModule{controller: controller}
}

func (m *UserModule) GetControllers() []router.Controller {
	return []router.Controller{
		router.Bind("POST", "users", "/", m.controller.Create),
		router.Bind("GET", "users", "/:id", m.controller.Get),
		router.Bind("PUT", "users", "/:id", m.controller.Update),
		router.Bind("DELETE", "users", "/:id", m.controller.Delete),
	}
}

func main() {
	module := []router.ControllerModule{
		NewUserModule(),
	}

	routers, err := router.GetRouters(module)
	if err != nil {
		return
	}

	err = routers.Run(":8081")
	if err != nil {
		panic(err)
	}
}
