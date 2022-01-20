package dath

import "math"

type CMYK struct {
	C, M, Y, K float64
}

func (cc *color) CMYK() *CMYK {
	r := &CMYK{}
	r.C, r.M, r.Y, r.K = rgb2cmyk(cc.r, cc.g, cc.b)
	return r
}

func cmyk2rgb(c, m, y, k float64) (r, g, b float64) {
	r = math.Round((1-c)*(1-k)*10) / 10
	g = math.Round((1-m)*(1-k)*10) / 10
	b = math.Round((1-y)*(1-k)*10) / 10
	return
}

func rgb2cmyk(r, g, b float64) (c, m, y, k float64) {
	k = 1 - math.Max(r, math.Max(g, b))
	c = (1 - r - k) / (1 - k)
	m = (1 - g - k) / (1 - k)
	y = (1 - b - k) / (1 - k)
	return
}
