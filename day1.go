package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(fileName string) []string {
	lines := make([]string, 0)
	openedFile, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		log.Fatalf("Error opening the file %s\n", err)
		return lines
	}
	scanr := bufio.NewScanner(openedFile)
	for scanr.Scan() {
		lines = append(lines, scanr.Text())
	}
	if err := scanr.Err(); err != nil {
		fmt.Printf("The error when reading the file is %s\n", err)
		return lines
	}
	return lines
}

func main() {
	fmt.Println("The output of readfile is " + readFile("day1.txt")[0])
}
