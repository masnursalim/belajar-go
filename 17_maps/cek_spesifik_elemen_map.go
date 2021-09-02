package main
import "fmt"

func main() {
	var a = map[string]string{"brand":"Ford", "model":"Mustang", "year":"1964", "day":""}

	val1, ok1 := a["brand"]	// cek for existing key and its value
	val2, ok2 := a["color"] // cek non-existing key and value
	val3, ok3 := a["day"] //cek for existing key and its value
	_, ok4 := a["model"] //cek only existing key	

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
}