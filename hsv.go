package dath

import "math"

// HSV struct contains the converted values from a Color
type HSV struct {
	H, S, V float64
}

// HSV takes a Color and returns a HSV struct
func (c *Color) HSV() *HSV {
	h := &HSV{}
	h.H, h.S, h.V = rgb2hsv(c.r, c.g, c.b)
	return h
}

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
