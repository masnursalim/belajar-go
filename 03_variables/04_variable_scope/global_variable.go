package main
import "fmt"

/* global variable declaration */
var g int

func main() {
	/* local variable declaration */
	var a, b int
	a = 10
	b = 20

	g = a + b
	fmt.Printf("Nilai dari a = %d, b = %d, g = %d\n", a, b, g)
}