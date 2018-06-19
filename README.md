## grpc
Golang grpc以及grpc-gateway的简单使用

## 编译proto
- 安装ProtocolBuffers以及gate-gateway，可以查找网站相关资料
- 编译google.api
```
cd protos
protoc --go_out=plugins=grpc:"." "google/api/*.proto"
```

- 编译*.proto
```
cd protos
protoc --go_out=plugins=grpc:"." ./*.proto
```

- 编译*.proto.gw.go
```
cd protos
protoc --grpc-gateway_out=logtostderr=true:"." "*.proto"
```

## 配置使用
- 拷贝conf/conf.toml.exmpale为conf/conf.toml，并配置数据库信息
- 执行服务端程序：go run server.go
- 执行客户端程序：go run client.go  
- 执行gateway程序: go run gateway.go

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

## 性能测试

```
goos: windows
goarch: amd64
pkg: grpc/tests
Benchmark_GetArticleInfo-8          2000           1014170 ns/op
Benchmark_GetArticleList-8           300           4816491 ns/op
Benchmark_GetCategoryInfo-8         2000            734985 ns/op
Benchmark_GetCategoryList-8         2000            738987 ns/op
PASS
ok      grpc/tests      8.441s
```