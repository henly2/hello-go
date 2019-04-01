package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	AA = 0
)

type (
	Reader interface {
		Read() (string, error)
	}

	LocalFile struct {
		Path string
	}
	RemoteFile struct {
		Url string
	}
)

func (tmp *LocalFile) Read() (string, error) {
	b, err := ioutil.ReadFile(tmp.Path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (tmp *RemoteFile) Read() (string, error) {
	resp, err := http.Get(tmp.Url)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func BuildReader(mode int, param string) (Reader, error) {
	if mode == 1 {
		return &LocalFile{Path: param}, nil
	} else if mode == 2 {
		return &RemoteFile{Url: param}, nil
	}

	return nil, fmt.Errorf("not support mode = %d", mode)
}

type Person04 struct {
	name string
	sex  byte
	age  int
}

func (tmp *Person04) PrintInfo() {
	fmt.Println(tmp.name, tmp.sex, tmp.age)
}
func PrintInfo(tmp *Person04) {
	fmt.Println(tmp.name, tmp.sex, tmp.age)
}

type Student01 struct {
	Person04

	id   int
	addr string
}

func main() {
	a := Person04{"ls", 'm', 6}
	a.PrintInfo()

	PrintInfo(&a)

	ff := PrintInfo
	ff(&a)

	p := Student01{a, 666, "aaa"}
	p.PrintInfo()

	var c Reader
	var err error

	c, err = BuildReader(1, "mmmmmm")
	if err != nil {
		return
	}
	fmt.Println(c.Read())

}
