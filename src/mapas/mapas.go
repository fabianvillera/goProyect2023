package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "ff000",
		"verde": "00ff00",
		"azul":  "0000ff",
	}
	fmt.Println(colors) // para llamar a todos los datos de nuestro mapa con sus respectivos valores 
	fmt.Println(colors["rojo"])
	// agregar un nuevo elemento 
	colors["negro"] = "00000" // para agregar un nuevo elemento por lo general debemos realizar los siguientes pasos para poder encontrarnos con el mismo primero debemos definir el nombre de nuestro mapa luego debemos definir el valor de la llave a la cual deseamos colocar y el valor esto entre comillas dado que si no se hace de esta manera pasara un error 

// 	// verificacion de la existencia de valores 
//       valor ,ok :=[]
//     fmt.Println(valor, ok)
	
// 
}