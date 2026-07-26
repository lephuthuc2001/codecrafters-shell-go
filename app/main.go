package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func splitCommandInput(input string) (command string, commandArguments string) {
	input = strings.Trim(input, "\n")
	input = strings.TrimSpace(input)

	inputFields := strings.Fields(input)

	return inputFields[0], strings.Join(inputFields[1:], " ")
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
			fmt.Println(commandArguments)
		default:
			fmt.Printf("%v: command not found", command)

			fmt.Print("\n")
		}
	}
}
