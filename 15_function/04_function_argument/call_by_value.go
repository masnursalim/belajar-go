package main
import "fmt"

func swap (x, y int) int  {
	var temp int

	temp = x
	x = y
	y = temp

	return temp
}

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a %d\n", a)
	fmt.Printf("Before swap, value of b %d\n", b)

	swap(a,b)

	fmt.Printf("After swap, value of a %d\n", a)
	fmt.Printf("After swap, value of b %d\n", b)

}