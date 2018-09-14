# Expose the compatibility test CLI.
# The CLI offers the following features:
#   * Generate test data: docker run --mount type=bind,source="$(pwd)"/samples,target=/samples stratumn/go-chainscript:latest generate /samples/go-samples.json
#   * Validate test data: docker run --mount type=bind,source="$(pwd)"/samples,target=/samples stratumn/go-chainscript:latest validate /samples/js-samples.json

FROM golang:alpine

RUN apk add make
RUN apk add git

RUN mkdir /samples

WORKDIR /go/src/github.com/stratumn/go-chainscript
ADD . .

RUN make deps

WORKDIR /go/src/github.com/stratumn/go-chainscript/cmd
RUN go build -o chainscript-cli
RUN mv chainscript-cli /go/bin

ENTRYPOINT [ "/go/bin/chainscript-cli" ]