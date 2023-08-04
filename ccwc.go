package main

import (
	"flag"
)

func main() {
	cFlag := flag.Bool("c", false, "the command -c")
	flag.Parse()
	fileName := flag.Arg(0)

	if *cFlag {
		println("-c was provided")
	} else {
		println("-c was NOT provided")
	}

	println(fileName)
}
