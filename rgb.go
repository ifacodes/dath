package dath

import (
	"math"

	d "github.com/shopspring/decimal"
)

// LAB struct contains the converted values from a Color
type RGB struct {
	R, G, B d.Decimal
}

// RGB takes a Color and returns a RGB struct
func (c *Color) RGB() *RGB {
	return &RGB{
		R: c.r.Mul(d.NewFromFloat(255.0)).Round(2),
		G: c.g.Mul(d.NewFromFloat(255.0)).Round(2),
		B: c.b.Mul(d.NewFromFloat(255.0)).Round(2),
	}
}

func rgb2xyz(r, g, b d.Decimal) (x, y, z float64) {
	rf, _ := r.Float64()
	gf, _ := g.Float64()
	bf, _ := b.Float64()
	if rf > 0.04045 {
		rf = math.Pow((rf+0.055)/1.055, 2.4)
	} else {
		rf = rf / 12.92
	}
	if gf > 0.04045 {
		gf = math.Pow((gf+0.055)/1.055, 2.4)
	} else {
		gf = gf / 12.92
	}
	if bf > 0.04045 {
		bf = math.Pow((bf+0.055)/1.055, 2.4)
	} else {
		bf = bf / 12.92
	}
	x = (rf * sRGB[0][0]) + (gf * sRGB[0][1]) + (bf * sRGB[0][2])
	y = (rf * sRGB[1][0]) + (gf * sRGB[1][1]) + (bf * sRGB[1][2])
	z = (rf * sRGB[2][0]) + (gf * sRGB[2][1]) + (bf * sRGB[2][2])
	return
}

func calculateHue(max, chroma, r, g, b d.Decimal) (h d.Decimal) {
	if chroma.Equal(d.NewFromFloat(0.0)) {
		h = chroma
	} else if max.Equal(r) {
		if g.LessThan(b) {
			h = d.NewFromFloat(6.0)
		}
		h = h.Add(g.Sub(b).Div(chroma))

	} else if max.Equal(g) {
		h = d.NewFromFloat(2.0).Add(b.Sub(r).Div(chroma))
	} else if max == b {
		h = d.NewFromFloat(4.0).Add(r.Sub(g).Div(chroma))
	}
	h = h.Mul(d.NewFromFloat(60.0))
	return
}

func rgb2hsv(r, g, b d.Decimal) (h, s, v d.Decimal) {
	v = d.Max(r, g, b)
	chroma := v.Sub(d.Min(r, g, b))
	h = calculateHue(v, chroma, r, g, b)
	if !v.Equal(d.New(0, 0)) {
		s = chroma.Div(v)
	} else {
		s = d.New(0, 0)
	}
	return
}
func rgb2hsl(r, g, b d.Decimal) (h, s, l d.Decimal) {
	max := d.Max(r, g, b)
	chroma := max.Sub(d.Min(r, g, b))
	l = max.Sub(chroma.Div(d.NewFromFloat(2.0)))
	h = calculateHue(max, chroma, r, g, b)
	if l.Equal(d.New(0, 0)) || l.Equal(d.New(1, 0)) {
		s = d.New(0, 0)
	} else {
		s = max.Sub(l).Div(d.Min(l, d.New(1, 0).Sub(l)))
	}
	return
}
