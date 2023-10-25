module github.com/shanehowearth/cloudblog/readarticle

go 1.12

replace github.com/shanehowearth/cloudblog => /home/shane/GoLang/src/github.com/shanehowearth/cloudblog

require (
	github.com/golang/protobuf v1.5.3
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/rafaeljusto/redigomock v0.0.0-20191009171320-7be4dd8da67a
	github.com/stretchr/testify v1.8.3
	golang.org/x/net v0.17.0 // indirect
	google.golang.org/grpc v1.56.3
)
