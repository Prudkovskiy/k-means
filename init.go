package main

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type FileReader struct {
	FileName string
	Matrix   [][]int
	Kmax     int
}

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

// распаковываем содержимое файла в матрицу смежности
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
	var allinfo []map[int](map[int]int)
	err = yaml.Unmarshal(data[start:], &allinfo)
	if err != nil {
		return err
	}

	n := len(allinfo) // размер матрицы смежности
	// заполняем ее нулями
	for i := 0; i < n; i++ {
		r.Matrix = append(r.Matrix, []int{})
		for j := 0; j < n; j++ {
			r.Matrix[i] = append(r.Matrix[i], 0)
		}
	}
	// дополняем весами ребер
	for _, adjacency := range allinfo {
		for node, spisok := range adjacency {
			for friendnode, weight := range spisok {
				r.Matrix[node-1][friendnode-1] = weight
			}
		}
	}

	return err
}

// func main() {
// 	k := new(FileReader)
// 	k.FileName = "in.txt"
// 	b, err := k.Read()
// 	err = k.Unpack(b)

// 	fmt.Println(err, k.Matrix)
// }
