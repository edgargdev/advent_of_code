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
	assertEquals(6, scratchcard.gameNumbers[2], t)
}

func TestScratchcardKnowsWhatNumbersWon(t *testing.T) {
	scratchcard := parseLineForScratchcard("Card 2: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	scratchcard2 := parseLineForScratchcard("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83")
	actual := scratchcard.getWinners()
	actual2 := scratchcard2.getWinners()

	assertEquals(8, len(scratchcard.gameNumbers), t)
	assertEquals(4, len(actual), t)

	assertEquals(8, len(scratchcard2.gameNumbers), t)
	assertEquals(1, len(actual2), t)
}

func TestScratchcardKnowsItHasNoWinners(t *testing.T) {
	scratchcard := parseLineForScratchcard("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11")
	actual := scratchcard.getWinners()

	assertEquals(8, len(scratchcard.gameNumbers), t)
	assertEquals(0, len(actual), t)
}

func assertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if actual != expected {
		t.Fatalf("Expected %v but got %v ", expected, actual)
	}
}
