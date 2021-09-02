package main
import "fmt"

func main() {
	var x = 5
	fmt.Println(x)

	x = 5
	x += 3
	fmt.Println(x)

	x = 5
	x -= 3
	fmt.Println(x)

	x = 5
	x *= 3
	fmt.Println(x)

	x = 5
	x /= 3
	fmt.Println(x)

	x = 5
	x %= 3
	fmt.Println(x)

	x = 5
	fmt.Printf("x is %b\n", x)
	x &= 3
	fmt.Printf("x is %b\n", x)

	x = 5
	fmt.Printf("x is %b\n", x)
	x |= 3
	fmt.Printf("x is %b\n", x)

	x = 5
	fmt.Printf("x is %b\n", x)
	x ^= 3
	fmt.Printf("x is %b\n", x)

	x = 5
	fmt.Printf("x is %b\n", x)
	x >>= 3
	fmt.Printf("x is %b\n", x)

	x = 5
	fmt.Printf("x is %b\n", x)
	x <<= 3
	fmt.Printf("x is %b\n", x)




}