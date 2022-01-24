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

func swap(number1 *int, number2 *int) {
	*number1 = *number2
	*number2 = *number1
}

func main() {
	//In order to refer to the variable in the memory, use the " & " in front of the variable name
	/*	name:= "Adaoma"
		gender(&name)
		fmt.Println(name)
	*/

	//Example2:
	value := 1.5
	square(&value)
	fmt.Println(value)

	//Example3:
	/*	val1:= 5
		val2:= 0
		swap(&val1, &val2)
		fmt.Println(val1, val2)
	*/
}
