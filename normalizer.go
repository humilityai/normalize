package normalize

import (
	"github.com/humilityai/sam"
)

// Normalizer ...
type Normalizer interface {
	Scale(float64) float64
	Unscale(float64) float64
	RescaleSlice(sam.SliceFloat64)
	ScaledSlice(sam.SliceFloat64) sam.SliceFloat64
	UnscaledSlice(sam.SliceFloat64) sam.SliceFloat64
}

// Normalizers ...
type Normalizers []Normalizer
