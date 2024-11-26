package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi selection sort untuk mengurutkan array secara ascending
func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Untuk Input jumlah daerah
	fmt.Print("Masukkan jumlah daerah: ")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil || n <= 0 {
		fmt.Println("Input tidak valid: jumlah daerah harus berupa angka positif.")
		return
	}

	for i := 0; i < n; i++ {
		// Untuk Input jumlah rumah dan nomor rumah
		fmt.Printf("Masukkan jumlah rumah dan nomor rumah untuk daerah %d: ", i+1)
		scanner.Scan()
		parts := strings.Fields(scanner.Text())

		// Untuk Konversi jumlah rumah
		m, err := strconv.Atoi(parts[0])
		if err != nil || m <= 0 {
			fmt.Println("Input tidak valid: jumlah rumah harus berupa angka positif.")
			return
		}

		// Untuk Membaca nomor rumah
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			rumah[j], err = strconv.Atoi(parts[j+1])
			if err != nil {
				fmt.Println("Input tidak valid: nomor rumah harus berupa angka.")
				return
			}
		}

		// Untuk Mengurutkan nomor rumah
		selectionSort(rumah)

		// Untuk Menampilkan hasil pengurutan
		fmt.Printf("Nomor rumah yang sudah diurutkan untuk daerah %d: ", i+1)
		for _, r := range rumah {
			fmt.Printf("%d ", r)
		}
		fmt.Println()
	}
}