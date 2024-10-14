package main

import "fmt"

func main() {
	// perulangan untuk meminta input dua bilangan
	for j := 0; j < 2; j++ {
		var bilangan int

		// Meminta input bilangan
		fmt.Print("Bilangan: ")
		fmt.Scan(&bilangan)

		// Mencari faktor
		fmt.Print("Faktor: ")
		for i := 1; i <= bilangan; i++ {
			if bilangan%i == 0 {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()

		// Menentukan bilangan prima
		var jumlahFaktor int = 0
		for i := 1; i <= bilangan; i++ {
			if bilangan%i == 0 {
				jumlahFaktor++
			}
		}

		if jumlahFaktor == 2 {
			fmt.Println("Prima: true")
		} else {
			fmt.Println("Prima: false")
		}

		fmt.Println() 
	}
}
