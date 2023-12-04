package main

import (
	"testing"
)

func TestCalibrateValueSimpleCase(t *testing.T) {
	input := "1abc2"

	output := CalibrateValue(input)
	if output != 12 {
		t.Fatalf("Expected 12, but got %v ", output)
	}
}

func TestCalibrateValueOtherCase(t *testing.T) {
	input := "pqr3stu8vwx"

	output := CalibrateValue(input)
	expected := 38
	if output != expected {
		t.Fatalf("Expected %v but got %v ", expected, output)
	}
}

func TestCalibrateValueWithMultipleIntegersInString(t *testing.T) {
	input := "a1b2c3d4e5f"

	output := CalibrateValue(input)
	expected := 15
	if output != expected {
		t.Fatalf("Expected %v but got %v ", expected, output)
	}
}

func TestCalibrateValueWithSingleInteger(t *testing.T) {
	input := "treb7uchet"

	output := CalibrateValue(input)
	expected := 77
	if output != expected {
		t.Fatalf("Expected %v but got %v ", expected, output)
	}
}
