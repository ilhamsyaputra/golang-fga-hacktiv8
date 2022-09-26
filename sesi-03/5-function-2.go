// PREDEFINED RETURN VALUE

package main

import (
	"fmt"
	"math"
)

func calculate(diameter float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(diameter/2, 2)
	circumference = math.Pi * diameter
	return
}

func main() {
	var diameter float64 = 15

	area, circumference := calculate(diameter)

	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)
}
