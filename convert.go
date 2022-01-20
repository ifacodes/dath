package dath

var (
	d65     = [3]float64{0.95047, 1.0, 1.08883}
	sRGB    = [3][3]float64{{0.4124564, 0.3575761, 0.1804375}, {0.2126729, 0.7151522, 0.0721750}, {0.0193339, 0.1191920, 0.9503041}}
	sRGBInv = [3][3]float64{{3.2404542, -1.5371385, -0.4985314}, {-0.9692660, 1.8760108, 0.0415560}, {0.0556434, -0.2040259, 1.0572252}}
)

func (c *color) RGB() (r, g, b int) {
	return
}

func (c *color) RGBA() (r, g, b, a int) {
	return
}

func hslCalculateHue(max, chroma, r, g, b float64) (h float64) {
	if chroma == 0.0 {
		h = 0.0
	} else if max == r {
		if g < b {
			h = 6.0
		}
		h += (g - b) / chroma

	} else if max == g {
		h = 2.0 + ((b - r) / chroma)
	} else if max == b {
		h = 4.0 + ((r - g) / chroma)
	}
	h *= 60.0
	return
}
