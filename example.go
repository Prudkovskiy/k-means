package main

import "fmt"

func main() {
	reader := new(FileReader)
	reader.FileNameIn = "./test_data/in2.txt"
	reader.FileNameOut = "./test_data/out2.txt"
	bytes, _ := reader.Read()
	reader.Unpack(bytes)

	fw := new(FloydWarshall)
	matr, _ := fw.Run(reader.Matrix)

	pam := new(PAM)
	pam.Kmax = reader.Kmax
	pam.Distances = matr

	clust, ff := pam.swap(1)
	visuality := make(map[int][]int, len(clust))
	i := 1
	for _, val := range clust {
		visuality[i] = make([]int, 0, len(val))
		for _, node := range val {
			visuality[i] = append(visuality[i], node+1)
		}
		i++
	}
	fmt.Println(visuality, ff)

	clust, ff = pam.swap(2)
	visuality = make(map[int][]int, len(clust))
	i = 1
	for _, val := range clust {
		visuality[i] = make([]int, 0, len(val))
		for _, node := range val {
			visuality[i] = append(visuality[i], node+1)
		}
		i++
	}
	fmt.Println(visuality, ff)

	clust, ff = pam.swap(3)
	visuality = make(map[int][]int, len(clust))
	i = 1
	for _, val := range clust {
		visuality[i] = make([]int, 0, len(val))
		for _, node := range val {
			visuality[i] = append(visuality[i], node+1)
		}
		i++
	}
	fmt.Println(visuality, ff)

	clusters, weightFunction := pam.RunPAM()
	visuality = make(map[int][]int, len(clusters))
	i = 1
	for _, val := range clusters {
		visuality[i] = make([]int, 0, len(val))
		for _, node := range val {
			visuality[i] = append(visuality[i], node+1)
		}
		i++
	}
	data, err := reader.Pack(visuality)
	if err != nil {
		panic(err)
	}
	reader.Write(data)
	fmt.Println("\n", visuality, weightFunction)
}
