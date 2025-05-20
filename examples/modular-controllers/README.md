# Modular Controllers Example

This example demonstrates the **recommended approach** when using `gin-stack`: defining each route as an independent controller with proper metadata.

Each controller implements the `Controller` interface with two methods:

- `GetBind()` – declares HTTP method, path, group, version, and optional middlewares
- `Execute(ctx *gin.Context)` – handles the request logic

---

## :file_folder: Structure

```txt
modular-controllers/
├── controllers/
│   ├── create_user_controller.go
│   └── get_user_controller.go
├── main.go
└── README.md
```