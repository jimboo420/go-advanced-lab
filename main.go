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

func main() {
	fmt.Println("Hello, World!")
}
