// Package dath provides types and method for
// handling color data in a variety of color spaces
package dath

// BUG(ifamakes): Conversion can be off when going back and forth due to the unreliablity of floats.

//"github.com/shopspring/decimal"

type ColorOption func(c *Color)

// Color contains the base color values used for converting between spaces
type Color struct {
	r, g, b float64
}

// NewColor returns a new Color from the provided color values.
// It's recommended to use this funciton to create your colors in case implementation changes.
func NewColor() *Color {
	color := &Color{}
	return color
}

func (c *Color) clip() {
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

// Takes RGB values and return a Color
func (c *Color) FromRGB(r, g, b int) *Color {
	c.r = float64(r) / 255.0
	c.g = float64(g) / 255.0
	c.b = float64(b) / 255.0
	c.clip()
	return c
}

// Takes CMYK values and returns a Color
func (cc *Color) FromCMYK(c, m, y, k float64) *Color {
	cc.r, cc.g, cc.b = cmyk2rgb(c, m, y, k)
	cc.clip()
	return cc
}

// Takes HSL values and returns a Color
func (c *Color) FromHSL(h, s, l float64) *Color {
	c.r, c.g, c.b = hsl2rgb(h, s, l)
	c.clip()
	return c
}

// Takes HSV values and returns a Color
func (c *Color) FromHSV(h, s, v float64) *Color {
	c.r, c.g, c.b = hsv2rgb(h, s, v)
	c.clip()
	return c
}

// Takes LAB values and returns a Color
func (c *Color) FromLAB(l, a, b float64) *Color {
	x, y, z := lab2xyz(l, a, b)
	c.r, c.g, c.b = xyz2rgb(x, y, z)
	c.clip()
	return c
}

// Takes LUV values and returns a Color
func (c *Color) FromLUV(l, u, v float64) *Color {
	x, y, z := luv2xyz(l, u, v)
	c.r, c.g, c.b = xyz2rgb(x, y, z)
	c.clip()
	return c
}

// Takes HCL values and returns a Color
func (cc *Color) FromHCL(h, c, l float64) *Color {
	cc.r, cc.g, cc.b = hcl2rgb(h, c, l)
	cc.clip()
	return cc
}
