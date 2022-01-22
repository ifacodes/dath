package dath

import (
	d "github.com/shopspring/decimal"
)

// HSV struct contains the converted values from a Color
type HSV struct {
	H, S, V d.Decimal
}

// HSV takes a Color and returns a HSV struct
func (c *Color) HSV() *HSV {
	h := &HSV{}
	h.H, h.S, h.V = rgb2hsv(c.r, c.g, c.b)
	return h
}

func hsv2rgb(h, s, v d.Decimal) (r, g, b d.Decimal) {
	f := func(n d.Decimal) d.Decimal {
		k := n.Add(h.Div(d.NewFromFloat(60.0))).Mod(d.NewFromFloat(6.0))
		return v.Sub(v.Mul(s).Mul(d.Max(d.NewFromFloat(0.0), d.Min(d.NewFromFloat(1.0), d.NewFromFloat(4.0).Sub(k), k))))
	}
	r = f(d.NewFromFloat(5.0)).Round(2)
	g = f(d.NewFromFloat(3)).Round(2)
	b = f(d.NewFromFloat(1)).Round(2)
	return
}
