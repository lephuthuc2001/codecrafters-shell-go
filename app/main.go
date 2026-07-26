package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("$ ")

	reader :=  bufio.NewReader(os.Stdin)
	command, error := reader.ReadString('\n');

	if error != nil {
		log.Fatal("xx")
	}

	
	fmt.Printf("%v: command not found",strings.Trim(command,"\n"))
}
