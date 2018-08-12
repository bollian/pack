package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	usage = "Usage: pack <dimensions> <sphere_diameter>\nExample: pack 36x40x24 5"
)

func main() {
	var exitCode int = program()
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func program() int {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, usage)
		return 1
	}

	var split []string = strings.Split(os.Args[1], "x")
	if len(split) != 3 {
		fmt.Fprintln(os.Stderr, usage)
		return 1
	}

	var dimensions [3]float64
	var tmp float64
	var err error
	for i, s := range split {
		if tmp, err = strconv.ParseFloat(s, 32); err == nil {
			dimensions[i] = tmp
		} else {
			fmt.Fprintln(os.Stderr, usage)
			return 1
		}
	}

	diameter, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		fmt.Fprintln(os.Stderr, usage)
		return 1
	}

	fmt.Println("Before pack functions: ", dimensions, diameter)

	fmt.Println("Max HCP: ", packHCP(dimensions, diameter))
	fmt.Println("Cubic: ", packCubic(dimensions, diameter))
	fmt.Println("Square Pyramid: ", packSquarePyramid(dimensions, diameter))

	return 0
}

func permutate(arr [3]float64) [][3]float64 {
	return [][3]float64{
		[3]float64{arr[0], arr[1], arr[2]},
		[3]float64{arr[0], arr[2], arr[1]},
		[3]float64{arr[1], arr[0], arr[2]},
		[3]float64{arr[1], arr[2], arr[0]},
		[3]float64{arr[2], arr[0], arr[1]},
		[3]float64{arr[2], arr[1], arr[0]},
	}
}
