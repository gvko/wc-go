package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func checkFilePathGiven(filePath string) {
	filePath = strings.TrimSpace(filePath)
	if filePath == "" {
		fmt.Println("Error: file name must be provided")
		os.Exit(1)
	}
}

func getFileSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	checkErr(err)
	return fileInfo.Size()
}

func main() {
	cFlag := flag.Bool("c", false, "the command -c")
	flag.Parse()
	filePath := flag.Arg(0)
	checkFilePathGiven(filePath)
	
	_, err := os.ReadFile(filePath)
	checkErr(err)

	if *cFlag {
		fileSize := getFileSize(filePath)
		println(fileSize, filePath)
	}
}
