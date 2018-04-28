# go-web-server docker file.
FROM golang:1.10 AS builder
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
RUN mkdir -p /go/src/angular-go-web-app/go
WORKDIR /go/src/angular-go-web-app/go
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app ./
ENTRYPOINT ["./app"] 
EXPOSE 9090