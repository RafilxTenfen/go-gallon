package main

import (
	"fmt"
)

func main() {
	bottleGroup := gallons.CreateBottleGroup(2)

	fmt.Printf("bottle %+v", bottleGroup)
}
