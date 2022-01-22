package dath

/* // HCL struct contains the converted values from a Color
type HCL struct {
	H, C, L d.Decimal
}

// HCL takes a Color and returns a HCL struct
func (cc *Color) HCL() *HCL {
	hcl := &HCL{}
	hcl.H, hcl.C, hcl.L = rgb2hcl(cc.r, cc.g, cc.b)
	return hcl
}

func rgb2hcl(r, g, b d.Decimal) (h, c, l d.Decimal) {
	x, y, z := rgb2xyz(r, g, b)
	l, u, v := xyz2luv(x, y, z)
	c = math.Sqrt(math.Pow(u, 2) + math.Pow(v, 2))
	h = math.Atan2(v, u) * (180 / math.Pi)
	if h < 0 {
		h += 360.0
	}
	return
}

func hcl2rgb(h, c, l d.Decimal) (r, g, b d.Decimal) {
	x, y, z := luv2xyz(l, c*math.Cos(h*(math.Pi/180.0)), c*math.Sin(h*(math.Pi/180.0)))
	r, g, b = xyz2rgb(x, y, z)
	return
}

func (cc *Color) LCH() (l, c, h d.Decimal) {
	hcl := cc.HCL()
	return hcl.L, hcl.C, hcl.H
} */
