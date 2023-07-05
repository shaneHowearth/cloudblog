module github.com/shanehowearth/cloudblog/readarticle

go 1.12

replace github.com/shanehowearth/cloudblog => /home/shane/GoLang/src/github.com/shanehowearth/cloudblog

require (
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/rafaeljusto/redigomock v0.0.0-20191009171320-7be4dd8da67a
	github.com/stretchr/testify v1.8.1
	google.golang.org/grpc v1.53.0
)
