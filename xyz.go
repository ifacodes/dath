package dath

import "math"

type XYZ struct {
	X, Y, Z float64
}

func xyz2rgb(x, y, z float64) (r, g, b float64) {
	r = (x * sRGBInv[0][0]) + (y * sRGBInv[0][1]) + (z * sRGBInv[0][2])
	g = (x * sRGBInv[1][0]) + (y * sRGBInv[1][1]) + (z * sRGBInv[1][2])
	b = (x * sRGBInv[2][0]) + (y * sRGBInv[2][1]) + (z * sRGBInv[2][2])
	if r <= 0.0031308 {
		r = 12.92 * r
	} else {
		r = 1.055*math.Pow(r, (1.0/2.4)) - 0.055
	}
	if g <= 0.0031308 {
		g = 12.92 * g
	} else {
		g = 1.055*math.Pow(g, (1.0/2.4)) - 0.055
	}
	if b <= 0.0031308 {
		b = 12.92 * b
	} else {
		b = 1.055*math.Pow(b, (1.0/2.4)) - 0.055
	}
	r = math.Round((r * 100)) / 100
	g = math.Round((g * 100)) / 100
	b = math.Round((b * 100)) / 100
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
