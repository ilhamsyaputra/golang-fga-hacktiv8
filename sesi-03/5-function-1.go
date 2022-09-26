package main

import (
	"fmt"
	"math"
)

func main() {
	diameter := 15

	area, circumference := calculate(float64(diameter))
	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)
}

func calculate(diameter float64) (float64, float64) {

	area := math.Pi * math.Pow(diameter/2, 2)
	circumference := math.Pi * diameter

	return area, circumference
}
