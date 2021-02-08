FROM golang:1.15.7-buster

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/handlers
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/sqlite
RUN go build
CMD ["go-web-template"]
