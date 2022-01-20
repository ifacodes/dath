// Package dath provides types and method for
// handling color data in a variety of color spaces
package dath

import "log"

// BUG(ifamakes): Conversion can be off when going back and forth due to the unreliablity of floats.

//"github.com/shopspring/decimal"

type ColorOption func(c *color)

type color struct {
	r, g, b float64
}

func NewColor(c ...ColorOption) *color {
	color := &color{}
	for _, opt := range c {
		opt(color)
	}
	return color
}

func FromRGB(r, g, b int) ColorOption {
	return func(c *color) {
		c.r = float64(r) / 255.0
		c.g = float64(g) / 255.0
		c.b = float64(b) / 255.0
	}
}

func FromCMYK(c, m, y, k float64) ColorOption {
	return func(cc *color) {
		cc.r, cc.g, cc.b = cmyk2rgb(c, m, y, k)
	}
}

func FromHSL(h, s, l float64) ColorOption {
	return func(c *color) {
		c.r, c.g, c.b = hsl2rgb(h, s, l)
	}
}

func FromHSV(h, s, v float64) ColorOption {
	return func(c *color) {
		c.r, c.g, c.b = hsv2rgb(h, s, v)
	}
}

func FromLUV(l, u, v float64) ColorOption {
	return func(c *color) {
		x, y, z := luv2xyz(l, u, v)
		log.Println(x, y, z)
		c.r, c.g, c.b = xyz2rgb(x, y, z)
	}
}

func FromHCL(h, c, l float64) ColorOption {
	return func(cc *color) {
		cc.r, cc.g, cc.b = hcl2rgb(h, c, l)
	}
}
