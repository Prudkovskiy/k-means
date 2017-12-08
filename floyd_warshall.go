package main

import (
	"fmt"
)

// FloydWarshall алгоритм Флойда-Уоршелла,
// поиск кротчайших путей между всеми парами вершин
type FloydWarshall struct {
	shortest [][][]float64
}

// Run - процедура поиска кротчайших путей
func (r *FloydWarshall) Run(matrix [][]float64) ([][]float64, error) {
	n := len(matrix)
	m := len(matrix[0])
	if n != m {
		return nil, fmt.Errorf("the matrix is not quadratic")
	}

	// O(n^3) time
	r.shortest = make([][][]float64, n+1)
	for i := 0; i < n+1; i++ {
		r.shortest[i] = make([][]float64, n+1)
		for j := 0; j < n+1; j++ {
			r.shortest[i][j] = make([]float64, n+1)
		}
	}

	// O(n^2) time
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			r.shortest[0][i][j] = matrix[i-1][j-1]
		}
	}

	// O(n^3) time
	for x := 1; x < n+1; x++ {
		for u := 1; u < n+1; u++ {
			for v := 1; v < n+1; v++ {
				cur := &r.shortest[x][u][v]
				pred := r.shortest[x-1][u][v]
				shortUX := r.shortest[x-1][u][x]
				shortXV := r.shortest[x-1][x][v]

				if pred < shortUX+shortXV {
					*cur = pred
				} else {
					*cur = shortUX + shortXV
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = r.shortest[n][i+1][j+1]
		}
	}

	return matrix, nil
}
