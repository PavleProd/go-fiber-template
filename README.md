# Go Fiber Web Server Template

A starter template for web-server in Go using **Fiber**, **Cobra**, and **Viper**

## External Modules
- [Fiber v3](https://github.com/gofiber/fiber) – HTTP web framework
- [Cobra](https://github.com/spf13/cobra) – CLI framework, provides `serve` command for starting web server
- [Viper](https://github.com/spf13/viper) – Load and parse config files/env vars
- [Zerolog](https://github.com/rs/zerolog) – Structured logging
- [invopop/validation](https://github.com/invopop/validation) – Lightweight request/config validation

## Features
- Fiber HTTP Server with example routes, error handling and graceful shutdown
- CLI commands management using Cobra. Root command with common logic and Serve for Web Server
- Config management via Viper, including config validation via invopop/validation
- Dependency Injection Container structure

## Setup

### Prerequisites

- **Go** ≥ 1.23
- **Task** 
  ```bash
  go install github.com/go-task/task/v3/cmd/task@latest
  ```
### Installation

```bash
git clone https://github.com/PavleProd/go-fiber-template.git
cd go-fiber-template

go mod download
```

## Commands
- `task run`: Runs `serve` command in temporary binary
- `task build` Builds binaries into `/bin` folder and Runs `serve` command
- `task test`: Runs tests with formatted output
- `task lint`: Runs code linting with premade configuration (`.golangci.yml`)

These commands can be easily used as GitHub Actions

## Example Endpoint
```bash
curl -X POST "http://localhost:8080/example"   -H "Content-Type: application/json"   -d '{"message":"hello world"}'
```

Response: `200 OK`
