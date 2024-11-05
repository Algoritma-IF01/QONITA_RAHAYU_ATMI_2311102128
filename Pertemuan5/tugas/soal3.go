package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
// Batas maksimum elemen array
const NMAX int = 127 

// tipe data tabel untuk menyimpan karakter
type tabel [NMAX]rune 

// Mengisi array `t` dengan karakter dari input, dan mengabaikan akhiran " ."
func isiArray(t *tabel, n *int) {
	var input string
	fmt.Print("Teks : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	input = strings.ToUpper(input)   
	input = strings.TrimSuffix(input, " .") 

	for i, char := range input {
		if i >= NMAX {
			break
		}
		t[i] = char
		*n++
	}
}

// Mencetak isi array `t` sebanyak `n` karakter
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

// Membalik urutan elemen dalam array `t`
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Memeriksa apakah array `t` adalah palindrom
func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)
	fmt.Print("Reverse teks : ")
	balikanArray(&tab, n)
	cetakArray(tab, n)

	balikanArray(&tab, n)
	if palindrome(tab, n) {
		fmt.Println("Palindrom ? true")
	} else {
		fmt.Println("Palindrom ? false")
	}
}
