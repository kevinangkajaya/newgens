package test

import (
	"newgens/src"
	"testing"
)

func TestReadFilesMt202(t *testing.T) {
	data1, err := src.ReadLines("files/202MEP_52A_57A_58A.txt")
	if err != nil {
		t.Fatal(err)
	}

	data2, err := src.ReadLines("files/202MEP_52D_57D_58A.txt")
	if err != nil {
		t.Fatal(err)
	}

	for _, datum1 := range data1 {
		t.Log(datum1)
	}
	for _, datum2 := range data2 {
		t.Log(datum2)
	}
}
