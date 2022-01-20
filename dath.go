package dath

//"github.com/shopspring/decimal"

type ColorOption func(c *color)

func FromRGB(r, g, b int) ColorOption {
	return func(c *color) {
		c.r = float64(r) / 255.0
		c.g = float64(g) / 255.0
		c.b = float64(b) / 255.0
	}
}

func FromHSL(h, s, l float64) ColorOption {
	return func(c *color) {
		c.r, c.g, c.b = hsl2rgb(h, s, l)
	}
}

func FromHSV(h, s, v float64) ColorOption {
	return func(c *color) {
		c.r, c.g, c.b = hsv2rgb(h, s, v)
	}
}

func FromLUV(l, u, v float64) ColorOption {
	return func(c *color) {
		x, y, z := luv2xyz(l, u, v)
		c.r, c.g, c.b = xyz2rgb(x, y, z)
	}
}

func FromHCL(h, c, l float64) ColorOption {
	return func(cc *color) {
		cc.r, cc.g, cc.b = hcl2rgb(h, c, l)
	}
}

func NewColor(c ...ColorOption) (*color, error) {
	color := &color{}
	for _, opt := range c {
		opt(color)
	}
	return color, nil
}

/* func easeInOut(x float64) float64 {
	return -(math.Cos(math.Pi*x) - 1) / 2
} */

/* func Gradient(a *HSL, b *HSL, v float64) *HSL {
	c := &HSL{}
	v2 := easeInOut(v)
	if (b.H - a.H) > 180 {
		c.H = a.H + 360
		c.H = math.Mod((1-v2)*c.H+v2*b.H, 360.0)
	}
	if (b.H - a.H) <= 180 {
		c.H = a.H + v*(b.H-a.H)
	}
	c.S = math.Max(0.0, math.Min((1-v)*a.S+v*b.S, 1.0))
	c.L = math.Max(0.0, math.Min((1-v)*a.L+v*b.L, 1.0))
	return c
}

func Gradient(a *LUV, b *LUV, v float64) *LUV {
	c := &LUV{}
	v = easeInOut(v)

	c.L = a.L + (b.L-a.L)*v
	c.U = a.U + (b.U-a.U)*v
	c.V = a.V + (b.V-a.V)*v
	return c
}
*/
