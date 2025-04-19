package multiplication

type Number interface {
	int64 | float64
}

func Product[T Number](a, b T) T {
	return a * b
}
