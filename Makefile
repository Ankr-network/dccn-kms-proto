# googleapis: github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
# here api path is /home/mobius/project
.PHONY: proto
proto:
	protoc --proto_path=/home/mobius/project:. --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. kv/kv.proto
