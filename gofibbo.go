package gofibbo

func Fibbo(n int64) int64 {
	if n < 3 {
		return 1
	}
	return Fibbo(n-1) + Fibbo(n-2)
}
