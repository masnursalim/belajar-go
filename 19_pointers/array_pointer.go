package main
import "fmt"

const MAX int = 3

func main() {
	a := []int{100,200,300}

	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]	// assign address of integer
	}

	// looping array pointer
	for i = 0; i < MAX; i++ {
		fmt.Printf("Nilai dari ptr[%d] = %d\n", i, *ptr[i])
	}
}