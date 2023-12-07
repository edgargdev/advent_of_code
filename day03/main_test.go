package day03

import (
	"testing"
)

func TestPart1Answer(t *testing.T) {
	FindPartNumbersAdjacentToASymbol("resources/input.txt")
}

func TestSmallTxt(t *testing.T) {
	answer := FindPartNumbersAdjacentToASymbol("resources/small.txt")
	if answer != 4362 {
		t.Fatalf("Expected 4362 not %v\n", answer)
	}
}

func TestRandomTxt(t *testing.T) {
	FindPartNumbersAdjacentToASymbol("resources/random.txt")
}

func TestRandomTxtMedium(t *testing.T) {
	FindPartNumbersAdjacentToASymbol("resources/random_medium.txt")
}

func TestAround(t *testing.T) {
	FindPartNumbersAdjacentToASymbol("resources/around.txt")
}

func TestLeftBottomeCorner(t *testing.T) {
	FindPartNumbersAdjacentToASymbol("resources/left_bottom_corner.txt")
}

func TestGearRationWithLBC(t *testing.T) {
	FindPartNumbersAdjacentToASymbol("resources/left_bottom_corner.txt")
}

func TestGearRationWithRandom(t *testing.T) {
	FindPartNumbersAdjacentToASymbol("resources/random.txt")
}

func TestGearRationWithRandomSmall(t *testing.T) {
	FindPartNumbersAdjacentToASymbol("resources/small_random.txt")
}

func TestJustNumbers(t *testing.T) {
	FindPartNumbersAdjacentToASymbol("resources/just_numbers.txt")
}
