package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("$ ")

	reader :=  bufio.NewReader(os.Stdin)
	command, error := reader.ReadString('\n');

	if error != nil {
		log.Fatal("xx")
	}

	fmt.Printf("%v: command not found",command)
}
