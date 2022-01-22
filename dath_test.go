package dath

import (
	"fmt"
	"testing"
)

func TestCreateColors(t *testing.T) {
	c := NewColor().FromRGB(255, 0, 0)
	if (*c != Color{1.0, 0.0, 0.0}) {
		t.Errorf("c == %#v; want color{1.0, 0.0, 0.0}", *c)
	}
	c2 := NewColor().FromHSL(225.0, 1.0, 0.8)
	if (*c2 != Color{0.6, 0.7, 1.0}) {
		t.Errorf("c == %#v; want color{0.6, 0.7, 1.0}", *c2)
	}
	c3 := NewColor().FromHSV(225.0, 0.4, 1.0)
	if (*c3 != Color{0.6, 0.7, 1.0}) {
		t.Errorf("c == %#v; want color{0.6, 0.7, 1.0}", *c3)
	}
	c4 := NewColor().FromLUV(50.93, 23.57, -110.91)
	if (*c4 != Color{0.64, 0.29, 0.94}) {
		t.Errorf("c == %#v; want color{0.64, 0.29, 0.94}", *c4)
	}
	c5 := NewColor().FromHCL(4.96, 104.91, 50.51)
	if (*c5 != Color{0.84, 0.26, 0.37}) {
		t.Errorf("c == %#v; want color{0.83, 0.25, 0.36}", *c5)
	}
	c6 := NewColor().FromCMYK(0.40, 0.32, 0.0, 0.0)
	if (*c6 != Color{0.6, 0.7, 1.0}) {
		t.Errorf("c == %#v; want color{0.6, 0.7, 1.0}", *c6)
	}
}

func ExampleInterpolate() {
	c1 := NewColor().FromHSL(225.0, 1.0, 0.8)
	c2 := NewColor().FromRGB(6, 185, 183)

	result := Interpolate(c1, c2)

	fmt.Printf("result == %v\n", *result)
	// Output: result == {0.41 0.72 0.86}
}

func TestInterpolate(t *testing.T) {
	c1 := NewColor().FromHSL(225.0, 1.0, 0.8)
	c2 := NewColor().FromRGB(6, 185, 183)

	result := Interpolate(c1, c2, 0.0)

	if *result != *c1 {
		t.Errorf("result == %#v; want %#v", *result, *c1)
	}
}
