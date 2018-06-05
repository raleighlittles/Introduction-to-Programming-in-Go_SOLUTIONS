package main

import "fmt"
import "math"

func main() {
	problem1()
	problem2()
}

func problem1() {
	fmt.Println(`Methods are called by object (part of a class), functions can be called on their own.`)
}

func problem2() {
	fmt.Println(`Embedded anonymous fields are for when you have a 'is-a' relationship with the overall Type
		and its field, whereas normal fields are when you have an 'has-a' relationship.`)
}

// problem3 begin
type Shape interface {
	area() float64 // already implemented in textbook
	perimeter() float64
}

// from page 106
type MultiShape struct {
	shapes []Shape
}

// part 1 of task
func (m *MultiShape) perimeter() float64 {
	var peri float64
	for _, shape := range m.shapes {
		peri += shape.perimeter()
	}
	return peri
}

// from page 102
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// from page 97
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// this is what the atual implementation part is
func (r *Rectangle) perimeter() float64 {
	peri := r.x1 + r.x2 + r.y1 + r.y2
	return peri
}

type Circle struct {
	x float64
	y float64
	r float64
}

// this is what the actual implementation part is
func (c *Circle) perimeter() float64 {
	peri := 2 * math.Pi * c.r
	return peri
}
