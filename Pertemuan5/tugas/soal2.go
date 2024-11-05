package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel untuk nama klub
	var klubA, klubB string
	fmt.Print("Masukkan nama klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scan(&klubB)

	// Meminta jumlah pertandingan dari pengguna
	var jumlahPertandingan int
	fmt.Print("Masukkan jumlah pertandingan: ")
	fmt.Scan(&jumlahPertandingan)

	// Untuk menyimpan hasil pertandingan
	var hasil []string
	pertandinganSelesai := false 

	// Untuk memasukkan skor berdasarkan jumlah pertandingan
	for i := 1; i <= jumlahPertandingan; i++ {
		var skorA, skorB int

		fmt.Printf("Masukkan skor %s untuk pertandingan %d: ", klubA, i)
		fmt.Scan(&skorA)
		fmt.Printf("Masukkan skor %s untuk pertandingan %d: ", klubB, i)
		fmt.Scan(&skorB)

		// Mengecek jika skor negatif
		if skorA < 0 || skorB < 0 {
			fmt.Println("\nSkor tidak valid. Pertandingan selesai.")
			pertandinganSelesai = true
			break 
		}

		// Menentukan pemenang berdasarkan skor
		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
	}

	// Menampilkan hasil pertandingan yang valid sebelum skor negatif
	fmt.Println("\nDaftar hasil pertandingan:")
	for i, winner := range hasil {
		if winner == "Draw" {
			fmt.Printf("Hasil %d: Draw\n", i+1)
		} else {
			fmt.Printf("Hasil %d: %s\n", i+1, winner)
		}
	}

	// Menampilkan pesan akhir jika pertandingan dihentikan karena skor negatif
	if pertandinganSelesai {
		fmt.Println("Pertandingan selesai.")
	}
}
