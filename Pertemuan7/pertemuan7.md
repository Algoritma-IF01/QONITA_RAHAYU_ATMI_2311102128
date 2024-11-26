# <h1 align="center">Laporan Praktikum Modul 12 & 13-PENGURUTAN DATA</h1>

<h1 align="center">QONITA RAHAYU ATMI-2311102128</h1>

<h1>A. Soal Latihan</h1>

### 1. Latihan 1 Selection Sort

```go
package main

import "fmt"

type arrInt [4321]int

func selectionSort1(T *arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi n bilangan bulat
	   F.S. array T terurut secara ascending atau membesar dengan SELECTION SORT */
	for i := 0; i < n-1; i++ {
		// Inisialisasi indeks minimum
		idx_min := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[idx_min] {
				idx_min = j
			}
		}
		// Tukar elemen T[i] dengan T[idx_min] jika perlu
		if idx_min != i {
			T[i], T[idx_min] = T[idx_min], T[i]
		}
	}
}

func main() {
	// Contoh penggunaan
	var T arrInt
	n := 5
	T[0], T[1], T[2], T[3], T[4] = 64, 34, 25, 12, 22

	fmt.Println("Array sebelum diurutkan:", T[:n])
	selectionSort1(&T, n)
	fmt.Println("Array setelah diurutkan:", T[:n])
}
```

### Output Screenshot:

![Output ss_Latihan1!](assets/ss_latihan1.png)

### 2. Latihan 2 Selection Sort Struct

```go
package main

import "fmt"

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func selectionSort2(T *arrMhs, n int) {
	/* I.S. terdefinisi array T yang berisi n data mahasiswa
	   F.S. array T terurut secara ascending berdasarkan ipk dengan
	   menggunakan algoritma SELECTION SORT */

	var idx_min int
	var temp mahasiswa

	for i := 0; i < n-1; i++ {
		// Inisialisasi indeks minimum
		idx_min = i

		// Cari elemen dengan IPK terkecil di subarray [i+1, n-1]
		for j := i + 1; j < n; j++ {
			if T[j].ipk < T[idx_min].ipk {
				idx_min = j
			}
		}

		// Tukar elemen di indeks i dengan elemen di idx_min jika perlu
		if idx_min != i {
			temp = T[i]
			T[i] = T[idx_min]
			T[idx_min] = temp
		}
	}
}

func main() {
	// Contoh data mahasiswa
	var T arrMhs
	T[0] = mahasiswa{"Alice", "123", "A", "Teknik Informatika", 3.8}
	T[1] = mahasiswa{"Bob", "124", "B", "Sistem Informasi", 3.2}
	T[2] = mahasiswa{"Charlie", "125", "A", "Teknik Informatika", 3.5}
	T[3] = mahasiswa{"Diana", "126", "B", "Sistem Informasi", 3.9}
	n := 4

	fmt.Println("Data mahasiswa sebelum diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}

	selectionSort2(&T, n)

	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan IPK:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}
}
```

### Output Screenshot:

![Output ss_Latihan2!](assets/ss_latihan2.png)


### 3. Latihan 3 Insertion Sort

```go
package main

import "fmt"

type arrInt [4321]int

func insertionSort1(T *arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi n bilangan bulat
	   F.S. array T terurut secara mengecil (descending) dengan INSERTION SORT */
	var temp, i, j int

	for i = 1; i < n; i++ {
		temp = T[i] // Simpan elemen ke-i
		j = i       // Inisialisasi indeks pembanding

		// Geser elemen-elemen sebelumnya yang lebih kecil dari temp
		for j > 0 && temp > T[j-1] {
			T[j] = T[j-1]
			j--
		}

		// Tempatkan temp pada posisi yang sesuai
		T[j] = temp
	}
}

func main() {
	// Contoh penggunaan
	var T arrInt
	n := 5
	T[0], T[1], T[2], T[3], T[4] = 22, 12, 34, 64, 25

	fmt.Println("Array sebelum diurutkan:", T[:n])
	insertionSort1(&T, n)
	fmt.Println("Array setelah diurutkan secara descending:", T[:n])
}
```

### Output Screenshot:

![Output ss_Latihan3!](assets/ss_latihan3.png)

### 4. Latihan 4 Insertion Sort Struct

```go
package main

import "fmt"

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func insertionSort2(T *arrMhs, n int) {
	/* I.S. terdefinisi array T yang berisi n data mahasiswa
	   F.S. array T terurut secara mengecil (descending) berdasarkan nama
	   dengan menggunakan algoritma INSERTION SORT */
	var temp mahasiswa
	var i, j int

	for i = 1; i < n; i++ {
		temp = T[i] // Simpan elemen ke-i
		j = i       // Inisialisasi indeks pembanding

		// Geser elemen-elemen sebelumnya
		for j > 0 && temp.nama > T[j-1].nama {
			T[j] = T[j-1]
			j--
		}

		// Tempatkan temp pada posisi yang sesuai
		T[j] = temp
	}
}

func main() {
	// Contoh data mahasiswa
	var T arrMhs
	T[0] = mahasiswa{"Charlie", "125", "A", "Teknik Informatika", 3.5}
	T[1] = mahasiswa{"Alice", "123", "A", "Teknik Informatika", 3.8}
	T[2] = mahasiswa{"Bob", "124", "B", "Sistem Informasi", 3.2}
	T[3] = mahasiswa{"Diana", "126", "B", "Sistem Informasi", 3.9}
	n := 4

	fmt.Println("Data mahasiswa sebelum diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}

	insertionSort2(&T, n)

	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan nama (descending):")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}
}
```

### Output Screenshot:

![Output ss_Latihan4!](assets/ss_latihan4.png)

<h1>B. Tugas</h1>

### 1. tugas 1 Selection Sort

```go
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
```

### Output Screenshot:

![Output ss_soal1!](assets/ss_soal1.png)

### 2. tugas 2 Selection Sort

```go
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
```

### Output Screenshot:

![Output ss_soal2!](assets/ss_soal2.png)

### 3. Tugas 3 Insertion Sort

```go
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
```

### Output Screenshot:

![Output soal3!](assets/ss_soal3_1.png)

### 4. Tugas 4 Insertion Sort

```go
package main

import (
	"fmt"
)

const nMax = 7919

// Struct Buku
type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

// Struct DaftarBuku
type DaftarBuku struct {
	Pustaka  []Buku
	nPustaka int
}

// Fungsi untuk DaftarkanBuku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data buku ke-%d (ID Judul Penulis Penerbit Eksemplar Tahun Rating):\n", i+1)
		fmt.Scan(&buku.ID, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Eksemplar, &buku.Tahun, &buku.Rating)
		pustaka.Pustaka = append(pustaka.Pustaka, buku)
	}
	pustaka.nPustaka = n
}

// Fungsi untuk CetakTerfavorit
func CetakTerfavorit(pustaka DaftarBuku) {
	if pustaka.nPustaka == 0 {
		fmt.Println("Tidak ada buku di perpustakaan.")
		return
	}
	terfavorit := pustaka.Pustaka[0]
	for i := 1; i < len(pustaka.Pustaka); i++ {
		buku := pustaka.Pustaka[i]
		if buku.Rating > terfavorit.Rating {
			terfavorit = buku
		}
	}
	fmt.Printf("Buku terfavorit: %s oleh %s (%s, %d) - Rating: %d\n", terfavorit.Judul, terfavorit.Penulis, terfavorit.Penerbit, terfavorit.Tahun, terfavorit.Rating)
}

// Fungsi untuk UrutBuku
func UrutBuku(pustaka *DaftarBuku) {
	// Menggunakan Insertion Sort untuk mengurutkan berdasarkan rating
	for i := 1; i < len(pustaka.Pustaka); i++ {
		key := pustaka.Pustaka[i]
		j := i - 1
		for j >= 0 && pustaka.Pustaka[j].Rating < key.Rating {
			pustaka.Pustaka[j+1] = pustaka.Pustaka[j]
			j--
		}
		pustaka.Pustaka[j+1] = key
	}
}

// Fungsi untuk Cetak5Terbaru
func Cetak5Terbaru(pustaka DaftarBuku) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < pustaka.nPustaka; i++ {
		buku := pustaka.Pustaka[i]
		fmt.Printf("%s oleh %s (%s, %d) - Rating: %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating
func CariBuku(pustaka DaftarBuku, rating int) {
	// Menggunakan Binary Search untuk mencari rating
	low, high := 0, len(pustaka.Pustaka)-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka.Pustaka[mid].Rating == rating {
			buku := pustaka.Pustaka[mid]
			fmt.Printf("Buku ditemukan: %s oleh %s (%s, %d) - Rating: %d, Eksemplar: %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating, buku.Eksemplar)
			return
		}
		if pustaka.Pustaka[mid].Rating < rating {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku

	var n int
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	// Daftarkan buku
	DaftarkanBuku(&pustaka, n)

	// Cetak buku terfavorit
	CetakTerfavorit(pustaka)

	// Urutkan buku berdasarkan rating
	UrutBuku(&pustaka)

	// Cetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(pustaka)

	// Cari buku berdasarkan rating
	var rating int
	fmt.Print("Masukkan rating yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
```

### Output Screenshot:

![Output soal4!](assets/ss_soal4_1.png)

