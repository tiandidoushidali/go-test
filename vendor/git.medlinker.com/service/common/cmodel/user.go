package cmodel

const (
	// 患者
	UserTypePatient = iota + 1
	// 医生
	UserTypeDoctor
	// 经纪人
	UserTypeBroker
)

type User struct {
	// 用户ID：med_primary.user.id
	Id uint32
	// 用户类型
	Type uint32
	// 用户姓名
	FullName string
}
