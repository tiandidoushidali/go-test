package test_mongodb

type PatientArchivesAssay struct {
	IllId int64 `json:"illId" bson:"illId"`
	IllName string `json:"illName" bson:"illName"`
	WhetherRelapse int `json:"whetherRelapse" bson:"whetherRelapse"`
	RelapseTime int64 `json:"relapseTime" bson:"relapseTime"`
	TreatmentProcess [] struct{
		Id int64 `json:"id" bson:"id"`
		TreatmentEffect string `json:"treatmentEffect" bson:"treatmentEffect"`
		TreatmentEffectTime int64 `json:"treatmentEffectTime" bson:"treatmentEffectTime"`
		TreatmentEnd int64 `json:"treatmentEnd" bson:"treatmentEnd"`
		TreatmentStart int64 `json:"treatmentStart" bson:"treatmentStart"`
		TreatmentPlan []string `json:"treatmentPlan" bson:"treatmentPlan"`
	}
}

