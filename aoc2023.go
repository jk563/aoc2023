// Package for solvinf AoC 2023 challenges
package main

import (
	"bufio"
	"fmt"
	"os"

	"jamiekelly.com/aoc2023/internal/day1"
	"jamiekelly.com/aoc2023/internal/day2"
)

func main() {
	day1Input := fileToLines("assets/input_day1.txt")
	day1Part1Solution := day1.SumCalibrationValuesDigits(day1Input)
	fmt.Printf("Day 1, Part 1: %d\n", day1Part1Solution)
	day1Part2Solution := day1.SumCalibrationValuesDigitsAndWords(day1Input)
	fmt.Printf("Day 1, Part 2: %d\n", day1Part2Solution)

	fmt.Println()

	day2Input := fileToLines("assets/input_day2.txt")
	day2Part1Solution := day2.SumPossibleGames(day2Input)
	fmt.Printf("Day 2, Part 1: %d\n", day2Part1Solution)
	day2Part2Solution := day2.SumPowerMinimumCubes(day2Input)
	fmt.Printf("Day 2, Part 2: %d\n", day2Part2Solution)
}

func fileToLines(filePath string) []string {
	file, _ := os.Open(filePath)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
