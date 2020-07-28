package normalize

import "github.com/humilityai/sam"

// ZScoreNormalizer ...
type ZScoreNormalizer struct {
	Mean   float64
	StdDev float64
}

// NewZscoreNormalizer ...
func NewZscoreNormalizer(mean, stddev float64) Normalizer {
	return ZScoreNormalizer{
		Mean:   mean,
		StdDev: stddev,
	}
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
func (s ZScoreNormalizer) RescaleSlice(slice sam.SliceFloat64) {
	for i, v := range slice {
		slice[i] = s.Scale(v)
	}
}

// ScaledSlice ...
func (s ZScoreNormalizer) ScaledSlice(slice sam.SliceFloat64) sam.SliceFloat64 {
	scaledData := make(sam.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Scale(v)
	}

	return scaledData
}

// UnscaledSlice ...
func (s ZScoreNormalizer) UnscaledSlice(slice sam.SliceFloat64) sam.SliceFloat64 {
	scaledData := make(sam.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Unscale(v)
	}

	return scaledData
}
