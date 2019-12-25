FROM golang:1.13.1
WORKDIR /go/src/app
RUN go get firebase.google.com/go
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux
COPY ./src .
COPY ./src /go/src/github.com/cglotr/koran/server/src
RUN go install ./...
CMD ["app"]
