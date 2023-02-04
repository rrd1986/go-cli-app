# go-cli-app
[![Go Report Card](https://goreportcard.com/badge/github.com/rrd1986/go-cli-app)](https://goreportcard.com/report/github.com/rrd1986/go-cli-app)

## Go modules required for this repo
* go get -u github.com/spf13/cobra@latest
* go install github.com/spf13/cobra-cli@latest
* go install github.com/go-task/task/v3/cmd/task@latest

## Code generation
Code generation our CLI application with the cobra init command followed by the package. The command had generated the application with the correct file structure and imports:

`cobra-cli init`

then to add the get command

`cobra-cli add get`

## Build the app

`$ go build -o bin/go-cli-app main.go`

or

`$ task build`

## Run the app

`$ ./bin/go-cli-app get`

or

`$ task run`

## Test the app

```
rashmiranjandibyajyoti@Rashmis-MacBook-Pro ~/goWorkspace/src/github.com/rrd1986/go-cli-app $% task run 
task: [run] GOFLAGS=-mod=mod go run main.go
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  go-cli-app [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Downloads a sample pdf file from web
  help        Help about any command

Flags:
  -h, --help     help for go-cli-app
  -t, --toggle   Help message for toggle

Use "go-cli-app [command] --help" for more information about a command.
rashmiranjandibyajyoti@Rashmis-MacBook-Pro ~/goWorkspace/src/github.com/rrd1986/go-cli-app $% 
```





