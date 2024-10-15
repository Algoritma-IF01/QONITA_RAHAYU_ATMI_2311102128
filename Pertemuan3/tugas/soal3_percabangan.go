package main

import (
	"fmt"
)

func hitungBiayaPos(berat int) (kg, sisaGram, biayaKg, biayaSisaGram int, biayaDitambahkan bool) {
	const biayaPerKg = 10000
	kg = berat / 1000
	sisaGram = berat % 1000
	biayaKg = kg * biayaPerKg
	biayaDitambahkan = true 

	// jika berat total >= 10kg, sisa gram tidak menambah biaya total
	if kg >= 10 {
		biayaSisaGram = sisaGram * 5
		biayaDitambahkan = false    
	} else {
		// menghitung biaya untuk sisa gram jika berat kurang dari 10kg
		if sisaGram >= 500 {
			biayaSisaGram = sisaGram * 5
		} else {
			biayaSisaGram = sisaGram * 15
		}
	}
	return
}

func main() {
	for contoh := 1; contoh <= 3; contoh++ {
		var berat int
		fmt.Printf("\nContoh #%d", contoh)
		fmt.Print("\nInput Berat untuk (gram): ")
		fmt.Scan(&berat)

		kg, sisaGram, biayaKg, biayaSisaGram, biayaDitambahkan := hitungBiayaPos(berat)
		totalBiaya := biayaKg
		if biayaDitambahkan {
			totalBiaya += biayaSisaGram
		}

		// Tampilkan output sesuai dengan format di gambar
		fmt.Printf("Berat parsel (gram): %d\n", berat)
		fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisaGram)
		fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
	}
}
