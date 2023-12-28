package ee

func EvalLazy[T any](condition bool, ifTrue, ifFalse func() T) T {
	if condition {
		return ifTrue()
	} else {
		return ifFalse()
	}
}

func EvalEager[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	} else {
		return ifFalse
	}
}
