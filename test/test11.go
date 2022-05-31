package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
1
  2 3 4 5   6
  7 8 9 10  11

*/
type testStr struct {
	Id int64
}

func setId(p testStr) {
	var a int64 = 123
	b := a
	p.Id = b
}

func test11(arr []int64) {
	arr2 := make([]int64, len(arr))
	copy(arr2, arr)
	arr2[1] = 200
	arr2 = nil
}

const creations = 20_0_00_000000

func main() {
	mp := map[string]interface{}{
		"c":3,
		"d":4,
		"bcd":4,
		"a":1,
		"b":2,
		"abc":3,
	}

	mpKeys := make([]string, 0)
	for k, _ := range mp {
		mpKeys = append(mpKeys, k)
	}
	sort.Slice(mpKeys, func(i, j int) bool {
		return mpKeys[i] < mpKeys[j]
	})
	var build strings.Builder
	for _, v := range mpKeys {
		build.WriteString(v)
		build.WriteString("=")
		//vv, _ := json.Marshal(v)
		//build.WriteString(string(vv))
		var val string
		mpv := mp[v]
		switch mpv.(type) {
		case string:
			val = mpv.(string)
		case float64:
			ft := mpv.(float64)
			val = strconv.FormatFloat(ft, 'f', -1, 64)
		case float32:
			ft := mpv.(float32)
			val = strconv.FormatFloat(float64(ft), 'f', -1, 64)
		case int64:
			it := mpv.(int64)
			val = strconv.FormatInt(it, 10)
		case int32:
			it := mpv.(int32)
			val = strconv.Itoa(int(it))
		case int16:
			it := mpv.(int16)
			val = strconv.Itoa(int(it))
		case int8:
			it := mpv.(int8)
			val = strconv.Itoa(int(it))
		case []byte:
			val = string(mpv.([]byte))
		default:
			nval, _ := json.Marshal(mpv)
			val = string(nval)
		}
		build.WriteString(val)
	}

	fmt.Println(mpKeys, build.String())
	return

	fmt.Println(fmt.Sprintf("%x", 11))
	fmt.Println(-1 ^ -1<<5)
	return
	fmt.Println("----", creations+2)
	ts := testStr{Id: 1}
	setId(ts)
	arr := []int64{1, 2, 3, 4, 5}

	test11(arr)

	arr1 := arr[1:]
	arr1[1] = 100
}
