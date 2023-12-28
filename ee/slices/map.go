package slices

func Map[T, S any](in []T, f func(T) S) []S {
	out := make([]S, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}
