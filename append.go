package main
import "fmt"
func main(){
	s :=[]float64{2,5,7,3,8,10}
	fmt.Println(s)
	s = append(s,5,9)
	fmt.Println(s)
	//copy
	s1 := []float64{60,40,20}
	fmt.Println(s1)
	copy(s1,s)
	fmt.Println(s)
	fmt.Println(s1)

}
