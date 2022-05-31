module go-test

go 1.14

require (
	git.medlinker.com/golang/xerror v1.0.3
	git.medlinker.com/service/common v1.2.8
	git.medlinker.com/service/grpcwrapper v1.4.0
	github.com/BurntSushi/toml v0.3.1
	github.com/beanstalkd/go-beanstalk v0.1.0
	github.com/bwmarrin/snowflake v0.3.0
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/go-kratos/kratos v1.0.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4 // indirect
	github.com/influxdata/influxdb v1.7.9
	github.com/influxdata/influxdb-client-go/v2 v2.3.1-0.20210518120617-5d1fff431040
	github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	github.com/jinzhu/copier v0.3.4
	github.com/jinzhu/gorm v1.9.16
	github.com/lithammer/shortuuid v2.0.3+incompatible
	github.com/olivere/elastic/v7 v7.0.32
	github.com/opentracing/opentracing-go v1.2.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/common v0.26.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/serialx/hashring v0.0.0-20200727003509-22c0c7ab6b1b
	github.com/shopspring/decimal v1.3.1
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/goconvey v1.6.4
	github.com/speps/go-hashids v1.0.0
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tealeg/xlsx v1.0.5
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/xwb1989/sqlparser v0.0.0-20180606152119-120387863bf2
	go.mongodb.org/mongo-driver v1.4.4
	go.opentelemetry.io/otel v1.7.0 // indirect
	go.opentelemetry.io/otel/internal/metric v0.27.0 // indirect
	go.opentelemetry.io/otel/metric v0.30.0 // indirect
	go.opentelemetry.io/otel/oteltest v0.20.0 // indirect
	go.uber.org/zap v1.17.0
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.3.4
	gorm.io/gorm v1.23.5
)

replace (
	github.com/satori/go.uuid => github.com/satori/go.uuid v1.2.0
	//github.com/ugorji/go => github.com/ugorji/go v1.1.2-0.20180728093225-eeb0478a81ae
	google.golang.org/grpc => google.golang.org/grpc v1.15.0
)
