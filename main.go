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

// DoubleValue takes an integer and doubles it
// Question: Will this modify the original variable? Why or why not?
// Answer: No, because Go passes arguments by value. The function receives a copy of the value,
// so modifications only affect the local copy, not the original variable.
func DoubleValue(x int) int {
	return x * 2
}

// DoublePointer takes a pointer to an integer and doubles the value it points to
// Question: Will this modify the original variable? Why or why not?
// Answer: Yes, because the function receives a pointer (memory address) to the original variable.
// Modifying the value at that address affects the original variable.
func DoublePointer(x *int) {
	*x = *x * 2
}

// CreateOnStack creates a local variable and returns its value
// This variable stays on the stack because it's returned by value
func CreateOnStack() int {
	x := 42
	return x // Value returned, variable stays on stack
}

// CreateOnHeap creates a local variable and returns a pointer to it
// This variable escapes to the heap because the pointer is returned
func CreateOnHeap() *int {
	x := 42
	return &x // Pointer returned, variable escapes to heap
}

// SwapValues swaps two values and returns them
// Does not use pointers - returns new values
func SwapValues(a, b int) (int, int) {
	return b, a
}

// SwapPointers swaps the values that two pointers point to
// Uses pointers to modify original values
func SwapPointers(a, b *int) {
	*a, *b = *b, *a
}

// AnalyzeEscape demonstrates escape analysis by calling both stack and heap functions
func AnalyzeEscape() {
	stackVal := CreateOnStack()
	heapPtr := CreateOnHeap()

	fmt.Printf("Stack value: %d\n", stackVal)
	fmt.Printf("Heap value via pointer: %d\n", *heapPtr)
}

/*
Escape Analysis Explanation:

Which variables escaped to the heap?
- The variable 'x' in CreateOnHeap() escapes to the heap because we return a pointer to it.

Why did they escape?
- When a function returns a pointer to a local variable, the Go compiler determines that the variable
  must outlive the function's stack frame. Since the pointer will be used after the function returns,
  the variable cannot remain on the stack (which gets cleaned up when the function returns).

What does "escapes to heap" mean?
- "Escapes to heap" means the variable is allocated on the heap instead of the stack. The heap is
  a region of memory used for dynamic allocation where objects can persist beyond the lifetime of
  the function that created them. This allows the variable to be accessed safely even after the
  function that created it has returned.

- Variables that don't escape (like in CreateOnStack()) are allocated on the stack, which is faster
  but only exists during the function's execution.

- The Go compiler performs escape analysis automatically to determine where to allocate variables
  for optimal performance and memory safety.
*/

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("\n====== Process Information ======")
	ExploreProcess()

	fmt.Println("\n====== Math Operations ======")
	// Factorial demonstrations
	fact0, err := Factorial(0)
	if err != nil {
		fmt.Printf("Factorial(0) error: %v\n", err)
	} else {
		fmt.Printf("Factorial(0) = %d\n", fact0)
	}

	fact5, err := Factorial(5)
	if err != nil {
		fmt.Printf("Factorial(5) error: %v\n", err)
	} else {
		fmt.Printf("Factorial(5) = %d\n", fact5)
	}

	fact10, err := Factorial(10)
	if err != nil {
		fmt.Printf("Factorial(10) error: %v\n", err)
	} else {
		fmt.Printf("Factorial(10) = %d\n", fact10)
	}

	// Prime number checks
	prime17, err := IsPrime(17)
	if err != nil {
		fmt.Printf("IsPrime(17) error: %v\n", err)
	} else {
		fmt.Printf("IsPrime(17) = %t\n", prime17)
	}

	prime20, err := IsPrime(20)
	if err != nil {
		fmt.Printf("IsPrime(20) error: %v\n", err)
	} else {
		fmt.Printf("IsPrime(20) = %t\n", prime20)
	}

	prime25, err := IsPrime(25)
	if err != nil {
		fmt.Printf("IsPrime(25) error: %v\n", err)
	} else {
		fmt.Printf("IsPrime(25) = %t\n", prime25)
	}

	// Power calculations
	power2_8, err := Power(2, 8)
	if err != nil {
		fmt.Printf("Power(2, 8) error: %v\n", err)
	} else {
		fmt.Printf("Power(2, 8) = %d\n", power2_8)
	}

	power5_3, err := Power(5, 3)
	if err != nil {
		fmt.Printf("Power(5, 3) error: %v\n", err)
	} else {
		fmt.Printf("Power(5, 3) = %d\n", power5_3)
	}

	fmt.Println("\n====== Closure Demonstration ======")
	// Independent counters
	counter1 := MakeCounter(0)
	counter2 := MakeCounter(100)

	fmt.Printf("Counter1: %d\n", counter1())
	fmt.Printf("Counter1: %d\n", counter1())
	fmt.Printf("Counter2: %d\n", counter2())
	fmt.Printf("Counter1: %d\n", counter1())
	fmt.Printf("Counter2: %d\n", counter2())

	// Multipliers
	doubler := MakeMultiplier(2)
	tripler := MakeMultiplier(3)
	testNumber := 7

	fmt.Printf("Doubler(%d) = %d\n", testNumber, doubler(testNumber))
	fmt.Printf("Tripler(%d) = %d\n", testNumber, tripler(testNumber))

	fmt.Println("\n====== Higher-Order Functions ======")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original: %v\n", numbers)

	// Apply - square all numbers
	squared := Apply(numbers, func(x int) int {
		return x * x
	})
	fmt.Printf("Squared: %v\n", squared)

	// Filter - even numbers
	evens := Filter(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Printf("Evens: %v\n", evens)

	// Reduce - sum all numbers
	sum := Reduce(numbers, 0, func(acc, curr int) int {
		return acc + curr
	})
	fmt.Printf("Sum: %d\n", sum)

	// Compose - double then add 10
	doubleThenAdd10 := Compose(
		func(x int) int { return x + 10 },
		func(x int) int { return x * 2 },
	)
	result := doubleThenAdd10(5)
	fmt.Printf("Double then add 10 of 5: %d\n", result)

	fmt.Println("\n====== Pointer Demonstration ======")
	// Swap values demonstration
	a, b := 5, 10
	fmt.Printf("Before SwapValues: a=%d, b=%d\n", a, b)
	newA, newB := SwapValues(a, b)
	fmt.Printf("After SwapValues: a=%d, b=%d (originals unchanged)\n", a, b)
	fmt.Printf("Returned values: newA=%d, newB=%d\n", newA, newB)

	// Swap pointers demonstration
	x, y := 15, 25
	fmt.Printf("Before SwapPointers: x=%d, y=%d\n", x, y)
	SwapPointers(&x, &y)
	fmt.Printf("After SwapPointers: x=%d, y=%d (originals modified)\n", x, y)

	fmt.Println("\n====== Escape Analysis ======")
	AnalyzeEscape()
}
