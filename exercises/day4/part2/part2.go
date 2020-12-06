package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	var validPassportCount int

	totalLines := len(fileTextLines)

	for i, line := range fileTextLines {
		personData = personData + " " + line
		if line == "" || i+1 == totalLines {
			personData = strings.TrimSpace(personData)
			fmt.Println(personData)

		}
	}
	fmt.Printf("Valid Passports: %d\n", validPassportCount)
}
