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