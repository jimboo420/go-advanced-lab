package main

import (
	"errors"
	"fmt"
	"os"
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

// Apply applies a function to each element of a slice and returns a new slice
func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = operation(num)
	}
	return result
}

// Filter returns a new slice containing only elements where predicate returns true
func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range nums {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

// Reduce reduces a slice to a single value using the operation function
func Reduce(nums []int, initial int, operation func(accumulator, current int) int) int {
	accumulator := initial
	for _, num := range nums {
		accumulator = operation(accumulator, num)
	}
	return accumulator
}

// Compose returns a new function that is the composition of f and g (f(g(x)))
func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// ExploreProcess demonstrates process exploration and memory addresses
func ExploreProcess() {
	fmt.Println("\n=== Process Exploration ===")
	
	// Get current process ID
	pid := os.Getpid()
	fmt.Printf("Current Process ID: %d\n", pid)
	fmt.Println("- Process ID (PID) is a unique identifier assigned by the operating system to each running process")
	
	// Get parent process ID
	ppid := os.Getppid()
	fmt.Printf("Parent Process ID: %d\n", ppid)
	fmt.Println("- Process isolation ensures each process runs in its own protected memory space")
	
	// Create a slice of integers
	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice contents: %v\n", data)
	
	// Print the memory address of the slice header
	fmt.Printf("Slice header address: %p\n", &data)
	fmt.Println("- The slice header contains metadata (length, capacity, pointer to underlying array)")
	
	// Print the memory address of the first element in the slice
	if len(data) > 0 {
		fmt.Printf("First element address: %p\n", &data[0])
		fmt.Println("- This is the actual memory location where the first element is stored")
	}
	
	fmt.Println("- Process isolation means other processes cannot access these memory addresses")
	fmt.Println("- Each process operates in its own virtual address space for security and stability")
}

func main() {
	fmt.Println("Hello, World!")
	ExploreProcess()
}
