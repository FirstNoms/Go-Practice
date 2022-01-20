package main

import "fmt"

func main() {
	//Array Declaration
	//	VariableFunc || VariableName || ArrayLength || VariableType
	/*	var num [5]int
		num[4] = 100
		fmt.Println(num)
	*/

	//Example2: Assigning values to an array and finding its total and average
	/*	var fig[5] float64
		fig[0] = 98
		fig[1]= 93
		fig[2]= 77
		fig[3]= 82
		fig[4] = 83

		var total float64 = 0
		for i:=0; i<len(fig); i++{
			total += fig[i]
		}
		average:= total/float64(len(fig))
		fmt.Printf("%.1f is the Total\n", total)
		fmt.Printf("%.1f is the Average",average)
	*/

	//Example3: Using the enhanced for loop.
	valu := [10]float64{98, 93, 77, 82, 83}
	total := float64(0)
	for _, value := range valu {
		total += value
	}
	average := total / float64(len(valu))
	fmt.Printf("%.f is the total\n", total)
	fmt.Printf("%.1f is the average\n", average)
}
