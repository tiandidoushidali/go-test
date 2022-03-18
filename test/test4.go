package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"sync/atomic"
	"time"
)

var atomicValue *atomic.Value
var confPath string
func init() {
	atomicValue = &atomic.Value{}
	flag.StringVar(&confPath, "conff", "", "default config path")
}

type Setter struct {

}

type Confff struct{
	Application struct{
		Name string `toml:"name"`
		Env int `toml:"env"`
	} `toml:"application"`
}

func (set *Setter) Set(str string) error {
	fmt.Println("------str-----\n", str)

	cff := new(Confff)


	_, err := toml.Decode(str, &cff)
	if err != nil {
		panic(err)
	}

	fmt.Println(cff.Application.Name, cff.Application.Env)
	atomicValue.Store(cff)
	return nil
}


func main() {
	flag.Parse()
	setter := &Setter{}
	paladin.Init()
	paladin.Watch("test4.toml", setter)

	t := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-t.C:
			conff := atomicValue.Load().(*Confff)
			fmt.Println("---", conff.Application.Name, "---", conff.Application.Env)
		}
	}
	str := ""
	strArr  := make([]string, 0)
	err := json.Unmarshal([]byte(str), &strArr)
	if err == nil {
		fmt.Println(strArr)
	}
	fmt.Println(strArr)

	val := &atomic.Value{}
	val.Store(map[string]int{"aa":1, "bb":2})
	v := val.Load()
	if vv, ok := v.(map[string]int); ok {
		fmt.Println("-----", vv)
	}


	mp := make(map[string]map[string][]map[string]string)
	mppa := make(map[string]string)
	mppa["aaa1"] = "aaa1"
	mppb := make(map[string]string)
	mppb["bbb1"] = "bbb1"

	mmp := make(map[string][]map[string]string, 0)
	mmp["ab"] = append(mmp["ab"], mppa, mppb)
	mp["mp"] = mmp
}
