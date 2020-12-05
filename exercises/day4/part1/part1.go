package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 4 - Part 1")

	readFile, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	var personData string

	for _, line := range fileTextLines {
		personData = personData + " " + line
		if line == "" {
			fmt.Println(personData)
			kvCount := 0
			for _, char := range personData {
				if char == ':' {
					kvCount++
				}
			}
			fmt.Printf("Total KV Pairs: %d\n", kvCount)
			personData = ""
		}
	}
}
