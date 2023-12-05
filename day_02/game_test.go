package day02

import (
	"testing"
)

func TestGameKnowsItsID(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	game := BuildGame(input)
	actual := game.id
	expected := 1

	assertEquals(actual, expected, t)
}

func TestGameKnowsItsIdOtherThanTwo(t *testing.T) {
	input := "Game 22: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	game := BuildGame(input)
	actual := game.id
	expected := 22

	assertEquals(actual, expected, t)
}

func TestGameKnowsColorsItHas(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

	game := BuildGame(input)
	actualID := game.id
	actualBlue := game.numberOfBlueDice
	actualRed := game.numberOfRedDice
	actualGreen := game.numberOfGreenDice

	expectedID := 1
	expectedBlue := 6
	expectedRed := 4
	expectedGreen := 2

	assertEquals(actualID, expectedID, t)
	assertEquals(actualBlue, expectedBlue, t)
	assertEquals(actualRed, expectedRed, t)
	assertEquals(actualGreen, expectedGreen, t)
}


func assertEquals(actual interface{}, expected interface{}, t *testing.T) {
	if actual != expected {
		t.Fatalf("Expected %v but got %v ", expected, actual)
	}
}

