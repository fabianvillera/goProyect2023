package main

import "fmt"

func main() {
	var arra1 = [3]int{12, 4, 5}
	println(arra1) // esta es la primera manera de declarar una tupla la cual hace uso de la variable en su estado normal

	arr2 := [2]int{2, 3}
	fmt.Println(arr2) // esta es la segunda manera de nosotros llamar a la funcion con la cual podemos emplear el uso de la forma de inicializar la funcion :=

	var cars = [4] string{"a","b","c","d"}
	println(cars) // se pueden usar los distintos tipos de tipos de datos en los arreglos recordemos que si se presentara mas de un tipod debemos especificarlo en la propia tabla 

	arr1:=[5]int{1:19,2:1}
	fmt.Println(arr1) // de esta manera tendremos dos tipos donde solo especificamos dos posciciones de todas las existentes recordemos que los arreglos en este lenguaje al igual que py comienzan por el 0 

}

