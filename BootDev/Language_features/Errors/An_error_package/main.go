package main

import (
	"errors"
)

// var err error = errors.New("something is  wrong")
// we can also declare the err like this as a variable

func divide(x, y float64) (float64, error) {
	if y == 0 {
        return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}
