package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	//Load input into array of strings
	data, err := parseInput()

	if err != nil {
		fmt.Println("Unable to extract values from input file:", err)
		return
	}

	//Test array of strings
	output, err := testCards(data)

	if err != nil {
		fmt.Println("Error testing cards:", err)
		return
	}

	for _, result := range output {
		fmt.Println(result)
	}
}

func parseInput() ([]string, error) {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Unable to open input:", err)
		return nil, nil
	}

	defer file.Close()

	var ccList []string

	s := bufio.NewScanner(file)

	for s.Scan() {
		ccNum := s.Text()
		ccList = append(ccList, ccNum)
	}

	return ccList, nil
}

func testCards(ccList []string) ([]string, error) {

	var output []string
	//Loop over cards to begin testing
	for _, cc := range ccList {

		//Build regex test first character and separator requirements for credit card
		pattern := `^[4,5,6][0-9, -]*$`
		regex, err := regexp.Compile(pattern)

		if err != nil {
			fmt.Println("Error building regex:", err)
			return nil, nil
		}

		//If card does not pass regex, it is invalid, move to next
		if !regex.MatchString(cc) {
			output = append(output, "Invalid")
			continue
		}

		//Testing Digit counts
		digitCount := 0
		for _, char := range cc {
			if unicode.IsDigit(char) {
				digitCount++
			}
		}

		//If 16 digits were not present, invalidate and move on
		if digitCount != 16 {
			output = append(output, "Invalid")
			continue
		}

		//Testing if hyphenated, are the hyphens valid
		hyphen := '-'
		if strings.ContainsRune(cc, hyphen) {
			//If the cc number contained a hyphen, test that it is properly segmented
			pattern := `^[\d]{4}-[\d]{4}-[\d]{4}-[\d]{4}$`
			regex, err := regexp.Compile(pattern)

			if err != nil {
				fmt.Println("Error building regex 2:", err)
				return nil, nil
			}

			//If card does not pass regex, it is invalid, move to next
			if !regex.MatchString(cc) {
				output = append(output, "Invalid")
				continue
			}
		}

		var dupeCount int = -1
		var lastDigit int = -1
		//Testing consecutive digits
		for _, char := range cc {
			//This test ignores everything but digit
			if !unicode.IsDigit(char) {
				continue
			}

			intVal := int(char)

			if lastDigit == intVal {
				dupeCount++
			} else {
				lastDigit = intVal
				dupeCount = 1
			}

			if dupeCount >= 4 {
				output = append(output, "Invalid")
			}
		}

		//If to this point, the cc has passed all tests
		output = append(output, "Valid")
	}

	return output, nil
}
