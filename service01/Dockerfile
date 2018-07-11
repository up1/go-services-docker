FROM golang:alpine as builder
WORKDIR /go
COPY . /go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app main

FROM scratch
WORKDIR /root/
COPY --from=builder /go/app .
CMD ["./app"]
