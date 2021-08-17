FROM golang:alpine AS builder

WORKDIR $GOPATH/src/github.com/avgalani/test-api

COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/test-api

FROM scratch

COPY --from=builder /go/bin/test-api /go/bin/test-api

EXPOSE 8080

ENTRYPOINT ["/go/bin/test-api"]
