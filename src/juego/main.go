package main

import (
	"fmt"
	"math/rand"
	
)

// creacion de funciones que toman valores
func main() {
jugar()
	// fmt.Println(rand.Intn(100))// es el rango  de el rando y la forma de llamarlo 
}

func jugar() {
	numeroAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int

	const maximo = 10
	for intentos < maximo {
		intentos++
		fmt.Printf("ingrese un numero(intentos restantes: %d):", maximo - intentos + 1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numeroAleatorio {
			fmt.Println("felicidades")
			return
		} else if numIngresado < numeroAleatorio {
			fmt.Println("el numero a adivinar es mayor ")
		} else if numIngresado > numeroAleatorio {
			fmt.Println(" el numero a adivinar es menor")
		}
	}
	fmt.Println("se acabaron el numero de intentos: ", numeroAleatorio)
}

//         saludo := hello ("alex")
// 		fmt.Println(saludo)
// 		sum,mult := calc(4,5)
// 		fmt.Println( " la suma es: ", sum)
// 		fmt.Println("la multiplicacion es: ", mult)
// }
// func hello (name string) string{
// 	return "hola"+ name
// }
// // funcion de calculadora solo una suma

// func  calc( a,b int) (int, int){
// 	suma := a+b
// 	mult := a*b
// return suma, mult
// }

//  for i := 1; i <= 10; i++{
// 	fmt.Println(i)
// 	// este es un codigo para parar el bucle
// 	if i == 5{
// 		break // continue elimina el valor que tenemos en la declaracion
// 	}
//  }

// // ciclo for
// var i int // inicializacion de la varible i
// for i <= 10 { // regla de orden para la variable
// 	fmt.Println(i) // declaraciond e impresion de la variable
// 	i++ // incremento
// }
// }

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// metodo de uso para elif o arboles de deciciones
// func main(){
// 	// if t := time.Now(); t.Hour() < 12 {
// 	// 	fmt.Println("dia ")
// 	// }else {
// 	// 	fmt.Println("noche")
// 	// }
// 	// declaracion swith
// 	os := runtime.GOOS
// 	switch os {
// 	case "windows":
// 		fmt.Println("go run windows")
// 	case "linux":
// 		fmt.Println("go run linux")
// 	default:
// 		fmt.Println("otro sistma os ")
// 	}

// 	// esta tambien es otra manera de corre el codigo que teniamos anteriormente
// 	switch t := time.Now(); {
// 	case t.Hour() < 12:
// 		fmt.Println(" dia")
// 	default:
// 		println("noche ")
// 	}
// }// en este caso estamos usando una forma de el if else que es llamada casos en los cuales colocamos los posibles casos y de estos dependeran las acciones que tomaremos
