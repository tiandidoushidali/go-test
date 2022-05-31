package main

func main() {

	////通过grpc 库 建立一个连接
	//conn, err := grpc.Dial("172.20.17.225:8000", grpc.WithInsecure())
	//if err != nil {
	//	return
	//}
	//defer conn.Close()
	////通过刚刚的连接 生成一个client对象。
	//c := pro_med.NewMessageClient(conn)
	//resp, err := c.MessageGetHistory(context.Background(), &pro_med.MessageGetHistoryReq{
	//	GroupId:              12770939,
	//	Limit:                50,
	//	Start:                0,
	//	Sort:                 1,
	//})
	//if err != nil {
	//	fmt.Println("----", err)
	//}
	//fmt.Println("----resp.List----", len(resp.List))
	//for _, v := range resp.List {
	//	data := v
	//	decoded, _ := base64.StdEncoding.DecodeString(data)
	//	//fmt.Println("---decoded-----", decoded)
	//	msg := &message.Message{}
	//	proto.Unmarshal(decoded, msg)
	//	if strings.Contains(msg.GetJson().GetContent(), "2204249163150638") {
	//
	//		b, _ := json.Marshal(msg)
	//		fmt.Println(string(b))
	//	}
	//}
	//fmt.Println("-----resp-----", resp.Index)
}
