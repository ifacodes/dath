package dath

import "math"

func hsv2rgb(h, s, v float64) (r, g, b float64) {
	f := func(n float64) float64 {
		k := math.Mod(n+(h/60.0), 6.0)
		return v - v*s*math.Max(0.0, math.Min(1.0, math.Min(4.0-k, k)))
	}
	r = f(5)
	g = f(3)
	b = f(1)
	return
}
