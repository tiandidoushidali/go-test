package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("/usr/local/works/go/src/go-test/module/test-csv/project.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	rs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range rs {
		fmt.Println(fmt.Sprintf("k:%d, v:%v", k, v))
	}

}
