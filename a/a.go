package main
import (
	"github.com/go-modules-by-example-staging/submodulestest/b"
	"fmt"
)
const Name = b.Name
func main() {
	fmt.Println(Name)
}
