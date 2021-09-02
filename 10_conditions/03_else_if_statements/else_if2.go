package main
import "fmt"

func main() {
	x := 30

	if x >= 10 {
		fmt.Println("x is greater than or equal to 10")
	}else if x > 20 {
		fmt.Println("x is greater than to 20")
	}else{
		fmt.Println("x is less than to 10")
	}
}