package main
import "fmt"
func main(){
	//to access the 4 th element
	x := [5]int{1,5,0,6,22}
	fmt.Println(x[3])
	fmt.Println(x[2:5])
	//to access element in slice
	y := []int{4,7,75}
	fmt.Println(y[2])
	//3 here is the length of the slice
	z := make([]int, 3,9)
	fmt.Println(len(z))
	//to find the smallest number in the list of slice
	a := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17}
	min := a[0]
	//loc := 1
	for i := 1; i < len(a) ; i++{
		if a[i] < min{
			min = a[i]
			//loc = i + 1

		}
		//fmt.Println(min)
	}
	fmt.Println(min)



}
