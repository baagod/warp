package warp

func Clamp[T int | int64 | float64](value, min, max T) int {
	if value < min {
		return int(min)
	} else if value > max {
		return int(max)
	}
	return int(value)
}
