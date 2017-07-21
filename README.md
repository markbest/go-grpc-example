# go-grpc-example
golang下使用grpc的一个简单例子

# 编译proto文件方法
protoc --go_out=plugins=grpc:"server" "protos/article.proto"

# 测试执行
- go run server.go (服务端)
- go run client.go (客户端)
