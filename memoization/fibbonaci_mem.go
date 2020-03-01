package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin *bufio.Reader

type options uint

const (
	fibonnaci = iota + 1
	fibonnaciMem
	fibonnaciMem2
	end
)

type fibFunc func(uint) int

func init() {

	stdin = bufio.NewReader(os.Stdin)
}

func menu() options {
	option := 0
	for option < 1 || option > 4 {
		fmt.Println("1. Fibonnaci")
		fmt.Println("2. Fibonnaci Mem")
		fmt.Println("3. Fibonnaci Mem")
		fmt.Println("4. Exit")
		fmt.Printf("\nChoose an option....:")
		if _, err := fmt.Fscanf(stdin, "%d", &option); err != nil {
			// In case of not introducing a number
			option = 0
		}
		stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left

	}
	return options(option)
}

func main() {
Loop:
	for {
		option := menu()
		switch option {
		case fibonnaci:
			var n uint
			fmt.Printf("\nEnter a number for fibonnaci:")
			if _, err := fmt.Fscanf(stdin, "%d", &n); err != nil {
			}
			stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left
			fmt.Printf("\nThe fibonnaci of %d is: %d\n", n, fib(n))
		case fibonnaciMem:
			var n uint
			fmt.Printf("\nEnter a number for fibonnaci:")
			if _, err := fmt.Fscanf(stdin, "%d", &n); err != nil {
			}
			stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left
			fmt.Printf("\nThe fibonnaci with Mem of %d is: %d\n", n, fibMemoize(n))
		case fibonnaciMem2:
			var n uint
			fmt.Printf("\nEnter a number for fibonnaci:")
			if _, err := fmt.Fscanf(stdin, "%d", &n); err != nil {
			}
			stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left
			fmt.Printf("\nThe fibonnaci with Mem of %d is: %d\n", n, fibMemoize2(n))
		default:
			break Loop
		}
	}
	os.Exit(0)
}

// Normal fibonnaci
func fib(n uint) int {
	if n < 2 {
		return int(n)
	}
	return fib(n-1) + fib(n-2)
}

func memoize(f fibFunc) fibFunc {
	cache := map[uint]int{
		0: 0,
		1: 1,
	}
	return func(n uint) int {
		if _, ok := cache[n]; !ok {
			cache[n] = f(n)
		}
		return cache[n]
	}
}

func fibMemoize(n uint) int {
	var f fibFunc
	// The code within memoize is the same as fib, but the recursion calls to f
	// that uses the memoized version
	f = memoize(func(n uint) int {
		if n < 2 {
			return int(n)
		}
		return f(n-1) + f(n-2)
	})

	return f(n)
}

func fibMemoize2(n uint) int {
	cache := make(map[uint]int)
	var f fibFunc
	f = func(n uint) int {
		if n < 2 {
			cache[n] = int(n)
		}
		if _, ok := cache[n]; !ok {
			cache[n] = f(n-1) + f(n-2)
		}
		return cache[n]
	}
	return f(n)
}
