package dath

import (
	"math"

	d "github.com/shopspring/decimal"
)

// LAB struct contains the converted values from a Color
type LAB struct {
	L, A, B d.Decimal
}

// LAB takes a Color and returns a LAB struct
func (c *Color) LAB() *LAB {
	lab := &LAB{}
	x, y, z := rgb2xyz(c.r, c.g, c.b)
	l, a, b := xyz2lab(x, y, z)
	lab.L, lab.A, lab.B = d.NewFromFloatWithExponent(l, -2), d.NewFromFloatWithExponent(a, -2), d.NewFromFloatWithExponent(b, -2)
	return lab
}

func lab2xyz(l, a, b float64) (x, y, z float64) {
	fy := (l + 16) / 116
	fz := fy - (b / 200)
	fx := (a / 500) + fy
	var zr float64
	if math.Pow(fz, 3.0) > 0.008856 {
		zr = math.Pow(fz, 3.0)
	} else {
		zr = (116*fz - 16) / 903.3
	}
	var yr float64
	if l > 903.3*0.008856 {
		yr = math.Pow((l+16)/116, 3.0)
	} else {
		yr = l / 903.3
	}
	var xr float64
	if math.Pow(fx, 3.0) > 0.008856 {
		xr = math.Pow(fx, 3.0)
	} else {
		xr = (116*fx - 16) / 903.3
	}
	x = xr * d65[0]
	y = yr * d65[1]
	z = zr * d65[2]
	return
}
