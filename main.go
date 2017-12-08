package main

import (
	"fmt"
	"os"
)

func main() {
	reader := new(FileReader)
	reader.FileNameIn = string(os.Args[1])  // входной файл
	reader.FileNameOut = string(os.Args[2]) // выходной файл
	bytes, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = reader.Unpack(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fw := new(FloydWarshall)
	matr, err := fw.Run(reader.Matrix)
	if err != nil {
		fmt.Println(err)
		return
	}

	pam := new(PAM)
	pam.Kmax = reader.Kmax
	pam.Distances = matr

	clusters, _ := pam.RunPAM()
	data := reader.Pack(clusters)
	err = reader.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
