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
