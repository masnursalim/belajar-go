package main
import "fmt"

func main() {
	// simple copy
	myslice1 := []int{1, 2, 3, 4, 5}
	myslice2 := myslice1[:len(myslice1)]
	copy(myslice2, myslice1)

	fmt.Println(myslice2)

	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))
  
	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
  
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}