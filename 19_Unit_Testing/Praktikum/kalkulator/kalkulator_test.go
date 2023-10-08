package kalkulator
import (
    "testing"
)
func TestSubtract(t *testing.T) {
    result := Subtract(3, 2)
    if result != 1 {
        t.Errorf("Expected 1, but got %f", result)
    }
}
func TestAdd(t *testing.T) {
    result := Add(3, 2)
    if result != 5 {
        t.Errorf("Expected 5, but got %f", result)
    }
}
func TestDivide(t *testing.T) {
    result, err := Divide(6, 2)
    if err != nil {
        t.Errorf("Expected no error, but got %v", err)
    }
    if result != 3 {
        t.Errorf("Expected 3, but got %f", result)
    }

    _, err = Divide(6, 0)
    if err == nil {
        t.Errorf("Expected an error, but got none")
    }
}
func TestMultiply(t *testing.T) {
    result := Multiply(3, 2)
    if result != 6 {
        t.Errorf("Expected 6, but got %f", result)
    }
}