package main

import (
	"bufio"
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

func getFileLinesCount(file *os.File) int {
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

func main() {
	cFlag := flag.Bool("c", false, "the command -c")
	lFlag := flag.Bool("l", false, "the command -l")
	flag.Parse()

	filePath := flag.Arg(0)
	checkFilePathGiven(filePath)

	file, err := os.Open(filePath)
	checkErr(err)
	defer file.Close()

	if *cFlag {
		fileSize := getFileSize(filePath)
		println("   ", fileSize, filePath)
	}
	if *lFlag {
		lineCount := getFileLinesCount(file)
		println("   ", lineCount, filePath)
	}
}
