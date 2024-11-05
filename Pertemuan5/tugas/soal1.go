package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menampilkan semua elemen dari array
func tampilkanArray(arr []int) {
	fmt.Println("Isi array:", arr)
}

// Fungsi untuk menampilkan elemen dengan indeks ganjil
func tampilkanIndeksGanjil(arr []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen dengan indeks genap
func tampilkanIndeksGenap(arr []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen dengan indeks kelipatan x
func tampilkanKelipatanX(arr []int, x int) {
	fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menghapus elemen pada indeks tertentu
func hapusElemen(arr []int, indeks int) []int {
	return append(arr[:indeks], arr[indeks+1:]...)
}

// Fungsi untuk menghitung rata-rata elemen dalam array
func hitungRataRata(arr []int) float64 {
	total := 0
	for _, nilai := range arr {
		total += nilai
	}
	return float64(total) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi
func hitungStandarDeviasi(arr []int) float64 {
	rataRata := hitungRataRata(arr)
	var total float64
	for _, nilai := range arr {
		total += math.Pow(float64(nilai)-rataRata, 2)
	}
	return math.Sqrt(total / float64(len(arr)))
}

// Fungsi untuk menghitung frekuensi dari suatu bilangan tertentu
func hitungFrekuensi(arr []int, bilangan int) int {
	jumlah := 0
	for _, nilai := range arr {
		if nilai == bilangan {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	arr := make([]int, N)
	fmt.Print("Masukkan ", N, " bilangan bulat: ")
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	for {
		// Menampilkan menu
		fmt.Println("\nPilih opsi:")
		fmt.Println("1. Tampilkan semua elemen array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("5. Menghapus elemen pada indeks tertentu")
		fmt.Println("6. Menghitung rata-rata elemen")
		fmt.Println("7. Menghitung standar deviasi elemen")
		fmt.Println("8. Menghitung frekuensi dari bilangan tertentu")
		fmt.Println("9. Keluar")
		fmt.Print("Pilihan Anda: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanArray(arr)
		case 2:
			tampilkanIndeksGanjil(arr)
		case 3:
			tampilkanIndeksGenap(arr)
		case 4:
			var x int
			fmt.Print("Masukkan angka kelipatan x: ")
			fmt.Scan(&x)
			tampilkanKelipatanX(arr, x)
		case 5:
			var indeksHapus int
			fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
			fmt.Scan(&indeksHapus)
			arr = hapusElemen(arr, indeksHapus)
			tampilkanArray(arr)
		case 6:
			rataRata := hitungRataRata(arr)
			fmt.Printf("Rata-rata dari elemen array: %.2f\n", rataRata)
		case 7:
			standarDeviasi := hitungStandarDeviasi(arr)
			fmt.Printf("Standar deviasi dari elemen array: %.2f\n", standarDeviasi)
		case 8:
			var bilangan int
			fmt.Print("Masukkan bilangan untuk menghitung frekuensinya: ")
			fmt.Scan(&bilangan)
			frekuensi := hitungFrekuensi(arr, bilangan)
			fmt.Printf("Frekuensi dari bilangan %d: %d kali\n", bilangan, frekuensi)
		case 9:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
