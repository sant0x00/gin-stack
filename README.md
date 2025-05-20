# Gin Stack

**Gin Stack** is a lightweight and structured routing layer built on top of the [Gin](https://github.com/gin-gonic/gin) web framework.  
It helps organize your API by encouraging modular controller definitions, route grouping, versioning, and optional middleware support.

This library provides a clean abstraction for building scalable APIs while reducing boilerplate and improving maintainability.

---

## Features

- [x] Clean controller-based routing interface
- [x] Support for grouping by API version and domain
- [x] Optional middleware injection per controller
- [x] Flexible handler binding (`Bind()` function) for quick routes
- [x] Easy to extend and integrate into existing Gin apps

---

## Installation

```bash
go get github.com/sant0x00/gin-stack
```

## Project Structure

```
.
├── router.go             # Core router logic
├── examples/             # Usage examples (modular and flat)
│   ├── modular-controllers/
│   └── flat-controller-binding/
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

## Examples

Explore the `examples/` folder to see how to use Gin Stack in different scenarios:

- [examples/modular-controllers](examples/modular-controllers/README.md): A recommended structure where each controller is defined in its own file, implementing the `Controller` interface.
  Ideal for large or scalable APIs.
- [examples/flat-controller-binding](examples/flat-controller-binding/README.md): A simplified approach using Bind() for directly binding handlers.
  Useful for prototyping, health checks, or small applications.

## Contributors

- [sant0x00](https://github.com/sant0x00)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.