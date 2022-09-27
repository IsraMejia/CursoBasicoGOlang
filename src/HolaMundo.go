package main

import (
	"fmt"
)

func main() {
	//declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.141516

	//fmt  paquete  que se encarga de mostrar salidas en consola
	fmt.Println("Hello World, Hola mundo GO")
	fmt.Println("Salu2 rex \t Salu2 coco")

	//fmt.Println(pi, pi2)
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//declaracion de variables
	base := 12 //creadas en este momento
	base = 10  // asignacion posterior
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Cero Values nan or nun, Golang asigna un valor por defecto
	var a int     //0
	var b float64 //0
	var c string  //''
	var d bool    //false
	fmt.Println("Valores por defecto en int, float, string, bool")
	fmt.Println(a, b, c, d)

	//Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado //Se crea la variable y asigna su tipo de forma automatica
	fmt.Println("Area del cuadrado: ", areaCuadrado)
}
