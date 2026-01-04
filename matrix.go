package main

import "fmt"

// Soal 3: Perkalian dan Trace
func SolveSoal3(a, b [][]int, n int) {
	// Inisialisasi matriks hasil R
	r := make([][]int, n)
	for i := range r {
		r[i] = make([]int, n)
	}

	// Perkalian Matriks
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				r[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	// Hitung Trace
	trace := 0
	for i := 0; i < n; i++ {
		trace += r[i][i]
	}

	fmt.Println("Hasil Matriks R:")
	for _, row := range r {
		fmt.Println(row)
	}
	fmt.Printf("Trace: %d\n", trace)
}

// Soal 4: Transformasi Baris dan Nilai Maks
func SolveSoal4(m [][]int, n int) {
	fmt.Println("Matrix M (Generated):")
	for _, row := range m {
		fmt.Println(row)
	}

	// Tukar baris 0 dengan baris N-1
	m[0], m[n-1] = m[n-1], m[0]
	fmt.Printf("Menukar Baris 0 dan %d...\n", n-1)

	fmt.Println("Matrix M Terkini:")
	maxVal := m[0][0]
	posX, posY := 0, 0
	for i := 0; i < n; i++ {
		fmt.Println(m[i])
		for j := 0; j < n; j++ {
			if m[i][j] > maxVal {
				maxVal = m[i][j]
				posX, posY = i, j
			}
		}
	}
	fmt.Printf("Nilai Maksimum: %d ditemukan di Posisi (%d,%d)\n", maxVal, posX, posY)
}
