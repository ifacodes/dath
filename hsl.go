package dath

import "math"

func hsl2rgb(h, s, l float64) (r, g, b float64) {
	f := func(n float64) float64 {
		a := s * math.Min(l, 1.0-l)
		k := math.Mod(n+(h/30.0), 12.0)
		return l - a*math.Max(-1.0, math.Min(1.0, math.Min(k-3.0, 9.0-k)))
	}
	r = math.Round((f(0) * 100)) / 100
	g = math.Round((f(8) * 100)) / 100
	b = math.Round((f(4) * 100)) / 100
	return
}