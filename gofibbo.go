package gofibbo

var memo [100]int64

func Fibbo(n int64) int64 {
	if n < 3 {
		return 1
	}
	if memo[n] == 0 {
		memo[n] = Fibbo(n-1) + Fibbo(n-2)
	}
	return memo[n]
}

var digit [100000]int

func LastDigit(n int) int {
	if n < 3 {
		return 1
	}
	if digit[n] == 0 {
		digit[n] = LastDigit(n-1) + LastDigit(n-2)
	}
	return digit[n]
}
