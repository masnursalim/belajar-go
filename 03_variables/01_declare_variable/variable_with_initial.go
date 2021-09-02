package main
import "fmt"

func main() {
	var student1 string = "john"	// type is string
	var student2 = "Jane"			// type is inferred
	x := 1							// type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}