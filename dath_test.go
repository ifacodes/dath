package dath

import (
	"fmt"

	d "github.com/shopspring/decimal"
)

/* func TestCreateColors(t *testing.T) {
	c := NewColor().FromRGB(255.0, 0.0, 0.0)
	cc := &Color{d.NewFromFloat(1.0), d.NewFromFloat(0.0), d.NewFromFloat(0.0)}
	if *c != *cc {
		t.Errorf("c == %#v\n; want %#v", *c, *cc)
	}
	c2 := NewColor().FromHSL(225.0, 1.0, 0.8)
	if (*c2 != Color{d.NewFromFloat(0.6), d.NewFromFloat(0.7), d.NewFromFloat(1.0)}) {
		t.Errorf("c == %#v; want color{0.6, 0.7, 1.0}", *c2)
	}
	c3 := NewColor().FromHSV(225.0, 0.4, 1.0)
	if (*c3 != Color{d.NewFromFloat(0.6), d.NewFromFloat(0.7), d.NewFromFloat(1.0)}) {
		t.Errorf("c == %#v; want color{0.6, 0.7, 1.0}", *c3)
	}
	c4 := NewColor().FromLUV(50.93, 23.57, -110.91)
	if (*c4 != Color{d.NewFromFloat(0.64), d.NewFromFloat(0.29), d.NewFromFloat(0.94)}) {
		t.Errorf("c == %#v; want color{0.64, 0.29, 0.94}", *c4)
	}
	/* 	c5 := NewColor().FromHCL(4.96, 104.91, 50.51)
	   	if (*c5 != Color{0.84, 0.26, 0.37}) {
	   		t.Errorf("c == %#v; want color{0.83, 0.25, 0.36}", *c5)
	   	}
	c6 := NewColor().FromCMYK(0.40, 0.32, 0.0, 0.0)
	if (*c6 != Color{d.NewFromFloat(0.6), d.NewFromFloat(0.7), d.NewFromFloat(1.0)}) {
		t.Errorf("c == %#v; want color{0.6, 0.7, 1.0}", *c6)
	}
} */

func ExampleInterpolate() {
	c1 := NewColor().FromHSL(225.0, 1.0, 0.8)
	c2 := NewColor().FromRGB(6, 185, 183)

	result := Interpolate(c1, c2)

	fmt.Printf("result == %s\n", result)
	// Output: result == {0.41 0.72 0.86}
}

func ExampleColor_FromRGB() {
	c := NewColor().FromRGB(0, 255, 0)

	fmt.Printf("result == %s\n", c)
	// Output: result == {0 1 0}
}

func ExampleColor_FromHSL() {
	c := NewColor().FromHSL(145, 0.5, 0.2)

	fmt.Printf("result == %s\n", c)
	// Output: result == {0.1 0.3 0.18}
}

func ExampleColor_FromHSV() {
	c := NewColor().FromHSV(59, 0.4, 1)

	fmt.Printf("result == %s\n", c)
	// Output: result == {1 0.99 0.6}
}

func ExampleColor_FromCMYK() {
	c := NewColor().FromCMYK(0.40, 0.32, 0.0, 0.0)

	fmt.Printf("result == %s\n", c)
	// Output: result == {0.6 0.68 1}
}

func ExampleColor_FromLAB() {
	c := NewColor().FromLAB(60.53, 2.79, 50.20)

	fmt.Printf("result == %s\n", c)
	// Output: result == {0.69 0.56 0.21}
}

func ExampleColor_FromLUV() {
	c := NewColor().FromLUV(50.93, 23.57, -110.91)

	fmt.Printf("result == %s\n", c)
	// Output: result == {0.64 0.29 0.94}
}

func ExampleColor_RGB() {
	c := &Color{d.NewFromFloat(0.0), d.NewFromFloat(1.0), d.NewFromFloat(0.0)}
	fmt.Printf("result == %s\n", *c.RGB())
	// Output: result == {0 255 0}
}

func ExampleColor_HSL() {
	c := &Color{d.NewFromFloat(0.1), d.NewFromFloat(0.3), d.NewFromFloat(0.18)}
	fmt.Printf("result == %s\n", *c.HSL())
	// Output: result == {144 0.5 0.2}
}

func ExampleColor_HSV() {
	c := &Color{d.NewFromFloat(1), d.NewFromFloat(0.99), d.NewFromFloat(0.6)}
	fmt.Printf("result == %s\n", *c.HSV())
	// Output: result == {58.5 0.4 1}
}

func ExampleColor_CMYK() {
	c := &Color{d.NewFromFloat(0.6), d.NewFromFloat(0.7), d.NewFromFloat(1)}
	fmt.Printf("result == %s\n", *c.CMYK())
	// Output: result == {0.4 0.3 0 0}
}

func ExampleColor_LAB() {
	c := &Color{d.NewFromFloat(0.69), d.NewFromFloat(0.56), d.NewFromFloat(0.21)}
	fmt.Printf("result == %s\n", *c.LAB())
	// Output: result == {60.85 2.81 50.32}
}

func ExampleColor_LUV() {
	c := &Color{d.NewFromFloat(0.64), d.NewFromFloat(0.29), d.NewFromFloat(0.94)}
	fmt.Printf("result == %s\n", *c.LUV())
	// Output: result == {50.65 23.43 -111.19}
}
