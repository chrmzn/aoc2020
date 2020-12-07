package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseBirthYear(birthYearString string) (birthYear int, err error) {
	_, err = regexp.MatchString(`\d{4}`, birthYearString)
	if err != nil {
		return birthYear, err
	}
	birthYear, err = strconv.Atoi(birthYearString)
	if 1920 <= birthYear && birthYear <= 2002 {
		return birthYear, nil
	}
	return birthYear, errors.New("Birth year needs to be between 1920 and 2002")
}

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
			personFields := strings.Split(personData, " ")

			for _, field := range personFields {
				fields := strings.Split(field, ":")
				fieldName, fieldValue := fields[0], fields[1]
				fmt.Printf("%s:%s\n", fieldName, fieldValue)
			}
			personData = ""
		}
	}
	fmt.Printf("Valid Passports: %d\n", validPassportCount)
}
