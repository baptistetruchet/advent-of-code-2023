package utils

func ChanToSlice[C any](c <-chan C) []C {
	s := make([]C, 0)
	for i := range c {
		s = append(s, i)
	}
	return s
}
