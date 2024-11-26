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
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scanln(&n)

	// Untuk Validasi jumlah baris
	if n <= 0 || n >= 1000 {
		fmt.Println("Jumlah baris harus antara 1 dan 999.")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for i := 1; i <= n; i++ {
		fmt.Printf("Masukkan angka untuk baris %d: ", i)
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)

		var ganjil, genap []int

		// Untuk memisahkan angka ganjil dan genap 
		for i := 0; i < len(parts); i++ {
			part := parts[i]
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Input harus berupa angka.")
				return
			}
			if num%2 == 0 {
				genap = append(genap, num)
			} else {
				ganjil = append(ganjil, num)
			}
		}

		// untuk mengerutkan angka ganjil dan genap
		selectionSort(ganjil)
		selectionSort(genap)

		// Untuk menampilkan hasil nilai ganjil dan genap
		for i := 0; i < len(ganjil); i++ {
			fmt.Printf("%d ", ganjil[i])
		}
		for i := 0; i < len(genap); i++ {
			fmt.Printf("%d ", genap[i])
		}
		fmt.Println()
	}
}