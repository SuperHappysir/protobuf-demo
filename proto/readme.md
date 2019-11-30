### 安装编译器
```bash
cd /tmp

wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.0/protoc-3.11.0-linux-x86_64.zip

mkdir -p protoc

unzip protoc-3.11.0-linux-x86_64.zip -d protoc/

ln -s $(pwd)/protoc/bin/protoc  /usr/local/bin/
```

### go使用protobuf需要安装一些其他包
```bash
go get -u google.golang.org/grpc					# 安装 gRPC 框架
go get -u github.com/golang/protobuf/protoc-gen-go	# 安装 Go 版本的 protobuf 编译器
```

### 生成go struct
```bash
# protoc 编译器的 grpc 插件会处理 service 字段定义的 UserInfoService
# 使 service 能编码、解码 message
protoc -I . --go_out=plugins=grpc:. ./user.proto
```
