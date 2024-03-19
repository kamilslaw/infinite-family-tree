package utils

func MapValues[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func PickAnyKey[M ~map[K]V, K comparable, V any](m M) K {
	for k := range m {
		return k
	}
	panic("map is empty")
}
