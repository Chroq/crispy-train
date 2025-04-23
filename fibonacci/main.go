package main

import (
	"fmt"
	"math/big"
	"time"
)

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func FibonacciLoop(n int) int {
	if n <= 1 {
		return n
	}

	var n1, n2 = 0, 1
	for i := 2; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}

func FibonacciDynamic(n int) int {
	if n <= 1 {
		return n
	}
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

var cache = make(map[int]*big.Int, 1000)

func FibonacciMemo(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	if val, exists := cache[n]; exists {
		return val
	}

	cache[n] = new(big.Int).Add(FibonacciMemo(n-1), FibonacciMemo(n-2))
	return cache[n]
}

func main() {
	//start := time.Now()
	//i:= 1000
	//fmt.Printf("%d :  %d (%d ms) \n", i, FibonacciMemo(i), time.Since(start).Milliseconds())

	for i := 0; i < 10000; i += 1 {
		start := time.Now()
		fmt.Printf("%d :  %d (%d ms) \n", i, FibonacciMemo(i), time.Since(start).Milliseconds())
	}

	/* n := 40 */
	/* println("FibonacciRecursive:", FibonacciRecursive(n)) */
	/* println("FibonacciMemo:", FibonacciMemo(n)) */
	/* println("FibonacciLoop:", FibonacciLoop(n)) */
	/* println("FibonacciDynamic:", FibonacciDynamic(n)) */
}
