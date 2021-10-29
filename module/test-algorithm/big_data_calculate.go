package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func strReverse(param string) string {
	if len(param) == 0 {
		return param
	}

	buffer := bytes.Buffer{}
	for index := len(param) - 1; index >= 0; index -- {
		buffer.WriteString(fmt.Sprintf("%c", param[index]))
	}
	return buffer.String()
}
/**
大数计算
 */
func add(paraA, paraB string) string {
	paraA = strReverse(paraA)
	paraB = strReverse(paraB)
	fmt.Println("a,b:", paraA, paraB)
	lenA, lenB := len(paraA), len(paraB)
	if lenA == 0 {
		return paraB
	}
	if lenB == 0 {
		return paraA
	}

	lenBase := 0
	if lenA > lenB {
		lenBase = lenA
	} else {
		lenBase = lenB
	}

	overFlow := 0
	buffer := bytes.Buffer{}
	for indexBase := 0; indexBase < lenBase; indexBase ++ {
		itemA := 0
		if lenA > indexBase {
			itemA, _ = strconv.Atoi(string(paraA[indexBase]))
		}

		itemB := 0
		if lenB > indexBase {
			itemB, _ = strconv.Atoi(string(paraB[indexBase]))
		}

		sum := itemA + itemB + overFlow
		if sum / 10 > 0 {
			overFlow = sum / 10
			sum = sum - overFlow * 10
		} else {
			overFlow = 0
		}
		fmt.Println("----", itemA, itemB, overFlow, sum)
		buffer.WriteString(strconv.Itoa(sum))
	}
	if overFlow > 0 {
		buffer.WriteString(fmt.Sprintf("%d", overFlow))
	}
	return buffer.String()
}

func add2(paraA, paraB string) string {
	lenA, lenB := len(paraA), len(paraB)
	if lenA == 0 {
		return paraB
	}
	if lenB == 0 {
		return paraA
	}

	lenBase := 0
	if lenA > lenB {
		lenBase = lenA
	} else {
		lenBase = lenB
	}

	indexA, indexB := lenA - 1, lenB - 1
	overFlow := 0
	result := ""
	for indexBase := 0; indexBase < lenBase; indexBase ++ {
		itemA := 0
		if lenA > indexBase {
			itemA, _ = strconv.Atoi(string(paraA[indexA]))
			indexA --
		}

		itemB := 0
		if lenB > indexBase {
			itemB, _ = strconv.Atoi(string(paraB[indexB]))
			indexB --
		}

		sum := itemA + itemB + overFlow
		if sum / 10 > 0 {
			overFlow = sum / 10
			sum = sum - overFlow * 10
		} else {
			overFlow = 0
		}
		fmt.Println("----", itemA, itemB, overFlow, sum)
		result = strconv.Itoa(sum) + result
	}
	if overFlow > 0 {
		result = strconv.Itoa(overFlow) + result
	}
	return result
}

func main() {
	fmt.Println("====", 5/3)
	sum := add("99999999999999999999999999999999", "99999999999999999999999999999999")
	fmt.Println("sum:", strReverse(sum))
	fmt.Println("big sum:", add2("99999999999999999999999999999999", "99999999999999999999999999999999"))
	fmt.Println("reverse:", strReverse("12345643"))
	//fmt.Println("ssssssss:", 99999999999999999999999999999999 + 99999999999999999999999999999999)
}
