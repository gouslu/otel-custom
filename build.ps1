# protoc -I=proto --go_out=pb --go-grpc_out=pb proto/*.proto
# go build

docker build -t otel-custom -f ./Dockerfile .
