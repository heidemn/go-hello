package math

import (
	"fmt"
	"reflect"
	"testing"
)

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected == actual {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
}

func TestSquare(t *testing.T) {
	AssertEqual(t, 25.0, Square(5.0))
	AssertEqual(t, 2.25, Square(1.5))
	AssertEqual(t, 0.25, Square(0.5))

	AssertEqual(t, 0, SquareInt(0))
	AssertEqual(t, 1, SquareInt(1))
	AssertEqual(t, 4, SquareInt(2))

	var tests = []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{-1, 1},
		{2, 4},
		{-2, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("SquareInt(%d) == %d", tt.input, tt.expected), func(t *testing.T) {
			AssertEqual(t, tt.expected, SquareInt(tt.input))
		})
	}
}
