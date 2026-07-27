package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

func handleCommandType(command string) {
	if isBuiltInCommand(command) {
		fmt.Printf("%v is a shell builtin", command)
	} else {
		path, err := exec.LookPath(command)

		if err != nil {
			fmt.Printf("%v: not found", command)
		} else {
			fmt.Printf("%v is %v", command, path)
		}
	}

	fmt.Print("\n")
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
			handleCommandType(commandArguments[0])
		default:
			fmt.Printf("%v: command not found", command)
			fmt.Print("\n")
		}
	}
}
