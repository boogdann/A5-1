package main

import (
	a51 "2/internal/a51/v2"
	"2/internal/app"
)

const (
	filename     = "tests/textfile1.txt"
	saveFilename = "tests/cipher/textfile1.method_1.save.txt"
)

func main() {
	a := app.New(a51.Method1, 128, filename, 123456789)
	a.Run()

	if err := a.Save(); err != nil {
		panic(err)
	}
}
