package main
import "fmt"

func main() {
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("Length = %d\n", len(myslice1))
	fmt.Printf("Capacity = %d\n", cap(myslice1))

	// membuat slice tidak menyertakan capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("Length = %d\n", len(myslice2))
	fmt.Printf("Capacity = %d\n", cap(myslice2))
}