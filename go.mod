module github.com/jdxj/video

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-micro/v3 v3.0.0-beta.2.0.20200911124113-3bb76868d194
	github.com/micro/micro/v3 v3.0.0-beta.4
	github.com/prometheus/client_golang v1.7.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20200909081042-eff7692f9009 // indirect
	golang.org/x/tools v0.0.0-20200117065230-39095c1d176c // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.3.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
