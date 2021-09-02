package main
import "fmt"

func main() {
	var x float64 = 1.7e+308
	var y = 10.5
	fmt.Printf("Type : %T, Value: %v\n", x, x)
	fmt.Printf("Type : %T, Value: %v\n", y, y)
}