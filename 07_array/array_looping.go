package main
import "fmt"

func main() {
	var x[5]int

	// var i,j int

	for i := 0; i < len(x); i++ {
		x[i] = i+10
	}

	for j := 0; j < len(x); j++ {
		fmt.Printf("Elemen [%d] = %d\n", j, x[j])
	}
}