

.PHONY: grpc
# 编译pb协议文件
grpc:
	protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative tayzhangtest.proto
