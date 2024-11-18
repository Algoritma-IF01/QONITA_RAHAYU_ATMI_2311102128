package main

import (
	"fmt"
)

func main() {
	// Input jumlah ikan (x) dan kapasitas wadah (y)
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	// Input berat ikan secara satu per satu
	weights := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	// Mengelompokkan ikan ke dalam wadah
	var containers [][]float64
	for i := 0; i < x; i += y {
		end := i + y
		if end > x {
			end = x
		}
		containers = append(containers, weights[i:end])
	}

	// Menghitung total berat ikan di setiap wadah
	totalWeights := []float64{}
	for _, container := range containers {
		total := 0.0
		for _, weight := range container {
			total += weight
		}
		totalWeights = append(totalWeights, total)
	}

	// Menghitung rata-rata berat ikan di setiap wadah
	averageWeights := []float64{}
	for _, container := range containers {
		total := 0.0
		for _, weight := range container {
			total += weight
		}
		averageWeights = append(averageWeights, total/float64(len(container)))
	}

	// Output total berat ikan di setiap wadah dengan format yang diminta
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, total := range totalWeights {
		fmt.Printf("Berat wadah ke-%d: %.2f\n", i+1, total)
	}

	// Output rata-rata berat ikan di setiap wadah
	totalSum := 0.0
	for _, total := range totalWeights {
		totalSum += total
	}
	averageTotal := totalSum / float64(len(totalWeights))
	fmt.Printf("Berat rata-rata ikan di setiap wadah: %.2f\n", averageTotal)
}
