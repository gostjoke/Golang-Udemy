package pointer

import (
	"math"
)

func Power2(a *float64) {
	*a = math.Pow(*a, 2)
}

func Power3(b *float64) {
	*b = math.Pow(*b, 3)
}
