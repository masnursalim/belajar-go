package main
import (
	"fmt"
	"math"
)

/* define circle */
type Circle struct {
	x,y,radius float64
}

type Person struct {
	firstName,lastName string
}

/* define a method for cirlce */
func(circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func(p Person) displayFullName() string {
	return p.firstName + " "+p.lastName
}

func main() {
	circle := Circle{x:0, y:0, radius:5}
	fmt.Printf("Circle area : %v\n", circle.area())

	person := Person{firstName:"Nursalim", lastName:"Al Farizi"}
	fmt.Printf("Nama lengkap : %v\n", person.displayFullName())

}