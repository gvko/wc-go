package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	cFlag := flag.Bool("c", false, "the command -c")
	flag.Parse()

	fileName := flag.Arg(0)
	fileName = strings.TrimSpace(fileName)
	if fileName == "" {
		fmt.Println("Error: file name must be provided")
		os.Exit(1)
	}

	if *cFlag {
		println("-c was provided")
	} else {
		println("-c was NOT provided")
	}

	println(fileName)
}
