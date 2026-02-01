package main

import (
	"errors"
	"fmt"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >=2")
	}
	if n == 2 {
		return true, nil
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	if exponent == 0 {
		return 1, nil
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

// MakeCounter returns a function that increments and returns a counter
// Each call to the returned function increments the counter by 1
// Multiple counters are independent
func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

// MakeMultiplier returns a function that multiplies its input by the captured factor
func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// MakeAccumulator returns three functions that share the same captured state
// add increases the accumulator, subtract decreases it, get returns current value
func MakeAccumulator(initial int) (func(int), func(int), func() int) {
	accumulator := initial
	add := func(x int) {
		accumulator += x
	}
	subtract := func(x int) {
		accumulator -= x
	}
	get := func() int {
		return accumulator
	}
	return add, subtract, get
}

func main() {
	fmt.Println("Hello, World!")
}
