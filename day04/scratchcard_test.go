package day04

import "testing"

func TestScratchcardKnowsItsID(t *testing.T) {
	scratchcard := parseLineForScratchcard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")

	assertEquals(1, scratchcard.id, t)
}

func TestScratchcardKnowsItsWinningNumbers(t *testing.T) {
	scratchcard := parseLineForScratchcard("Card 2: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")

	assertEquals(2, scratchcard.id, t)
	assertEquals(5, len(scratchcard.winningNumbers), t)
	assertEquals(41, scratchcard.winningNumbers[0], t)
	assertEquals(48, scratchcard.winningNumbers[1], t)
	assertEquals(17, scratchcard.winningNumbers[4], t)
}

func TestScratchcardKnowsWinningNumbersLonger(t *testing.T) {
	scratchcard := parseLineForScratchcard("Card 3: 41 48 83 86 17 11 23 | 83 86  6 31 17  9 48 53")

	assertEquals(3, scratchcard.id, t)
	assertEquals(7, len(scratchcard.winningNumbers), t)
}

func TestScratchcardKnowsItsGameNumbers(t *testing.T) {
	scratchcard := parseLineForScratchcard("Card 2: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")

	assertEquals(8, len(scratchcard.gameNumbers), t)
}

func assertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if actual != expected {
		t.Fatalf("Expected %v but got %v ", expected, actual)
	}
}
