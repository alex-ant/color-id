package color

import "fmt"

const (
	intensityHigh uint8 = 255
	intensityLow  uint8 = 0
)

// Color specifies an RGB color.
type Color struct {
	R uint8
	G uint8
	B uint8
}

func getHex(i uint8) string {
	s := fmt.Sprintf("%X", i)

	if len(s) == 1 {
		s = "0" + s
	}

	return s
}

// Hex returns a hexadecimal representation of color.
func (c Color) Hex() string {
	return fmt.Sprintf("#%s%s%s", getHex(c.R), getHex(c.G), getHex(c.B))
}

type paletteColor struct {
	High   [2]float32
	Dim    []float32
	Bright []float32
}

// Palette contains RGB palette configuration.
type Palette struct {
	r paletteColor
	g paletteColor
	b paletteColor
}

// NewPalette returns new color palette.
func NewPalette() *Palette {
	return &Palette{
		r: paletteColor{
			High:   [2]float32{0, 0.25},
			Dim:    []float32{0.25, 0.42},
			Bright: []float32{0.92, 1},
		},
		g: paletteColor{
			High:   [2]float32{0.25, 0.67},
			Dim:    []float32{0.67, 0.83},
			Bright: []float32{0.08, 0.25},
		},
		b: paletteColor{
			High:   [2]float32{0.67, 1},
			Dim:    []float32{},
			Bright: []float32{0.5, 0.67},
		},
	}
}

// GetColor returns the corresponding RGB color at the given [0,1] point on the
// linear palette.
func (p *Palette) GetColor(point float32) Color {
	if point < 0 || point > 1 {
		return Color{}
	}

	return Color{
		R: p.r.getIntensity(point),
		G: p.g.getIntensity(point),
		B: p.b.getIntensity(point),
	}
}

func (c *paletteColor) getIntensity(point float32) uint8 {
	if point >= c.High[0] && point <= c.High[1] {
		return intensityHigh
	}

	if len(c.Bright) == 2 && point >= c.Bright[0] && point <= c.Bright[1] {
		return uint8((point - c.Bright[0]) * float32(intensityHigh) / (c.Bright[1] - c.Bright[0]))
	}

	if len(c.Dim) == 2 && point >= c.Dim[0] && point <= c.Dim[1] {
		return -uint8((point - c.Dim[0]) * float32(intensityHigh) / (c.Dim[1] - c.Dim[0]))
	}

	return intensityLow
}
