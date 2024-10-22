package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (a, b) dan (c, d)
func jarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

// Fungsi untuk menentukan apakah titik (x, y) berada di dalam lingkaran
func didalam(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

// Fungsi untuk menentukan posisi titik terhadap dua lingkaran
func posisiTitik(cx1, cy1, r1, cx2, cy2, r2, x, y int) string {
	dalamLingkaran1 := didalam(cx1, cy1, r1, x, y)
	dalamLingkaran2 := didalam(cx2, cy2, r2, x, y)

	// untuk menentukan posisi titik didalam atau di luar lingkaran
	if dalamLingkaran1 && dalamLingkaran2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if dalamLingkaran1 {
		return "Titik di dalam lingkaran 1"
	} else if dalamLingkaran2 {
		return "Titik di dalam lingkaran 2"
	}
	return "Titik di luar lingkaran 1 dan 2"
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int

	for i := 1; i <= 4; i++ {

		// input untuk lingkaran 1
		fmt.Printf("Lingkaran %d\n", i)
		fmt.Print("Masukkan pusat dan jari-jari lingkaran 1 : ")
		fmt.Scanf("%d %d %d\n", &cx1, &cy1, &r1)

		// input untuk lingkaran 2
		fmt.Print("Masukkan pusat dan jari-jari lingkaran 2 : ")
		fmt.Scanf("%d %d %d\n", &cx2, &cy2, &r2)

		// input untuk koordinat titik
		fmt.Print("Masukkan koordinat titik : ")
		fmt.Scanf("%d %d\n", &x, &y)

		// untuk menampilkan hasil posisi titik terhadap kedua lingkaran
		fmt.Println(posisiTitik(cx1, cy1, r1, cx2, cy2, r2, x, y))
		fmt.Println()
	}
}
