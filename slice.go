package main
import "fmt"
func main(){
	//var slice []int
	x := []float64{4,5,2,8,6}        /*appending from one slice to another*/
	z := append(x,3,10)
	fmt.Println(x,z)
	a := [5]float64{1,5,0,44,66}
	slice1 := a[0:]
	fmt.Println(a)
	fmt.Println(slice1)
	 var y int
	y = 6
	fmt.Println(y)
        /*copying from one slice to another*/
	sli := []float64{20,30,50}
	sli1 := make([]float64,4)
	copy(sli1,sli)
	fmt.Println(sli,sli1)

}
