package main
import "fmt"
func average(sf ... float64)float64{
	total := 0.0
	for _,v := range sf {
		total += v
	}
	return total / float64(len(sf))

}
func main()  {
   data :=[] float64{43,43,23,54,54}
   n := average(data...)
   fmt.Println(n)
	
}







// package main

// import "fmt"

// func main() {
// 	var x  int =10
// 	fmt.Println(x)
// 	editar(&x)

// 	// var x int = 10
// 	// var p *int = &x
// 	// fmt.Println(&x)
// 	// fmt.Println(p)
// }
// func editar (x*int){
// 	*x = 20   
// }