package main
import "fmt"

func factorial(x float64) (y float64)  {
	if x > 0 {
		y = x * factorial(x-1)
	} else {
		y = 1
	}

	return
}

func main() {
	fmt.Println(factorial(4))
}