package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[INFO] Usage: <prog> <filename>")
		os.Exit(0)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("[Error] " + err.Error())
	}
	defer f.Close()

	fmt.Println("[INFO] Starting compilation for: ", os.Args[1])
	Analyze(f)
}
