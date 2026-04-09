package main

import (
	"fmt"
	"unsafe"

	"rsc.io/quote"
)

const iva float64 = 0.15

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

	//numeros
	numeroEntero := 22

	fmt.Printf("esta es una suma dentro de un  print %d + %d = %d\n", numeroEntero, 1, numeroEntero+1)

	//srings

	cadena := "esta es una cadena de texto!!!"

	fmt.Println(cadena)

	//arrays

	arrayStatic := [2]string{"HOLA", "MUNDO"}
	fmt.Printf("este es un mensa desde un array, %s %s\n", arrayStatic[0], arrayStatic[1])

	//mapas

	diccionario := map[string]int{
		"1": 1724492841,
		"2": 1728192711,
	}

	fmt.Printf("El usurio 1 tiene un ci con el número %d \n", diccionario["1"])

	//sumas enteras

	entero1 := 12
	entero2 := 1110

	suma := entero1 + entero2

	fmt.Printf("La suma entre el %d y %d es igual a: %d \n", entero1, entero2, suma)

	//suma flotante

	floatante1 := 12.15
	flotante2 := 1110.131

	sumaFloat := floatante1 + flotante2

	fmt.Printf("Esta si es una suma de flotante entre %f y %f es igual a: %f \n", floatante1, flotante2, sumaFloat)
	fmt.Printf("Y esta es un uso de una variable constante: %f y un numero floante %f, sumando da iigual a : %f\n", iva, flotante2, flotante2+iva)

	var array1 [3]int
	array1[0] = 2
	array1[1] = 6
	array1[2] = 123

	fmt.Printf("Array 1: %v\n", array1)

	fmt.Println("El tamaño del array1 es :", len(array1))
	fmt.Printf("Tipo: %T, byets: %d , bits: %d\n", array1, unsafe.Sizeof(array1), unsafe.Sizeof(array1)*8)

	slice := array1[1:2]
	fmt.Println(slice)
}
