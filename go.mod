module github.com/jdxj/video

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20200909081042-eff7692f9009 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.3.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
