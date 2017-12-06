package main

import (
	"math"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// FileReader работает с файлом входных данных
type FileReader struct {
	FileName    string
	FileNameOut string
	Matrix      [][]float64 // матрица смежности
	Kmax        int         // верхняя граница кластеризации
}

// Read считывает данные из файла
func (r *FileReader) Read() ([]byte, error) {
	file, err := os.Open(r.FileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// get the file size
	stat, _ := file.Stat()
	// read the file
	data := make([]byte, stat.Size())
	file.Read(data)

	return data, err
}

func (r *FileReader) Write(data []byte) error {
	file, err := os.Create(r.FileNameOut)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return err
}

// Pack запаковывает пришедшую мапу кластеров в слайс байт
func (r *FileReader) Pack(clust map[int][]int) ([]byte, error) {
	data, err := yaml.Marshal(clust)
	if err != nil {
		return nil, err
	}
	return data, err
}

// Unpack распаковывает содержимое файла в матрицу смежности и Kmax
func (r *FileReader) Unpack(data []byte) error {
	// доходим до конца первой строки
	var start int
	for k, i := range data {
		if i == 10 {
			start = k
			break
		}
	}
	// инициализируем границу кластеризации первой строкой файла
	k, err := strconv.Atoi(string(data[:start-1]))
	if err != nil {
		return err
	}
	r.Kmax = k

	// достаем списки смежности из файла в формате yaml
	var allinfo []map[int](map[int]float64)
	err = yaml.Unmarshal(data[start:], &allinfo)
	if err != nil {
		return err
	}

	// выделяем место под матрицу смежности n*n
	n := len(allinfo)
	r.Matrix = make([][]float64, n)
	for i := 0; i < n; i++ {
		r.Matrix[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				r.Matrix[i][j] = math.Inf(1)
			} else {
				r.Matrix[i][j] = 0
			}
		}
	}

	// дополняем ее весами ребер
	for _, adjacency := range allinfo {
		for node, spisok := range adjacency {
			for friendnode, weight := range spisok {
				r.Matrix[node-1][friendnode-1] = weight
			}
		}
	}

	return err
}
