FROM golang:1.13.1
WORKDIR /go/src/app
COPY ./src .
COPY ./src /go/src/github.com/arikama/koran/server/src
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["app"]
