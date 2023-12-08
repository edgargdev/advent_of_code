package day02

import (
	"log"
	"strconv"
	"strings"
)

type Game struct {
	id                int
	numberOfBlueDice  int
	numberOfRedDice   int
	numberOfGreenDice int
}

func BuildGame(inputString string) Game {
	gameStringSplit := strings.Split(inputString, ": ")
	games := gameStringSplit[1]
	// log.Printf("Games: %q\n", games)
	revealedDiceInstances := strings.Split(games, "; ")

	gameIdString := strings.Split(gameStringSplit[0], " ")[1]
	possibleDice := map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}

	for _, v := range revealedDiceInstances {
		sameColoredDice := strings.Split(v, ", ")
		// log.Printf("Currently reading sameColoredDice %v\n", sameColoredDice)
		for _, x := range sameColoredDice {
			quantityAndColor := strings.Split(x, " ")
			color := quantityAndColor[1]
			numberString := quantityAndColor[0]
			// log.Printf("This hand had %q %q dice\n", numberString, color)
			quantity, err := strconv.Atoi(numberString)
			if err != nil {
				log.Printf("ERROREE %q", err)
			}
			// log.Printf("possibleDiceColor comparing %v and %v", possibleDice[color], quantity)
			if possibleDice[color] < quantity {
				// log.Printf("assinging %q to %v", color, quantity)
				possibleDice[color] = quantity
			}
		}
	}

	if i, err := strconv.Atoi(gameIdString); err == nil {
		return Game{
			id:                i,
			numberOfBlueDice:  possibleDice["blue"],
			numberOfRedDice:   possibleDice["red"],
			numberOfGreenDice: possibleDice["green"],
		}
	}
	return Game{id: 0}
}
