package main
import (
	"fmt"
	"strings"
)

func main() {
	csv := []string{"Hello","World","Nursalim"}
	sep := "|"

	fmt.Println(strings.Join(csv,sep))
}