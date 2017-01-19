GO_LDFLAGS=-ldflags " -w"

TAG=dev
PREFIX=shenshouer/kubeadm

build: clean
	@echo "🐳 $@"
	@go build -a -installsuffix cgo ${GO_LDFLAGS} .

linux: clean
	@echo "🐳 $@"
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ${GO_LDFLAGS} .

image: clean
	@echo "🐳 $@"
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ${GO_LDFLAGS} .
	@docker build -t $(PREFIX):$(TAG) .
	@docker push $(PREFIX):$(TAG)
	
clean:
	@rm -f kubeadm