
GITCOMMIT=$(shell git describe --match 'v[0-9]*' --dirty='.m' --always)
BUILDTIME=$(shell date -u '+%Y%m%d.%I%M%S%p')
VERSION=0.0.1
GO_LDFLAGS=-ldflags "-X `go list ./version`.Version=$(VERSION) -X `go list ./version`.BUILDTIME=$(BUILDTIME) -X `go list ./version`.GITCOMMIT=$(GITCOMMIT) -w"

TAG=dev
PREFIX=dhub.yunpro.cn/shenshouer/fngrok

build: ## build the go packages
	@echo "🐳 $@"
	@mkdir -p ./bin/client && mkdir -p ./bin/server
	@go build -a -installsuffix cgo ${GO_LDFLAGS} -o ./bin/client/fngrok ./main/client/main.go
	@go build -a -installsuffix cgo ${GO_LDFLAGS} -o ./bin/server/fngrok ./main/server/main.go

image: clean
	@echo "🐳 $@"
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ${GO_LDFLAGS} -o ./bin/server/fngrok ./main/server/main.go
	@docker build -t $(PREFIX):$(TAG) .
	@docker push $(PREFIX):$(TAG)
	
clean:
	@rm -f distribute