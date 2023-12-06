package day03

import "testing"

func TestGridNumberCanReadNumber(t *testing.T) {
	actual := ParseLineForGridNumbers("467..114..")

	assertEquals(467, actual[0].value, t)
	assertEquals(114, actual[1].value, t)
}

func TestGridNumberCanReadOtherNumbers(t *testing.T) {
	actual := ParseLineForGridNumbers("..35..633.")

	assertEquals(35, actual[0].value, t)
	assertEquals(633, actual[1].value, t)
}

func TestGridNumberKnowsItsInitialPosition(t *testing.T) {
	actual := ParseLineForGridNumbers("..35..633.")

	assertEquals(2, actual[0].initialPosition, t)
	assertEquals(6, actual[1].initialPosition, t)
}

func assertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if actual != expected {
		t.Fatalf("Expected %v but got %v ", expected, actual)
	}
}
