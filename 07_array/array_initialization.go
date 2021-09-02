package main
import "fmt"

func main() {
	arr1 := [5]int{}			// belum diinisialisai
	arr2 := [5]int{1,2,3}		// partial inisialisasi
	arr3 := [5]int{1,2,3,4,5}	// full initialization

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

}