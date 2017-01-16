package main
import "fmt"
func main(){
	i :=1
	for i=1; i < 100;i += 1 {
		if i % 3 ==0 && i % 5 == 0{
			fmt.Println("FIZZBUZZ")
		}else if i % 3 == 0{
			fmt.Println("FIZZ")
		}else if i % 5 == 0{
			fmt.Println("BUZZ")

		}else {
			fmt.Println(i)
			//i = i + 1
		}

	}
	fmt.Println(i)


}
