package dath

import (
	d "github.com/shopspring/decimal"
)

// HSL struct contains the converted values from a Color
type HSL struct {
	H, S, L d.Decimal
}

// HSL takes a Color and returns a HSL struct
func (c *Color) HSL() *HSL {
	hsl := &HSL{}
	hsl.H, hsl.S, hsl.L = rgb2hsl(c.r, c.g, c.b)
	return hsl
}

func hsl2rgb(h, s, l d.Decimal) (r, g, b d.Decimal) {
	f := func(n d.Decimal) d.Decimal {
		a := s.Mul(d.Min(l, d.NewFromFloat(1.0).Sub(l)))
		k := n.Add(h.Div(d.NewFromFloat(30.0))).Mod(d.NewFromFloat(12.0))
		return l.Sub(a.Mul(d.Max(d.NewFromFloat(-1.0), d.Min(d.NewFromFloat(1.0), k.Sub(d.NewFromFloat(3.0)), d.NewFromFloat(9.0).Sub(k)))))
	}
	r = f(d.NewFromFloat(0.0)).Round(2)
	g = f(d.NewFromFloat(8.0)).Round(2)
	b = f(d.NewFromFloat(4.0)).Round(2)
	return
}
