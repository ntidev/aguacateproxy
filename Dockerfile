FROM golang:1.8 AS build-env
WORKDIR /go/src/app
COPY . .
RUN go get -d -v 
RUN GOOS=linux GOARCH=amd64 go build -o aguacateproxy

# final stage
FROM debian:wheezy-20180831
VOLUME /etc/aguacateproxy
COPY --from=build-env /go/src/app/aguacateproxy /bin/
CMD ["aguacateproxy"]