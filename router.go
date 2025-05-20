package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type Controller interface {
	GetBind() ControllerBind
	Execute(ctx *gin.Context)
}

type ControllerBind struct {
	Name         string
	Method       string
	Version      string
	Group        string
	RelativePath string
	Middlewares  []gin.HandlerFunc
}

type ControllerModule interface {
	GetControllers() []Controller
}

func GetRouters(modules []ControllerModule) *gin.Engine {
	engine := gin.Default()

	api := engine.Group("/api")
	groups := map[string]*gin.RouterGroup{}

	for _, module := range modules {
		for _, controller := range module.GetControllers() {
			bind := controller.GetBind()

			if strings.TrimSpace(bind.RelativePath) == "" {
				fmt.Printf("Controller %s has no relative path\n", bind.Name)
				continue
			}

			version := strings.Trim(bind.Version, "/")
			group := strings.Trim(bind.Group, "/")
			path := strings.Trim(bind.RelativePath, " ")

			var base string

			if version != "" {
				base += "/" + version
			}

			if group != "" {
				base += "/" + group
			}

			if base != "" {
				base += "/"
			}
			fullGroup, exists := groups[base]
			if !exists {
				fullGroup = api.Group(base)
				groups[base] = fullGroup
			}

			handlers := append(bind.Middlewares, controller.Execute)

			switch strings.ToUpper(bind.Method) {
			case http.MethodGet:
				fullGroup.GET(path, handlers...)
			case http.MethodPost:
				fullGroup.POST(path, handlers...)
			case http.MethodPut:
				fullGroup.PUT(path, handlers...)
			case http.MethodDelete:
				fullGroup.DELETE(path, handlers...)
			case http.MethodPatch:
				fullGroup.PATCH(path, handlers...)
			case http.MethodOptions:
				fullGroup.OPTIONS(path, handlers...)
			case http.MethodHead:
				fullGroup.HEAD(path, handlers...)
			default:
				fmt.Printf("Unsupported method %s for controller %s\n", bind.Method, bind.Name)
				os.Exit(1)
			}
		}
	}

	return engine
}
