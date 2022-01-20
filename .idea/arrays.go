package main

import "fmt"

func main() {
	//Array Declaration
	//	VariableFunc || VariableName || ArrayLength || VariableType
	var num [5]int
	num[4] = 100
	fmt.Println(num)

	//Example2: Assigning values to an array and finding its total and average
	var fig [5]float64
	fig[0] = 98
	fig[1] = 93
	fig[2] = 77
	fig[3] = 82
	fig[4] = 83

	var total float64 = 0
	for i := 0; i < len(fig); i++ {
		total += fig[i]
	}
	average := total / float64(len(fig))
	fmt.Printf("%.1f is the Total\n", total)
	fmt.Printf("%.1f is the Average", average)

	//Example3: Using the enhanced for loop.
	valu := [10]float64{98, 93, 77, 82, 83}
	totall := float64(0)
	for _, value := range valu {
		totall += value
	}
	average = totall / float64(len(valu))
	fmt.Printf("%.f is the total\n", totall)
	fmt.Printf("%.1f is the average\n", average)

	//Example4: Still on enhanced for loop in strings.
	names := [5]string{"Nomso", "Tobi", "Janet", "Bless", "Love"}
	for _, value := range names {
		fmt.Println(value)
	}

	//Example5: Still using enhanced for loop.
	numbs := [10]float64{20, 34, 45, 56, 67, 77, 89, 98, 76, 32}
	var sum float64
	var product float64 = 1.0
	var averrage float64
	for _, value := range numbs {
		sum += value
		product *= value
		averrage = sum / float64(len(numbs))
	}
	fmt.Printf("%.2f is sum of the array\n", sum)
	fmt.Printf("%.2f is product of the array\n", product)
	fmt.Printf("%.2f is the average of the array\n", averrage)

	//Example6: Slicing an array.
	/*	numbs = []float64{20, 34, 45, 56, 67, 77, 89, 98, 76, 32}
		one := numbs[3:8]
		fmt.Println(one)
	*/

	//Example7: Slicing and apending an array
	nums := []float64{20, 34, 45, 56, 67, 77, 89, 98, 76, 32}
	two := append(nums, 500, 100, 200)
	fmt.Println(nums, two)

	//Example8:Using the Slice function suing the MAKE keyword in an array:
	numbe := []int{2, 4, 6, 8, 10}
	slice1 := make([]int, 3) //Creates a new int array of length 3
	copy(slice1, numbe)      //this copies the values of the oll array into the new array.
	fmt.Println(numbe, slice1)

	//Example9: Still on slicing and copying an array element into a new array,finding the sum of the new array, appenindind new elements into the new array:
	array := []int{20, 40, 60, 80, 100}
	array2 := make([]int, 4)
	copy(array2, array)
	array3 := append(array2, 100, 120, 140)
	//finding the sum of the arry2
	var summ int
	for _, value := range array2 {
		summ += value
	}
	fmt.Println(array, array2, "The sum of array2 is ", sum, array3)

	//Example10: Given the array below: find arry[2:5]?
	/*	arry:=[6]string{"a","b","c","d","e","f"}
		fmt.Println(arry[2:5])
	*/

	//Example11: find the smallest number in this list.
	figu := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	smallest := 100
	largest := 0
	for i := 0; i <= len(figu)-1; i++ {
		if figu[i] < smallest {
			smallest = figu[i]
		}
		if figu[i] > largest {
			largest = figu[i]
		}
	}
	fmt.Println(figu)
	fmt.Println("The smallest number is ", smallest)
	fmt.Println("The largest is ", largest)

}
