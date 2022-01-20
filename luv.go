package dath

import "math"

func luv2xyz(l, u, v float64) (x, y, z float64) {
	f := func(a, b, c, d float64) float64 {
		return (d - b) / (a - c)
	}
	if l > (903.3 * 0.008856) {
		y = math.Pow((l+16)/116, 3)
	} else {
		y = l / 903.3
	}
	u0 := 4.0 * d65[0] / (d65[0] + 15.0*d65[1] + 3.0*d65[2])
	v0 := 9.0 * d65[1] / (d65[0] + 15.0*d65[1] + 3.0*d65[2])
	d := y * ((39.0 * l / (v + 13*l*v0)) - 5)
	a := (1.0 / 3.0) * ((52.0 * l / (u + 13*l*u0)) - 1)
	b := -5.0 * y
	x = f(a, b, -(1.0 / 3.0), d)
	z = (x * a) + b
	return
}
