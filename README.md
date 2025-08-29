# Go Fiber Web Server Template

A starter template for web-server in Go using **Fiber**, **Cobra**, and **Viper**

## External Modules
- [Fiber](https://github.com/gofiber/fiber) – HTTP web framework, fast and minimal
- [Cobra](https://github.com/spf13/cobra) – CLI framework, provides `serve` command
- [Viper](https://github.com/spf13/viper) – Load and parse config files/env vars
- [Zerolog](https://github.com/rs/zerolog) – Structured logging
- [invopop/validation](https://github.com/invopop/validation) – Lightweight request/config validation

## Features
- Fiber HTTP Server with example routes, error handling and graceful shutdown
- CLI commands management using Cobra. Root command with common logic and Serve for Web Server
- Config management via Viper, including config validation via invopop/validation
- Dependency Injection Container structure

## Commands
- **GolangCI-Lint** – Code linting with premade configuration (`task lint`)
- **gotestsum** – Test runner with clean output (`task test`)

## Installation
```bash
git clone https://github.com/PavleProd/go-fiber-template.git
cd go-fiber-template

go mod download
```
## Usage
Run the server:
```bash
go run . serve
```

Build a binary:
```bash
go build -o server .
./server serve
```

## Example Endpoint
```bash
curl -X POST "http://localhost:8080/example"   -H "Content-Type: application/json"   -d '{"message":"hello world"}'
```

Response: `200 OK`
