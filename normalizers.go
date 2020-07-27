package normalize

import (
	"github.com/humilityai/slices"
)

// Normalizer ...
type Normalizer interface {
	Scale(float64) float64
	Unscale(float64) float64
	RescaleSlice(slices.SliceFloat64)
	ScaledSlice(slices.SliceFloat64) slices.SliceFloat64
	UnscaledSlice(slices.SliceFloat64) slices.SliceFloat64
}

// MinMaxNormalizer ...
type MinMaxNormalizer struct {
	Max, Min float64
}

// MeanNormalizer ...
type MeanNormalizer struct {
	Max, Min, Mean float64
}

// ZScoreNormalizer ...
type ZScoreNormalizer struct {
	Mean   float64
	StdDev float64
}

// Normalizers ...
type Normalizers []Normalizer

// NewMinMaxNormalizer ...
func NewMinMaxNormalizer(min, max float64) Normalizer {
	return MinMaxNormalizer{
		Max: max,
		Min: min,
	}
}

// NewMeanNormalizer ...
func NewMeanNormalizer(max, min, mean float64) Normalizer {
	return MeanNormalizer{
		Max:  max,
		Min:  min,
		Mean: mean,
	}
}

// NewZscoreNormalizer ...
func NewZscoreNormalizer(mean, stddev float64) Normalizer {
	return ZScoreNormalizer{
		Mean:   mean,
		StdDev: stddev,
	}
}

// NewMinMaxNormalizerFromData ....
func NewMinMaxNormalizerFromData(data slices.SliceFloat64) Normalizer {
	min, max := data.Bounds()

	return MinMaxNormalizer{
		Max: max,
		Min: min,
	}
}

// NewMeanNormalizerFromData ....
func NewMeanNormalizerFromData(data slices.SliceFloat64) Normalizer {
	min, max := data.Bounds()

	return MeanNormalizer{
		Max:  max,
		Min:  min,
		Mean: data.Avg(),
	}
}

// Scale ...
func (s MinMaxNormalizer) Scale(value float64) float64 {
	numerator := value - s.Min
	denominator := s.Max - s.Min

	return numerator / denominator
}

// Unscale ...
func (s MinMaxNormalizer) Unscale(value float64) float64 {
	return (value * (s.Max - s.Min)) + s.Min
}

// RescaleSlice ...
func (s MinMaxNormalizer) RescaleSlice(slice slices.SliceFloat64) {
	for i, v := range slice {
		slice[i] = s.Scale(v)
	}
}

// ScaledSlice ...
func (s MinMaxNormalizer) ScaledSlice(slice slices.SliceFloat64) slices.SliceFloat64 {
	scaledData := make(slices.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Scale(v)
	}

	return scaledData
}

// UnscaledSlice ...
func (s MinMaxNormalizer) UnscaledSlice(slice slices.SliceFloat64) slices.SliceFloat64 {
	scaledData := make(slices.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Unscale(v)
	}

	return scaledData
}

// Scale ...
func (s MeanNormalizer) Scale(value float64) float64 {
	numerator := value - s.Mean
	denominator := s.Max - s.Min

	return numerator / denominator
}

// Unscale ...
func (s MeanNormalizer) Unscale(value float64) float64 {
	return (value * (s.Max - s.Min)) + s.Mean
}

// RescaleSlice ...
func (s MeanNormalizer) RescaleSlice(slice slices.SliceFloat64) {
	for i, v := range slice {
		slice[i] = s.Scale(v)
	}
}

// ScaledSlice ...
func (s MeanNormalizer) ScaledSlice(slice slices.SliceFloat64) slices.SliceFloat64 {
	scaledData := make(slices.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Scale(v)
	}

	return scaledData
}

// UnscaledSlice ...
func (s MeanNormalizer) UnscaledSlice(slice slices.SliceFloat64) slices.SliceFloat64 {
	scaledData := make(slices.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Unscale(v)
	}

	return scaledData
}

// Scale ...
func (s ZScoreNormalizer) Scale(value float64) float64 {
	numerator := value - s.Mean
	denominator := s.StdDev

	return numerator / denominator
}

// Unscale ...
func (s ZScoreNormalizer) Unscale(value float64) float64 {
	return (value * (s.StdDev)) + s.Mean
}

// RescaleSlice ...
func (s ZScoreNormalizer) RescaleSlice(slice slices.SliceFloat64) {
	for i, v := range slice {
		slice[i] = s.Scale(v)
	}
}

// ScaledSlice ...
func (s ZScoreNormalizer) ScaledSlice(slice slices.SliceFloat64) slices.SliceFloat64 {
	scaledData := make(slices.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Scale(v)
	}

	return scaledData
}

// UnscaledSlice ...
func (s ZScoreNormalizer) UnscaledSlice(slice slices.SliceFloat64) slices.SliceFloat64 {
	scaledData := make(slices.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Unscale(v)
	}

	return scaledData
}
