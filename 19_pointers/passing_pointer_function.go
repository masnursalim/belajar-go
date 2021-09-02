package main
import "fmt"

func main() {
	var a int = 10
	var b int = 20

	fmt.Printf("Sebelum swap, nilai dari a : %d\n", a)
	fmt.Printf("Sebelum swap, nilai dari b : %d\n", b)

	swap(&a, &b)
	fmt.Printf("Setelah swap, nilai dari a : %d\n", a)
	fmt.Printf("Setelah swap, nilai dari b : %d\n", b)

}

func swap (x *int, y *int)  {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}