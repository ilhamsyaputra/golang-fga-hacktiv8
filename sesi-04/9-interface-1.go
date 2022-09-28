package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

type cube struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (kubus cube) area() float64 {
	return kubus.length * kubus.length * 6
}

func (kubus cube) perimeter() float64 {
	return kubus.length * 12
}

func (kubus cube) volume() float64 {
	return math.Pow(kubus.length, 3)
}

func calculate(text string, entity shape) {
	fmt.Printf("%s area: %v \n", text, entity.area())
	fmt.Printf("%s perimeter: %v \n", text, entity.perimeter())
}

func main() {
	var circle1 shape = circle{radius: 7}
	var rectangle1 shape = rectangle{height: 2, width: 4}
	var cube1 shape = cube{length: 4}

	fmt.Printf("Value variable circle1: %+v \n", circle1)
	fmt.Printf("Type of circle1: %T \n", circle1)
	fmt.Println()
	fmt.Printf("Value variable rectangle1: %+v \n", rectangle1)
	fmt.Printf("Type of rectangle1: %T \n", rectangle1)
	fmt.Println()
	fmt.Printf("Value variable cube1: %+v \n", cube1)
	fmt.Printf("Type of cube1: %T \n", cube1)

	fmt.Println()
	calculate("Circle", circle1)
	fmt.Println()
	calculate("Rectangle", rectangle1)
	fmt.Println()
	calculate("Cube", cube1)
	fmt.Println()

	// karena method volume dari cube tidak didefinisikan di interface shape, maka untuk menjalankan method volume() adalah dengan:
	fmt.Printf("%s volume: %v \n", "Cube", cube1.(cube).volume())

	// kita dapat melakukan pengecekan apakah suatu method sudah di definisikan untuk suatu struct // ASSERTION
	// Contoh tidak ada:
	var circle2 shape = circle{radius: 9}

	value, ok := circle2.(circle).volume() // AKAN TERJADI ERROR DAN TIDAK BISA DI COMPILE KARENA CIRCLE TIDAK MEMILIKI METHOD VOLUME()

	if ok {
		fmt.Printf("%s volume: %v \n", "Circle", value)
	}
}
