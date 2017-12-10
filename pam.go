package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

// PAM is partitioning around medoids
type PAM struct {
	Distances [][]float64
	Kmax      int
}

// O(n*k),
// where k - number of clusters, n - number of nodes in graph
func (p *PAM) clustering(medoids []int) (map[int][]int, float64, int) {
	var f float64 // целевая функция (сумма кротчайших расстояний "медоид-не медоид")

	// эти переменные нужны для фазы build,
	// когда идет поиск вершины с max расстоянием до ближайшего медоида
	max := math.Inf(-1)
	var nextMedoid int

	n := len(p.Distances)
	k := len(medoids)
	clusters := make(map[int][]int, k)
	for _, i := range medoids {
		clusters[i] = append(clusters[i], i)
	}

	for node := 0; node < n; node++ {
		min := math.Inf(1)
		medoid := -1
		_, ok := clusters[node] // проверяем, является ли node медоидом
		if !ok {
			for _, currMedoid := range medoids {
				if p.Distances[node][currMedoid] < min {
					min = p.Distances[node][currMedoid]
					medoid = currMedoid
				}
			}
			clusters[medoid] = append(clusters[medoid], node)
			f += min

			if min > max {
				max = min
				nextMedoid = node
			}
		}
	}
	return clusters, f, nextMedoid
}

// O(k*k*n)
func (p *PAM) build(k int) []int {
	n := len(p.Distances)
	medoids := make([]int, 0, k) // len = 0, capacity = k
	// произвольно выбираем вершну в качестве первого медоида
	rand.Seed(time.Now().UnixNano()) // задаем старт генератора, привязан ко времени
	medoids = append(medoids, rand.Intn(n))

	// найдем оставшиеся k-1 медоидов
	for i := 1; i < k; i++ {
		_, _, newMedoid := p.clustering(medoids)
		medoids = append(medoids, newMedoid)
	}
	p.clustering(medoids)

	return medoids
}

// O(number_of_iterations*k*n*(k*n))
func (p *PAM) swap(k int) (map[int][]int, float64) {
	n := len(p.Distances)
	medoids := p.build(k)
	notMedoids := make([]int, n-k)

	var idx int
	// O(n*k)
	for i := 0; i < n; i++ {
		// проверка является ли вершина i медоидом
		isMedoid := false
		for _, medoid := range medoids { // O(k)
			if i == medoid {
				isMedoid = true
			}
		}
		// если вершина не является медоидом, то добавляем ее в немедоиды
		if !isMedoid {
			notMedoids[idx] = i
			idx++
		}
	}

	optimalMedoids := medoids
	_, optimalMin, _ := p.clustering(optimalMedoids)

	// main PAM loop - O(number_of_iterations*k*n*(k*n))
	stop := false // условие останова
	for !stop {
		// O(k*n*(k*n))
		for medIdx, medoid := range optimalMedoids {
			ch := make(chan map[float64][]int, len(notMedoids))
			wg := &sync.WaitGroup{}
			_, iterMin, _ := p.clustering(optimalMedoids) // значение целевой функции на данной итерации

			for notMedIdx, notMedoid := range notMedoids {
				newMedoids := make([]int, k)
				copy(newMedoids, optimalMedoids)
				newMedoids[medIdx] = notMedoid

				wg.Add(1)
				go func(out chan<- map[float64][]int, notMedIdx, notMed, med int) {
					defer wg.Done()
					_, sum, _ := p.clustering(newMedoids)
					inf := []int{notMedIdx, notMed, med}
					var inform = map[float64][]int{sum: inf}
					out <- inform
				}(ch, notMedIdx, notMedoid, medoid)
			}
			wg.Wait()
			close(ch)
			// находим те медоид/немедоид, с которым целевая функция максимально уменьшилась
			var notMedIdx, changeNotMed, changeMed = -1, -1, -1
			for inform := range ch {
				for mins, inf := range inform {
					if mins < iterMin {
						iterMin = mins
						notMedIdx = inf[0]
						changeNotMed = inf[1]
						changeMed = inf[2]
					}
				}
			}
			if notMedIdx != -1 {
				optimalMedoids[medIdx] = changeNotMed
				notMedoids[notMedIdx] = changeMed
			}
		}
		_, currMin, _ := p.clustering(optimalMedoids)
		if optimalMin == currMin {
			stop = true
		} else {
			optimalMin = currMin
		}
	}
	clusters, weightFunc, _ := p.clustering(optimalMedoids)

	return clusters, weightFunc
}

// RunPAM выбирает оптимальное число кластеров для данной матрицы смежностей
// и разбивает граф на оптимальное число кластеров
// O(Kmax*O(swap))
func (p *PAM) RunPAM() (map[int][]int, float64) {
	predClusters, predWeightFunc := p.swap(1)
	optimalClusters := predClusters
	optimalWeightFunc := predWeightFunc
	var currClusters map[int][]int
	var currWeightFunc float64
	var nextWeightFunc float64
	var min float64 = 1

	for k := 2; k < p.Kmax; k++ {
		currClusters, currWeightFunc = p.swap(k)
		_, nextWeightFunc = p.swap(k + 1)
		optimal := math.Abs(currWeightFunc-nextWeightFunc) /
			math.Abs(predWeightFunc-currWeightFunc)
		if optimal < min {
			min = optimal
			optimalClusters = currClusters
			optimalWeightFunc = currWeightFunc
		}
		predClusters, predWeightFunc = currClusters, currWeightFunc
	}
	// если алгоритм дошел до Kmax-1, то надо проверить,
	// не будет ли оптимальнее разбить на Kmax,
	// то есть сравнить значения целевых функций
	if len(optimalClusters) == p.Kmax-1 {
		currClusters, currWeightFunc = p.swap(p.Kmax)
		if currWeightFunc < optimalWeightFunc {
			optimalClusters = currClusters
			optimalWeightFunc = currWeightFunc
		}
	}
	return optimalClusters, optimalWeightFunc
}
