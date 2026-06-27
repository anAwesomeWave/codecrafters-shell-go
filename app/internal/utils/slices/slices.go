package slices

func Map[S ~[]F, F, T any](s S, f func(F) T) []T {
	ret := make([]T, 0, len(s))

	for _, elem := range s {
		ret = append(ret, f(elem))
	}

	return ret
}

func MapE[S ~[]F, F, T any](s S, f func(F) (T, error)) ([]T, error) {
	ret := make([]T, 0, len(s))

	for _, elem := range s {
		mapped, err := f(elem)
		if err != nil {
			return nil, err
		}

		ret = append(ret, mapped)
	}

	return ret, nil
}

func Filter[S ~[]F, F any](s S, f func(F) bool) S {
	ret := make(S, 0, len(s))

	for _, elem := range s {
		if f(elem) {
			ret = append(ret, elem)
		}
	}

	return ret
}

func FilterE[S ~[]F, F any](s S, f func(F) (bool, error)) (S, error) {
	ret := make(S, 0, len(s))

	for _, elem := range s {
		ok, err := f(elem)
		if err != nil {
			return nil, err
		}

		if ok {
			ret = append(ret, elem)
		}
	}

	return ret, nil
}
