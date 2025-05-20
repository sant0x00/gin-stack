# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

-

### Changed

-

### Removed

-

### Fixed

-

## [1.0.0] - 2025-05-20

### Added

- Created support for route grouping by API version and domain (`/api/:version/:group`).
- Added `Controller` and `ControllerModule` interfaces to enforce consistent controller behavior.
- Implemented support for custom middlewares per controller.
- Added example folder with two approaches: modular controllers and `Bind()`-style handlers.
- Added dedicated README for each example, explaining the design and structure.
- Introduced `ControllerBind` struct for explicit route declarations.

### Changed

- Refactored `GetRouters` to return an error instead of terminating the process with `os.Exit()`.
- Extracted logic for base path generation and route registration into helper functions (`buildBasePath`, `registerRouter`) to improve modularity and testability.
- Improved naming conventions (`bound` → `Bound`) for consistency with Go export rules.
- Reorganized `examples/` folder to better reflect the two main usage styles of the library.

### Removed

- Removed `os.Exit` calls from routing logic to avoid side effects and improve debuggability.

### Fixed

- Handled edge cases for empty or malformed route paths.
- Improved error handling for unsupported HTTP methods during route registration.
