package main

import ("fmt"
	"rsc.io/quote"
)


 func main() {
	fmt.Println("hola mundo ")
	
	fmt.Println(quote.Hello() )
 }

// const Pi float32 = 3.14 // esta es una constante -- se pueden declarar y no usar a diferencia de las variables las cuales no se puede dado que se tiene que usar
// const(
// 	x = 3.12
// 	y =001
// 	// se pueden declarar las variables y no usarlas 
// )
// const(
// 	Domingo =iota +1 // aumenta en uno los valores 
// 	Lunes 
// 	Martes 
// 	Miercoles
// 	Jueves
// 	Viernes
// 	Sabado

// )

// 	// declaracion de variables
// firstName , lastName , age := "alex" , "reld" , 8 
// age = 30 // declaracion de nueva variable 
// // var es global mientras := es por bloque en la funcion 
// fmt.Println(firstName,lastName,age)
// fmt.Println(Pi)
// fmt.Println(x,y)
}