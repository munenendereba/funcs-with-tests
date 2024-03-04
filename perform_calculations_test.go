package main

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		a, b any
		want float64
		err  error
	}{
		{3, 4, 7.0, nil},
		{"3", "4", 7.0, nil},
		{"3x", 4, 0.0, errors.New("parsing")},
		{true, "4", 0.0, errors.New("parsing")},
		{"6", "9e4", 90006, nil},
		{7, "9 p", 0.0, errors.New("parsing")},
		{3.0, "7", 10.0, nil},
		{float32(9), 8.7, 17.7, nil},
	}

	for _, test := range tests {
		a, b := test.a, test.b
		testname := fmt.Sprintf(`%s,%s`, a, b)

		t.Run(testname, func(t *testing.T) {
			result, err := AddTwoNumbers(a, b)

			if err != nil {
				if !strings.Contains(err.Error(), test.err.Error()) {
					t.Errorf(`AddTwoNumbers(%s,%s) = Expected "%s", but got "%s"`, a, b, test.err.Error(), err.Error())
				}
			}

			if result != test.want {
				t.Fatalf(`AddTwoNumbers(%s,%s) = want %f but got %f`, a, b, test.want, result)
			}

		})
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddTwoNumbers(20399332.99, "45e9")
	}
}
