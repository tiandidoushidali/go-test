module go-test

go 1.14

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/beanstalkd/go-beanstalk v0.1.0
	github.com/coreos/etcd v3.3.10+incompatible
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/go-kratos/kratos v1.0.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4 // indirect
	github.com/influxdata/influxdb v1.7.9
	github.com/influxdata/influxdb-client-go/v2 v2.3.1-0.20210518120617-5d1fff431040
	github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	github.com/jinzhu/copier v0.3.4
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/lithammer/shortuuid v2.0.3+incompatible
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/common v0.26.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/serialx/hashring v0.0.0-20200727003509-22c0c7ab6b1b
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/goconvey v1.6.4
	github.com/speps/go-hashids v1.0.0
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tealeg/xlsx v1.0.5
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	go.etcd.io/bbolt v1.3.3 // indirect
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.39.0
)

replace (
	github.com/satori/go.uuid => github.com/satori/go.uuid v1.2.0
	github.com/ugorji/go => github.com/ugorji/go v1.1.2-0.20180728093225-eeb0478a81ae
	google.golang.org/grpc => google.golang.org/grpc v1.15.0
)
