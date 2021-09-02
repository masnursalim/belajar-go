package main
import "fmt"

func main() {
	// crate a slice
	numbers := []int {0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	fmt.Println("numbers[1:4] :",numbers[1:4])

	/* missing lower bound implies 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* missing upper bound implies len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)
}

func printSlice(slice []int)  {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)
}