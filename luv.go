package dath

import (
	"math"

	d "github.com/shopspring/decimal"
)

// LUV struct contains the converted values from a Color
type LUV struct {
	L, U, V d.Decimal
}

// LUV takes a Color and returns a LUV struct
func (c *Color) LUV() *LUV {
	luv := &LUV{}
	x, y, z := rgb2xyz(c.r, c.g, c.b)
	l, u, v := xyz2luv(x, y, z)
	if math.IsNaN(l) {
		l = 0.0
	}
	if math.IsNaN(u) {
		u = 0.0
	}
	if math.IsNaN(v) {
		v = 0.0
	}
	luv.L, luv.U, luv.V = d.NewFromFloatWithExponent(l, -2), d.NewFromFloatWithExponent(u, -2), d.NewFromFloatWithExponent(v, -2)
	return luv
}

func luv2xyz(l, u, v float64) (x, y, z float64) {
	f := func(a, b, c, d float64) float64 {
		return (d - b) / (a - c)
	}
	if l > (903.3 * 0.008856) {
		y = math.Pow((l+16.0)/116.0, 3.0)
	} else {
		y = l / 903.3
	}
	u0 := 4.0 * d65[0] / (d65[0] + 15.0*d65[1] + 3.0*d65[2])
	v0 := 9.0 * d65[1] / (d65[0] + 15.0*d65[1] + 3.0*d65[2])
	d := y * ((39.0 * l / (v + 13*l*v0)) - 5.0)
	a := (1.0 / 3.0) * (((52.0 * l) / (u + 13*l*u0)) - 1.0)
	b := -5.0 * y
	c := -(1.0 / 3.0)
	x = f(a, b, c, d)
	z = (x * a) + b
	return
}
