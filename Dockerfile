FROM golang:1.10-alpine3.7 as builder
WORKDIR /go/src/github.com/gophercon

RUN apk --no-cache add git
RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.lock Gopkg.toml ./ 
RUN dep ensure -vendor-only

COPY . ./
RUN go build

FROM alpine:3.7
COPY --from=builder /go/src/github.com/gophercon ./
CMD ["./gophercon"]