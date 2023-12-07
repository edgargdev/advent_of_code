package day03

import (
	"testing"
)


// func TestPart1Answer(t *testing.T) {
// 	FindPartNumbersAdjacentToASymbol("input.txt")
// }
//
// func TestSmallTxt(t *testing.T) {
// 	answer := FindPartNumbersAdjacentToASymbol("small.txt")
// 	if answer != 4362 {
// 		t.Fatalf("Expected 4362 not %v\n", answer)
// 	}
// }
//
// func TestRandomTxt(t *testing.T) {
// 	answer := FindPartNumbersAdjacentToASymbol("random.txt")
// 	if answer != 795 {
// 		t.Fatalf("Expected 795 not %v\n", answer)
// 	}
// }
//
// func TestRandomTxtMedium(t *testing.T) {
// 	answer := FindPartNumbersAdjacentToASymbol("random_medium.txt")
// 	if answer != 10797 {
// 		t.Fatalf("Expected 10797 not %v\n", answer)
// 	}
// }

func TestAround(t *testing.T) {
	answer := FindPartNumbersAdjacentToASymbol("around.txt")
	if answer != 0 {
		t.Fatalf("WRONG")
	}
}

func TestLeftBottomeCorner(t *testing.T) {
	actual := FindPartNumbersAdjacentToASymbol("left_bottom_corner.txt")
	if actual != 2985 {
		t.Fatalf("Expected 2985 got %v\n", actual)
	}
}
