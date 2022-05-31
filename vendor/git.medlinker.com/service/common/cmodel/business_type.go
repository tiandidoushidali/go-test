package cmodel

import "git.medlinker.com/service/common/ecode/errors"

// 业务类型
type BusinessType uint8

const (
	// 随诊
	BusinessTypeFollowUp BusinessType = iota + 1
	// 问诊
	BusinessTypeInquiry
	// 开药门诊
	BusinessTypeMedicine
	// 患者自购
	BusinessTypeBuySelf
)

var atoBusinessTypeMapping = map[string]BusinessType{
	"follow-up": BusinessTypeFollowUp,
	"inquiry":   BusinessTypeInquiry,
	"medicine":  BusinessTypeMedicine,
	"buy-self":  BusinessTypeBuySelf,
}

var businessTypeToAMapping = map[BusinessType]string{}

func init() {
	for k, v := range atoBusinessTypeMapping {
		businessTypeToAMapping[v] = k
	}
}

func (bt BusinessType) String() string {
	if str, ok := businessTypeToAMapping[bt]; ok {
		return str
	}

	return "unknown"
}

func AToBusinessType(str string) (BusinessType, error) {
	if v, ok := atoBusinessTypeMapping[str]; ok {
		return v, nil
	}

	return 0, errors.New("unknown business-type")
}
