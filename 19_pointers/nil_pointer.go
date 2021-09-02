package main
import "fmt"

func main() {
	var ptr *int

	fmt.Println("Nilai dari variable ptr : ",ptr)

	if ptr == nil {
		fmt.Println("Nilai variable prt adalah nil")
	}
}