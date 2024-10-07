package main

import (
	"fmt"
)

func main() {
	const phi = 3.1415926535
	var r int

	fmt.Print("Jejari: ")
	fmt.Scan(&r)

	jariJari := float64(r)

	volume := (4.0 / 3.0) * phi * (jariJari * jariJari * jariJari)

	luas := 4 * phi * (jariJari * jariJari)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}