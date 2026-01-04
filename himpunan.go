package main

import "fmt"

// Fungsi pembantu untuk mengecek apakah elemen ada di slice
func contains(set []int, element int) bool {
	for _, v := range set {
		if v == element {
			return true
		}
	}
	return false
}

// Irisan (A ∩ C)
func intersection(a, c []int) []int {
	var result []int
	for _, v := range a {
		if contains(c, v) {
			result = append(result, v)
		}
	}
	return result
}

// Gabungan (A ∪ B)
func union(a, b []int) []int {
	result := a
	for _, v := range b {
		if !contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// Selisih (A \ B)
func difference(a, b []int) []int {
	var result []int
	for _, v := range a {
		if !contains(b, v) {
			result = append(result, v)
		}
	}
	return result
}

// Selisih Simetri (A ⊕ B)
func symmetricDifference(a, b []int) []int {
	diff1 := difference(a, b)
	diff2 := difference(b, a)
	return union(diff1, diff2)
}

// Soal 1: R = ((A ⊕ B) \ C) ∪ (A ∩ C)
func SolveSoal1(a, b, c []int, n int) {
	symDiff := symmetricDifference(a, b) // (A ⊕ B)
	firstPart := difference(symDiff, c)  // ((A ⊕ B) \ C)
	secondPart := intersection(a, c)     // (A ∩ C)
	r := union(firstPart, secondPart)    // Gabungan hasil akhir

	fmt.Printf("A: %v B: %v C: %v\n", a, b, c)
	fmt.Printf("Hasil Operasi R: %v\n", r)

	// Filter: Genap dan < (N/4)
	limit := n / 4
	var filtered []int
	for _, v := range r {
		if v%2 == 0 && v < limit {
			filtered = append(filtered, v)
		}
	}
	fmt.Printf("Hasil Filter (Genap < %d): %v\n", limit, filtered)
	fmt.Printf("Total Elemen: %d\n", len(filtered))
}

// Soal 2: Pasangan Subset x + y > K
func SolveSoal2(s []int, k int) {
	fmt.Printf("Set S: %v | Target K: %d\n", s, k)
	count := 0
	fmt.Println("Subset 2-Elemen (Sum > K):")
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			sum := s[i] + s[j]
			if sum > k {
				count++
				fmt.Printf("%d. {%d, %d} (Sum=%d)\n", count, s[i], s[j], sum)
			}
		}
	}
	fmt.Printf("Total: %d Pasangan\n", count)
}
