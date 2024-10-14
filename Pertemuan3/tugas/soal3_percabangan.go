package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 3; i++ {
		var beratParsel, kg, gram int
		var biayaPerKg, biayaSisa, totalBiaya int

		// Input berat parsel dalam gram
		fmt.Printf("Contoh #%d:\n", i)
		fmt.Print("Berat parsel (gram): ")
		fmt.Scan(&beratParsel)

		// berat dalam kg dan sisa gram
		kg = beratParsel / 1000
		gram = beratParsel % 1000

		//  biaya per kg 
		biayaPerKg = kg * 10000

		// biaya sisa berdasarkan gram
		if gram > 0 && kg < 10 {
			if gram <= 500 {
				biayaSisa = gram * 5
			} else {
				biayaSisa = gram * 15
			}
		} else {
			biayaSisa = 0 
		}

		// total biaya
		totalBiaya = biayaPerKg + biayaSisa

		// Output 
		fmt.Printf("Detail berat: %d kg + %d gram\n", kg, gram)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKg, biayaSisa)
		fmt.Printf("Total biaya: Rp. %d\n\n", totalBiaya)
	}

	fmt.Println("Proses selesai.")
}