package main

import (
	"fmt"

	"github.com/gustavosinacio/bootdev-go/structs"
)

func main() {
	car := structs.Car{
		Model:   "myModel",
		Make:    "myMake",
		Doors:   5,
		Mileage: 13089,
	}

	fmt.Println(car)
}
