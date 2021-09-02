package main
import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World!"
	fmt.Println(strings.Replace(str, "World", "Go", 1))
}