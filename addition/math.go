package addition

type Number interface {
	int64 | float64
}

func Add[T Number](a, b T) T {
	return a + b
}
