package fizzbuzz

import (
	"strconv"
)

/*
Solve function converts an given integer array of length n starting with 1 to a string
if it is divisible by 3 the integer should be converted to Fizz
if it is divisible by 5 the integer should be converted to Buzz
if it is divisible by 3 AND 5 integer should be converted to FizzBuzz
if it is not divisible by 3 or 5, integer should not be converted
*/
func Solve(puzzleLength int) string {
	return decodeEachFizzBuzzPuzzleNumber(puzzleLength, "")
}

func decodeEachFizzBuzzPuzzleNumber(numbersLeft int, decodedNumbers string) string {
	if numbersLeft == 0 {
		return decodedNumbers
	}
	nextDecodedNumber := decodeFizzBuzzPuzzleNumber(numbersLeft)
	return decodeEachFizzBuzzPuzzleNumber(numbersLeft-1, nextDecodedNumber+decodedNumbers)
}

func decodeFizzBuzzPuzzleNumber(n int) string {
	if isDivisibleByThree(n) && isDivisibleByFive(n) {
		return "FizzBuzz"
	}
	if isDivisibleByThree(n) {
		return "Fizz"
	}
	if isDivisibleByFive(n) {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

func isDivisibleByThree(n int) bool {
	return areNumbersDivisible(n, 3)
}

func isDivisibleByFive(n int) bool {
	return areNumbersDivisible(n, 5)
}

func areNumbersDivisible(source, n int) bool {
	return source%n == 0
}
