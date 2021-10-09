package main

// 剑指 Offer 10- I. 斐波那契数列

//func fib(n int) int {
//	if n == 1 {
//		return 1
//	} else if n == 0 {
//		return 0
//	}
//
//	return fib(n-1) + fib(n-2)
//}

func fib(n int) int {
	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	}

	a := 0
	b := 1
	c := 0
	for i := 1; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c % 1000000007
}

func numWays(n int) int {
	switch n {
	case 2:
		return 2
	case 1:
		return 1
	case 0:
		return 1
	}

	return numWays(n-1) + numWays(n-2)
}
