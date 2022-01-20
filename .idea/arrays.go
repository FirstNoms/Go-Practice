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
	/*	valu := [10]float64{98, 93, 77, 82, 83}
		total := float64(0)
		for _, value := range valu {
			total += value
		}
		average := total / float64(len(valu))
		fmt.Printf("%.f is the total\n", total)
		fmt.Printf("%.1f is the average\n", average)
	*/

	//Example4: Still on enhanced for loop in strings.
	/*	names := [5]string{"Nomso", "Tobi","Janet","Bless", "Love"}
		for _, value:=range names{
			fmt.Println(value)
		}
	*/

	//Example5: Still using enhanced for loop.
	/*	numbs:=[10]float64{20,34,45,56,67,77,89,98,76,32}
		var sum float64
		var product float64 = 1.0
		var average float64
		for _, value :=range numbs{
			sum+= value
			product *= value
			average = sum/float64(len(numbs))
		}
		fmt.Printf("%.2f is sum of the array\n", sum)
		fmt.Printf("%.2f is product of the array\n", product)
		fmt.Printf("%.2f is the average of the array\n", average)
	*/

	//Example6: Slicing an array.
	/*	numbs:=[]float64{20,34,45,56,67,77,89,98,76,32}
		one:= numbs[3:8]
		fmt.Println(one)
	*/

	//Example7: Slicing and apending an array
	nums := []float64{20, 34, 45, 56, 67, 77, 89, 98, 76, 32}
	two := append(nums, 500, 100, 200)
	fmt.Println(nums, two)
}
