package main

import (
	"fmt"
)

func main() {
	var menit, jam, biaya int
	var member bool
	var voucher string

	fmt.Print("Masukkan Durasi Jam: ")
	fmt.Scanln(&jam)
	fmt.Print("Masukkan Durasi Menit: ")
	fmt.Scanln(&menit)
	fmt.Print("Apakah Member (true/false): ")
	fmt.Scanln(&member)
	fmt.Print("Voucher: ")
	fmt.Scanln(&voucher)

	if member {
		biaya = 3500
	} else {
		biaya = 5000
	}
	TotalMenit := jam*60 + menit

	sebelumDiskon := (TotalMenit / 60) * biaya
	sisaMenit := TotalMenit % 60

	if sisaMenit > 10 {
		sebelumDiskon += (biaya * sisaMenit) / 60
	}
	setelahDiskon := sebelumDiskon

	if TotalMenit > 180 {
		digitTerakhir := voucher[len(voucher)-1]
		if digitTerakhir == '5' || digitTerakhir == '6' {
			setelahDiskon = int(float64(sebelumDiskon) * 0.9)
		}
	}

	fmt.Printf("Biaya sewa setelah diskon adalah: Rp %d.00\n", setelahDiskon)
}
