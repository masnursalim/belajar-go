package main
import (
	"fmt"
	"reflect"
)

func main() {
	age := 27.5
	fmt.Printf("Tipe dari variable age tanpa reflect %T\n", age)
	fmt.Println("Tipe dari variable age dengan reflect", reflect.TypeOf(age))
}