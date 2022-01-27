package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Print(i, " ")
		i++
	}

	//Example2: Print out the Odd and Even Numbers
	for j := 10; j >= 1; j-- {
		if j%2 == 0 {
			fmt.Println(j, "even")
		} else {
			fmt.Println(j, "odd")
		}

		//Example3: Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9,
		for i := 1; i <= 100; i++ {
			if i%3 == 0 {
				fmt.Print(i, " ")
			}
		}

		//Example4: Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five. Print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz".
		for i := 1; i <= 100; i++ {
			if i%3 == 0 {
				fmt.Println("Fizz")
			}
			if i%5 == 0 {
				fmt.Println("Buzz")
			}
			if i%3 == 0 && i%5 == 0 {
				fmt.Println("FizzBuzz")
			}
		}
	}
}
