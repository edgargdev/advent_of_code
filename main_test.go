package day_01

import (
	"testing"
)

func TestCalibrateValueSimpleCase(t *testing.T) {
	input := "1abc2"

	output := CalibrateValue(input)
	expected := 12

	assertEquals(output, expected, t)
}

func TestCalibrateValueOtherCase(t *testing.T) {
	input := "pqr3stu8vwx"

	output := CalibrateValue(input)
	expected := 38

	assertEquals(output, expected, t)
}

func TestCalibrateValueWithMultipleIntegersInString(t *testing.T) {
	input := "a1b2c3d4e5f"

	output := CalibrateValue(input)
	expected := 15

	assertEquals(output, expected, t)
}

func TestCalibrateValueWithSingleInteger(t *testing.T) {
	input := "treb7uchet"

	output := CalibrateValue(input)
	expected := 77

	assertEquals(output, expected, t)
}

func TestCalibrateValueWithWordNumnbers(t *testing.T) {
	input := "two1nine"

	output := CalibrateValue(input)
	expected := 29

	assertEquals(output, expected, t)
}

func assertEquals(actual interface{}, expected interface{}, t *testing.T) {
	if actual != expected {
		t.Fatalf("Expected %v but got %v ", expected, actual)
	}
}
