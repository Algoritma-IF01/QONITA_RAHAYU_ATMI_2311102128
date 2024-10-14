package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64

	for { 
		// Meminta input dari pengguna untuk kedua kantong
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		// Jika total berat kedua kantong lebih dari 150 kg, hentikan program
		if kantong1+kantong2 > 150 {
			fmt.Println("Proses selesai, total berat lebih dari 150 kg.")
			break 
		}

		// Jika salah satu berat kantong negatif, program berhenti
		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Berat kantong tidak boleh negatif. Proses dihentikan.")
			break 
		}

		// Menghitung selisih berat antara kedua kantong
		selisih := math.Abs(kantong1 - kantong2)
		// Mengecek apakah selisih berat lebih dari atau sama dengan 9
		oleng := selisih >= 9 

		// Output hasil apakah sepeda akan oleng atau tidak
		if oleng {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}
