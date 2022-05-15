USER ?= dhf0820
NS ?= dhf0820
TAG ?= latest
ARC= amd64
TEST ?= dhf0820
PROD ?= vertisoft
VERSION ?= $(TAG)
ARCH ?= $(ARC)
IMG_NAME ?= release
IMAGE_NAME ?= $(IMG_NAME)_$(ARCH)
IMAGE_TEST_NAME ?= $(IMG_NAME)_test_$(ARCH)
#LINUX_IMAGE_NAME ?= release

SERVER_OUT := "cmd/server"
CLIENT_OUT := "cmd/client"

##API_OUT := "api/api.pb.go"
PB_OUT := pkg/pb/*.pb.go
API_REST_OUT := "api/api.pb.gw.go"
API_SWAG_OUT := "api/api.swagger.json"
PKG := "gitlab.com/dhf0820/ids_release_service"
SRC := /Users/dhf/Dropbox/1-work/ROI/Release
SERVER_PKG_BUILD := "${PKG}/cmd/server"
CLIENT_PKG_BUILD := "${PKG}/cmd/client"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
BINARY=$(IMG_NAME)_$(ARCH)
DOCKER_NAME=$(IMG_NAME)_$(ARCH)
LINUX_NAME=$(IMG_NAME)
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=release
# BINARY_UNIX=$(BINARY_NAME)_linux
# BINARY_MAC=$(IMAGE_NAME)
# MAC_IMAGE_NAME= $(BINARY_MAC)_mac


.PHONY: all api server client cert pb1

all: server client

pb1:
    @protoc -I proto proto/*.proto --go_out=plugins=grpc:./pb

#api/api.pb.go:
#protoc -I ./ --proto_path=./ --go_out=./ pkg/proto/*.proto
api:
	@protoc -I ./protobufs/ \
		--proto_path=./ \
		--go_out=plugins=grpc:./ \
		./protobufs/*.proto


#	@protoc -I ./protobufs \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--proto_path=./ \
		--go_out=plugins=grpc:./ \
		./protobufs/*.proto



#api: api/api.pb.go api/api.pb.gw.go api/api.swagger.json ## Auto-generate grpc go sources

dep: ## Get the dependencies
	@go get -v -d ./...

tidy: # add all new includes
	@go mod tidy


build:
	go build -o $(BINARY_NAME)
	#$(GOBUILD) -o $(BINARY_NAME)

docker-build:
	docker build -t $(NS)/$(IMG_NAME):$(VERSION) -f Dockerfile .

docker-push: # push to docker
	docker push $(NS)/$(IMG_NAME):$(VERSION)

mac:
	ARCH=arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY) -v

linux_test:
	go build -o $(BINARY_NAME)
	docker build -t $(TEST)/$(IMG_NAME):$(VERSION) -f Dockerfile .
	docker push $(NS)/$(IMG_NAME):$(VERSION)


local_linux:
	go build -o $(BINARY_NAME)
	docker build -t $(TEST)/$(IMG_NAME):$(VERSION) -f Dockerfile .



test:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY) -v
	docker build -t $(TEST)/$(DOCKER_NAME):$(VERSION) -f Dockerfile_$(ARCH) .
	docker push $(TEST)/$(DOCKER_NAME):$(VERSION)

prod:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY) -v
	docker build -t $(PROD)/$(DOCKER_NAME):$(VERSION) -f Dockerfile_$(ARCH) .
	docker push $(PROD)/$(DOCKER_NAME):$(VERSION)

	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	# docker build -t $(PROD)/$(LINUX_IMAGE_NAME):$(VERSION) -f Dockerfile .
	# docker push $(PROD)/$(LINUX_IMAGE_NAME):$(VERSION)

release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	docker build -t $(NS)/$(LINUX_IMAGE_NAME):$(VERSION) -f Dockerfile .
	docker push $(NS)/$(LINUX_IMAGE_NAME):$(VERSION)

build-linux:
	ARCH=amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY) -v

docker-build:
	docker buildx -t $(NS)/$(LINUX_IMAGE_NAME):$(VERSION) -f Dockerfile_$(ARCH) .

docker-push: # push to docker
	docker push $(NS)/$(LINUX_IMAGE_NAME):$(VERSION)

mac:
	ARCH=arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) $(GOBUILD) -o $(BINARY) -v


server: dep api ## Build the binary file for server
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o release
	#@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

client: dep api ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(PB_OUT) $(API_REST_OUT) $(API_SWAG_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run-server:
	go run main.go -port 8080

run-server-tls:
	go run cmd/server/main.go -port 9901 -tls server

run-server-mutual-tls:
	go run cmd/server/main.go -port 7777 -tls mutual

run-client-do3:
	go run src/client/main.go  -address docker-test.vertisoft.com -port 8080

run-client:
	go run src/client/main.go  -address localhost -port 9001

run-client-do-test:
	go run src/client/main.go  -address 161.35.229.137 -port 30001

run-client-tls:
	go run cmd/client/main.go  -address 0.0.0.0:7777 -tls server

run-client-mutual-tls:
	go run cmd/client/main.go  -address 0.0.0.0:7777 -tls mutual

cert:
	cd cert; ./gen.sh; cd ..