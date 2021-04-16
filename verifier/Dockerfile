FROM golang:1.16.3 AS builder

# Copy files to build
COPY . /go/src/go.cron
WORKDIR /go/src/go.cron/verifier
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o verifier .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/go.cron/verifier .
CMD [ "./verifier" ]
