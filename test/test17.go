package main

import (
	"fmt"
	"os"
)

/**
from user reference 为 9 的用户id
*/
type SystemUserId int64
/**
 * 系统消息
 */
const SYSTEM SystemUserId = 0

/**
 * 业务处理 此角色不下发消息，只接收Command
 */
const BUSINESS SystemUserId = 1

/**
 * 新的朋友
 */
const NEW_FRIEND SystemUserId = 2

/**
 * 医生助手
 */
const DOCTOR_ASSISTANT SystemUserId = 3

/**
 * 动态相关
 */
const TREND SystemUserId = 4

/**
 * 订单相关
 */
const ORDER SystemUserId = 5

//事项提醒
const BUSINESS_NOTI SystemUserId = 6

//天使谷
const AngelVally SystemUserId = 7

//钱包
const WALLET SystemUserId = 8

// 人事科小薪
const PERSONNEL_SALARY SystemUserId = 11

// 药剂科小药
const PHARMACY_MEDICINE SystemUserId = 12

// 小组
const GROUP SystemUserId = 13

// 开药门诊
const PRESCRIBE_OUTPATIENT SystemUserId = 14

// 健康咨询
const HEALTH_ADVICE SystemUserId = 15

// 患者报道
const PATIENT_REPORT SystemUserId = 16

// 医联经纪人
const MED_BROKER SystemUserId = 17

func GetSystemUserId(userId int64) SystemUserId {
	return SystemUserId(userId)
}

/**
获取系统名字
*/
func (s SystemUserId) GetSystemName() string {
	switch s {
	case SYSTEM: return "系统通知"
	case BUSINESS: return "业务处理"
	case NEW_FRIEND: return "新的朋友"
	case DOCTOR_ASSISTANT: return "医生助手"
	case TREND: return "动态相关"
	case ORDER: return "订单相关"
	case BUSINESS_NOTI: return "事项提醒"
	case AngelVally: return "羽毛通知"
	case WALLET: return "钱包通知"
	case PERSONNEL_SALARY: return "人事科小薪"
	case PHARMACY_MEDICINE: return "药剂科小药"
	case GROUP: return "小组"
	case PRESCRIBE_OUTPATIENT: return "开药门诊"
	case HEALTH_ADVICE: return "健康咨询"
	case PATIENT_REPORT: return "患者报道"
	default: return ""
	}
}

func (s SystemUserId) GetHomeUrl() string {
	switch s {
	case GROUP: return "/medrn?moduleName=Group&routeName=GRecommand"
	case PRESCRIBE_OUTPATIENT: return "/patient/workbench"
	case HEALTH_ADVICE: return "/hospital/inquiryupcoming"
	case PATIENT_REPORT: return "/medrn?moduleName=Patient&routeName=PatientReport"
	case MED_BROKER:
		switch os.Getenv("DEPLOY_ENV") {
		case "pro", "fat": return "/link?url=https%3A%2F%2Fweb.medlinker.com%2Fdoctor%2FappH5%2FmyDrugBroker"
		default: return "/link?url=https%3A%2F%2Fweb-qa.medlinker.com%2Fdoctor%2FappH5%2FmyDrugBroker"
		}
	default: return ""
	}
}
func main() {
	fmt.Println(os.Getenv("GRADLE_HOME"))
	fmt.Println("-----", GetSystemUserId(17).GetHomeUrl())
}
