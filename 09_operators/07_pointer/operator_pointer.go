package main
import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* example of type operator */
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", ptr)

	// alamat a disimpan ke ptr
	ptr = &a
	fmt.Printf("Nilai dari a : %d\n", a)
	fmt.Println(&a)
	fmt.Printf("Nilai dari *ptr : %d\n", *ptr)
	fmt.Println(ptr)

	a = 10;
	fmt.Println(*ptr)
	fmt.Println(&a)
}