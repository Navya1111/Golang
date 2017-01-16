package main
import "fmt"
func main(){
     arr := []float64{ 30,50,40,60}
	fmt.Println(avg(arr))
	fmt.Println(swap("hello","hi"))

}
func avg (arr []float64) float64{
	val := 0.0
	for _ ,i:= range arr{
		val = val + i
	}
	return (val / float64(len(arr)))
}
func swap(a,b string) (string , string){
	return b,a
}