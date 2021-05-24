#protoc grpcapi.proto --go-grpc_out==plugins=grpc:.
#install protoc tool
#go get -u google.golang.org/protobuf/cmd/protoc-gen-go
#go get u google.golang.org/grpc
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
    topic.proto