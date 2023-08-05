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

func getFileWordCount(file *os.File) int {
	wordCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := scanner.Text()
		wordCount += len(strings.Split(words, " "))
	}
	return wordCount
}

func getFileCharCount(file *os.File) int {
	charCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		charCount += len(line)
	}
	return charCount
}

func main() {
	cFlag := flag.Bool("c", false, "the command -c")
	lFlag := flag.Bool("l", false, "the command -l")
	wFlag := flag.Bool("w", false, "the command -w")
	mFlag := flag.Bool("m", false, "the command -m")
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
	if *wFlag {
		wordCount := getFileWordCount(file)
		println("   ", wordCount, filePath)
	}
	if *mFlag {
		charCount := getFileCharCount(file)
		println("   ", charCount, filePath)
	}
}
