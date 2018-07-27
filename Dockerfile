FROM golang:1.10.2
LABEL maintainer "viyancs"

EXPOSE 8282
WORKDIR /currency
COPY . .

RUN export GOPATH=$PWD
RUN go get -u github.com/go-gem/gem
RUN go get -u github.com/astaxie/beego/orm
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/davecgh/go-spew/spew
RUN go build -o binmain main.go

CMD ["./binmain"]
