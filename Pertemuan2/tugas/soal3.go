package main

import (
	"fmt"
)

func main() {
	var angka1, angka2, angka3, angka4, angka5 int
	fmt.Print("Masukkan 5 angka integer antara 32-127: ")
	fmt.Scan(&angka1, &angka2, &angka3, &angka4, &angka5)

	fmt.Print("Karakter dari angka-angka yang dimasukkan: ")
	fmt.Printf("%c%c%c%c%c\n", angka1, angka2, angka3, angka4, angka5)

	var input string
	fmt.Print("Masukkan 3 karakter abjad ASCII: ")
	fmt.Scan(&input)

	if len(input) > 3 {
		fmt.Println("Error: tidak boleh lebih dari 3 karakter abjad ASCII.")
		return
	}

	fmt.Print("Karakter abjad ASCII yang dimasukkan dan setelahnya: ")
	fmt.Printf("%c%c%c\n", input[0]+1, input[1]+1, input[2]+1)
}
