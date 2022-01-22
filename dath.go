// Package dath provides types and method for
// handling color data in a variety of color spaces
package dath

// BUG(ifamakes): Conversion can be off when going back and forth due to the unreliablity of floats.

import (
	"fmt"

	d "github.com/shopspring/decimal"
)

// Color contains the base color values used for converting between spaces
type Color struct {
	r, g, b d.Decimal
}

// NewColor returns a new Color from the provided color values.
// It's recommended to use this funciton to create your colors in case implementation changes.
func NewColor() *Color {
	color := &Color{}
	return color
}

func (c *Color) String() string {
	return fmt.Sprintf("{%s %s %s}", c.r.String(), c.g.String(), c.b.String())

}

func (c *Color) clip() {
	if c.r.LessThan(d.New(0, 0)) {
		c.r = d.New(0, 0)
	} else if c.r.GreaterThan(d.New(255, 0)) {
		c.r = d.New(255, 0)
	}
	if c.g.LessThan(d.New(0, 0)) {
		c.g = d.New(0, 0)
	} else if c.g.GreaterThan(d.New(255, 0)) {
		c.g = d.New(255, 0)
	}
	if c.b.LessThan(d.New(0, 0)) {
		c.b = d.New(0, 0)
	} else if c.b.GreaterThan(d.New(255, 0)) {
		c.b = d.New(255, 0)
	}
}

// Takes RGB values and return a Color
func (c *Color) FromRGB(r, g, b int) *Color {
	c.r = d.NewFromFloatWithExponent(float64(r)/255.0, -2)
	c.g = d.NewFromFloatWithExponent(float64(g)/255.0, -2)
	c.b = d.NewFromFloatWithExponent(float64(b)/255.0, -2)
	c.clip()
	return c
}

// Takes CMYK values and returns a Color
func (cc *Color) FromCMYK(c, m, y, k float64) *Color {
	cc.r, cc.g, cc.b = cmyk2rgb(d.NewFromFloat(c), d.NewFromFloat(m), d.NewFromFloat(y), d.NewFromFloat(k))
	cc.clip()
	return cc
}

// Takes HSL values and returns a Color
func (c *Color) FromHSL(h, s, l float64) *Color {
	c.r, c.g, c.b = hsl2rgb(d.NewFromFloat(h), d.NewFromFloat(s), d.NewFromFloat(l))
	c.clip()
	return c
}

// Takes HSV values and returns a Color
func (c *Color) FromHSV(h, s, v float64) *Color {
	c.r, c.g, c.b = hsv2rgb(d.NewFromFloat(h), d.NewFromFloat(s), d.NewFromFloat(v))
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

/* // Takes HCL values and returns a Color
func (cc *Color) FromHCL(h, c, l float64) *Color {
	cc.r, cc.g, cc.b = hcl2rgb(d.NewFromFloat(h), d.NewFromFloat(c), d.NewFromFloat(l))
	cc.clip()
	return cc
} */
