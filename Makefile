# Copyright 2024, Tencent Inc.
# Author: yijian
# Date: 2024/02/02

all: mooon_auth_example

mooon_auth_example: mooon_auth.go
	go build -o $@ $<

.PHONY: rpc tidy clean fetch

clean:
	rm -fr mooon_auth_example

rpc:
	goctl rpc protoc mooon_auth.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=. --style=go_zero

tidy:
	go mod tidy

fetch: # 强制用远程仓库的覆盖本地，运行时需指定分支名，如：make fetch branch=main
	git fetch --all&&git reset --hard origin/$$branch