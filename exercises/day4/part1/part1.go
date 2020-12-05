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
			if kvCount == 8 {
				validPassportCount++
				fmt.Println("VALID!")
			} else if kvCount == 7 && !strings.Contains(personData, "cid:") {
				validPassportCount++
				fmt.Println("VALID!")
			}
			fmt.Println("")
			personData = ""
		}
	}
	fmt.Printf("Valid Passports: %d\n", validPassportCount)
}
