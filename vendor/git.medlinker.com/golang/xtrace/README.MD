# xtrace

分布式链路跟踪SDK。

## 特点

1. 支持层级`SpanId`
1. 支持链路`TraceId`
1. 支持自定义`Key-Value`传递
1. 支持通过`http.Header/*http.Header/context.Context`参数创建`Trace`对象
1. 支持通过`grpc.OutgoingContext/grpc.IncomingContext`参数创建`Trace`对象
1. 支持通过`Trace`对象创建`HTTPHeader/context.Context/grpcContext`对象，用于不同接口参数传递


## 使用

### 1. 创建`Trace`对象
```go
// 通过HTTP Header创建Trace对象，
// 自动读取Header中的链路信息
trace := xtrace.GetTrace(header)
```
```go
// 通过标准库context.Context创建Trace对象，
// 该ctx接口对象也支持grpc请求传递过来的context.Context对象，
// 内部会自动识别ctx类型并自动读取链路参数
trace := xtrace.GetTrace(ctx)
```
### 2. 创建`HTTP Header`请求参数
```go
// 将链路信息序列化并生成HTTP Header Map，
// 用于下一步HTTP请求中将链路信息通过Header方式提交过去
header := trace.HTTPHeader()
```

### 3. 创建`context.Context`请求参数
```go
// StdContext生成一个标准库的context.Context对象，
// 用于方法上下文之间传递链路信息
ctx := trace.StdContext()
```
```go
// GRPCContext生成一个GRPC的context.Context对象，
// 用于发起GRPC请求的时候传递链路信息。
// 注意：该ctx也支持标准库方式读取链路信息。
ctx := trace.GRPCContext()
```

### 4. 自定义链路参数
```go
// 创建Trace对象
trace := xtrace.GetTrace(ctxOrHeader)
// 设置自定义链路参数
trace.TraceMap.Set("UserId", 123456)
// 创建用于GRPC请求的上下文接口对象
ctx := trace.GRPCContext()

// 将ctx使用到GRPC请求...
```

### 5. 其他方法
通过`http.Header/*http.Header/context.Context`参数创建相对应所需的链路信息传递对象，并且从这些参数中获取指定键名的参数。
```go
func CreateHTTPHeader(ctxOrHeader interface{}) map[string]string 
func CreateStdContext(ctxOrHeader interface{}) context.Context 
func CreateGRPCContext(ctxOrHeader interface{}) context.Context 
func GetValueFrom(ctxOrHeader interface{}, key string) string
```





