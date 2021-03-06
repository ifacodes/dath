package dath

import (
	"math"

	d "github.com/shopspring/decimal"
)

func xyz2rgb(x, y, z float64) (r, g, b d.Decimal) {
	rf := (x * sRGBInv[0][0]) + (y * sRGBInv[0][1]) + (z * sRGBInv[0][2])
	gf := (x * sRGBInv[1][0]) + (y * sRGBInv[1][1]) + (z * sRGBInv[1][2])
	bf := (x * sRGBInv[2][0]) + (y * sRGBInv[2][1]) + (z * sRGBInv[2][2])
	if rf <= 0.0031308 {
		rf = 12.92 * rf
	} else {
		rf = 1.055*math.Pow(rf, (1.0/2.4)) - 0.055
	}
	if gf <= 0.0031308 {
		gf = 12.92 * gf
	} else {
		gf = 1.055*math.Pow(gf, (1.0/2.4)) - 0.055
	}
	if bf <= 0.0031308 {
		bf = 12.92 * bf
	} else {
		bf = 1.055*math.Pow(bf, (1.0/2.4)) - 0.055
	}
	if math.IsNaN(rf) {
		rf = 0.0
	}
	if math.IsNaN(gf) {
		gf = 0.0
	}
	if math.IsNaN(bf) {
		bf = 0.0
	}
	r = d.NewFromFloatWithExponent(rf, -2)
	g = d.NewFromFloatWithExponent(gf, -2)
	b = d.NewFromFloatWithExponent(bf, -2)
	return
}

func xyz2luv(x, y, z float64) (l, u, v float64) {
	yr := y / d65[1]
	if yr > 0.008856 {
		l = (116 * math.Cbrt(yr)) - 16
	} else {
		l = 903.3 * yr
	}
	uu := 4 * x / (x + 15*y + 3*z)
	vv := 9 * y / (x + 15*y + 3*z)
	ur := 4 * d65[0] / (d65[0] + 15*d65[1] + 3*d65[2])
	vr := 9 * d65[1] / (d65[0] + 15*d65[1] + 3*d65[2])
	u = 13 * l * (uu - ur)
	v = 13 * l * (vv - vr)
	return
}

func xyz2lab(x, y, z float64) (l, a, b float64) {
	xr := x / d65[0]
	yr := y / d65[1]
	zr := z / d65[2]
	f := func(x float64) float64 {
		if x > 0.00856 {
			return math.Cbrt(x)
		} else {
			return (903.3*x + 16) / 116
		}
	}
	fy := f(yr)
	l = 116*fy - 16
	a = 500 * (f(xr) - fy)
	b = 200 * (fy - f(zr))
	return
}
