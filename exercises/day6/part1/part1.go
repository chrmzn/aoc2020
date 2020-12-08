package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 6 - Part 1")

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
	var groupCount int
	var personCount int

	totalLines := len(fileTextLines)

	for i, line := range fileTextLines {
		if line != "" || i+1 == totalLines {
			personCount++
		}
		personData = personData + " " + line
		if line == "" || i+1 == totalLines {
			personData = strings.TrimSpace(personData)
			personData = strings.Replace(personData, " ", "", -1)

			fmt.Println(personData)
			counterMap := make(map[rune]int)
			for _, char := range personData {
				counterMap[char]++
			}
			fmt.Println(personCount)
			fmt.Println(counterMap)
			for _, element := range counterMap {
				if element == personCount {
					groupCount++
				}
			}

			personData = ""
			personCount = 0
		}
	}
	fmt.Println(groupCount)
}
