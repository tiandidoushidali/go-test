package config

type GrapeCfgData struct {
	Nodes []Node `gconv:"node"`
}
type Node struct {
	Name string `gconv:"name"`
	Ws   string `gconv:"ws"`
	Tcp  string `gconv:"tcp"`
}

var (
	NodeHubMap = map[string]Node{}
	instance   = new(GrapeCfgData)
)

//func GrapeCfg() *GrapeCfgData {
//	if instance == nil {
//		instance = new(GrapeCfgData)
//
//	}
//	return instance
//}

//func (p *DispatcherCfgData) Init(cfgStr string) (err error) {
//	_, err = toml.DecodeFile(cfgStr, _instance)
//	return
//}

func init() {

}