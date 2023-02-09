# protoc -I=proto --go_out=pb --go-grpc_out=pb proto/*.proto
# go build

param ([switch] $Debug)

if ($Debug) {
        docker build -t otel-custom -f ./Dockerfile.debug .
}
else {
    docker build -t otel-custom -f ./Dockerfile .
}