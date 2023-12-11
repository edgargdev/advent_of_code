package day04

import "testing"

func TestCalculateAllScratchcardPoints(t *testing.T) {
	GetPointsFromAllScratchcards("resources/input.txt")
}

func TestCalculateAllScratchcardPointsSmall(t *testing.T) {
	GetPointsFromAllScratchcards("resources/small.txt")
}
