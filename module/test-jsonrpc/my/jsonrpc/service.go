package jsrpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/pkg/log"
	"go/token"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

type JsonRpcService struct {
	serviceMap sync.Map
}

type JsonRpcSvr interface {
	SvrName() string
}

type JsonRpcReq struct {
	Id int `json:"id"`
	Method string `json:"method"`
	Params JsonRpcParams `json:"params"`
	JsonRpc string `json:"jsonrpc"`
}

//type JsonRpcReq struct {
//	Id int `json:"id"`
//	Method string `json:"method"`
//	Params []interface{} `json:"params"`
//	JsonRpc string `json:"jsonrpc"`
//}

type JsonRpcParams struct {
	AppId string `json:"appId"`
	TimeStamp int64 `json:"timeStamp"` // 时间戳
	Sign string `json:"sign"` // 签名
	Payload map[string]interface{} `json:"payload"`
}

type JsonRpcOk struct {
	Id int `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result interface{} `json:"result"`
}

type JsonRpcErr struct {
	Id int `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Error struct{
		Code int64 `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

var jsonContentType = []string{"application/json; charset=utf-8"}

// Render render it to http response writer.
func (render JsonRpcOk) Render(w http.ResponseWriter) error {
	jsonBytes, err := json.Marshal(render)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

// WriteContentType write content-type to http response writer.
func (render JsonRpcOk) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = jsonContentType
	}
}

// Render render it to http response writer.
func (render JsonRpcErr) Render(w http.ResponseWriter) error {
	jsonBytes, err := json.Marshal(render)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

// WriteContentType write content-type to http response writer.
func (render JsonRpcErr) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = jsonContentType
	}
}

func New() *JsonRpcService {
	s := &JsonRpcService{}

	// 将jsonrpc 服务进行注册
	s.register(NewBudgetJsonRpc())

	return s
}


var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

type methodType struct {
	sync.Mutex // protects counters
	method reflect.Method
	ArgType reflect.Type
	ReplyType reflect.Type
	numCalls uint
}

type JsonRpcServer struct {
	name string // name of service
	rcvr reflect.Value // receiver of methods for the service
	typ reflect.Type // type of the receiver
	method map[string]*methodType  // register method
}

type JsonRpcCall struct {
	ServiceName string // 调用的服务名
	Method string // 方法名
	Args interface{}
	Reply interface{}
	Error error
}

func (server *JsonRpcService) Call(serverName, serviceMethod string, args string) (resp interface{}, err error){
	var build strings.Builder
	build.WriteString(strings.ToUpper(serviceMethod[:1]))
	build.WriteString(serviceMethod[1:])
	serviceMethod = build.String()

	call := new(JsonRpcCall)
	call.ServiceName = serverName
	call.Method = serviceMethod
	call.Args = args

	fmt.Println("----serverName----", serverName)
	svrInfc, ok := server.serviceMap.Load(serverName)
	if !ok {
		log.Error("JsonRpc Call server.serviceMap.Load(%s) fail svrInfc:%+v, ok:%v", serverName, svrInfc, ok)
		return
	}
	svr, _ := svrInfc.(*JsonRpcServer)
	var md *methodType
	if md, ok = svr.method[serviceMethod]; !ok {
		log.Error("JsonRpc svr.method[%s] is not exists", serviceMethod)
		err = errors.New("method not allow")
		return
	}
	methodFunc := md.method.Func


	fmt.Println("----args :", md.ArgType, md.ArgType.Elem())
	req := reflect.New(md.ArgType.Elem()).Interface()
	err = json.Unmarshal([]byte(args), &req)
	if err != nil {
		// 参数错误
		return
	}
	fmt.Println("----req can set", md.ArgType, req)
	//resp := reflect.New(md.ReplyType)
	//req.Set(reflect.ValueOf(&BudgetJsonRpcAddReq{
	//	A: 1,
	//	B: 2,
	//}))
	//returnsValues := methodFunc.Call([]reflect.Value{svr.rcvr, reflect.ValueOf(context.Background()), reflect.ValueOf(&BudgetJsonRpcAddReq{
	//	A: 1,
	//	B: 2,
	//})})
	returnsValues := methodFunc.Call([]reflect.Value{svr.rcvr, reflect.ValueOf(context.Background()), reflect.ValueOf(req)})
	fmt.Println("----returnsValues-----", returnsValues)

	if e, ok := returnsValues[1].Interface().(error); !ok {
		err = e
	}
	return returnsValues[0].Interface(), err
}

func (server *JsonRpcService) register(rcvr JsonRpcSvr) error {
	fmt.Println("-------jsonrpc register rcvr: ", rcvr)
	s := new(JsonRpcServer)
	s.typ = reflect.TypeOf(rcvr)
	s.rcvr = reflect.ValueOf(rcvr)
	sname := reflect.Indirect(s.rcvr).Type().Name()
	fmt.Println("-----sname----", sname)
	if sname == "" {
		s := "jsonRpc.Register: no service name for type" + s.typ.String()
		return errors.New(s)
	}

	if !token.IsExported(sname) {
		s := "rpc.Register: type " + sname + " is not exported"
		return errors.New(s)
	}
	s.name = rcvr.SvrName()
	fmt.Println("-----sname2----", s.name)

	s.method = suitableMethods(s.typ, true)

	fmt.Println("-----s.method---", s.method)
	if len(s.method) == 0 {
		str := ""
		method := suitableMethods(reflect.PtrTo(s.typ), false)
		if len(method) != 0 {
			str = "rpc.Register: type " + sname + " has no exported methods of suitable type (hint: pass a pointer to value of that type)"
		} else {
			str = "rpc.Register: type " + sname + " has no exported methods of suitable type"
		}
		log.Error(str)
		return errors.New(str)
	}

	log.Info("jsonRpc register start LoadOrStore (sname:%s, s:%+v)", s.name, s)
	if act, ok := server.serviceMap.LoadOrStore(s.name, s); ok {
		log.Info("jsonRpc register call server.serviceMap.LoadOrStore act:%+v, ok:%v ", act, ok)
		return errors.New("jsonRpc: service already defined: \" + sname")
	}
	ss, ok := server.serviceMap.Load(s.name)
	fmt.Println("------", ss, ok)
	return nil
}


func suitableMethods(typ reflect.Type, reportErr bool) map[string]*methodType {
	methods := make(map[string]*methodType)
	for m := 0; m < typ.NumMethod(); m ++ {
		method := typ.Method(m)
		mtype := method.Type
		mname := method.Name
		// Method must be export.
		if method.PkgPath != "" {
			continue
		}
		fmt.Println("-----params ln-----", mtype.NumIn())
		// Method needs three ins: receiver, *args, *reply
		// *jsrpc.BudgetJsonRpc context.Context *jsrpc.BudgetJsonRpcAddReq
		if mtype.NumIn() != 3 {
			if reportErr {
				log.Error("rpc.Register: method %q has %d input parameters; needs exactly three\n", mname, mtype.NumIn())
			}
			continue
		}

		// 请求：第三个参数为指针
		argType := mtype.In(2) //req
		if argType.Kind() != reflect.Ptr {
			if reportErr {
				log.Error("rpc.Register: req type of method %q is not a pointer: %q\n", mname, argType)
			}
			continue
		}
		// Reply 类型必须导出
		// 方法需要一个输出
		// *jsrpc.BudgetJsonRpcAddResp error
		if mtype.NumOut() != 2 {
			if reportErr {
				log.Error("rpc.Register: method %q has %d output parameters; needs exactly one\n", mname, mtype.NumOut())
			}
			continue
		}
		replyType := mtype.Out(0) //resp
		if replyType.Kind() != reflect.Ptr {
			if reportErr {
				log.Error("rpc.Register: req type of method %q is not a pointer: %q\n", mname, replyType)
			}
			continue
		}
		// 方法的返回类型必须是error
		if returnType := mtype.Out(1); returnType != typeOfError {
			if reportErr {
				log.Error("rpc.Register: return type of method %q is %q, must be error\n", mname, returnType)
			}
			continue
		}
		methods[mname] = &methodType{
			method:    method,
			ArgType:   argType,
			ReplyType: replyType,
		}
	}
	return methods
}