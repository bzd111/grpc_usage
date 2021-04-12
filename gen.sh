protoc -I proto proto/order_management.proto \
    -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
    --go_out=. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:.
