package dath

import (
	"testing"
)

func TestCreateColors(t *testing.T) {
	c, _ := NewColor(FromRGB(255, 0, 0))
	if (*c != color{1.0, 0.0, 0.0}) {
		t.Errorf("c == %#v; want color{1.0, 0.0, 0.0}", *c)
	}
	c2, _ := NewColor(FromHSL(225.0, 1.0, 0.8))
	if (*c2 != color{0.6, 0.7, 1.0}) {
		t.Errorf("c == %#v; want color{0.6, 0.7, 1.0}", *c2)
	}
	c3, _ := NewColor(FromHSV(225.0, 0.4, 1.0))
	if (*c3 != color{0.6, 0.7, 1.0}) {
		t.Errorf("c == %#v; want color{0.6, 0.7, 1.0}", *c3)
	}
	c4, _ := NewColor(FromLUV(50.93, 23.57, -110.91))
	if (*c4 != color{0.64, 0.29, 0.94}) {
		t.Errorf("c == %#v; want color{0.64, 0.29, 0.94}", *c4)
	}
	c5, _ := NewColor(FromHCL(4.96, 104.91, 50.51))
	if (*c5 != color{0.84, 0.26, 0.37}) {
		t.Errorf("c == %#v; want color{0.83, 0.25, 0.36}", *c5)
	}
}

/* func TestRGB2HSV(t *testing.T) {
	c1 := &RGB{153, 178, 255}
	c2 := &RGB{156, 137, 184}
	c3 := &RGB{240, 166, 202}
	if (*c1.RGBToHSV() != HSV{225, 0.4, 1}) {
		t.Errorf("c1 == %#v; want HSV{225, 0.4, 1}", *c1.RGBToHSV())
	}
	if (*c2.RGBToHSV() != HSV{264, 0.26, 0.72}) {
		t.Errorf("c2 == %#v; want HSV{264, 0.6, 0.72}", *c2.RGBToHSV())
	}
	if (*c3.RGBToHSV() != HSV{331, 0.31, 0.94}) {
		t.Errorf("c3 == %#v; want HSV{331, 0.31, 0.94}", *c3.RGBToHSV())
	}
}

func TestHSV2RGB(t *testing.T) {
	c1 := &HSV{264, 0.255, 0.72}
	c2 := &HSV{225, 0.4, 1}
	c3 := &HSV{331, 0.31, 0.94}
	if (*c1.HSVToRGB() != RGB{156, 137, 184}) {
		t.Errorf("c1 == %#v; want RGB{156, 137, 184}", *c1.HSVToRGB())
	}
	if (*c2.HSVToRGB() != RGB{153, 179, 255}) {
		t.Errorf("c2 == %#v; want RGB{153, 178, 255}", *c2.HSVToRGB())
	}
	if (*c3.HSVToRGB() != RGB{240, 165, 201}) {
		t.Errorf("c3 == %#v; want RGB{240, 166, 202}", *c3.HSVToRGB())
	}
}

func TestRGB2HSL(t *testing.T) {
	c1 := &RGB{240, 166, 202}
	c2 := &RGB{153, 178, 255}
	c3 := &RGB{156, 137, 184}
	if (*c1.RGBToHSL() != HSL{331, 0.71, 0.8}) {
		t.Errorf("c1 == %#v; want HSL{331, 0.71, 0.8}", *c1.RGBToHSL())
	}
	if (*c2.RGBToHSL() != HSL{225, 1, 0.8}) {
		t.Errorf("c2 == %#v; want HSL{225, 1, 0.8}", *c2.RGBToHSL())
	}
	if (*c3.RGBToHSL() != HSL{264, 0.25, 0.63}) {
		t.Errorf("c3 == %#v; want HSL{264, 0.24, 0.62}", *c3.RGBToHSL())
	}
} */

/* func TestRGB2XYZ(t*testing.T) {
	c1 := &RGB{240, 166, 202}
	c2 := &RGB{153, 178, 255}
	c3 := &RGB{156, 137, 184}
	if (*c1.RGBToXYZ() != XYZ{}) {
		t.Errorf("c1 == %#v; want HSL{331, 0.71, 0.8}", *c1.RGBToHSL())
	}
	if (*c2.RGBToXYZ() != HSL{225, 1, 0.8}) {
		t.Errorf("c2 == %#v; want HSL{225, 1, 0.8}", *c2.RGBToHSL())
	}
	if (*c3.RGBToXYZ() != HSL{264, 0.25, 0.63}) {
		t.Errorf("c3 == %#v; want HSL{264, 0.24, 0.62}", *c3.RGBToHSL())
	}
} */
