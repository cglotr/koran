FROM golang:1.13.1
WORKDIR /go/src/app
RUN go get -d -v firebase.google.com/go
RUN go get -d -v github.com/go-sql-driver/mysql
COPY ./src .
COPY ./src /go/src/github.com/cglotr/koran/server/src
RUN go install -v ./...
CMD ["app"]
