package main
import "fmt"
func main(){
	var x int
	var y float64
	fmt.Println("enter a number: " )
	fmt.Scanf("%d",&x)
	fmt.Print("enter the number of feet: ")
	fmt.Scanf("%d",&y)
	output := x * 2
	var a float64 = y * 0.3048
	fmt.Println(a)
	fmt.Println(output)
	output += 1
	fmt.Println(output)
	z := 10
	if z > 10{
		fmt.Println("Big")
	} else{
              fmt.Println("small")
	}

}
