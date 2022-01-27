package main

//Declaring an Interface
//An interface contains a "method set" to be implemented by another type.
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area += shape.area()
	}
	return area
}

//Interface can also be used as fields:
type MultiShape struct {
	shapes []Shape
}

func main() {
	//fmt.Println(totalArea(&radius, &height))
}
