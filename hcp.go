// Copyright 2018 Ian Boll. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
)

func packHCP(dimensions [3]float64, diameter float64) int {
	// diameter^2 = (diameter/2)^2 + triangleHeight^2
	// triangleHeight := math.Sqrt(math.Pow(diameter, 2.0) - math.Pow(diameter/2.0, 2.0))
	triangleHeight := calTriangleHeight(diameter)

	// diameter^2 = (diameter/2)^2 + pyramidOffset^2 + pyramidHeight^2           ***** Equation 1 *****
	// diameter^2 = 0^2 + (triangleHeight - pyramidOffset)^2 + pyramidHeight^2   ***** Equation 2 *****
	// (diameter/2)^2 + pyramidOffset^2 = (triangleHeight - pyramidOffset)^2     ***** Substitution using pyramidHeight^2 *****
	// (diameter/2)^2 + pyramidOffset^2 = triangleHeight^2 - 2*triangleHeight*pyramidOffset + pyramidOffset^2
	// (diameter/2)^2 = triangleHeight^2 - 2*triangleHeight*pyramidOffset
	// pyramidOffset = [triangleHeight^2 - (diameter/2)^2] / 2 / triangleHeight
	// pyramidOffset := (math.Pow(triangleHeight, 2) - math.Pow(diameter/2, 2)) / 2.0 / triangleHeight
	pyramidOffset := calPyramidOffset(diameter, triangleHeight)

	// diameter^2 = (diameter/2)^2 + pyramidOffset^2 + pyramidHeight^2
	// pyramidHeight = (0.75*diameter^2 - pyramidOffset^2)^0.5
	// pyramidHeight := math.Sqrt(0.75*math.Pow(diameter, 2) - math.Pow(pyramidOffset, 2))
	pyramidHeight := calPyramidHeight(diameter, pyramidOffset)

	fmt.Println("HCP initial: ", triangleHeight, pyramidOffset, pyramidHeight)

	var max int = -1
	for _, permutation := range permutate(dimensions) {
		pack := packHCPExplicit(permutation[0], permutation[1], permutation[2], diameter, triangleHeight, pyramidOffset, pyramidHeight)
		fmt.Println(permutation, "=", pack)
		if pack > max {
			max = pack
		}
	}
	return max
}

func packHCPExplicit(long float64, squished float64, height float64, diameter float64, triangleHeight float64, pyramidOffset float64, pyramidHeight float64) int {
	// calculate the number of spheres in odd and even-numbered rows
	longwaysCountOdd := int(long / diameter)
	longwaysCountEven := int((long - diameter/2) / diameter)
	// fmt.Println("Longway odd, even: ", longwaysCountOdd, longwaysCountEven)

	rows := int((squished-diameter)/triangleHeight) + 1
	// fmt.Println("Rows, odd: ", rows)

	var sheetCountOdd int = addDimensionAltered(rows, longwaysCountOdd, longwaysCountEven)

	rows = int((squished-diameter-pyramidOffset)/triangleHeight) + 1
	// fmt.Println("Rows, even: ", rows)
	var sheetCountEven int = addDimensionAltered(rows, longwaysCountEven, longwaysCountOdd)

	// calculate the number of sheets stacked vertically
	sheets := int((height-diameter)/pyramidHeight) + 1
	// fmt.Println("Sheets: ", sheets)
	return addDimensionAltered(sheets, sheetCountOdd, sheetCountEven)
}

// hcpAddDimension adds up alternating values of large and small newDimension times.
func addDimensionAltered(newDimension int, large int, small int) int {
	if newDimension&1 == 1 { // if the new dimension is odd
		return (newDimension/2+1)*large + newDimension/2*small
	}
	// the new dimension is even
	return (large + small) * newDimension / 2
}

func calTriangleHeight(diameter float64) float64 {
	// diameter^2 = (diameter/2)^2 + triangleHeight^2
	return math.Sqrt(math.Pow(diameter, 2.0) - math.Pow(diameter/2.0, 2.0))
}

func calPyramidOffset(diameter float64, triangleHeight float64) float64 {
	// diameter^2 = (diameter/2)^2 + pyramidOffset^2 + pyramidHeight^2           ***** Equation 1 *****
	// diameter^2 = 0^2 + (triangleHeight - pyramidOffset)^2 + pyramidHeight^2   ***** Equation 2 *****
	// (diameter/2)^2 + pyramidOffset^2 = (triangleHeight - pyramidOffset)^2     ***** Substitution using pyramidHeight^2 *****
	// (diameter/2)^2 + pyramidOffset^2 = triangleHeight^2 - 2*triangleHeight*pyramidOffset + pyramidOffset^2
	// (diameter/2)^2 = triangleHeight^2 - 2*triangleHeight*pyramidOffset
	// pyramidOffset = [triangleHeight^2 - (diameter/2)^2] / 2 / triangleHeight
	return (math.Pow(triangleHeight, 2) - math.Pow(diameter/2, 2)) / 2.0 / triangleHeight
}

func calPyramidHeight(diameter float64, pyramidOffset float64) float64 {
	// diameter^2 = (diameter/2)^2 + pyramidOffset^2 + pyramidHeight^2
	// pyramidHeight = (0.75*diameter^2 - pyramidOffset^2)^0.5
	return math.Sqrt(0.75*math.Pow(diameter, 2) - math.Pow(pyramidOffset, 2))
}
