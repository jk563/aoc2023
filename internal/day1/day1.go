package day1

import "fmt"
import "strconv"

// SumCalibrationValuesDigits takes a slice of input strings, extracts the calibration values, and adds them together. It assumes that only digits represent integers.
func SumCalibrationValuesDigits(input []string) int {
	total := 0

	for _, line := range input {
		total += calculateCalibrationValue(line)
	}

	return total
}

// SumCalibrationValuesDigitsAndWords takes a slice of input strings, extracts the calibration values, and adds them together. It assumes that digits represent integers and so can words.
func SumCalibrationValuesDigitsAndWords(input []string) int {
	total := 0

	for _, line := range input {
		line = translateLineToIntegers(line)
		total += calculateCalibrationValue(line)
	}

	return total
}

func calculateCalibrationValue(line string) int {
	digits := ""
	for _, codepoint := range line {
		integer, err := runeToInt(codepoint)
		if err == nil {
			digits += strconv.Itoa(integer)
		}
	}
	calibrationValue, _ := strconv.Atoi((string(digits[0]) + string(digits[len(digits)-1])))
	return calibrationValue
}

func runeToInt(r rune) (int, error) {
	if r < '0' || r > '9' {
		return 0, fmt.Errorf("Rune %v is not an int", r)
	}

	return int(r - '0'), nil
}

func translateLineToIntegers(line string) string {
	var newline string

	i := 0
	for i < len(line) {
		translatedSubstring, increment := integerStringOrCharacterScan(line[i:len(line)])
		newline += translatedSubstring
		i += increment
	}

	return newline
}

// Checks if the given substring starts with a spelled out integer
// Returns an integerOrCharacterString depending on if it finds a spelled out integer or not. If not it returns the first character of the input string. Also returns an increment to apply to begin the next search for an integer string
func integerStringOrCharacterScan(input string) (integerOrCharacterString string, increment int) {
	firstRune := input[0]
	inputLength := len(input)

	if firstRune == 'o' {
		if inputLength >= 3 && input[0:3] == "one" {
			return "1", 2
		}
	} else if firstRune == 't' {
		if inputLength >= 3 && input[0:3] == "two" {
			return "2", 2
		} else if inputLength >= 5 && input[0:5] == "three" {
			return "3", 4
		}
	} else if firstRune == 'f' {
		if inputLength >= 4 && input[0:4] == "four" {
			return "4", 3
		} else if inputLength >= 4 && input[0:4] == "five" {
			return "5", 3
		}
	} else if firstRune == 's' {
		if inputLength >= 3 && input[0:3] == "six" {
			return "6", 2
		} else if inputLength >= 5 && input[0:5] == "seven" {
			return "7", 4
		}
	} else if firstRune == 'e' {
		if inputLength >= 5 && input[0:5] == "eight" {
			return "8", 4
		}
	} else if firstRune == 'n' {
		if inputLength >= 4 && input[0:4] == "nine" {
			return "9", 3
		}
	}

	return string(firstRune), 1
}
