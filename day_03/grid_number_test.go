package day03

import "testing"

func TestGridNumberCanReadNumber(t *testing.T) {
	actual := ParseLineForGridNumbers("467..114..", 0)

	assertEquals(467, actual[0].value, t)
	assertEquals(114, actual[1].value, t)
}

func TestGridNumberCanReadOtherNumbers(t *testing.T) {
	actual := ParseLineForGridNumbers("..35..633.", 0)

	assertEquals(35, actual[0].value, t)
	assertEquals(633, actual[1].value, t)
}

func TestGridNumberKnowsItsInitialPosition(t *testing.T) {
	actual := ParseLineForGridNumbers("..35..633.", 0)

	assertEquals(2, actual[0].initialPosition, t)
	assertEquals(6, actual[1].initialPosition, t)
}

func TestGridNumberKnowsItsEndPoistion(t *testing.T) {
	actual := ParseLineForGridNumbers("..35..633.", 0)

	assertEquals(3, actual[0].endPosition, t)
	assertEquals(8, actual[1].endPosition, t)
}

func TestGridNumberWithRandomCharacters(t *testing.T) {
	actual := ParseLineForGridNumbers(".22.121..\n", 0)

	assertEquals(22, actual[0].value, t)
	assertEquals(1, actual[0].initialPosition, t)
	assertEquals(2, actual[0].endPosition, t)

	assertEquals(121, actual[1].value, t)
	assertEquals(4, actual[1].initialPosition, t)
	assertEquals(6, actual[1].endPosition, t)
}

func TestGridNumberWithThreeDifferentNumbersInOneLine(t *testing.T) {
	actual := ParseLineForGridNumbers(".2&&\\.121.2242.\n", 0)

	assertEquals(2, actual[0].value, t)
	assertEquals(1, actual[0].initialPosition, t)
	assertEquals(1, actual[0].endPosition, t)

	assertEquals(121, actual[1].value, t)
	assertEquals(6, actual[1].initialPosition, t)
	assertEquals(8, actual[1].endPosition, t)

	assertEquals(2242, actual[2].value, t)
	assertEquals(10, actual[2].initialPosition, t)
	assertEquals(13, actual[2].endPosition, t)
}

func assertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if actual != expected {
		t.Fatalf("Expected %v but got %v ", expected, actual)
	}
}
