package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		command, error := reader.ReadString('\n')

		if error != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", error)
			os.Exit(1)
		}

		trimmedCommand := strings.Trim(command, "\n")

		if trimmedCommand == "exit"{
			break
		}

		fmt.Printf("%v: command not found", trimmedCommand)

		fmt.Print("\n")
	}
}
