package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TransferPatientMessage struct {
	RecordId int64	`json:"id"`
	OldAssistantUserId uint32 `json:"old_assistant_user_id"`
	PatientIds []uint32	`json:""`
	DoctorIds []uint32
	AssistantUserId uint32	`json:"new_assistant_user_id"`
	Type int `json:"type"`	// 1-医生，2-患者
	TransferUserId []byte `json:"transfer_user_id"` // 转移用户ID(多个以,分隔)
}

type BinlogContent struct {
	BeforeInfo map[string]interface{}
	AfterInfo	map[string]interface{}
	Action string
	Schema string
	Table string
	Instance string
}

type Consumer struct {

	transferPatientMessage TransferPatientMessage
}

func (cm *Consumer) binlogContent2Message(binCtx BinlogContent) *Consumer {
	beforeInfo, _ := json.Marshal(binCtx.BeforeInfo) // map to json
	msg := TransferPatientMessage{}
	json.Unmarshal(beforeInfo, &msg)

	transferUserId := string(msg.TransferUserId)
	if transferUserId != "" {
		tuserStrIds := strings.Split(transferUserId, ",")
		tuserIds := make([]uint32, 0)

		for _, v := range tuserStrIds {
			tuserId, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				continue
			}
			tuserIds = append(tuserIds, uint32(tuserId))
		}

		switch msg.Type {
		case 1: // 医生
			msg.DoctorIds = tuserIds
			break
		case 2: // 患者
			msg.PatientIds = tuserIds
			break
		}
	}
	cm.transferPatientMessage = msg

	return cm
}

func (cm *Consumer) PrintMessage(trans TransferPatientMessage) {
	time.Sleep(1 * time.Second)
	fmt.Println("----1----:", trans.RecordId)
}

func (cm *Consumer) PrintMessage2(trans TransferPatientMessage) {
	time.Sleep(5 * time.Second)
	fmt.Println("----2----:", trans.RecordId)
}

func main() {
	cm := &Consumer{TransferPatientMessage{
		RecordId:           1,
		OldAssistantUserId: 1,
	}}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for i := 0 ; i < 100; i ++ {
			time.Sleep(100 * time.Millisecond)
			cm.binlogContent2Message(BinlogContent{
				BeforeInfo: map[string]interface{}{"id":time.Now().Unix()},
				AfterInfo:  map[string]interface{}{"id":time.Now().Unix()},
			})
			ts := TransferPatientMessage{
				RecordId:           time.Now().Unix(),
			}
			go func() {
				cm.PrintMessage(ts)
			}()

			go func() {
				cm.PrintMessage2(ts)
			}()
		}
		time.Sleep(15 * time.Second)
		wg.Done()
	}()

	//wg.Add(1)
	//go func() {
	//	t := time.NewTicker(2 * time.Second)
	//	for {
	//		select {
	//		case <-t.C:
	//			cm.PrintMessage()
	//		}
	//	}
	//	wg.Done()
	//}()

	fmt.Println("------=====:", cm.transferPatientMessage.RecordId)

	wg.Wait()


}