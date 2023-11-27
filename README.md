# Capita-challenge-server

## Installation

```shell
go install
go mod tidy
```

## Development commands

Runs the app locally

```shell
go run .
```

# Swagger commands

Install swagger

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

Add go binary folder to PATH

```shell
export PATH="~/go/bin:$PATH"
```

After changes to swag documentation

```shell
swag init --parseDependency -g main.go
```
