package SprintlyCLI

import (
	"fmt"
	"github.com/jonathanporta/go-sprintly-api"
)

func PrintProducts(products []SprintlyAPI.Product) {
	fmt.Println(products)
}
