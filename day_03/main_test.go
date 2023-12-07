package day03

import (
	"testing"
)


// func TestPart1Answer(t *testing.T) {
// 	FindPartNumbersAdjacentToASymbol("input.txt")
// }

func TestSmallTxt(t *testing.T) {
	answer := FindPartNumbersAdjacentToASymbol("small.txt")
	if answer != 4361 {
		t.Fatalf("Expected 4361 not %v\n", answer)
	}
}
