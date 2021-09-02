package main
import (
	"fmt"
	"strings"
)

func main() {
	str := "Nursalim,Teknik Informatika,Unindra,2018"
	var arrstr []string = strings.Split(str,",")
	fmt.Println("Panjang array : ",len(arrstr))
	
	for idx, val := range arrstr {
		fmt.Printf("Index : %d, Value : %v\n", idx, val)
	}
}