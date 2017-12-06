package main

import "fmt"

func main() {
	reader := new(FileReader)
	reader.FileName = "./data/in.txt"
	reader.FileNameOut = "./data/out.txt"
	bytes, _ := reader.Read()
	reader.Unpack(bytes)

	fw := new(FloydWarshall)
	matr, _ := fw.Run(reader.Matrix)
	pam := new(PAM)
	pam.Distances = matr
	medoids := pam.swap(2)
	clusters, weightFunction, _ := pam.clustering(medoids)
	data, err := reader.Pack(clusters)
	if err != nil {
		panic(err)
	}
	reader.Write(data)
	fmt.Println(weightFunction)
}
