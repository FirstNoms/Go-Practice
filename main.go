package main

import "fmt"

//first program in Go.
func main() {
	//fmt.Println("Hello World")

	//Collecting a number from a user and outputting it doubled.
	/* 	fmt.Print("Enter a Number: ")
	   	var userInput float64
	   	fmt.Scanf("%f", &userInput)
	   	output := userInput*2
	   	fmt.Println(output)
	*/

	//Example2
	/*	var value = 5; value+=1;
		fmt.Println(value)
	*/

	//Example3: Converting Fahrenheit into Celsius
	/*	fmt.Print("Enter the Temperature in Fahrenheit: ")
		var userInput float64
		fmt.Scanf("%f", &userInput)
		output:= ((userInput-32) * 5/9)
		fmt.Printf( "To convert %.1f to Celsius is %.1f m", userInput,output)
	*/

	//Example4: Converting feet into meters
	/*	fmt.Print("Enter your height in feet: ")
		var userInput float64
		fmt.Scanf("%f", &userInput)
		display:= (userInput * 0.3048)
		fmt.Printf("Converting %.1fft to Meters is = %.1f meter", userInput,display)
	*/

	//Example5:Write a complete program that calculates and prints the product of three integers
	/*	var firstNumber = 30
		var secondNumber = 50
		var thirdNumber = 20
		product:= firstNumber * secondNumber * thirdNumber
		sum:= firstNumber+ secondNumber+thirdNumber
		difference:= (firstNumber-secondNumber)-thirdNumber
		average := float64(sum)/3
		fmt.Printf("The product of the three numbers is %d\n",product)
		fmt.Printf("The sum of the three numbers is %d\n", sum)
		fmt.Printf("The difference of the three numbers is %d\n", difference)
		fmt.Printf("The average of the three numbers is %.2f\n", average)
	*/

	//Example6:Write a complete program that collects input from the user and calculates the product/sum/average of three integers
	/*	fmt.Print("Enter firstNumber: ")
		var firstNumber int64
		fmt.Scanf("%d", &firstNumber)
		fmt.Print("Enter secondNumber: ")
		var secondNumber int64
		fmt.Scanf("%d", &secondNumber)
		fmt.Print("Enter thirdNumber: ")
		var thirdNumber int64
		fmt.Scanf("%d", &thirdNumber)
		product:= firstNumber * secondNumber * thirdNumber
		sum:= firstNumber + secondNumber + thirdNumber
		average := float64(sum)/3
		fmt.Printf("Sum is %d\n", sum)
		fmt.Printf("Product is %d\n", product)
		fmt.Printf("Average is %.2f\n", average)
	*/

	//Example7:Write an application that asks the user to enter two integers,and displays the larger number followed by the words "is larger". If the numbers are equal, print the message "These numbers are equal"
	fmt.Printf("Enter the first Number")
}
