package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FindGamesWithPossibleRGB(red int, green int, blue int) {
	log.Printf("Looking for all possible games with following dice %v blue, %v red, %v green\n", blue, red, green)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var allGames []Game
	runningSum := 0
	partTwoRunningSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		if scanner.Text() != "" {
			allGames = append(allGames, BuildGame(scanner.Text()))
		}
	}

	for _, game := range allGames {
		partTwoRunningSum = partTwoRunningSum + (game.numberOfRedDice * game.numberOfBlueDice * game.numberOfGreenDice)
		if game.numberOfBlueDice <= blue && game.numberOfRedDice <= red && game.numberOfGreenDice <= green {
			log.Printf("Game %v is a possible game", game.id)

			runningSum += game.id
		}
	}

	fmt.Println("Part 1 answer")
	fmt.Println(runningSum)

	fmt.Println("Part 2 answer")
	fmt.Println(partTwoRunningSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
