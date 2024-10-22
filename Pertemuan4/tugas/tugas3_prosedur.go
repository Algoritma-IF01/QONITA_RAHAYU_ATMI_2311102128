package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Struct untuk menyimpan informasi peserta
type Peserta struct {
	nama       string
	scores     []string 
	totalSoal  int
	totalWaktu int
}

// Fungsi untuk menghitung total soal yang diselesaikan dan total waktu dari skor yang diberikan
func hitungSkor(scores []string) (totalSoal int, totalWaktu int) {
	for _, scoreStr := range scores {
		var score int
		// Mengonversi string ke int untuk perhitungan
		fmt.Sscanf(scoreStr, "%d", &score) 
		// jika skor kurang dari 301 menit
		if score < 301 {
			totalSoal++
			totalWaktu += score
		}
	}
	return totalSoal, totalWaktu
}

// Fungsi untuk menentukan pemenang berdasarkan jumlah soal dan total waktu
func cariPemenang(peserta []Peserta) Peserta {
	// Inisialisasi pemenang dengan peserta pertama
	pemenang := peserta[0]

	for _, p := range peserta[1:] {
		// Pemenang ditentukan dari jumlah soal yang lebih banyak
		if p.totalSoal > pemenang.totalSoal {
			pemenang = p
			// Jika soal sama, cek total waktu
		} else if p.totalSoal == pemenang.totalSoal && p.totalWaktu < pemenang.totalWaktu {
			pemenang = p
		}
	}

	return pemenang
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var peserta []Peserta
	pesertaKe := 1
	maxPeserta := 3

	fmt.Println("Data Peserta ")

	// Input untuk diulang hingga maksimal peserta ke-3
	for pesertaKe <= maxPeserta {
		if pesertaKe < maxPeserta {
			// untuk menampilkan hanya untuk peserta 1 dan 2
			fmt.Printf("Masukkan data peserta %d : ", pesertaKe)
		} else {
			// Tidak menampilkan saat peserta ke-3
			fmt.Print("")
		}

		scanner.Scan()
		line := scanner.Text()

		if strings.ToLower(line) == "selesai" {
			break
		}

		// Split data input per peserta
		data := strings.Fields(line)

		// untuk ambil data nama peserta
		nama := data[0]

		// untuk menyimpan skor sebagai string
		var scores []string
		for _, s := range data[1:] {
			// Menyimpan skor sebagai string langsung
			scores = append(scores, s)
		}

		// untuk menghitung total soal yang diselesaikan dan total waktu yang dibutuhkan
		totalSoal, totalWaktu := hitungSkor(scores)

		// untuk menambahkan peserta ke list peserta
		peserta = append(peserta, Peserta{nama, scores, totalSoal, totalWaktu})

		// Increment nomor peserta
		pesertaKe++
	}

	// Jika tidak ada peserta, keluar
	if len(peserta) == 0 {
		fmt.Println("Tidak ada peserta.")
		return
	}

	// untuk mencari pemenang
	pemenang := cariPemenang(peserta)

	// Output hasil
	fmt.Printf("\nPemeneng : %s %d %d\n", pemenang.nama, pemenang.totalSoal, pemenang.totalWaktu)
}
