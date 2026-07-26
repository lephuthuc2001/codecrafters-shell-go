package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var count int;
	for {
		if count == 0 {
			fmt.Print("$ ") 
		} else{
			fmt.Print("\n$ ") 
		}

	reader :=  bufio.NewReader(os.Stdin)
	command, error := reader.ReadString('\n');

	if error != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", error)
		os.Exit(1)
	}

	fmt.Printf("%v: command not found",strings.Trim(command,"\n"))

	count ++
	}
}
