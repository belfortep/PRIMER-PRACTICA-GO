package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	fmt.Println("Hola Mundo!")

	fmt.Println(uuid.New().String())

}
