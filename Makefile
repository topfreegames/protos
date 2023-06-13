.PHONY: setup eventsgateway maestro

# TODO(caio.rodrigues): Att to generate in new folder using new protoc-gen pattern
setup:
	@rm -rf $$GOPATH/src/google.golang.org/protobuf
	@go get -d -u google.golang.org/protobuf/cmd/protoc-gen-go
	@git -C $$GOPATH/src/google.golang.org/protobuf checkout v1.2.0 
	@go install google.golang.org/protobuf/cmd/protoc-gen-go

	@dep ensure

eventsgateway:
	@protoc -I eventsgateway/grpc/protobuf/ eventsgateway/grpc/protobuf/*.proto --go_out=plugins=grpc:eventsgateway/grpc/generated
	@mockgen --source eventsgateway/grpc/generated/events.pb.go --destination eventsgateway/grpc/mock/mock.go

maestro:
	@protoc --proto_path=maestro/grpc/protobuf/ maestro/grpc/protobuf/*.proto --go_out=plugins=grpc:maestro/grpc/generated --go_opt=paths=maestro/grpc/protobuf/*.proto maestro/grpc/protobuf/maestro_events.proto
	
	# @protoc -I maestro/grpc/protobuf/ maestro/grpc/protobuf/*.proto --go_out=plugins=grpc:maestro/grpc/generated

healthcheck:
	@protoc --go_out=./maestro/grpc/generated --go_opt=paths=source_relative   --go-grpc_out=./maestro/grpc/generated --go-grpc_opt=paths=source_relative maestro/grpc/protobuf/health-check.proto