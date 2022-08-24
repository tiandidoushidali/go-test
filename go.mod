module go-test

go 1.14

require (
	git.medlinker.com/golang/xerror v1.0.3
	git.medlinker.com/service/common v1.2.8
	git.medlinker.com/service/grpcwrapper v1.0.0
	github.com/BurntSushi/toml v0.3.1
	github.com/Shopify/sarama v1.23.1
	github.com/aws/aws-sdk-go v1.43.21 // indirect
	github.com/beanstalkd/go-beanstalk v0.1.0
	github.com/bwmarrin/snowflake v0.3.0
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/go-kratos/kratos v1.0.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4 // indirect
	github.com/influxdata/influxdb v1.7.9
	github.com/influxdata/influxdb-client-go/v2 v2.3.1-0.20210518120617-5d1fff431040
	github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	github.com/jinzhu/copier v0.3.4
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/lithammer/shortuuid v2.0.3+incompatible
	github.com/opentracing/opentracing-go v1.2.0
	github.com/panjf2000/gnet/v2 v2.1.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/common v0.26.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/serialx/hashring v0.0.0-20160406013744-cba55573cd94
	github.com/shopspring/decimal v1.3.1
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/assertions v1.1.1 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/speps/go-hashids v1.0.0
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/tealeg/xlsx v1.0.5
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	go.mongodb.org/mongo-driver v1.4.4
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20220803195053-6e608f9ce704 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)

replace (
	github.com/satori/go.uuid => github.com/satori/go.uuid v1.2.0
	//github.com/ugorji/go => github.com/ugorji/go v1.1.2-0.20180728093225-eeb0478a81ae
	google.golang.org/grpc => google.golang.org/grpc v1.15.0
)
