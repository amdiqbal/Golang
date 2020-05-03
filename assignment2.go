package main

import (
	"fmt"
)

func performArtithmetic() {
	fmt.Print("Enter First Number :")
	var varNumber1 int
	fmt.Scanln(&varNumber1) //store the first input

	fmt.Print("Please Enter the Second Number Greater than 0 \n")
	fmt.Print("Enter Second Number :")
	var varNumber2 int
	fmt.Scanln(&varNumber2)

	var varNumber3 int
	// Adding the two numbers
	varNumber3 = varNumber1 + varNumber2
	fmt.Println("Addition of Two Number is : ", varNumber3)

	// subracting the numers
	varNumber3 = varNumber1 - varNumber2
	fmt.Println("Subtraction of Two Number is : ", varNumber3)

	// Muliply the numbers
	varNumber3 = varNumber1 * varNumber2
	fmt.Println("Multiplication of Two Number is : ", varNumber3)

	// Dividing the numbers
	varNumber3 = varNumber1 / varNumber2
	fmt.Println("Division of Two Number is : ", varNumber3)
}

func main() {
	//Printing first line
	fmt.Println("Simple Calculator Program using Defer Function")
	fmt.Println("-----------------------------------------------")

	// defer function
	defer fmt.Println("-----------------------------------------------")
	defer fmt.Println("End of the Calculator Programming!!")

	//Getting input from user
	fmt.Println("Please fill the Following Informations to Continue the Actions:")

	performArtithmetic()
}
