package util

func EqualSet[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	setA := make(map[T]struct{}, len(a))
	for _, item := range a {
		setA[item] = struct{}{}
	}

	for _, item := range b {
		if _, exists := setA[item]; !exists {
			return false
		}
	}

	return true
}
