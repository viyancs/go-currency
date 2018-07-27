FROM golang:1.10.2
LABEL maintainer "viyancs"

EXPOSE 8282



ENV GOPATH /currency
WORKDIR $GOPATH
COPY cmd test models main.go ./

RUN go get -u github.com/go-gem/gem
RUN go get -u github.com/astaxie/beego/orm
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/davecgh/go-spew/spew
RUN go build -a -o binmain main.go

CMD ["./binmain"]
