package main

import (
	"archive/zip"
	"io"
	"os"
)

func main() {
	f, e := os.Create("/usr/local/works/go/src/go-test/module/test-file/test.zip")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	fzipw := zip.NewWriter(f)
	defer fzipw.Close()

	f1, e := os.Open("/usr/local/works/go/src/go-test/module/test-file/es-head.crx")
	if e != nil {
		panic(e)
	}
	defer f1.Close()

	f2, e := os.Open("/usr/local/works/go/src/go-test/module/test-file/du-5000.xlsx")
	if e != nil {
		panic(e)
	}
	defer f2.Close()

	fzipwio, e := fzipw.Create("crx/es-head.crx")
	if e != nil {
		panic(e)
	}
	io.Copy(fzipwio, f1)

	fzipwio2, e := fzipw.Create("xlsx/du-5000.xlsx")
	if e != nil {
		panic(e)
	}
	io.Copy(fzipwio2, f2)
}
