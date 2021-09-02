package main
import "fmt"

func main() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Println(a)

	a["year"] = "1970"	// update elemen
	a["color"] = "red"	// add new elemen
	fmt.Println(a)
}