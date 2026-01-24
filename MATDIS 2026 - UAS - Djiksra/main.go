package main

import (
	"fmt"
	"math"
)

// Edge merepresentasikan hubungan antar desa dan jaraknya
type Edge struct {
	Node   string
	Weight int
}

// Graph adalah kumpulan data desa dan jalannya
type Graph map[string][]Edge

// Dijkstra adalah fungsi utama untuk mencari jalur terpendek
func Dijkstra(graph Graph, start string, end string) (int, []string) {
	// 1. Inisialisasi jarak awal dengan nilai tak terhingga (infinity)
	distances := make(map[string]int)
	for node := range graph {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	// 2. Simpan jejak jalur yang dilewati
	prev := make(map[string]string)
	visited := make(map[string]bool)

	for len(visited) < len(graph) {
		// Cari node dengan jarak terkecil yang belum dikunjungi
		currNode := ""
		minDist := math.MaxInt32

		for node, dist := range distances {
			if !visited[node] && dist < minDist {
				minDist = dist
				currNode = node
			}
		}

		// Jika tidak ada node lagi yang bisa dijangkau
		if currNode == "" {
			break
		}

		visited[currNode] = true

		// Jika sudah sampai tujuan, kita bisa berhenti (opsional, tapi lebih efisien)
		if currNode == end {
			break
		}

		// Update jarak ke tetangga-tetangganya
		for _, edge := range graph[currNode] {
			newDist := distances[currNode] + edge.Weight
			if newDist < distances[edge.Node] {
				distances[edge.Node] = newDist
				prev[edge.Node] = currNode
			}
		}
	}

	// 3. Menyusun jalur dari hasil 'prev'
	path := []string{}
	curr := end
	for curr != "" {
		path = append([]string{curr}, path...)
		curr = prev[curr]
	}

	return distances[end], path
}

func main() {
	// Contoh Data Desa (Graph)
	// Kamu bisa mengubah nama desa dan jaraknya di sini
	desaGraph := Graph{
		"Desa A": {{"Desa B", 4}, {"Desa C", 2}},
		"Desa B": {{"Desa A", 4}, {"Desa C", 5}, {"Desa D", 10}},
		"Desa C": {{"Desa A", 2}, {"Desa B", 5}, {"Desa E", 3}},
		"Desa D": {{"Desa B", 10}, {"Desa E", 4}},
		"Desa E": {{"Desa C", 3}, {"Desa D", 4}},
	}

	start := "Desa A"
	end := "Desa D"

	jarak, rute := Dijkstra(desaGraph, start, end)

	// Menampilkan Hasil
	fmt.Println("=== Program Optimasi Jalur Terpendek ===")
	fmt.Printf("Titik Awal : %s\n", start)
	fmt.Printf("Tujuan     : %s\n", end)
	fmt.Println("---------------------------------------")

	if jarak == math.MaxInt32 {
		fmt.Println("Jalur tidak ditemukan.")
	} else {
		fmt.Printf("Total Jarak Terpendek: %d KM\n", jarak)
		fmt.Printf("Rute yang Dilewati   : ")
		for i, desa := range rute {
			fmt.Print(desa)
			if i < len(rute)-1 {
				fmt.Print(" -> ")
			}
		}
		fmt.Println()
	}
}

