package main

import (
	"math"
)

// PAM is partitioning around medoids
type PAM struct {
	Distance [][]float64
	Kmax     int
}

func (p *PAM) clustering(distances [][]float64, medoids []int) (map[int][]int, float64) {
	var f float64 // целевая функция (сумма кротчайших расстояний "медоид-не медоид")
	n := len(distances)
	k := len(medoids)
	clusters := make(map[int][]int, k)
	for _, i := range medoids {
		clusters[i] = append(clusters[i], i)
	}

	// O(n*k),
	// where k - number of clusters, n - number of nodes in graph
	for node := 0; node < n; node++ {
		min := math.Inf(1) // min = infinity
		medoid := -1
		_, ok := clusters[node] // проверяем, является ли node медоидом
		if !ok {
			for _, currMedoid := range medoids {
				if distances[node][currMedoid] < min {
					min = distances[node][currMedoid]
					medoid = currMedoid
				}
			}
			clusters[medoid] = append(clusters[medoid], node)
			f += min
		}
	}
	return clusters, f
}

// func (p *PAM) swap(distances [][]float64, medoids []int) map[int][]int {
// 	k := len(medoids)
// }

// func (p *PAM) build(k int, distance [][]float64) []int {
// 	n := len(distance)
// 	firstMedoid := rand.Intn(n)
// 	medoids := make([]int, k)

// 	return medoids, nil
// }
