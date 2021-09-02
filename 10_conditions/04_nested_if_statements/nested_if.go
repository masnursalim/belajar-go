package main
import "fmt"

func main() {
	num := 20

	if num >= 10 {
		fmt.Println("num is more than 10")

		if num > 15 {
			fmt.Println("num is also more than 15")
		}
	}else {
		fmt.Println("num is less than 10")
	}
}