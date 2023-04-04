package main

import "fmt"

//Fórmula: C = K - 273

var tempK = 150.0

func main() {

	tempC := tempK - 273

	fmt.Printf("O valor do ponto de ebulição da água em graus Celsius é %g ºC", tempC)

}
