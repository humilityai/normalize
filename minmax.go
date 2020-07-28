package normalize

import "github.com/humilityai/sam"

// MinMaxNormalizer ...
type MinMaxNormalizer struct {
	Max, Min float64
}

// NewMinMaxNormalizer ...
func NewMinMaxNormalizer(min, max float64) Normalizer {
	return MinMaxNormalizer{
		Max: max,
		Min: min,
	}
}

// NewMinMaxNormalizerFromData ....
func NewMinMaxNormalizerFromData(data sam.SliceFloat64) Normalizer {
	min, max := data.Bounds()

	return MinMaxNormalizer{
		Max: max,
		Min: min,
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
func (s MinMaxNormalizer) RescaleSlice(slice sam.SliceFloat64) {
	for i, v := range slice {
		slice[i] = s.Scale(v)
	}
}

// ScaledSlice ...
func (s MinMaxNormalizer) ScaledSlice(slice sam.SliceFloat64) sam.SliceFloat64 {
	scaledData := make(sam.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Scale(v)
	}

	return scaledData
}

// UnscaledSlice ...
func (s MinMaxNormalizer) UnscaledSlice(slice sam.SliceFloat64) sam.SliceFloat64 {
	scaledData := make(sam.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Unscale(v)
	}

	return scaledData
}
