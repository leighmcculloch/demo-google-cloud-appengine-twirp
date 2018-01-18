# Build the twirp server and generated code from the proto definition.
protoc: protoc-deps
	protoc --go_out=. --twirp_out=. ./*.proto

# Install protoc build dependencies.
protoc-deps:
	go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/twitchtv/twirp/protoc-gen-twirp
