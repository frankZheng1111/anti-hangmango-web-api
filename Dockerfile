FROM golang:latest

WORKDIR $GOPATH/src/anti-hangmango-web-api
COPY . $GOPATH/src/anti-hangmango-web-api
RUN go build .

ENTRYPOINT ["./anti-hangmango-web-api"]
