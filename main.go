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
	var valueOne = 30
	var valueTwo = 50
	var valueThree = 20
	product := valueOne * valueTwo * valueThree
	sum := valueOne + valueTwo + valueThree
	difference := (valueOne - valueTwo) - valueThree
	average := float64(sum) / 3
	fmt.Printf("The product of the three numbers is %d\n", product)
	fmt.Printf("The sum of the three numbers is %d\n", sum)
	fmt.Printf("The difference of the three numbers is %d\n", difference)
	fmt.Printf("The average of the three numbers is %.2f\n", average)

	//Example6:Write a complete program that collects input from the user and calculates the product/sum/average of three integers
	fmt.Print("Enter valueOne: ")
	var valueFour int64
	fmt.Scanf("%d", &valueFour)
	fmt.Print("Enter valueTwo: ")
	var valueFive int64
	fmt.Scanf("%d", &valueFive)
	fmt.Print("Enter valueThree: ")
	var valueSix int64
	fmt.Scanf("%d", &valueThree)

	product = int(valueFour * valueFive * valueSix)
	sum = int(valueFour + valueFive + valueSix)
	average = float64(sum) / 3
	fmt.Printf("Sum is %d\n", sum)
	fmt.Printf("Product is %d\n", product)
	fmt.Printf("Average is %.2f\n", average)

	//Example7:Write an application that asks the user to enter two integers,and displays the larger number followed by the words "is larger". If the numbers are equal, print the message "These numbers are equal"
	fmt.Printf("Enter your first Number: ")
	var valueSeven int
	fmt.Scanf("%d", &valueSeven)
	fmt.Printf("Enter your second Number: ")
	var valueEight int
	fmt.Scanf("%d", valueEight)
	if valueSeven > valueEight {
		fmt.Printf("%d is larger", valueOne)
	} else if valueEight > valueSeven {
		fmt.Printf("%d is larger")
	} else if valueSeven == valueEight {
		fmt.Printf("Both numbers are Equal")
	}

	//Example8: Write an application that reads an integer and determines and prints whether it’s odd or even.
	fmt.Printf("Enter a Number: ")
	var userInput int64
	fmt.Scanf("%d", &userInput)
	if userInput%2 == 0 {
		fmt.Printf("%d is an Even number", userInput)
	} else {
		fmt.Printf("%d is an Odd number", userInput)
	}

	//Example9: Calculate the Diameter, Circumference, Area of a Circle.
	fmt.Printf("Enter the Radius of the Circle: ")
	var radius int64
	fmt.Scanf("%d", &radius)
	diameter := 2 * radius
	circumference := 2 * (3.14159 * float64(radius))
	area := 3.14159 * float64(radius*radius)
	fmt.Printf("%d is the Diameter\n", diameter)
	fmt.Printf("%.2f is the Circumference\n", circumference)
	fmt.Printf("%.2f is the Area\n", area)

}
