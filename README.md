# Chapter-4-Demo

In this repo, we have the CLI code (generated from using the Cobra CLI) with a very simple Viper configuration for a basic calculator.

## Update config file:
The configuration for where to store the file is located in the `config.json` file.  Please update it to whatever path works for your environment.

First build the application using the following command:
```go build -o calculator main.go```

Run the application using the following:
```./calculator```
and the following CLI description, usage, and available commands will be printed:

```                
A basic calculator CLI

Usage:
  calculator [command]

Available Commands:
  add         Add value
  clear       Clear result
  completion  Generate the autocompletion script for the specified shell
  divide      Divide value
  help        Help about any command
  multiply    Multiply value
  subtract    Subtract value

Flags:
  -h, --help     help for calculator
  -t, --toggle   Help message for toggle

Use "calculator [command] --help" for more information about a command.
```

## Example usage:
To add 8 to the value:
```./calculator add 8```
To multiply the value against 2:
```./calculator multiply 2```