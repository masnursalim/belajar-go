package main
import "fmt"

func main() {
	x := 5

	fmt.Println(x < 5 && x < 10)
	fmt.Println(x < 5 || x < 4) 
	fmt.Println(!(x < 5 && x < 10))
}