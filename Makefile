# googleapis: github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
# here api path is /home/mobius/project
.PHONY: proto
proto:
	protoc --proto_path=. --go_out=plugins=grpc:. kv/kv.proto
	protoc --proto_path=/home/mobius/project:. --grpc-gateway_out=logtostderr=true:. kvhttp/kv_http.proto
