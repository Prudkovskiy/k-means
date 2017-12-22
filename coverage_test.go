package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	fileNameIn     string
	fileNameOut    string
	expectedResult string
}

func TestInit(t *testing.T) {
	cases := []TestCase{
		TestCase{"test_data/notExist.dat", "", "_"},
		TestCase{"test_data/echo.dat", "test_data/echo.ans", "_"},
		TestCase{"test_data/badKmax1.dat", "_", "_"},
		TestCase{"test_data/badKmax2.dat", "_", "_"},
		TestCase{"test_data/badJSON.dat", "_", "_"},
		TestCase{"test_data/badNode1.dat", "_", "_"},
		TestCase{"test_data/badNode2.dat", "_", "_"},
	}

	reader := new(FileReader)

	reader.FileNameIn = cases[0].fileNameIn
	reader.FileNameOut = cases[0].fileNameOut
	b, err := reader.Read()
	if err == nil {
		t.Errorf("[0] the input file doesn't exist")
	}
	err = reader.Write(b)
	if err == nil {
		t.Errorf("[0] bad output file name")
	}

	reader.FileNameIn = cases[1].fileNameIn
	reader.FileNameOut = cases[1].fileNameOut
	b, err = reader.Read()
	if err != nil {
		t.Errorf("[1] unexpected error in reading: %+v", err)
	}
	err = reader.Write(b)
	if err != nil {
		t.Errorf("[1] unexpected error in writing: %+v", err)
	}

	reader.FileNameIn = cases[2].fileNameIn
	b, err = reader.Read()
	err = reader.Unpack(b)
	if err == nil {
		t.Errorf("[2] uncorrect Kmax: first line in file (Kmax) not integer")
	}

	reader.FileNameIn = cases[3].fileNameIn
	b, err = reader.Read()
	err = reader.Unpack(b)
	if err == nil {
		t.Errorf("[3] uncorrect Kmax: Kmax > number of nodes in your graph")
	}

	reader.FileNameIn = cases[4].fileNameIn
	b, err = reader.Read()
	err = reader.Unpack(b)
	if err == nil {
		t.Errorf("[4] your data in input file not json format")
	}

	reader.FileNameIn = cases[5].fileNameIn
	b, err = reader.Read()
	err = reader.Unpack(b)
	if err == nil {
		t.Errorf("[5] you must type only numerical nodes")
	}

	reader.FileNameIn = cases[6].fileNameIn
	b, err = reader.Read()
	err = reader.Unpack(b)
	if err == nil {
		t.Errorf("[6] you must type only numerical friend nodes")
	}
}

func TestFloydWarshall(t *testing.T) {
	matrix := [][]float64{{1, 2}, {3, 4}, {5, 6}}
	fw := new(FloydWarshall)
	_, err := fw.Run(matrix)
	if err == nil {
		t.Errorf("the matrix is not quadratic")
	}
}

func TestPAM(t *testing.T) {
	cases := []TestCase{
		TestCase{"test_data/test1.dat", "test_data/test1.ans", `{"1":[1,2,3],"2":[4,5,6]}`},
		TestCase{"test_data/test2.dat", "test_data/test2.ans", `{"1":[1,2,5],"2":[3],"3":[4]}`},
		TestCase{"test_data/test3.dat", "test_data/test3.ans", `{"1":[1,2,3,4],"2":[5,6],"3":[7,8,9],"4":[10],"5":[11]}`},
	}

	reader := new(FileReader)
	fw := new(FloydWarshall)

	for caseNum, item := range cases {
		reader.FileNameIn = item.fileNameIn   // входной файл
		reader.FileNameOut = item.fileNameOut // выходной файл
		bytes, _ := reader.Read()
		reader.Unpack(bytes)

		matr, _ := fw.Run(reader.Matrix)

		pam := new(PAM)
		pam.Kmax = reader.Kmax
		pam.Distances = matr

		clusters, _ := pam.RunPAM()
		data := reader.Pack(clusters)
		reader.Write(data)

		if !reflect.DeepEqual(data, []byte(item.expectedResult)) {
			t.Errorf("[%d] wrong results: got \n%+v \nexpected \n%+v",
				caseNum, string(data), item.expectedResult)
		}
	}
}

/*
    to create test coverage:
	go test -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

*/
