package main

import (
	"fmt"
	"log"

	"github.com/matt-riley/go-template/cmd"
)

func main() {
	message, err := cmd.Hello()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
