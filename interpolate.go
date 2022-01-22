package dath

import (
	d "github.com/shopspring/decimal"
)

// InterpolateType offers an enum to select which color space to interpolate in.
type InterpolateType int64

const (
	None InterpolateType = iota
	UseRGB
	UseCYMK
	UseHSV
	UseHSL
	UseLUV
	UseHCL
	UseLAB
)

// Interpolate returns the interpolation of the given colors at a given ratio.
// vt specifies the ratio and the color space to interpolate within (float64 or InterpolateType).
// i.e. Interpolate(color1, color1, 0.75, UseHSV)
// The default ratio and color space are 0.5 and UseLUV
func Interpolate(c1 *Color, c2 *Color, vt ...interface{}) (c *Color) {
	var v d.Decimal
	var t InterpolateType
	if len(vt) > 0 {
		for _, n := range vt {
			switch n := n.(type) {
			case float64:
				v = d.NewFromFloat(n)
			case InterpolateType:
				t = n
			}
		}
	} else {
		v = d.NewFromFloat(0.5)
	}

	switch t {
	case UseRGB:
		c = mixRGB(c1, c2, v)
	case UseCYMK:
	case UseHSV:
		hsv := mixHSV(c1.HSV(), c2.HSV(), v)
		c = NewColor()
		c.r, c.g, c.b = hsv2rgb(hsv.H, hsv.S, hsv.V)
	case UseHSL:
		hsl := mixHSL(c1.HSL(), c2.HSL(), v)
		c = NewColor()
		c.r, c.g, c.b = hsl2rgb(hsl.H, hsl.S, hsl.L)
	case UseLAB:
		lab := mixLAB(c1.LAB(), c2.LAB(), v)
		lf, _ := lab.L.Float64()
		af, _ := lab.A.Float64()
		bf, _ := lab.B.Float64()
		c = NewColor().FromLAB(lf, af, bf)
	case UseHCL:
		fallthrough
	case UseLUV:
		fallthrough
	default:
		luv := mixLUV(c1.LUV(), c2.LUV(), v)
		lf, _ := luv.L.Float64()
		uf, _ := luv.U.Float64()
		vf, _ := luv.V.Float64()
		c = NewColor().FromLUV(lf, uf, vf)
	}
	return
}

func lerp(v1, v2, r d.Decimal) d.Decimal {
	// I don't know what will happen here... I'll have to find out.
	/* 	if math.IsNaN(v1) {
	   		v1 = 0.0
	   	}
	   	if math.IsNaN(v2) {
	   		v2 = 0.0
	   	} */
	return d.NewFromFloat(1.0).Sub(r).Mul(v1).Add(r.Mul(v2))
}

func hslOrhsv(h1, s1, o1, h2, s2, o2, v d.Decimal) (hh, ss, oo d.Decimal) {
	if h2.Sub(h1).GreaterThan(d.NewFromFloat(180.0)) {
		hh = h1.Add(d.NewFromFloat(360.0))
		hh = d.NewFromFloat(1.0).Sub(v).Mul(hh).Add(v).Mul(h2).Mod(d.NewFromFloat(360.0))
	}
	if h2.Sub(h1).LessThanOrEqual(d.NewFromFloat(180.0)) {
		hh = h1.Add(v.Mul(h2.Sub(h1)))
	}
	ss = d.Max(d.NewFromFloat(0.0), d.Min(lerp(s1, s2, v), d.NewFromFloat(1.0)))
	oo = d.Max(d.NewFromFloat(0.0), d.Min(lerp(o1, o2, v), d.NewFromFloat(1.0)))
	return
}

func mixLUV(c1 *LUV, c2 *LUV, v d.Decimal) *LUV {
	new := &LUV{}
	new.L = lerp(c1.L, c2.L, v)
	new.U = lerp(c1.U, c2.U, v)
	new.V = lerp(c1.V, c2.V, v)
	return new
}

func mixLAB(c1 *LAB, c2 *LAB, v d.Decimal) *LAB {
	new := &LAB{}
	new.L = lerp(c1.L, c2.L, v)
	new.A = lerp(c1.A, c2.A, v)
	new.B = lerp(c1.B, c2.B, v)
	return new
}

func mixRGB(c1 *Color, c2 *Color, v d.Decimal) *Color {
	new := &Color{}
	new.r = lerp(c1.r, c2.r, v)
	new.g = lerp(c1.g, c2.g, v)
	new.b = lerp(c1.b, c2.b, v)
	return new
}

func mixHSV(c1 *HSV, c2 *HSV, v d.Decimal) *HSV {
	hsv := &HSV{}
	hsv.H, hsv.S, hsv.V = hslOrhsv(c1.H, c1.S, c1.V, c2.H, c2.S, c2.V, v)
	return hsv
}
func mixHSL(c1 *HSL, c2 *HSL, v d.Decimal) *HSL {
	hsl := &HSL{}
	hsl.H, hsl.S, hsl.L = hslOrhsv(c1.H, c1.S, c1.L, c2.H, c2.S, c2.L, v)
	return hsl
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
