package main

import (
	"fmt"
)

func main() {
	var celsius float64

	fmt.Print("Masukkan Temperatur Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := int((celsius * 9 / 5) + 32)

	reamur := int(celsius * 4 / 5)

	kelvin := int(celsius + 273.15)

	fmt.Printf("Temperatur Celsius: %d\n", int(celsius))
	fmt.Printf("Derajat Reamur: %d\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %d\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %d\n", kelvin)
}
