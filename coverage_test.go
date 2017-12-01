package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	k := new(FileReader)
	k.FileName = "in.txt"
	b, err := k.Read()
	if err != nil {
		t.Errorf("unexpected error ", err)
	}
	err = k.Unpack(b)
	if err != nil {
		t.Errorf("unexpected error ", err)
	}
}

/*
    to create test coverage:
	go test -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

*/
