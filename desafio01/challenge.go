package main

import "fmt"

var ebullitionK = 373.2

func main() {
	tempK := ebullitionK
	tempC := tempK - 273

	fmt.Printf("A temperatura em Kelvin é %g e em Celcius é %g°C", tempK, tempC)
}
