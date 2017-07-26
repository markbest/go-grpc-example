## go-grpc-example
golang下使用grpc的一个简单例子

## 编译proto文件方法
- protobuf的相关配置网上很多，这里不再讲述
- protoc --go_out=plugins=grpc:"server" "protos/article.proto"

## 使用方法
- 拷贝conf/conf.toml.exmpale为conf/conf.toml，并配置数据库信息
- 执行服务端程序：go run server.go
- 执行客户端程序：go run client.go
