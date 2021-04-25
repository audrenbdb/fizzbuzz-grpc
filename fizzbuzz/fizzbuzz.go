package fizzbuzz

import (
	"strconv"
)

/*
Solve converts an array of int to a string
if it is divisible by 3 the integer should be converted to Fizz
if it is divisible by 5 the integer should be converted to Buzz
if it is divisible by 3 AND 5 integer should be converted to FizzBuzz
if it is not divisible by 3 or 5, integer should not be converted
*/
func Solve(puzzle []int32) string {
	return decodeEachFizzBuzzPuzzleItem(puzzle, "")
}

func decodeEachFizzBuzzPuzzleItem(itemsLeft []int32, decodedItems string) string {
	if len(itemsLeft) == 0 {
		return decodedItems
	}
	nextDecodedItem := decodeFizzBuzzPuzzleItem(itemsLeft[0])
	return decodeEachFizzBuzzPuzzleItem(itemsLeft[1:], decodedItems+nextDecodedItem)
}

func decodeFizzBuzzPuzzleItem(n int32) string {
	if isDivisibleByThree(n) && isDivisibleByFive(n) {
		return "FizzBuzz"
	}
	if isDivisibleByThree(n) {
		return "Fizz"
	}
	if isDivisibleByFive(n) {
		return "Buzz"
	}
	return strconv.Itoa(int(n))
}

func isDivisibleByThree(n int32) bool {
	return areNumbersDivisible(n, 3)
}

func isDivisibleByFive(n int32) bool {
	return areNumbersDivisible(n, 5)
}

func areNumbersDivisible(source, n int32) bool {
	return source%n == 0
}
