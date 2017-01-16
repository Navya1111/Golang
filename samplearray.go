package main
import "fmt"
func main() {
	/*x := [5]float64{20,30,40,50,60}
	for  e, b :=  range x {
		fmt.Println(e,b)
		//fmt.Println(x)
	}*/
	y := 5
	z := 2
	c := y
	y = z
	z = c
	fmt.Println(z)
	fmt.Println(y)
	a, b := swap(5, 60)
	fmt.Println(a, b)
	d := make(map[string]string)
	d["water"] = "hi"
	d["air"] = "hello"
	d["snow"] = "how"
	d["drop"] = "are"
	d["grow"] = "you"
	if one, can := d["air"]; can {
		fmt.Println(one, can)
	}
	//fmt.Println(one, can)
	f := []float64 {5,9,20,67,89,24}
	val := 0.0
	for range f {
		val += f
	}
	fmt.Println(val/ float64(len(f)))



}

func swap(a int ,b int) (int,int){
     return b,a
}



