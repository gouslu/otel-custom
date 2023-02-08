FROM golang:1.19-bullseye

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# RUN go install github.com/go-delve/delve/cmd/dlv@latest

# # Build
RUN go build -gcflags="all=-N -l" -o /otel-custom

# # This is for documentation purposes only.
# # To actually open the port, runtime parameters
# # must be supplied to the docker command.
# # EXPOSE 8080 40000

# # (Optional) environment variable that our dockerised
# # application can make use of. The value of environment
# # variables can also be set via parameters supplied\
# # to the docker command on the command line.
# # ENV HTTP_PORT=8081

RUN cp ./config.yaml /config.yaml

RUN ls /

# # Run
CMD [ "/otel-custom", "--config", "/conf/config.yaml" ]
# CMD ["/go/bin/dlv", "--listen=:40001", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/otel-emitter"]

# CMD ["go", "run", "main.go"]

# CMD ["/go/bin/dlv", "exec", "--listen=:40000", "--headless", "--api-version=2", "/otel-emitter"]