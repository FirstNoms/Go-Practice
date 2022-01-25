package main

import (
	"fmt"
	"math"
)

/*-->Structs are a type in Go which contains names fields
  -->To call on the fields of the Struct, we'll create a new insance of the struct object
*/
//type Female struct{
//	name string
//	age int64
//	height float64
//}
//Example2:
//This is a type Circle with its instance variables
type Circle struct {
	lenght float64
	height float64
	radius float64
}

//This is a function cirleArea that takes in a variable of type Circle. This enables the variable(oval) have a copy of everything in type Circle and returns a float.
//func circleArea(oval Circle) float64{
// return math.Pi * oval.radius * oval.radius
//}

func (oval *Circle) area() float64 {
	return math.Pi * oval.radius * oval.radius
}

func main() {
	/*	firstLady := Female{"Ada", 20,6.10}
		fmt.Println(firstLady.name)
	*/

	//Example2:
	oval := &Circle{8.0, 6.9, 19.0}
	fmt.Println(oval.area())

}
