package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	test := sum(5, 4, 3, 2, 1)
	expected := 15

	if test != expected {
		t.Error("Expected value:", expected, "Returned value:", test)
	}
}

func TestShouldSubtractCorrect(t *testing.T) {
	test := subtract(10, 5, 2)
	expected := 3

	if test != expected {
		t.Error("Expected value:", expected, "Returned value:", test)
	}
}

func TestShouldMultiplyCorrect(t *testing.T) {
	test := multiply(1, 2, 5)
	expected := 10

	if test != expected {
		t.Error("Expected value:", expected, "Returned value:", test)
	}
}

func TestShouldDivideCorrect(t *testing.T) {
	test := divide(100, 2, 2)
	expected := 25

	if test != expected {
		t.Error("Expected value:", expected, "Returned value:", test)
	}
}
