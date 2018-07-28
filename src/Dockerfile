FROM golang:1.10.2
LABEL maintainer "viyancs"

EXPOSE 8282

RUN apt-get install tzdata
ENV TZ Asia/Jakarta

ENV GOPATH /currency
WORKDIR $GOPATH/src
COPY . .

RUN go get -u github.com/go-gem/gem
RUN go get -u github.com/astaxie/beego/orm
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/davecgh/go-spew/spew
RUN go build -a -o binmain main.go

CMD ["./binmain"]
