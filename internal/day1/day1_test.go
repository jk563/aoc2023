// Solves AoC 2023 Day 1
package day1

import "testing"

func TestCalculateCalibrationValuesDigitsExample(t *testing.T) {
	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	expected := 142
	actual := SumCalibrationValuesDigits(input)

	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestCalculateCalibrationValuesDigitsAndWordsExample(t *testing.T) {
	input := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	expected := 281
	actual := SumCalibrationValuesDigitsAndWords(input)

	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

// Coverage fails for five as it isn't in the example. So we repeat with a five inserted in a place that doesn't matter
func TestCalculateCalibrationValuesDigitsAndWordsExampleWithAFive(t *testing.T) {
	input := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightfiveseven2", "zoneight234", "7pqrstsixteen"}
	expected := 281
	actual := SumCalibrationValuesDigitsAndWords(input)

	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
