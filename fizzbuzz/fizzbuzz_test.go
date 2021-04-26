package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveFizzBuzz(t *testing.T) {
	var tests = []struct {
		puzzleLength int
		outputString string
	}{
		{
			puzzleLength: 1,
			outputString: "1",
		},
		{
			puzzleLength: 2,
			outputString: "12",
		},
		{
			puzzleLength: 3,
			outputString: "12Fizz",
		},
		{
			puzzleLength: 5,
			outputString: "12Fizz4Buzz",
		},
		{
			puzzleLength: 15,
			outputString: "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz",
		},
		{
			puzzleLength: 16,
			outputString: "12Fizz4BuzzFizz78FizzBuzz11Fizz1314FizzBuzz16",
		},
	}

	for _, test := range tests {
		result := Solve(test.puzzleLength)
		assert.Equal(t, test.outputString, result)
	}
}
