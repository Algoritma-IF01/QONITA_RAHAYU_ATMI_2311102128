package main

import "fmt"

// Fungsi untuk mencetak deret bilangan sesuai aturan tertentu
func cetakDeret(n int) {
    // Jika nilai n belum mencapai 1, terus berlanjut
    for n != 1 {
        // Cetak bilangan n
        fmt.Printf("%d ", n) 
        // Apakah n habis dibagi 2 atau tidak
        if n%2 == 0 { 
            // Jika n habis dibagi 2, maka n dibagi 2   
            n = n / 2
        // Ketika n tidak habis dibagi 2
        } else { 
            // Maka n dikali 3ditambah 1
            n = 3*n + 1
        }
    }
    // Cetak angka terakhir yaitu 1
    fmt.Printf("1\n") 
}

func main() {
    var n int

    // Meminta pengguna memasukkan angka
    fmt.Print("Masukkan Bilangan: ")
    // Membaca input dari pengguna
    fmt.Scanln(&n) 

    // Ketika n lebih besar dari 0 dan n kurang dari 1 juta maka melakukan fungsi cetak deret
    if n > 0 && n < 1000000 {
        cetakDeret(n) 
    } else {
        fmt.Println("Bilangan harus lebih kecil dari 1000000 dan bilangan positif.") 
    }
}
