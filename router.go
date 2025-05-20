package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Controller defines the interface that each controller must implement.
// It includes methods to retrieve the controller's binding information and execute the controller logic.
type Controller interface {
	GetBind() ControllerBind  // GetBind Returns the binding details of the controller (method, path, etc.)
	Execute(ctx *gin.Context) // Execute Executes the controller logic, handling the request
}

// ControllerBind holds the information necessary for binding a controller to a route.
// It contains the HTTP method, path, version, group, and any middlewares associated with the controller.
type ControllerBind struct {
	Name         string            // Name The name of the controller
	Method       string            // Method The HTTP method (GET, POST, etc.)
	Version      string            // Version The version of the API (e.g., "v1")
	Group        string            // Group The group or category of the route (e.g., "users")
	RelativePath string            // RelativePath The relative path for the route (e.g., "/create")
	Middlewares  []gin.HandlerFunc // Middlewares to apply to this route
}

// ControllerModule represents a module that contains multiple controllers.
// It defines the method `GetControllers()` to return all the controllers in that module.
type ControllerModule interface {
	GetControllers() []Controller // GetControllers Returns a slice of controllers for this module
}

// GetRouters configures and returns a Gin engine instance with all routes from the provided modules.
// It processes each controller, extracts its binding information, and registers the corresponding routes.
func GetRouters(modules []ControllerModule) (*gin.Engine, error) {
	engine := gin.Default()

	api := engine.Group("/api")
	groups := map[string]*gin.RouterGroup{}

	for _, module := range modules {
		for _, controller := range module.GetControllers() {
			bind := controller.GetBind()

			if strings.TrimSpace(bind.RelativePath) == "" {
				return nil, fmt.Errorf("controller %s has no relative path", bind.Name)
			}

			base, err := buildBasePath(bind.Version, bind.Group)
			if err != nil {
				return nil, err
			}

			fullGroup, exists := groups[base]
			if !exists {
				fullGroup = api.Group(base)
				groups[base] = fullGroup
			}

			handlers := append(bind.Middlewares, controller.Execute)

			if err := registerRouter(fullGroup, bind.Method, bind.RelativePath, handlers); err != nil {
				return nil, fmt.Errorf("failed to register route %s %s: %w", bind.Method, bind.RelativePath, err)
			}
		}
	}

	return engine, nil
}

// buildBasePath constructs the base path for a route group using version and group.
// If both are empty, it returns an empty string.
func buildBasePath(version, group string) (string, error) {
	version = strings.Trim(version, "/")
	group = strings.Trim(group, "/")

	if version == "" && group == "" {
		return "", nil
	}

	return "/" + strings.Join([]string{version, group}, "/"), nil
}

// registerRouter registers a route with the specified HTTP method and path under the provided router group.
func registerRouter(group *gin.RouterGroup, method, path string, handlers []gin.HandlerFunc) error {
	switch method {
	case http.MethodGet:
		group.GET(path, handlers...)
	case http.MethodPost:
		group.POST(path, handlers...)
	case http.MethodPut:
		group.PUT(path, handlers...)
	case http.MethodDelete:
		group.DELETE(path, handlers...)
	case http.MethodPatch:
		group.PATCH(path, handlers...)
	case http.MethodOptions:
		group.OPTIONS(path, handlers...)
	case http.MethodHead:
		group.HEAD(path, handlers...)
	default:
		return fmt.Errorf("unsupported method %s", method)
	}
	return nil
}

// Bind creates a new controller binding for a specific method, group, path, and handler.
// It performs basic validation to ensure the method and path are not empty.
func Bind(method, group, path string, handler gin.HandlerFunc) Controller {
	if strings.TrimSpace(method) == "" {
		fmt.Println("Method is required")
		return nil
	}

	if strings.TrimSpace(path) == "" {
		fmt.Println("Path is required")
		return nil
	}

	return &Bound{
		bind: ControllerBind{
			Method:       method,
			Group:        group,
			RelativePath: path,
		},
		handler: handler,
	}
}

// Bound is an implementation of the Controller interface, representing a bound controller with specific
// method, path, and associated handler.
type Bound struct {
	bind    ControllerBind
	handler gin.HandlerFunc
}

// GetBind returns the binding details for the controller (method, path, etc.)
func (c *Bound) GetBind() ControllerBind {
	return c.bind
}

// Execute executes the controller's handler function for the given request context.
func (c *Bound) Execute(ctx *gin.Context) {
	c.handler(ctx)
}
