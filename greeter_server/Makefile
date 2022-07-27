
.PHONY: server
# 编译pb协议文件

IMAGE := test-app-grpc-server
TAG :=v1.0.0
REGISTRY := ccr.ccs.tencentyun.com/tiemsdev

server:
	echo "==========  Make Build  =========="
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	echo "========== Docker Build =========="
	docker build -t $(REGISTRY)/$(IMAGE):$(TAG) -f Dockerfile .
	echo "====== Docker Push (ccr.ccs) ======"
	docker push $(REGISTRY)/$(IMAGE):$(TAG)