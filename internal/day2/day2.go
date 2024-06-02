package day2

import "fmt"
import "strings"
import "strconv"

// SumPossibleGames takes raw string input and adds up the IDs of the games that would have been possible
func SumPossibleGames(input []string) int {
	sumPossibleGames := 0

	for _, gameInput := range input {
		game := parseGame(gameInput)
		if game.isPossible() {
			sumPossibleGames += game.ID
		}
	}

	return sumPossibleGames
}

// SumPowerMinimumCubes works out the minimum numbers of each colour of cube for each game to be possible and returns the sum  of red, green, and blue cubes multiplied together for each game
func SumPowerMinimumCubes(input []string) int {
	sumPower := 0

	for _, gameInput := range input {
		game := parseGame(gameInput)
		minimumRed := 0
		minimumGreen := 0
		minimumBlue := 0

		for _, reveal := range game.reveals {
			if reveal.red > minimumRed {
				minimumRed = reveal.red
			}
			if reveal.green > minimumGreen {
				minimumGreen = reveal.green
			}
			if reveal.blue > minimumBlue {
				minimumBlue = reveal.blue
			}
		}

		sumPower += minimumRed * minimumGreen * minimumBlue
	}

	return sumPower
}

type game struct {
	ID      int
	reveals []reveal
}

type reveal struct {
	red   int
	blue  int
	green int
}

func (g game) isPossible() bool {
	for _, reveal := range g.reveals {
		if reveal.red > 12 || reveal.green > 13 || reveal.blue > 14 {
			return false
		}
	}

	return true
}

func parseGame(input string) game {
	gameInput := strings.Split(input, ":")

	var gameID int
	gameID, err := strconv.Atoi(strings.Split(gameInput[0], " ")[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to parse gameID. Input: %s", input))
	}

	var reveals []reveal
	for _, revealsInput := range strings.Split(gameInput[1], ";") {
		var reveal reveal
		for _, revealInput := range strings.Split(revealsInput, ",") {
			colorInput := strings.Split(revealInput, " ")
			amount, err := strconv.Atoi(colorInput[1])
			if err != nil {
				panic(fmt.Sprintf("Unable to parse amount. Input: %s", colorInput[1]))
			}
			switch colorInput[2] {
			case "red":
				reveal.red = amount
			case "green":
				reveal.green = amount
			case "blue":
				reveal.blue = amount
			}
		}
		reveals = append(reveals, reveal)
	}
	return game{
		ID:      gameID,
		reveals: reveals,
	}
}
