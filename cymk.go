package dath

import (
	d "github.com/shopspring/decimal"
)

var (
	o = d.New(1, 0)
)

// CYMK struct contains the converted values from a Color
type CMYK struct {
	C, M, Y, K d.Decimal
}

// CMYK takes a Color and returns a CYMK struct
func (cc *Color) CMYK() *CMYK {
	r := &CMYK{}
	r.C, r.M, r.Y, r.K = rgb2cmyk(cc.r, cc.g, cc.b)
	return r
}

func cmyk2rgb(c, m, y, k d.Decimal) (r, g, b d.Decimal) {
	r = o.Sub(c).Mul(o.Sub(k))
	g = o.Sub(m).Mul(o.Sub(k))
	b = o.Sub(y).Mul(o.Sub(k))
	return
}

func rgb2cmyk(r, g, b d.Decimal) (c, m, y, k d.Decimal) {
	k = o.Sub(d.Max(r, d.Max(g, b)))
	c = o.Sub(r.Sub(k)).Div(o.Sub(k))
	m = o.Sub(g.Sub(k)).Div(o.Sub(k))
	y = o.Sub(b.Sub(k)).Div(o.Sub(k))
	return
}
