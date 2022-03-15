package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 || height < 0 {
		return 0, fmt.Errorf("The height and width must be positive")
	} else {
		area := width * height
		return area / 10, nil
	}

}

func main() {
	var amount, total float64
	amount, err := paintNeeded(5.2, 3.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f litres needed\n", amount)
		total += amount
	}

	other_amount, err := paintNeeded(-5, 3.5)
	if err != nil {
		log.Fatal(err)
	} else {
		amount += other_amount
	}
	fmt.Printf("%0.2f litres needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f litres needed\n", total)
}
