package main
import "fmt"

/* global variable declaration */
var a int = 20

func main() {
	
	/* local variable declaration in main function */
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("Nilai dari a di main() = %d\n", a)
	c = sum(a,b)
	fmt.Printf("Nilai dari c di main() = %d\n", c)
}

func sum(a, b int) int {
	fmt.Printf("value of a in sum() = %d\n",  a);
    fmt.Printf("value of b in sum() = %d\n",  b);
	return a + b
}