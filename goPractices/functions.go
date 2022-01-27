package main

import "fmt"

//FUNCTIONS OUTSIDE MAIN
func f() (int, int) {
	return 5, 6
}

//Variadic Functions: that takes in as many arguments or a list of varying arguments
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

//VARIADIC FUNCTION 2
func smallestAndLargest(numbers ...int) int {
	largest := 0
	for _, number := range numbers {
		if largest < number {
			largest = number
		}
	}
	fmt.Println(largest, "is the largest number")
	return largest
}

func first() {
	fmt.Println("I'm Chinomso")
}
func second() {
	fmt.Println("Whats yor name?")
}

//Function that takes an Array as an argument.
func sum(array1 []float64) float64 {
	var newSum float64
	for _, value := range array1 {
		newSum += value
	}
	return float64(newSum)
}

//Function that determines an Odd and Even number
func half(number *int64) int64 {
	if *number%2 == 0 {
		fmt.Println("Even number")
		return 1
	} else {
		fmt.Println("Odd number")
		return 0
	}
}

//Recurssion function: This is when a function calls itself inside itself repeatedly till its condition is met.
func factorial(number uint) uint {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {
	x, y := f()
	fmt.Println(x, y)

	xs := []int{1, 2, 3, 7, 8, 9, 7}
	fmt.Println(add(xs...))

	//Example1: function inside function main (CLOSURE): i.e function calling a function.
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(20, 50))

	//Example2: Demonstrating DEFER, PANIC and RECOVER function keywords in Go, from the functions outside the main.
	defer second()
	first()

	//Example3: sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in GO?
	array1 := []float64{20, 30, 40, 50, 60, 23, 43, 56, 67, 89}
	array2 := make([]float64, 5)
	copy(array2, array1)
	fmt.Println(array2)
	fmt.Println(sum(array2))

	//Example4: Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true)
	fmt.Println("Enter any Number: ")
	var userInput int64
	fmt.Scanf("%d", &userInput)
	(half(&userInput))

	//fmt.Println(half(120))

	//Exaample5: Panic and Recover functions
	//The call to Panic stops the execution of the next function recover()
	//The recover  stops the panic and returns the value that was passed to call the panic.
	panic("STOP IT")
	str := recover()
	fmt.Println(str)

	//Example6: Defer Panic and Recover:
	// using the defer function allows
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("STOP IT")

	//Example7: Recussive call. This is when a function calls itself inside itself repeatedly till its condition is met.
	fmt.Println(factorial(5))

	//Example8:Write a function with one variadic parameter that finds the greatest number in a list of numbers.
	fmt.Println(smallestAndLargest(2, 4, 6, 8, 10, 12, 14, 16, 18, 20))

}
