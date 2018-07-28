#!/bin/bash
#creating at 28/07/2018
#created by viyancs
mkdir go-currency
cd go-currency
git clone https://github.com/viyancs/go-currency.git .
cd src
docker build -t viyancs/golang-mysql-api --no-cache=true .
docker run -p 3306:3306 --name mysqldb -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test_go -e MYSQL_ROOT_HOST=% -d mysql:5.7
docker run -p 8282:8282 --link mysqldb:mysqldb --name currency-go -d viyancs/golang-mysql-api
docker exec -ti currency-go sh -c "cd cmd/schema && go build -a -o bintask task.go && ./bintask orm syncdb -force && ./bintask sample"
