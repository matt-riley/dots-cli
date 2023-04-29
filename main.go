package main

import (
	"fmt"

	"github.com/matt-riley/dots-cli/cmd/dots"
)

func main() {
	message := dots.Hello()
	fmt.Println(message)
}
