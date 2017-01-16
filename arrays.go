package main
import "fmt"
func main(){
	var x [5]float64
	x[0] = 43
	x[1] = 50
	x[2] = 79
	x[3] = 11
	x[4] = 2
	var avg float64
	for i:= 0; i < 5 ; i++{
		avg = avg + x[i]

	}
	output := avg / 5
	fmt.Println(output)
}
