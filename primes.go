package main

import (
	"fmt"
	"math"
)

func IsDigit(number float64) bool {
	return number == float64(int(number))
}

func FindMp(multiple int) [2]float64 {
	var i int = 1
	for {
		value := (multiple + (i * i))
		root := math.Sqrt(float64(value))
		if IsDigit(root) == true {
			checkSums := [2]float64{root, float64(i)}
			fmt.Println(i)
			return checkSums
		}
		i += 1
	}
}

func SolveSim(a [2]float64) []float64 {
	x := ((a[0] * 2) + (a[1] * 2)) / 2
	y := (a[0] * 2) - x

	return []float64{x, y}
}
func main() {
	fmt.Println(SolveSim(FindMp(377)))
}
