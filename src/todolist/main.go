package main

import "fmt"

func main() {
	var a = [...]int{5 , 20, 30, 40, 5}
	
	//  modificar los valores de la matriz 
	 a[0] = 10
	 a[1]= 4
	 // asi con los demas elementos 
    fmt.Println(a) 

// for  in 
for index,  value := range a {
	fmt.Printf("indicd :%d, valor : %d\n",index,value)
}
// si deseamos crear una matriz en dos d
var matriz = [3] [3] int{{1,1,1} ,{2,2,2} , {3,3,3}}
fmt.Println(matriz)
}
// creation of matriz 