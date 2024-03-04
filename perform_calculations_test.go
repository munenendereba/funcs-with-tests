package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	num1, num2 := 3, 4
	want := 7.0

	result, _ := AddTwoNumbers(num1, num2)

	if want != result {
		t.Fatalf(`AddTwoNumbers(%d,%d) want match for %f, but got %f`, num1, num2, want, result)
	}
}

func TestAddTwoNumbersStrings(t *testing.T) {
	num1 := "6"
	num2 := "9e3"

	result, err := AddTwoNumbers(num1, num2)

	if err != nil {
		t.Fatalf(err.Error())
	}

	want := 9006.

	if want != result {
		t.Fatalf(`AddTwoNumbers(%s,%s) want match for %f, but got %f`, num1, num2, want, result)
	}
}

func TestAddTwoNumbersInvalidString(t *testing.T) {
	num1 := "6"
	num2 := "9 p"

	_, err := AddTwoNumbers(num1, num2)

	if err == nil {
		t.Fatalf(`AddTwoNumbers(%s,%s) want an error`, num1, num2)
	}

}
