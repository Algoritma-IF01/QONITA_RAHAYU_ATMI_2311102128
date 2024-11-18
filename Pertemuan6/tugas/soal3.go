package main

import (
	"fmt"
)

// Tipe data untuk array berat balita
type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	// Inisialisasi nilai awal minimum dan maksimum
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	// Iterasi untuk mencari nilai minimum dan maksimum
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0

	// Menjumlahkan semua berat balita
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	// Mengembalikan hasil rata-rata
	return total / float64(n)
}

func main() {
	var berat arrBalita    
	var jumlah int         
	var min, max float64   
	
	// Input jumlah balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&jumlah)

	// Input berat balita
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Hitung nilai minimum, maksimum, dan rata-rata
	hitungMinMax(berat, jumlah, &min, &max)
	rata := rerata(berat, jumlah)

	// Tampilkan hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
