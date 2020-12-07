package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseYear(yearString string, minValue, maxValue int) (year int, err error) {
	_, err = regexp.MatchString(`\d{4}`, yearString)
	if err != nil {
		return year, err
	}
	year, err = strconv.Atoi(yearString)
	if minValue <= year && year <= maxValue {
		return year, nil
	}
	return year, fmt.Errorf("year needs to be between %d and %d", minValue, maxValue)
}

func parseBirthYear(birthYearString string) (birthYear int, err error) {
	return parseYear(birthYearString, 1920, 2002)
}

func parseIssueYear(issueYearString string) (issueYear int, err error) {
	return parseYear(issueYearString, 2010, 2020)
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
