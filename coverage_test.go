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
		TestCase{"test_data/notExist.txt", "", "_"},
		TestCase{"test_data/in.txt", "test_data/out.txt", "_"},
		TestCase{"test_data/badKmax.txt", "_", "_"},
		TestCase{"test_data/badYAML.txt", "_", "_"},
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
		t.Errorf("[2] first line in file (Kmax) not integer")
	}

	reader.FileNameIn = cases[3].fileNameIn
	b, err = reader.Read()
	err = reader.Unpack(b)
	if err == nil {
		t.Errorf("[3] your data in input file not yaml format")
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
		TestCase{"test_data/in1.txt", "out1.txt",
			`1:
- 3
- 1
- 2
2:
- 5
- 4
- 6
`,
		},
		TestCase{"test_data/in2.txt", "out2.txt",
			`1:
- 3
- 1
- 2
2:
- 5
- 4
- 6
3:
- 7
- 8
- 9
4:
- 10
`,
		},
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
		// reader.Write(data)

		if !reflect.DeepEqual(data, []byte(item.expectedResult)) {
			t.Errorf("[%d] wrong results: got \n%+v expected \n%+v",
				caseNum, string(data), item.expectedResult)
		}
	}
}

/*
    to create test coverage:
	go test -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

*/
