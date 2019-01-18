package color

import (
	"math/rand"
	"sort"
	"time"
)

var workingPalette *Palette

func init() {
	workingPalette = NewPalette()
}

type colorItem struct {
	color Color
	point float32
	id    string
}

// Set contains colorset info.
type Set struct {
	generated []colorItem
}

// NewSet returns a new colorset.
func NewSet() *Set {
	return new(Set)
}

func (s *Set) sortGenerated() {
	sort.Slice(s.generated, func(i, j int) bool {
		return s.generated[i].point < s.generated[j].point
	})
}

func (s *Set) populateGenerated(point float32, id string) Color {
	c := workingPalette.GetColor(point)

	s.generated = append(s.generated, colorItem{
		color: c,
		point: point,
		id:    id,
	})

	s.sortGenerated()

	return c
}

// GetColor generates a color for the given ID which can be any string. The
// first color is being generated randomly, and all the rest function calls
// return a set of colors evenly distributed among the palette scale.
// The function returns the came color if the ID has already been registered.
func (s *Set) GetColor(id string) Color {
	// Return an existing color if one has already been generated.
	for _, ci := range s.generated {
		if id == ci.id {
			return ci.color
		}
	}

	// Generate first color at a random palette position for the first time.
	if len(s.generated) == 0 {
		rp := randomPoint()

		return s.populateGenerated(rp, id)
	}

	if len(s.generated) == 1 {
		newPoint := s.generated[0].point + 0.5
		if newPoint > 1 {
			newPoint--
		}

		return s.populateGenerated(newPoint, id)
	}

	// Generate the next color by determining a point within the largest gap
	// on the linear palette scale.
	maxGapStart := s.generated[len(s.generated)-1].point
	maxGapEnd := s.generated[0].point
	maxGapWidth := 1 - maxGapStart + maxGapEnd

	if len(s.generated) == 2 && maxGapStart-maxGapEnd > maxGapWidth {
		maxGapStart = maxGapEnd
		maxGapWidth = maxGapStart - maxGapEnd
	}

	for i := 0; i < len(s.generated)-1; i++ {
		p1 := s.generated[i].point
		p2 := s.generated[i+1].point
		gap := p2 - p1

		if gap > maxGapWidth {
			maxGapWidth = gap
			maxGapStart = p1
		}
	}

	newPoint := maxGapStart + maxGapWidth/2
	if newPoint > 1 {
		newPoint--
	}

	return s.populateGenerated(newPoint, id)
}

func randomPoint() float32 {
	rand.Seed(time.Now().UnixNano())
	return float32(rand.Intn(1000)) / 1000
}
