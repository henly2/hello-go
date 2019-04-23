package util

import "testing"

func TestSaveReadData(t *testing.T) {
	SetSavedDir("./")

	fileName := "test.txt"

	// 1 string
	data := "abcd!@#$"
	err := SaveData(data, fileName)
	if err != nil {
		t.Fatal(err)
	}

	dataRead := ""
	err = ReadData(fileName, &dataRead)
	if err != nil {
		t.Fatal(err)
	}

	if data != dataRead{
		t.Fatal(data, "!=", dataRead)
	}

	// 2 struct
	type SS struct {
		A string `json:"a"`
		B int    `json:"b"`
	}
	s := SS{
		A: "ss-a",
		B: 1234,
	}
	err = SaveData(s, fileName)
	if err != nil {
		t.Fatal(err)
	}

	sRead := SS{}
	err = ReadData(fileName, &sRead)
	if err != nil {
		t.Fatal(err)
	}

	if s.A != sRead.A || s.B != sRead.B{
		t.Fatal(s, "!=", sRead)
	}

}
