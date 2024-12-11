package main

import (
	"fmt"
)

	func bilanganSempurna(x int) bool {
		jumlah := 0
	
		for i := 1; i < x; i++ {
			if x%i == 0 {
				jumlah += i
			}
		}
	
		return jumlah == x
	}
	
	func main() {
		var a, b int 
	
		fmt.Print("Masukkan nilai a: ")
		fmt.Scan(&a)
		fmt.Print("Masukkan nilai b: ")
		fmt.Scan(&b)
	
		if a > b {
			fmt.Println("Nilai a harus lebih kecil atau sama dengan nilai b")
			return
		}
	
		fmt.Printf("Bilangan antara %d dan %d: ", a, b)
	
		ditemukan := false
		for i := a; i <= b; i++ {
			if bilanganSempurna(i) {
				fmt.Printf("%d ", i)
				ditemukan = true
			}
		}
	
		if !ditemukan {
			fmt.Println("Tidak ada bilangan yang ditemukan.")
		} else {
			fmt.Println()
		}
}