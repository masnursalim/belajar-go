package main
import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* take the address of var */
	ptr = &a 

	/* take the address of ptr */
	pptr = &ptr

	fmt.Println("nilai dari a : ",a)
	fmt.Println("nilai dari ptr : ",*ptr)
	fmt.Println("nilai dari pptr : ",**pptr)

	fmt.Println("alamat dari a : ",&a)
	fmt.Println("alamat dari ptr : ",ptr)
	fmt.Println("alamat dari pptr : ",pptr)
}