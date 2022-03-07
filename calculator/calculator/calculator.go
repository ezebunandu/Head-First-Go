package Calculator

/*
A simple calculator that performs one calculation
as per the user's choice
*/

import (
	"fmt"
)

func Calc() {
	fmt.Println("Welcome to calculator")
	fmt.Println("*****************MAIN MENU*******************")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("*********************************************")
	var choice int

	fmt.Scan(&choice)
	var a, b int

	fmt.Println("Enter value of a: ")
	fmt.Scan(&a)
	fmt.Println("Enter value of b: ")
	fmt.Scan(&b)
	if choice == 1 {
		ans := a + b
		fmt.Println("Answer = ", ans)
	} else if choice == 2 {
		ans := a - b
		fmt.Println("Answer = ", ans)
	} else if choice == 3 {
		ans := a * b
		fmt.Println("Answer = ", ans)
	} else {
		ans := a / b
		fmt.Println("Answer =  ", ans)
	}
	fmt.Println(("bye"))

}
