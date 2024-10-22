package main

import "fmt"

// Fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Fungsi h(x) = x + 1
func h(x int) int {
	return x + 1
}

// Fungsi fogoh(x) = f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x)))
}

// Fungsi gohof(x) = g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x)))
}

// Fungsi hofog(x) = h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	for i := 1; i <= 3; i++ {
		var a, b, c int
		fmt.Print("Masukkan tiga bilangan (a b c): ")
		fmt.Scan(&a, &b, &c)

		// Cetak hasil sesuai dengan komposisi fungsi
		fmt.Printf("fogoh (%d)  : %d\n", a, fogoh(a))
		fmt.Printf("gohof (%d)  : %d\n", b, gohof(b))
		fmt.Printf("hofog (%d)  : %d\n", c, hofog(c))
	}
}
