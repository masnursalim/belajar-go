package main
import "fmt"

func main() {
	prices := []int {10, 20, 30}
	fmt.Printf("Original prices : %v\n", prices)

	prices[2] = 50
	fmt.Printf("After changes prices : %v\n", prices)

}