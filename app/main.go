package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func splitCommandInput(input string) (command string, commandArguments []string) {
	input = strings.Trim(input, "\n")
	input = strings.TrimSpace(input)

	inputFields := strings.Fields(input)

	return inputFields[0], inputFields[1:]
}

func isBuiltInCommand(command string) bool {
	builtInCommand := []string{
		"type",
		"echo",
		"exit",
	}

	return slices.Contains(builtInCommand, command)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		str, error := reader.ReadString('\n')

		if error != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", error)
			os.Exit(1)
		}

		command, commandArguments := splitCommandInput(str)

		switch command {
		case "exit":
			return
		case "echo":
			fmt.Println(strings.Join(commandArguments, " "))
		case "type":
			typeCheckCommand := commandArguments[0]

			if isBuiltInCommand(typeCheckCommand) {
				fmt.Printf("%v is a shell builtin", command)
			} else {
				fmt.Printf("%v: command not found", command)
			}
			fmt.Print("\n")

		default:
			fmt.Printf("%v: command not found", command)
			fmt.Print("\n")
		}
	}
}
