package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveFizzBuzz(t *testing.T) {
	var tests = []struct {
		inputArray   []int32
		outputString string
	}{
		{
			inputArray:   []int32{3},
			outputString: "Fizz",
		},
		{
			inputArray:   []int32{3, 5},
			outputString: "FizzBuzz",
		},
		{
			inputArray:   []int32{15},
			outputString: "FizzBuzz",
		},
		{
			inputArray:   []int32{2},
			outputString: "2",
		},
		{
			inputArray:   []int32{2, 3, 4, 5, 15, 6, 1},
			outputString: "2Fizz4BuzzFizzBuzzFizz1",
		},
	}

	for _, test := range tests {
		result := Solve(test.inputArray)
		assert.Equal(t, test.outputString, result)
	}
}
