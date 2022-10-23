# go-work-first

通过go work很方便能引用本地包, 从此, 可以轻而易举将gopath项目改成go mod项目

go 1.8版本之后开始, 支持go work

## 开始
### 1.go work常用命令

1. 每个文件夹中执行 `go mod init xxx`
```
$ go mod init go-work-first
$ cd controllers
$ go mod init controllers
$ cd models
$ go mod init models
```
2. 根目录适用`go work`
```
$ go work init .
$ go work use models
$ go work use controllers

```
3. 运行
```
go run main.go
``` 

### 2. gopath项目 => go mod 项目

结合上面的知识, 其实只需要把mod名前面加上项目名即可, 比如: 
``` shell
$ cd controllers
# 这里加上项目名
$ go mod init go-work-first/controllers
```
#### 一个改造的例子 [go-api](https://github.com/relax-space/go-api/commit/1b516e03354382044e8e57992758256f7f11f53d)
1. 修改引用 
2. 新增mod
3. 新增work






