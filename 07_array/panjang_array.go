package main
import "fmt"

func main() {
	arr1 := [4]string{"Volvo","BMW","Ford","Mazda"}
	arr2 := [...]int{1,2,3,4,5,6}

	fmt.Println("Panjang arr1 : ",len(arr1))
	fmt.Println("Panjang arr2 : ",len(arr2))
}