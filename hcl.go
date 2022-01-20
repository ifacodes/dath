package dath

import "math"

func (cc *color) HCL() (h, c, l float64) {
	return rgb2hcl(cc.r, cc.g, cc.b)
}

func rgb2hcl(r, g, b float64) (h, c, l float64) {
	x, y, z := rgb2xyz(r, g, b)
	l, u, v := xyz2luv(x, y, z)
	c = math.Sqrt(math.Pow(u, 2) + math.Pow(v, 2))
	h = math.Atan2(v, u) * (180 / math.Pi)
	if h < 0 {
		h += 360.0
	}
	return
}

func hcl2rgb(h, c, l float64) (r, g, b float64) {
	x, y, z := luv2xyz(l, c*math.Cos(h*(math.Pi/180.0)), c*math.Sin(h*(math.Pi/180.0)))
	r, g, b = xyz2rgb(x, y, z)
	return
}

func (cc *color) LCH() (l, c, h float64) {
	h, c, l = cc.HCL()
	return l, c, h
}
