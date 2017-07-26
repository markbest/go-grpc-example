## go-grpc-example
golang下使用grpc的一个简单例子

## 编译proto
- protobuf的相关配置网上很多，这里不再讲述
- protoc --go_out=plugins=grpc:"server" "protos/*.proto"

## 配置使用
- 拷贝conf/conf.toml.exmpale为conf/conf.toml，并配置数据库信息
- 执行服务端程序：go run server.go
- 执行客户端程序：go run client.go  

## 单元测试
```
go test -v
=== RUN   Test_GetArticleInfo
--- PASS: Test_GetArticleInfo (0.03s)
        article_test.go:26: golang
=== RUN   Test_GetArticleList
--- PASS: Test_GetArticleList (0.01s)
        article_test.go:41: 20
=== RUN   Test_GetCategoryInfo
--- PASS: Test_GetCategoryInfo (0.00s)
        category_test.go:22: golang
=== RUN   Test_GetCategoryList
--- PASS: Test_GetCategoryList (0.00s)
        category_test.go:37: 8
PASS
ok      grpc/tests      0.785s
```

