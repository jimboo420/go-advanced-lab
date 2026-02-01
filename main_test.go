package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{name: "factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "factorial of 10", input: 10, want: 3628800, wantErr: false},
		{name: "factorial of -1", input: -1, want: 0, wantErr: true},
		{name: "factorial of -5", input: -5, want: 0, wantErr: true},
		{name: "factorial of 3", input: 3, want: 6, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{name: "prime check of 2", input: 2, want: true, wantErr: false},
		{name: "prime check of 3", input: 3, want: true, wantErr: false},
		{name: "prime check of 5", input: 5, want: true, wantErr: false},
		{name: "prime check of 4", input: 4, want: false, wantErr: false},
		{name: "prime check of 9", input: 9, want: false, wantErr: false},
		{name: "prime check of 1", input: 1, want: false, wantErr: true},
		{name: "prime check of 0", input: 0, want: false, wantErr: true},
		{name: "prime check of -3", input: -3, want: false, wantErr: true},
		{name: "prime check of 11", input: 11, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPrime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{name: "2^0", base: 2, exponent: 0, want: 1, wantErr: false},
		{name: "5^0", base: 5, exponent: 0, want: 1, wantErr: false},
		{name: "0^1", base: 0, exponent: 1, want: 0, wantErr: false},
		{name: "0^5", base: 0, exponent: 5, want: 0, wantErr: false},
		{name: "2^3", base: 2, exponent: 3, want: 8, wantErr: false},
		{name: "3^4", base: 3, exponent: 4, want: 81, wantErr: false},
		{name: "5^-1", base: 5, exponent: -1, want: 0, wantErr: true},
		{name: "10^2", base: 10, exponent: 2, want: 100, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name           string
		start          int
		calls          int
		expectedValues []int
	}{
		{name: "counter starting at 0", start: 0, calls: 3, expectedValues: []int{1, 2, 3}},
		{name: "counter starting at 10", start: 10, calls: 2, expectedValues: []int{11, 12}},
		{name: "counter starting at -5", start: -5, calls: 4, expectedValues: []int{-4, -3, -2, -1}},
		{name: "counter single call", start: 100, calls: 1, expectedValues: []int{101}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := MakeCounter(tt.start)
			for i := 0; i < tt.calls; i++ {
				result := counter()
				expected := tt.expectedValues[i]
				if result != expected {
					t.Errorf("After %d calls: got %v, want %v", i+1, result, expected)
				}
			}
		})
	}

	// Test counter independence
	t.Run("independent counters", func(t *testing.T) {
		counter1 := MakeCounter(0)
		counter2 := MakeCounter(10)

		val1 := counter1()
		val2 := counter2()
		val3 := counter1()
		val4 := counter2()

		if val1 != 1 {
			t.Errorf("counter1 first call: got %v, want 1", val1)
		}
		if val2 != 11 {
			t.Errorf("counter2 first call: got %v, want 11", val2)
		}
		if val3 != 2 {
			t.Errorf("counter1 second call: got %v, want 2", val3)
		}
		if val4 != 12 {
			t.Errorf("counter2 second call: got %v, want 12", val4)
		}
	})
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name     string
		factor   int
		inputs   []int
		expected []int
	}{
		{name: "multiply by 2", factor: 2, inputs: []int{1, 5, 10, -3}, expected: []int{2, 10, 20, -6}},
		{name: "multiply by 1", factor: 1, inputs: []int{7, 0, -5}, expected: []int{7, 0, -5}},
		{name: "multiply by -3", factor: -3, inputs: []int{2, 4, -1}, expected: []int{-6, -12, 3}},
		{name: "multiply by 0", factor: 0, inputs: []int{100, 50}, expected: []int{0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiplier := MakeMultiplier(tt.factor)
			for i, input := range tt.inputs {
				result := multiplier(input)
				expected := tt.expected[i]
				if result != expected {
					t.Errorf("Multiply %d by %d: got %v, want %v", input, tt.factor, result, expected)
				}
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	tests := []struct {
		name           string
		initial        int
		operations     []struct {
			op   string
			val  int
			want int
		}
	}{
		{
			name:    "basic accumulation",
			initial: 0,
			operations: []struct {
				op   string
				val  int
				want int
			}{
				{op: "add", val: 5, want: 5},
				{op: "add", val: 3, want: 8},
				{op: "subtract", val: 2, want: 6},
				{op: "add", val: 10, want: 16},
			},
		},
		{
			name:    "starting at negative",
			initial: -10,
			operations: []struct {
				op   string
				val  int
				want int
			}{
				{op: "add", val: 15, want: 5},
				{op: "subtract", val: 3, want: 2},
				{op: "subtract", val: 10, want: -8},
			},
		},
		{
			name:    "mixed operations",
			initial: 100,
			operations: []struct {
				op   string
				val  int
				want int
			}{
				{op: "subtract", val: 25, want: 75},
				{op: "add", val: 50, want: 125},
				{op: "subtract", val: 125, want: 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, subtract, get := MakeAccumulator(tt.initial)
			
			// Test initial value
			current := get()
			if current != tt.initial {
				t.Errorf("Initial value: got %v, want %v", current, tt.initial)
			}

			// Test operations
			for i, op := range tt.operations {
				switch op.op {
				case "add":
					add(op.val)
				case "subtract":
					subtract(op.val)
				}
				
				current := get()
				if current != op.want {
					t.Errorf("After operation %d (%s %d): got %v, want %v", i+1, op.op, op.val, current, op.want)
				}
			}
		})
	}

	// Test multiple accumulators are independent
	t.Run("independent accumulators", func(t *testing.T) {
		add1, _, get1 := MakeAccumulator(0)
		add2, subtract2, get2 := MakeAccumulator(100)

		add1(5)
		subtract2(25)
		add1(10)
		add2(50)

		if get1() != 15 {
			t.Errorf("accumulator1: got %v, want 15", get1())
		}
		if get2() != 125 {
			t.Errorf("accumulator2: got %v, want 125", get2())
		}
	})
}

func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		operation func(int) int
		expected  []int
	}{
		{
			name: "square numbers",
			nums: []int{1, 2, 3, 4},
			operation: func(x int) int {
				return x * x
			},
			expected: []int{1, 4, 9, 16},
		},
		{
			name: "double numbers",
			nums: []int{1, 2, 3, 4},
			operation: func(x int) int {
				return x * 2
			},
			expected: []int{2, 4, 6, 8},
		},
		{
			name: "negate numbers",
			nums: []int{1, -2, 3, -4},
			operation: func(x int) int {
				return -x
			},
			expected: []int{-1, 2, -3, 4},
		},
		{
			name: "add 10",
			nums: []int{5, 10, 15},
			operation: func(x int) int {
				return x + 10
			},
			expected: []int{15, 20, 25},
		},
		{
			name:      "empty slice",
			nums:      []int{},
			operation: func(x int) int { return x },
			expected:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Apply(tt.nums, tt.operation)
			if len(result) != len(tt.expected) {
				t.Errorf("Apply() length = %v, want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Apply()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}

	// Test that original slice is not modified
	t.Run("original slice not modified", func(t *testing.T) {
		original := []int{1, 2, 3}
		Apply(original, func(x int) int { return x * 2 })
		if original[0] != 1 || original[1] != 2 || original[2] != 3 {
			t.Errorf("Apply() modified original slice: got %v, want [1, 2, 3]", original)
		}
	})
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name: "even numbers",
			nums: []int{1, 2, 3, 4, 5, 6},
			predicate: func(x int) bool {
				return x%2 == 0
			},
			expected: []int{2, 4, 6},
		},
		{
			name: "positive numbers",
			nums: []int{-5, -2, 0, 3, 7, -1},
			predicate: func(x int) bool {
				return x > 0
			},
			expected: []int{3, 7},
		},
		{
			name: "numbers greater than 10",
			nums: []int{5, 10, 15, 20, 25},
			predicate: func(x int) bool {
				return x > 10
			},
			expected: []int{15, 20, 25},
		},
		{
			name: "numbers divisible by 3",
			nums: []int{1, 3, 5, 6, 9, 10},
			predicate: func(x int) bool {
				return x%3 == 0
			},
			expected: []int{3, 6, 9},
		},
		{
			name:      "empty slice",
			nums:      []int{},
			predicate: func(x int) bool { return true },
			expected:  []int{},
		},
		{
			name: "no matches",
			nums: []int{1, 3, 5, 7},
			predicate: func(x int) bool {
				return x > 10
			},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Filter(tt.nums, tt.predicate)
			if len(result) != len(tt.expected) {
				t.Errorf("Filter() length = %v, want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Filter()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		initial   int
		operation func(accumulator, current int) int
		expected  int
	}{
		{
			name:    "sum numbers",
			nums:    []int{1, 2, 3, 4},
			initial: 0,
			operation: func(acc, curr int) int {
				return acc + curr
			},
			expected: 10,
		},
		{
			name:    "product of numbers",
			nums:    []int{1, 2, 3, 4},
			initial: 1,
			operation: func(acc, curr int) int {
				return acc * curr
			},
			expected: 24,
		},
		{
			name:    "max number",
			nums:    []int{3, 7, 2, 9, 1},
			initial: -1000,
			operation: func(acc, curr int) int {
				if curr > acc {
					return curr
				}
				return acc
			},
			expected: 9,
		},
		{
			name:    "min number",
			nums:    []int{3, 7, 2, 9, 1},
			initial: 1000,
			operation: func(acc, curr int) int {
				if curr < acc {
					return curr
				}
				return acc
			},
			expected: 1,
		},
		{
			name:    "empty slice returns initial",
			nums:    []int{},
			initial: 100,
			operation: func(acc, curr int) int {
				return acc + curr
			},
			expected: 100,
		},
		{
			name:    "string length simulation (using integers)",
			nums:    []int{1, 2, 3},
			initial: 0,
			operation: func(acc, curr int) int {
				return acc + curr*curr
			},
			expected: 14, // 1^2 + 2^2 + 3^2 = 1+4+9=14
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reduce(tt.nums, tt.initial, tt.operation)
			if result != tt.expected {
				t.Errorf("Reduce() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	tests := []struct {
		name     string
		f        func(int) int
		g        func(int) int
		input    int
		expected int
	}{
		{
			name: "double then square",
			f: func(x int) int {
				return x * x
			},
			g: func(x int) int {
				return x * 2
			},
			input:    5,
			expected: 100, // (5*2)^2 = 10^2 = 100
		},
		{
			name: "square then double",
			f: func(x int) int {
				return x * 2
			},
			g: func(x int) int {
				return x * x
			},
			input:    5,
			expected: 50, // (5^2)*2 = 25*2 = 50
		},
		{
			name: "add 3 then multiply by 4",
			f: func(x int) int {
				return x * 4
			},
			g: func(x int) int {
				return x + 3
			},
			input:    2,
			expected: 20, // (2+3)*4 = 5*4 = 20
		},
		{
			name: "negate then absolute value",
			f: func(x int) int {
				if x < 0 {
					return -x
				}
				return x
			},
			g: func(x int) int {
				return -x
			},
			input:    5,
			expected: 5, // abs(-5) = 5
		},
		{
			name: "identity functions",
			f: func(x int) int {
				return x
			},
			g: func(x int) int {
				return x
			},
			input:    42,
			expected: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composed := Compose(tt.f, tt.g)
			result := composed(tt.input)
			if result != tt.expected {
				t.Errorf("Compose()(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}

	// Test that composition order is correct (f(g(x)))
	t.Run("composition order", func(t *testing.T) {
		// f: x+1, g: x*2
		// f(g(x)) = (x*2)+1
		f := func(x int) int { return x + 1 }
		g := func(x int) int { return x * 2 }
		composed := Compose(f, g)
		result := composed(3) // should be (3*2)+1 = 7
		if result != 7 {
			t.Errorf("Compose order test: got %v, want 7", result)
		}
	})
}