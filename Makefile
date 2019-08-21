.PHONY: setup eventsgateway maestro

setup:
	@rm -rf $$GOPATH/src/github.com/golang/protobuf
	@go get -d -u github.com/golang/protobuf/protoc-gen-go
	@git -C $$GOPATH/src/github.com/golang/protobuf checkout v1.2.0 
	@go install github.com/golang/protobuf/protoc-gen-go
	@dep ensure

eventsgateway:
	@protoc -I eventsgateway/grpc/protobuf/ eventsgateway/grpc/protobuf/*.proto --go_out=plugins=grpc:eventsgateway/grpc/generated
	@mockgen --source eventsgateway/grpc/generated/events.pb.go --destination eventsgateway/grpc/mock/mock.go

maestro:
	@protoc -I maestro/grpc/protobuf/ maestro/grpc/protobuf/*.proto --go_out=plugins=grpc:maestro/grpc/generated
