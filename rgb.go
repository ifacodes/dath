package dath

import "math"

func rgb2xyz(r, g, b float64) (x, y, z float64) {
	if r > 0.04045 {
		r = math.Pow((r+0.055)/1.055, 2.4)
	} else {
		r = r / 12.92
	}
	if g > 0.04045 {
		g = math.Pow((g+0.055)/1.055, 2.4)
	} else {
		g = g / 12.92
	}
	if b > 0.04045 {
		b = math.Pow((b+0.055)/1.055, 2.4)
	} else {
		b = b / 12.92
	}
	x = (r * sRGB[0][0]) + (g * sRGB[0][1]) + (b * sRGB[0][2])
	y = (r * sRGB[1][0]) + (g * sRGB[1][1]) + (b * sRGB[1][2])
	z = (r * sRGB[2][0]) + (g * sRGB[2][1]) + (b * sRGB[2][2])
	return
}

func (c *color) HSV() *hsv {
	h := &hsv{}
	h.H, h.S, h.V = rgb2hsv(c.r, c.g, c.b)
	return h
}

func (c *color) HSL() (h, s, l float64) {
	return rgb2hsl(c.r, c.g, c.b)
}

func rgb2hsv(r, g, b float64) (h, s, v float64) {
	v = math.Max(r, math.Max(g, b))
	chroma := v - math.Min(r, math.Min(g, b))
	h = hslCalculateHue(v, chroma, r, g, b)
	if v != 0 {
		s = chroma / v
	} else {
		s = 0
	}
	return
}
func rgb2hsl(r, g, b float64) (h, s, l float64) {
	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	chroma := max - min
	l = max - (chroma / 2.0)
	h = hslCalculateHue(max, chroma, r, g, b)
	if l == 0 || l == 1 {
		s = 0
	} else {
		s = ((max - l) / math.Min(l, 1-l))
	}
	return
}
