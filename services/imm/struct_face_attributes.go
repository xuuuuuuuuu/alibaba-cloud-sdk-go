package imm

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// FaceAttributes is a nested struct in imm response
type FaceAttributes struct {
	Glasses           string       `json:"Glasses" xml:"Glasses"`
	MaskConfidence    float64      `json:"MaskConfidence" xml:"MaskConfidence"`
	Mask              string       `json:"Mask" xml:"Mask"`
	GlassesConfidence float64      `json:"GlassesConfidence" xml:"GlassesConfidence"`
	RaceConfidence    float64      `json:"RaceConfidence" xml:"RaceConfidence"`
	Beard             string       `json:"Beard" xml:"Beard"`
	Race              string       `json:"Race" xml:"Race"`
	BeardConfidence   float64      `json:"BeardConfidence" xml:"BeardConfidence"`
	FaceBoundary      FaceBoundary `json:"FaceBoundary" xml:"FaceBoundary"`
}