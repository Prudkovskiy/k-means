package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("you have not entered the input or output files")
		return
	}
	reader := new(FileReader)
	reader.FileNameIn = string(os.Args[1])  // входной файл
	reader.FileNameOut = string(os.Args[2]) // выходной файл
	bytes, err := reader.Read()
	if err != nil {
		fmt.Println("something go wrong: ", err)
		return
	}
	err = reader.Unpack(bytes)
	if err != nil {
		fmt.Println("something go wrong: ", err)
		return
	}

	fw := new(FloydWarshall)
	matr, _ := fw.Run(reader.Matrix)

	pam := new(PAM)
	pam.Kmax = reader.Kmax
	pam.Distances = matr

	clusters, _ := pam.RunPAM()
	data := reader.Pack(clusters)
	err = reader.Write(data)
	if err != nil {
		fmt.Println("something go wrong: ", err)
		return
	}
}
