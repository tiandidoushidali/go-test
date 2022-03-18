package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("test/medlinker/pro_med_calendar_2022-01-06-01.sql")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create("test/medlinker/pro_med_calendar_2022-01-06-01_new.sql")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	rd := bufio.NewReader(file)
	wt := bufio.NewWriter(newFile)
	for {
		line, err := rd.ReadString('\n') // 为结束符读入一行
		if err != nil || io.EOF == err {

			break
		}
		if strings.HasPrefix(line, "UNLOCK TABLES") || strings.HasPrefix(line, "LOCK TABLES ") {
			fmt.Println(line)
			continue
		}
		wt.WriteString(line)
	}
	wt.Flush()
}
