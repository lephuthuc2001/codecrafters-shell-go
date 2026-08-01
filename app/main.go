package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func splitCommandInput(input string) (command string, commandArguments string) {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		return
	}

	inputFields := strings.Fields(input)
	command = inputFields[0];

	commandArguments = strings.TrimSpace(input[len(command):])

	return command, commandArguments
}

func isBuiltInCommand(command string) bool {
	builtInCommand := []string{
		"type",
		"echo",
		"exit",
		"pwd",
		"cd",
	}

	return slices.Contains(builtInCommand, command)
}

func handleCommandType(commandTypeArguments string) {
	if len(commandTypeArguments) == 0 {
		fmt.Println("type command missing arguments")
		return
	}

	command := strings.Fields(commandTypeArguments)[0]

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
 
		if len(command) ==0 {
			continue
		}

		switch command {

		case "exit":
			os.Exit(0)
		case "echo":
			if strings.Contains(commandArguments,"'"){
				fmt.Println(strings.ReplaceAll(commandArguments,"'",""))
			} else {
				output := strings.Join(strings.Fields(commandArguments)," ")
				fmt.Println(output)
			}

		case "type":
			handleCommandType(commandArguments)
		case "pwd":
			pwd, err := os.Getwd()

			if err != nil {
				fmt.Fprintln(os.Stderr, "Error printing current directory:", error)
			}

			fmt.Println(pwd)
		case "cd":
			if len(commandArguments) == 0 {
				homePath, err := os.UserHomeDir()

				if err != nil {
					fmt.Fprintln(os.Stderr, "Error changing directory:", error)
				}

				os.Chdir(homePath)
				continue
			}

			dirToGo := strings.Fields(commandArguments)[0]

			if dirToGo == "~" {
				homePath, err := os.UserHomeDir()

				if err != nil {
					fmt.Fprintln(os.Stderr, "Error changing directory:", error)
				}

				os.Chdir(homePath)
				continue
			}

			_, err := os.Stat(dirToGo)

			if err != nil {
				fmt.Fprintf(os.Stderr, "cd: %v: No such file or directory", dirToGo)
				fmt.Println()

			} else {
				os.Chdir(dirToGo)
			}

		default:
			_, err := exec.LookPath(command)

			if err != nil {
				fmt.Printf("%v: not found", command)
				fmt.Print("\n")
				continue
			}

			out, err := exec.Command(command, strings.Fields(commandArguments)...).Output()
			fmt.Print(string(out))
		}
	}
}
