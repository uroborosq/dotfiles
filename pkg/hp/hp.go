package hp

func F[T, E any](t T, _ E) T {
	return t
}

func P[T any](t T, e error) T {
	if e != nil {
		panic(e)
	}

	return t
}
