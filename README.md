
Compile instructions for docker:

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o release

Install RPC:
   gRPC Tooling:
go get -u google.golang.org/grpc

   Protocol buffer tooling:
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go

protoc -I ./pkg/pb ./pkg/proto/release.proto --go_out=plugins=grpc:./pkg
