package main

//Functions outside the main
func f() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	/*	x,y:=f()
		fmt.Println(x,y)


		xs := []int{1,2,3,7,8,9,7}
		fmt.Println(add(xs...))

		//function inside function main: i.e function calling a function.
		add := func(x, y int) int {
			return x + y
		}
		fmt.Println(add(20,50))
	*/

}
