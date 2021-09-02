package main
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("             Hello World          ")
	fmt.Println(strings.TrimSpace("             Hello World          "))
	fmt.Println(strings.TrimSpace(" \t\n I love my country  \n\t\r\n")) 	
}