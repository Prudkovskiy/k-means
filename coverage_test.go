package main

import (
	"reflect"
	"testing"
)

type TestCaseInit struct {
	FileName string
}

func TestInitialization(t *testing.T) {
	cases := []TestCaseInit{
		TestCaseInit{"test_data/right.txt"},
		TestCaseInit{"test_data/notExistFile.txt"},
		TestCaseInit{"test_data/badYaml.txt"},
		TestCaseInit{"test_data/badKmax.txt"},
	}

	reader := new(FileReader)
	for caseNum, item := range cases {
		reader.FileNameIn
		u, err := (item.FileName)

		if item.IsError && err == nil {
			t.Errorf("[%d] expected error, got nil", caseNum)
		}
		if !item.IsError && err != nil {
			t.Errorf("[%d] unexpected error", caseNum, err)
		}
		if !reflect.DeepEqual(u, item.User) {
			t.Errorf("[%d] wrong results: got %+v, expected %+v",
				caseNum, u, item.User)
		}
	}

	reader := new(FileReader)
	reader.FileNameIn = "test_data/in.txt"
	bytes, err := reader.Read()
	if err != nil {
		t.Errorf("unexpected error: ", err)
	}
	err = reader.Unpack(bytes)
	if err != nil {
		t.Errorf("unexpected error ", err)
	}

	reader.FileNameIn = "notexist.txt"
	bytes, err = reader.Read()
	if err != nil {
		t.Errorf("expected error, it's OK! ", err)
	}

}

/*
    to create test coverage:
	go test -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

*/
