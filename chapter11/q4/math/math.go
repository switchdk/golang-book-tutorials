package math

// Average of a slice of float64
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Max Maximum of a slice of float64
func Max(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	max := xs[0]
	for _, i := range xs {
		if i > max {
			max = i
		}
	}
	return max
}

// Min Minimum value of a slice of float64
func Min(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	min := xs[0]
	for _, i := range xs {
		if i < min {
			min = i
		}
	}
	return min
}
