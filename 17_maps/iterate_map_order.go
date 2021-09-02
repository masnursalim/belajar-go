package main
import "fmt"

func main() {
	a := map[string]int{"one":1, "two":2, "three":3, "four":4}

	var b = []string{} //defining the order
	b = append(b, "one", "two", "three", "four")

	// loop with no order
	for k, v := range a {
		fmt.Printf("%v : %v\n", k, v)
	}

	fmt.Println()

	for _, element := range b {
		fmt.Printf("%v : %v\n", element, a[element])
	}
}