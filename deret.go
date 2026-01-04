package main

import (
	"fmt"
	"math"
)

// Soal 5: Relasi Rekurens Iteratif
func SolveSoal5(c1, c2, n int) {
	a := make([]int, n+1)
	a[0] = 1 // Sesuai deskripsi a0=1
	a[1] = 1 // Sesuai deskripsi a1=1

	fmt.Printf("INPUT: C1=%d, C2=%d, N=%d\n", c1, c2, n)
	fmt.Println("Proses Perhitungan:")
	fmt.Printf("Suku 0: %d\nSuku 1: %d\n", a[0], a[1])

	for i := 2; i <= n; i++ {
		a[i] = (c1 * a[i-1]) + (c2 * a[i-2])
		fmt.Printf("Suku %d: %d\n", i, a[i])
	}
	fmt.Printf("HASIL AKHIR Suku ke-%d: %d\n", n, a[n])
}

// Soal 6: Analisis Kedekatan Deret Geometri
func SolveSoal6(a float64, r float64, n int) {
	// Sum Berhingga Sn = a(1-r^n) / (1-r)
	sn := 0.0
	for i := 0; i < n; i++ {
		sn += a * math.Pow(r, float64(i))
	}

	// Sum Tak Hingga Sinf = a / (1-r)
	sInf := a / (1 - r)

	percentage := (sn / sInf) * 100

	fmt.Printf("Input Paket: a=%.2f, r=%.2f, N=%d\n", a, r, n)
	fmt.Printf("Sum Berhingga S(%d): %.2f\n", n, sn)
	fmt.Printf("Sum Tak Hingga S(inf): %.2f\n", sInf)
	fmt.Printf("Persentase Kedekatan: %.2f%%\n", percentage)
}
