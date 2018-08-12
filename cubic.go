// Copyright 2018 Ian Boll. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "math"

func packCubic(dimensions [3]float64, diameter float64) int {
	width := int(dimensions[0] / diameter)
	length := int(dimensions[1] / diameter)
	height := int(dimensions[2] / diameter)
	return width * length * height
}

func packSquarePyramid(dimensions [3]float64, diameter float64) int {
	var max int = -1
	for _, permutation := range permutate(dimensions) {
		pack := packSquarePyramidExplicit(permutation[0], permutation[1], permutation[2], diameter)
		if pack > max {
			max = pack
		}
	}
	return max
}

func packSquarePyramidExplicit(width float64, length float64, height float64, diameter float64) int {
	sheetHeight := math.Sqrt(math.Pow(diameter, 2) - 2.0*math.Pow(diameter/2.0, 2))
	sheets := int((height-diameter)/sheetHeight) + 1

	oddSheetWidth := int(width / diameter)
	oddSheetLength := int(length / diameter)
	evenSheetWidth := int((width - (diameter / 2.0)) / diameter)
	evenSheetLength := int((length - (diameter / 2.0)) / diameter)
	return addDimensionAltered(sheets, oddSheetWidth*oddSheetLength, evenSheetWidth*evenSheetLength)
}
