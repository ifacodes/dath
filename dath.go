// Package dath provides types and method for
// handling color data in a variety of color spaces
package dath

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
	color.clip()
	return color
}

func (c *color) clip() {
	if c.r < 0.0 {
		c.r = 0.0
	} else if c.r > 255.0 {
		c.r = 255.0
	}
	if c.g < 0.0 {
		c.g = 0.0
	} else if c.g > 255.0 {
		c.g = 255.0
	}
	if c.b < 0.0 {
		c.b = 0.0
	} else if c.b > 255.0 {
		c.b = 255.0
	}
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

func FromLAB(l, a, b float64) ColorOption {
	return func(c *color) {
		x, y, z := lab2xyz(l, a, b)
		c.r, c.g, c.b = xyz2rgb(x, y, z)
	}
}

func FromLUV(l, u, v float64) ColorOption {
	return func(c *color) {
		x, y, z := luv2xyz(l, u, v)
		c.r, c.g, c.b = xyz2rgb(x, y, z)
	}
}

func FromHCL(h, c, l float64) ColorOption {
	return func(cc *color) {
		cc.r, cc.g, cc.b = hcl2rgb(h, c, l)
	}
}
