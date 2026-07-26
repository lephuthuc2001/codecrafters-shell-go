package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func splitCommand(str string) (command string, commandArguments string) {
	str = strings.Trim(str, "\n")
	str = strings.TrimSpace(str)

	strFields := strings.Fields(str)

	return strFields[0], strings.Join(strFields[1:]," ")
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

		command, commandArguments  := splitCommand(str)

		if command == "exit"{
			break
		}

		if command == "echo" {
			fmt.Println(commandArguments)
			continue
		}


		fmt.Printf("%v: command not found", command)

		fmt.Print("\n")
	}
}
