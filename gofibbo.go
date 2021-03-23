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
