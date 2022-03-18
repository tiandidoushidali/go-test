package main

import (
	"bytes"
	"encoding/csv"
	"github.com/tealeg/xlsx"
	"os"
)

type StructA struct {
	Name string
}

type StructC struct {
	name string
}

type StructB struct {
	StructA
	StructC
}
func main() {
	bytesBuffer := &bytes.Buffer{}
	bytesBuffer.WriteString("\xEF\xBB\xBF")
	file := xlsx.NewFile()
	sht, _ := file.AddSheet("测试1")
	row := sht.AddRow()
	cell := row.AddCell()
	cell.Value = "aaaa"
	style := xlsx.NewStyle()
	alignment := xlsx.Alignment{
		Horizontal: "center",
		Vertical:   "center",
	}
	style.Alignment = alignment
	style.ApplyAlignment = true
	style.Border = xlsx.Border{
		LeftColor:   "red",
		RightColor:  "red",
		TopColor:    "red",
		BottomColor: "red",
	}
	style.Font = xlsx.Font{
		Size:      100,
		Color:     "red",
		Bold:      true,
	}
	style.ApplyFont = true
	style.ApplyBorder = true
	cell.SetStyle(style)
	file.AddSheet("测试2")
	//file.Write(bytesBuffer)
	file.Save("aa.xlsx")

	return

	f, err := os.Create("test.csv")//创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)//创建一个新的写入文件流
	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	w.WriteAll(data)//写入数据
	w.Flush()
}