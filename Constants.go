package main
import "fmt"
func main(){
	const num int = 60000
	fmt.Println(num)
	var num1 int = 500 * num
	fmt.Println(num1)
	d := num1 /num
	fmt.Print(d)
}