package main

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

func main() {
	n := 40
	println("FibonacciRecursive:", FibonacciRecursive(n))
	/* println("FibonacciLoop:", FibonacciLoop(n)) */
	/* println("FibonacciDynamic:", FibonacciDynamic(n)) */
}
