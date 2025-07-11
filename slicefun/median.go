package slicefun

import "sort"

func Median(x []float64) float64 {
	if len(x) == 0 {
		return 0.0
	}
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i]
	}
	return (x[i-1] + x[i]) / 2
}
