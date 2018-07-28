This project is REST Full API for managing currency rates, using docker with 2 container first currency-go container  for web app second is mysql container for database ,if you want to  deploy this app please follow this guide :

Quick Installation
===========
$ chmod +x install.sh
$ ./install.sh

Manual Installation
==========
- $ mkdir go-currency
- $ cd go-currency
- $ git clone https://github.com/viyancs/go-currency.git .
- $ cd src
- $ docker build -t viyancs/golang-mysql-api --no-cache=true . (making image)
- $ docker run --name mysqldb -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test_go -e  MYSQL_ROOT_HOST=% -d mysql:5.7 (run mysql container)
- $ docker run -p 8282:8282 --name currency-go viyancs/golang-mysql-api (run web app container that provide http REST API)
- $ docker exec -ti currency-go sh -c "cd cmd/schema && go build -a -o bintask task.go && ./bintask orm syncdb -force && ./bintask sample" (build and rebuild table with data sample)


API Docs Usage
==========
GET http://localhost:8282/exchange (to get all exchange)
GET http://localhost:8282/exchange/(:date) (to get exchange by date example :date = "2018-7-25")
POST http://localhost:8282/exchange (insert new exchange) with body params like this :
    - date => '2017-08-25'
    - rate => '100000000'
    - from => 'BTC'
    - to   => 'IDR'
POST http://localhost:8282/track (insert new track without exchang eand date)
    - from => 'WAVES'
    - to   => 'IDR'
PUT http://localhost:8282/exchange (delete track by {from} and {to} field)
    - from => 'WAVES'
    - to => 'IDR'
