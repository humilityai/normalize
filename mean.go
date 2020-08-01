// Copyright 2020 Hummility AI Incorporated, All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package normalize

import "github.com/humilityai/sam"

// MeanNormalizer ...
type MeanNormalizer struct {
	Max, Min, Mean float64
}

// NewMeanNormalizer ...
func NewMeanNormalizer(max, min, mean float64) Normalizer {
	return MeanNormalizer{
		Max:  max,
		Min:  min,
		Mean: mean,
	}
}

// NewMeanNormalizerFromData ....
func NewMeanNormalizerFromData(data sam.SliceFloat64) Normalizer {
	min, max := data.Bounds()

	return MeanNormalizer{
		Max:  max,
		Min:  min,
		Mean: data.Avg(),
	}
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
func (s MeanNormalizer) RescaleSlice(slice sam.SliceFloat64) {
	for i, v := range slice {
		slice[i] = s.Scale(v)
	}
}

// ScaledSlice ...
func (s MeanNormalizer) ScaledSlice(slice sam.SliceFloat64) sam.SliceFloat64 {
	scaledData := make(sam.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Scale(v)
	}

	return scaledData
}

// UnscaledSlice ...
func (s MeanNormalizer) UnscaledSlice(slice sam.SliceFloat64) sam.SliceFloat64 {
	scaledData := make(sam.SliceFloat64, len(slice), len(slice))
	for i, v := range slice {
		scaledData[i] = s.Unscale(v)
	}

	return scaledData
}
