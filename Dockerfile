FROM golang:1.16.3 AS builder

WORKDIR /go/src/github.com/pjtuxe/go.cron/
RUN go get -d -v golang.org/x/net/html

# Copy files to build
COPY ./src/app.go .
COPY go.mod .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/pjtuxe/go.cron/app .
CMD [ "./app" ]
