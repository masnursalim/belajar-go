package main
import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	fmt.Println(strings.HasSuffix(str, "World"))
}