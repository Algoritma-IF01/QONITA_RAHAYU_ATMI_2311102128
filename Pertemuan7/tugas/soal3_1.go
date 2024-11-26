package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk insertion sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// Untuk menggeser elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Fungsi untuk memeriksa apakah array memiliki jarak tetap
func isJarakTetap(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0
	}

	jarak := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != jarak {
			return false, 0
		}
	}
	return true, jarak
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan data :")

	// Untuk Membaca input
	scanner.Scan()
	line := scanner.Text()

	// Untuk Memisahkan input menjadi array string
	parts := strings.Fields(line)

	// Array untuk menyimpan bilangan non-negatif
	var arr []int
	for i := 0; i < len(parts); i++ {
		part := parts[i]
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Input tidak valid.")
			return
		}
		// Untuk menghentikan jika bilangan negatif ditemukan
		if num < 0 {
			break
		}
		arr = append(arr, num)
	}

	// Untuk sorting menggunakan insertion sort
	insertionSort(arr)

	// Untuk Tampilkan array setelah diurutkan
	fmt.Println("Array setelah diurutkan:", arr)

	// Untuk mengecek apakah array memiliki jarak tetap
	tetap, jarak := isJarakTetap(arr)
	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}