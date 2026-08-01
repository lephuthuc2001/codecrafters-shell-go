package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"unicode"
)

func splitCommandInput(input string) (command string, commandArguments []string) {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		return
	}

	inputFields := strings.Fields(input)
	command = inputFields[0]

	argsToken := strings.TrimSpace(input[len(command):])

	apostrophe := '\u0027'

	runesInput := []rune(input)

	containQuote := slices.Contains(runesInput, apostrophe)

	splitFunc := func(c rune) bool {
		if containQuote {
			return c == apostrophe
		}

		return unicode.IsSpace(c)
	}

	commandArguments = strings.FieldsFunc(argsToken, splitFunc)

	if containQuote {
		for i, val := range commandArguments {
			if len(strings.TrimSpace(val)) == 0 {
				commandArguments[i] = " "
			}
		}
	} else {
		argsLength := len(commandArguments)
		for i, val := range commandArguments {
			if i + 1 <= argsLength - 1 {
				commandArguments[i] = val + " "
			}
		}
	}

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

func handleCommandType(commandTypeArguments []string) {
	if len(commandTypeArguments) == 0 {
		fmt.Println("type command missing arguments")
		return
	}

	command := commandTypeArguments[0]

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

		if len(command) == 0 {
			continue
		}

		switch command {

		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(commandArguments, ""))
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

			dirToGo := commandArguments[0]

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

			out, err := exec.Command(command, commandArguments...).Output()
			fmt.Print(string(out))
		}
	}
}
