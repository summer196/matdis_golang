package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Inisialisasi seed

	// --- PARAMETER PAKET 31 ---
	const (
		S1_N               = 100
		S2_M               = 5
		S2_K               = 10
		S3_N               = 3
		S4_N               = 3
		S5_C1, S5_C2, S5_N = 2, 1, 8
		S6_A, S6_R, S6_N   = 5.0, 0.5, 10
	)

	fmt.Println("=== MATERI I: HIMPUNAN ===")
	// Generate Himpunan secara acak
	genSet := func(size int, limit int) []int {
		set := []int{}
		for len(set) < size {
			val := rand.Intn(limit) + 1
			if !contains(set, val) {
				set = append(set, val)
			}
		}
		return set
	}

	setA := genSet(5, S1_N)
	setB := genSet(5, S1_N)
	setC := genSet(5, S1_N)
	SolveSoal1(setA, setB, setC, S1_N)

	fmt.Println("\n--- Soal 2 ---")
	setS := genSet(S2_M, 20)
	SolveSoal2(setS, S2_K)

	fmt.Println("\n=== MATERI II: MATRIKS ===")
	genMatrix := func(n int) [][]int {
		m := make([][]int, n)
		for i := range m {
			m[i] = make([]int, n)
			for j := 0; j < n; j++ {
				m[i][j] = rand.Intn(10) + 1
			}
		}
		return m
	}

	matrixA := genMatrix(S3_N)
	matrixB := genMatrix(S3_N)
	SolveSoal3(matrixA, matrixB, S3_N)

	fmt.Println("\n--- Soal 4 ---")
	matrixM := genMatrix(S4_N)
	SolveSoal4(matrixM, S4_N)

	fmt.Println("\n=== MATERI III: FUNGSI DAN DERET ===")
	SolveSoal5(S5_C1, S5_C2, S5_N)

	fmt.Println("\n--- Soal 6 ---")
	SolveSoal6(S6_A, S6_R, S6_N)
}
