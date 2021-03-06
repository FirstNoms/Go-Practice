package main

import "fmt"

//Pointers refer to a location in memory where a value is stored rather than the value itself. A pointer is represented with an " * " in front of the data type.
func gender(name *string) {
	*name = "Nomso"
}

//Example2:
func square(value *float64) {
	*value = *value * *value
}

//Example3:
func swap(number1 *int, number2 *int) {
	var temp int
	temp = *number1
	*number1 = *number2
	*number2 = temp
}

func main() {
	//In order to refer to the variable in the memory, use the " & " in front of the variable name
	name := "Adaoma"
	gender(&name)
	fmt.Println(name)

	//Example2:
	value := 1.5
	square(&value)
	fmt.Println(value)

	//Example3: Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1
	val1 := 5
	val2 := 4
	swap(&val1, &val2)
	fmt.Println(val1, val2)

}
