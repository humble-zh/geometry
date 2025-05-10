# Geometry

这是一个用于几何计算的 Go 库，提供了矩形和圆形的面积、周长计算功能。

## 安装

```bash
go get github.com/humble-zh/geometry
```

## doc
```bash
go doc -all github.com/humble-zh/geometry
go doc github.com/humble-zh/geometry.Rectangle.Area
```

## test
- 文件名`<文件名>_test.go`
```bash
go test ./... -v      # ./... 递归子目录, -v 显示每个测试函数的执行结果
go test -cover ./...  # 显示测试覆盖率
go test -race ./...   # 检测竞态条件
go test -v -run Area ./   # 只运行匹配正则(Area字符串)的测试用例
```

## examples
- 文件名`geometry_test/example_<文件名>_test.go`
    - 函数`func Example<文件名>_<函数名>(){}`
- 文件名`main/main.go`
    - `go run examples/main/main.go`