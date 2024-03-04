package main

import (
	"errors"
	"strconv"
)

func AddTwoNumbers(i, j any) (float64, error) {
	num1, err1 := checkNumber(i)
	num2, err2 := checkNumber(j)

	if err1 != nil {
		return 0.0, err1
	}
	if err2 != nil {
		return 0.0, err2
	}

	return num1 + num2, nil
}

func checkNumber(i any) (float64, error) {
	switch t := i.(type) {
	case float64:
		return t, nil
	case float32:
		return float64(t), nil
	case int:
		return float64(t), nil
	case string:
		res, err := strconv.ParseFloat(t, 64)

		if err != nil {
			return 0.0, err
		}

		return res, nil
	default:
		return 0.0, errors.New("invalid value for calculation")
	}
}
