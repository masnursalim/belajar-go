package main
import "fmt"

func myFunc(x int, y int) (result int) {
	result = x + y
	return
	// or with below
	// return result  
}
func main() {
	fmt.Println(myFunc(10,5))
	result := myFunc(10,5)
	fmt.Println(result)
}