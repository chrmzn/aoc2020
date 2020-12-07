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
	match, err := regexp.MatchString(`^\d{4}$`, yearString)
	if err != nil {
		return year, err
	}
	if !match {
		return year, fmt.Errorf("failed to match year")
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

func parseExpYear(expYearString string) (expYear int, err error) {
	return parseYear(expYearString, 2020, 2030)
}

func parseHeightString(heightString string) (height int, scale string, err error) {
	r, err := regexp.Compile(`^((?P<cm>\d{3})cm|(?P<in>\d{2})in)`)
	if err != nil {
		return height, scale, err
	}
	match := r.FindStringSubmatch(heightString)
	if match == nil {
		return height, scale, fmt.Errorf("Failed to match height regex for %s", heightString)
	}
	result := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	if result["cm"] != "" {
		height, err = strconv.Atoi(result["cm"])
		if err != nil {
			return height, scale, err
		}
		if 150 <= height && height <= 193 {
			return height, "cm", nil
		}
	} else if result["in"] != "" {
		height, err = strconv.Atoi(result["in"])
		if err != nil {
			return height, scale, err
		}
		if 59 <= height && height <= 76 {
			return height, "in", nil
		}
	}
	return height, scale, fmt.Errorf("Not a valid height scale, please use in/cm")
}

func parseHairColour(hairColourString string) (colour string, err error) {
	r, err := regexp.Compile(`^#(?P<colour>[0-9a-f]{6})$`)
	if err != nil {
		return colour, err
	}
	match := r.FindStringSubmatch(hairColourString)
	if match == nil {
		return colour, fmt.Errorf("Failed to match colour regex for %s", hairColourString)
	}
	result := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result["colour"], nil
}

func parseEyeColour(eyeColourString string) (colour string, err error) {
	for _, colour := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if eyeColourString == colour {
			return colour, nil
		}
	}
	return "", fmt.Errorf("Colour %s is not part of the accepted list", eyeColourString)
}

func parsePassportNumber(passportNumberString string) (passportNumber string, err error) {
	match, err := regexp.MatchString(`^\d{9}$`, passportNumberString)
	if err != nil {
		return "", err
	}
	if match {
		return passportNumberString, nil
	}
	return "", fmt.Errorf("Failed to match %s", passportNumberString)
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
			fmt.Println("")
			personData = strings.TrimSpace(personData)
			fmt.Println(personData)
			personFields := strings.Split(personData, " ")

			personMap := make(map[string]string)
			for _, field := range personFields {
				fields := strings.Split(field, ":")
				fieldName, fieldValue := fields[0], fields[1]
				personMap[fieldName] = fieldValue
			}
			fmt.Println(personMap)
			if personMap["byr"] == "" {
				personData = ""
				continue
			} else {
				_, err := parseBirthYear(personMap["byr"])
				if err != nil {
					personData = ""
					continue
				}
			}

			if personMap["iyr"] == "" {
				personData = ""
				continue
			} else {
				_, err := parseIssueYear(personMap["iyr"])
				if err != nil {
					personData = ""
					continue
				}
			}

			if personMap["eyr"] == "" {
				personData = ""
				continue
			} else {
				_, err := parseExpYear(personMap["eyr"])
				if err != nil {
					personData = ""
					continue
				}
			}

			if personMap["hgt"] == "" {
				personData = ""
				continue
			} else {
				_, _, err := parseHeightString(personMap["hgt"])
				if err != nil {
					personData = ""
					continue
				}
			}

			if personMap["hcl"] == "" {
				personData = ""
				continue
			} else {
				_, err := parseHairColour(personMap["hcl"])
				if err != nil {
					personData = ""
					continue
				}
			}

			if personMap["ecl"] == "" {
				personData = ""
				continue
			} else {
				_, err := parseEyeColour(personMap["ecl"])
				if err != nil {
					personData = ""
					continue
				}
			}

			if personMap["pid"] == "" {
				personData = ""
				continue
			} else {
				_, err := parsePassportNumber(personMap["pid"])
				if err != nil {
					personData = ""
					continue
				}
			}

			fmt.Println("Valid Passport!!")
			validPassportCount++
			personData = ""
		}
	}
	fmt.Printf("Valid Passports: %d\n", validPassportCount)
}
