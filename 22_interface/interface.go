package main
import (
	"fmt"
	"math"
)

// define interface
type Shape interface {
	area() float64
}

// define circle struct
type Circle struct {
	x,y,radius float64
}

// define rectangle
type Rectangle struct {
	width, height float64
}

// define method for circle (implementation Shape.area())
func (circle Circle) area() float64  {
	return math.Pi * circle.radius * circle.radius
}

// define method for rectange (implementation Shape.are())
func (rectangle Rectangle) area() float64  {
	return rectangle.width * rectangle.height
}

// define method for shape
func getArea(shape Shape) float64  {
	return shape.area()
}

func main() {
	circle := Circle {x:0, y:0, radius:5}
	rectangle := Rectangle {width:10, height:10}

	fmt.Printf("Circle area : %.2f\n", circle.area())
	fmt.Printf("Rectangle area : %.2f\n", rectangle.area())

	// or
	fmt.Printf("Circle area : %.2f\n", getArea(circle))
	fmt.Printf("Rectangle area : %.2f\n", getArea(rectangle))
}