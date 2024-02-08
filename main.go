package main

import (
	"fmt"
)

func main() {
	var mode string
	fmt.Println("Enter the operation : ")
	fmt.Scan(&mode)
	var num1, num2 float64
	fmt.Println("Enter first number : ")
	fmt.Scan(&num1)
	fmt.Println("Enter second number : ")
	fmt.Scan(&num2)
	switch mode {
	case "+":
		sum(num1, num2)
	case "-":
		sustract(num1, num2)
	case "*":
		multiply(num1, num2)
	case "/":
		divide(num1, num2)
	default:
		fmt.Println("Invalid operator")
	}
}

func sum(num1, num2 float64) {
	var resultadosum float64
	resultadosum = num1 + num2
	fmt.Println("The result is: ", resultadosum)
	return
}

func sustract(num1, num2 float64) {
	var resultadosustract float64
	resultadosustract = num1 - num2
	fmt.Println("The result is: ", resultadosustract)
	return
}

func multiply(num1, num2 float64) {
	var resultadomul float64
	resultadomul = num1 * num2
	fmt.Println("The result is: ", resultadomul)
	return
}

func divide(num1, num2 float64) {
	switch num2 {
	case 0:
		fmt.Println("You cant divide by 0")
		return
	default:
		switch num1 {
		case 0:
			fmt.Println("The result is : 0")
			return
		default:
			var resultadodiv float64
			resultadodiv = num1 / num2
			fmt.Println("The result is: ", resultadodiv)
			return
		}
	}

}
