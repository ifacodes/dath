package dath

import "math"

// LAB struct contains the converted values from a Color
type RGB struct {
	R, G, B int
}

// RGB takes a Color and returns a RGB struct
func (c *Color) RGB() *RGB {
	return &RGB{
		R: int(c.r * 255.0),
		G: int(c.g * 255.0),
		B: int(c.b * 255.0),
	}
}

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

func calculateHue(max, chroma, r, g, b float64) (h float64) {
	if chroma == 0.0 {
		h = 0.0
	} else if max == r {
		if g < b {
			h = 6.0
		}
		h += (g - b) / chroma

	} else if max == g {
		h = 2.0 + ((b - r) / chroma)
	} else if max == b {
		h = 4.0 + ((r - g) / chroma)
	}
	h *= 60.0
	return
}

func rgb2hsv(r, g, b float64) (h, s, v float64) {
	v = math.Max(r, math.Max(g, b))
	chroma := v - math.Min(r, math.Min(g, b))
	h = calculateHue(v, chroma, r, g, b)
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
	h = calculateHue(max, chroma, r, g, b)
	if l == 0 || l == 1 {
		s = 0
	} else {
		s = ((max - l) / math.Min(l, 1-l))
	}
	return
}
