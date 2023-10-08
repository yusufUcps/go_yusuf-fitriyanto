package kalkulator

import "errors"

func Subtract(a, b float64) float64 {
    return a - b
}

func Add(a, b float64) float64 {
    return a + b
}

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("Cannot divide by zero")
    }
    return a / b, nil
}

func Multiply(a, b float64) float64 {
    return a * b
}


