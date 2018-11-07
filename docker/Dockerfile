FROM golang:1.8.3 as builder
RUN go get github.com/MrManz/AWS_TestApp
RUN go get github.com/go-martini/martini
WORKDIR /go/src/github.com/MrManz/AWS_TestApp
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/MrManz/AWS_TestApp .
CMD ["./hello"]
EXPOSE 3000